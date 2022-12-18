/*
Copyright 2022 Seldon Technologies Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package experiment

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestValidateExperiment(t *testing.T) {
	g := NewGomegaWithT(t)

	type test struct {
		name       string
		store      *ExperimentStore
		experiment *Experiment
		err        error
	}

	getStrPtr := func(val string) *string { return &val }
	tests := []test{
		{
			name: "valid",
			store: &ExperimentStore{
				modelBaselines: map[string]*Experiment{},
				experiments:    map[string]*Experiment{},
			},
			experiment: &Experiment{
				Name:    "a",
				Default: getStrPtr("model1"),
				Candidates: []*Candidate{
					{
						Name: "model1",
					},
					{
						Name: "model2",
					},
				},
			},
		},
		{
			name: "duplicate candidate and mirror",
			store: &ExperimentStore{
				modelBaselines: map[string]*Experiment{},
				experiments:    map[string]*Experiment{},
			},
			experiment: &Experiment{
				Name:    "a",
				Default: getStrPtr("model1"),
				Candidates: []*Candidate{
					{
						Name: "model1",
					},
					{
						Name: "model2",
					},
				},
				Mirror: &Mirror{
					Name: "model2",
				},
			},
			err: &ExperimentNoDuplicates{experimentName: "a", resource: "model2"},
		},
		{
			name: "duplicate candidate",
			store: &ExperimentStore{
				modelBaselines: map[string]*Experiment{},
				experiments:    map[string]*Experiment{},
			},
			experiment: &Experiment{
				Name:    "a",
				Default: getStrPtr("model1"),
				Candidates: []*Candidate{
					{
						Name: "model1",
					},
					{
						Name: "model2",
					},
					{
						Name: "model2",
					},
				},
			},
			err: &ExperimentNoDuplicates{experimentName: "a", resource: "model2"},
		},
		{
			name: "baseline already exists",
			store: &ExperimentStore{
				modelBaselines: map[string]*Experiment{"model1": {Name: "b"}},
				experiments:    map[string]*Experiment{},
			},
			experiment: &Experiment{
				Name:    "a",
				Default: getStrPtr("model1"),
				Candidates: []*Candidate{
					{
						Name: "model1",
					},
					{
						Name: "model2",
					},
				},
			},
			err: &ExperimentBaselineExists{experimentName: "a", name: "model1"},
		},
		{
			name: "baseline already exists but its this model so ignore",
			store: &ExperimentStore{
				modelBaselines: map[string]*Experiment{"model1": {Name: "a"}},
				experiments:    map[string]*Experiment{},
			},
			experiment: &Experiment{
				Name:    "a",
				Default: getStrPtr("model1"),
				Candidates: []*Candidate{
					{
						Name: "model1",
					},
					{
						Name: "model2",
					},
				},
			},
		},
		{
			name: "No Canidadates or mirrors",
			store: &ExperimentStore{
				modelBaselines: map[string]*Experiment{},
				experiments:    map[string]*Experiment{},
			},
			experiment: &Experiment{
				Name: "a",
			},
			err: &ExperimentNoCandidatesOrMirrors{experimentName: "a"},
		},
		{
			name: "No Canidadates but mirror",
			store: &ExperimentStore{
				modelBaselines: map[string]*Experiment{},
				experiments:    map[string]*Experiment{},
			},
			experiment: &Experiment{
				Name: "a",
				Mirror: &Mirror{
					Name: "model",
				},
			},
		},
		{
			name: "Default model is not candidate",
			store: &ExperimentStore{
				modelBaselines: map[string]*Experiment{},
				experiments:    map[string]*Experiment{},
			},
			experiment: &Experiment{
				Name:    "a",
				Default: getStrPtr("model1"),
			},
			err: &ExperimentDefaultNotFound{experimentName: "a", defaultResource: "model1"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.store.validate(test.experiment)
			if test.err != nil {
				g.Expect(err.Error()).To(Equal(test.err.Error()))
			} else {
				g.Expect(err).To(BeNil())
			}
		})
	}
}