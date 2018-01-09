/*
 * Copyright 2018 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package controller_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/projectriff/function-controller/pkg/controller"
	_ "github.com/projectriff/function-controller/pkg/controller"
	"github.com/projectriff/function-controller/pkg/controller/mocks"
)

func TestFreshTrackerTracksNothing(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	tracker := mocks.NewMockLagTracker(mockCtrl)
	deployer := mocks.NewMockDeployer(mockCtrl)

	topicInformer := mocks.NewMockTopicInformer(mockCtrl)
	functionInformer := mocks.NewMockFunctionInformer(mockCtrl)

	tracker.EXPECT().Compute()
	ctrl := controller.New(topicInformer, functionInformer, nil, deployer, tracker)

	println(ctrl)
	//	ctrl.Run(make(<-chan struct{}))
}
