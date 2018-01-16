// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package metrics

// Counters
const (
	KafkaPartitionStarted        = "kafka.partition.started"
	KafkaPartitionStopped        = "kafka.partition.stopped"
	KafkaPartitionMessagesIn     = "kafka.partition.messages-in"
	KafkaPartitionAckMgrDups     = "kafka.partition.duplicates"
	KafkaPartitionAckMgrSkipped  = "kafka.partition.skipped"
	KafkaPartitionAckMgrListFull = "kafka.partition.ackmgr.list-full-errors"
	KafkaPartitionGetAckIDErrors = "kafka.partition.get-ackid-errors"
	KafkaPartitionAckErrors      = "kafka.partition.ack-errors"
	KafkaPartitionNacks          = "kafka.partition.nacks"
	KafkaConsumerStarted         = "kafka.consumer.started"
	KafkaConsumerStopped         = "kafka.consumer.stopped"
)

// Gauges
const (
	KafkaPartitionLag          = "kafka.partition.lag"
	KafkaPartitionBacklog      = "kafka.partition.backlog"
	KafkaPartitionReadOffset   = "kafka.partition.read-offset"
	KafkaPartitionCommitOffset = "kafka.partition.commit-offset"
)