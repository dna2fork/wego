// Copyright © 2017 Makoto Ito
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package word2vec

import (
	"testing"
)

func TestSigmoid(t *testing.T) {
	table := newSigmoidTable()
	f := table.sigmoid(3)
	if !(f >= 0 || f <= 1) {
		t.Errorf("Expected range is 0 < sigmoid(x) < 1, but got %v", f)
	}
}
