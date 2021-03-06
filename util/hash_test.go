/* Copyright 2020 PhysarumSM Development Team
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

 package util_test

import (
	"testing"

	"github.com/PhysarumSM/common/util"
)

func TestIpfsHashBytes(test *testing.T) {
	testString := "Ivan, Michael, Andrew, Henry, and Thomas are pretty cool eh\n"
	expectedHash := "QmX37CLME3bh8oJyLy858CZino5cbRg7Bd1zf5qoen4aQn"

	hash, err := util.IpfsHashBytes([]byte(testString))
	if err != nil {
		test.Fatalf("IpfsHashBytes() failed with error:\n%v", err)
	}

	if hash != expectedHash {
		test.Errorf("Calculated hash %s does not match expected hash %s", hash, expectedHash)
	}
}