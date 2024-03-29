// Copyright 2018 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package aggfuncs

import (
	"bytes"
	"sync/atomic"

	"github.com/hanchuanchuan/goInception/expression"
	"github.com/hanchuanchuan/goInception/sessionctx"
	"github.com/hanchuanchuan/goInception/util/chunk"
	"github.com/hanchuanchuan/goInception/util/set"
	"github.com/pingcap/errors"
	"modernc.org/mathutil"
)

type baseGroupConcat4String struct {
	baseAggFunc

	sep    string
	maxLen uint64
	// According to MySQL, a 'group_concat' function generates exactly one 'truncated' warning during its life time, no matter
	// how many group actually truncated. 'truncated' acts as a sentinel to indicate whether this warning has already been
	// generated.
	truncated *int32
}

func (e *baseGroupConcat4String) AppendFinalResult2Chunk(sctx sessionctx.Context, pr PartialResult, chk *chunk.Chunk) error {
	p := (*partialResult4GroupConcat)(pr)
	if p.buffer == nil {
		chk.AppendNull(e.ordinal)
		return nil
	}
	chk.AppendString(e.ordinal, p.buffer.String())
	return nil
}

func (e *baseGroupConcat4String) truncatePartialResultIfNeed(sctx sessionctx.Context, buffer *bytes.Buffer) (err error) {
	if e.maxLen > 0 && uint64(buffer.Len()) > e.maxLen {
		i := mathutil.MaxInt
		if uint64(i) > e.maxLen {
			i = int(e.maxLen)
		}
		buffer.Truncate(i)
		if atomic.CompareAndSwapInt32(e.truncated, 0, 1) {
			if !sctx.GetSessionVars().StmtCtx.TruncateAsWarning {
				return expression.ErrCutValueGroupConcat.GenWithStackByArgs(e.args[0].String())
			}
			sctx.GetSessionVars().StmtCtx.AppendWarning(expression.ErrCutValueGroupConcat.GenWithStackByArgs(e.args[0].String()))
		}
	}
	return nil
}

type basePartialResult4GroupConcat struct {
	buffer *bytes.Buffer
}

type partialResult4GroupConcat struct {
	basePartialResult4GroupConcat
}

type groupConcat struct {
	baseGroupConcat4String
}

func (e *groupConcat) AllocPartialResult() PartialResult {
	return PartialResult(new(partialResult4GroupConcat))
}

func (e *groupConcat) ResetPartialResult(pr PartialResult) {
	p := (*partialResult4GroupConcat)(pr)
	p.buffer = nil
}

func (e *groupConcat) UpdatePartialResult(sctx sessionctx.Context, rowsInGroup []chunk.Row, pr PartialResult) (err error) {
	p := (*partialResult4GroupConcat)(pr)
	v, isNull, preLen := "", false, 0
	for _, row := range rowsInGroup {
		if p.buffer != nil && p.buffer.Len() != 0 {
			preLen = p.buffer.Len()
			p.buffer.WriteString(e.sep)
		}
		for _, arg := range e.args {
			v, isNull, err = arg.EvalString(sctx, row)
			if err != nil {
				return errors.Trace(err)
			}
			if isNull {
				break
			}
			if p.buffer == nil {
				p.buffer = &bytes.Buffer{}
			}
			p.buffer.WriteString(v)
		}
		if isNull {
			if p.buffer != nil {
				p.buffer.Truncate(preLen)
			}
			continue
		}
	}
	if p.buffer != nil {
		return e.truncatePartialResultIfNeed(sctx, p.buffer)
	}
	return nil
}

func (e *groupConcat) MergePartialResult(sctx sessionctx.Context, src, dst PartialResult) error {
	p1, p2 := (*partialResult4GroupConcat)(src), (*partialResult4GroupConcat)(dst)
	if p1.buffer == nil {
		return nil
	}
	if p2.buffer == nil {
		p2.buffer = p1.buffer
		return nil
	}
	p2.buffer.WriteString(e.sep)
	p2.buffer.WriteString(p1.buffer.String())
	return e.truncatePartialResultIfNeed(sctx, p2.buffer)
}

// SetTruncated will be called in `executorBuilder#buildHashAgg` with duck-type.
func (e *groupConcat) SetTruncated(t *int32) {
	e.truncated = t
}

// GetTruncated will be called in `executorBuilder#buildHashAgg` with duck-type.
func (e *groupConcat) GetTruncated() *int32 {
	return e.truncated
}

type partialResult4GroupConcatDistinct struct {
	basePartialResult4GroupConcat
	valsBuf *bytes.Buffer
	valSet  set.StringSet
}

type groupConcatDistinct struct {
	baseGroupConcat4String
}

func (e *groupConcatDistinct) AllocPartialResult() PartialResult {
	p := new(partialResult4GroupConcatDistinct)
	p.valsBuf = &bytes.Buffer{}
	p.valSet = set.NewStringSet()
	return PartialResult(p)
}

func (e *groupConcatDistinct) ResetPartialResult(pr PartialResult) {
	p := (*partialResult4GroupConcatDistinct)(pr)
	p.buffer, p.valSet = nil, set.NewStringSet()
}

func (e *groupConcatDistinct) UpdatePartialResult(sctx sessionctx.Context, rowsInGroup []chunk.Row, pr PartialResult) (err error) {
	p := (*partialResult4GroupConcatDistinct)(pr)
	v, isNull := "", false
	for _, row := range rowsInGroup {
		p.valsBuf.Reset()
		for _, arg := range e.args {
			v, isNull, err = arg.EvalString(sctx, row)
			if err != nil {
				return errors.Trace(err)
			}
			if isNull {
				break
			}
			p.valsBuf.WriteString(v)
		}
		if isNull {
			continue
		}
		joinedVals := p.valsBuf.String()
		if p.valSet.Exist(joinedVals) {
			continue
		}
		p.valSet.Insert(joinedVals)
		// write separator
		if p.buffer == nil {
			p.buffer = &bytes.Buffer{}
		} else {
			p.buffer.WriteString(e.sep)
		}
		// write values
		p.buffer.WriteString(joinedVals)
	}
	if p.buffer != nil {
		return e.truncatePartialResultIfNeed(sctx, p.buffer)
	}
	return nil
}

// SetTruncated will be called in `executorBuilder#buildHashAgg` with duck-type.
func (e *groupConcatDistinct) SetTruncated(t *int32) {
	e.truncated = t
}

// GetTruncated will be called in `executorBuilder#buildHashAgg` with duck-type.
func (e *groupConcatDistinct) GetTruncated() *int32 {
	return e.truncated
}
