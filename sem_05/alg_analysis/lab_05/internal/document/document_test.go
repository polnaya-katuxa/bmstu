package document

import (
	"lab_05/internal/rule"
	"reflect"
	"testing"
)

func TestDocument_applyRules(t *testing.T) {
	type fields struct {
		Name   string
		Text   string
		Tokens []string
	}
	type args struct {
		rules []rule.Rule
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "with 1 end rule sequence",
			fields: fields{
				Tokens: []string{"я", "ем", "кашку", "и", "т", "д"},
			},
			args: args{
				rules: []rule.Rule{
					{
						Option:   []string{"и", "т", "д"},
						Standard: "и так далее",
					},
				},
			},
			want: []string{"я", "ем", "кашку", "и так далее"},
		},
		{
			name: "with 1 start rule sequence",
			fields: fields{
				Tokens: []string{"т", "о", "я", "ем", "кашку"},
			},
			args: args{
				rules: []rule.Rule{
					{
						Option:   []string{"т", "о"},
						Standard: "таким образом",
					},
				},
			},
			want: []string{"таким образом", "я", "ем", "кашку"},
		},
		{
			name: "with 0 rule sequences",
			fields: fields{
				Tokens: []string{"я", "ем", "кашку", "и", "пью", "компотик"},
			},
			args: args{
				rules: []rule.Rule{
					{
						Option:   []string{"и", "т", "д"},
						Standard: "и так далее",
					},
				},
			},
			want: []string{"я", "ем", "кашку", "и", "пью", "компотик"},
		},
		{
			name: "with all rule sequences",
			fields: fields{
				Tokens: []string{"я", "ем", "кашку", "и", "пью", "компотик",
					"и", "т", "д", "и", "пр", "и", "т", "п"},
			},
			args: args{
				rules: []rule.Rule{
					{
						Option:   []string{"и", "т", "д"},
						Standard: "и так далее",
					},
					{
						Option:   []string{"и", "т", "п"},
						Standard: "и тому подобное",
					},
					{
						Option:   []string{"и", "пр"},
						Standard: "и прочее",
					},
				},
			},
			want: []string{"я", "ем", "кашку", "и", "пью", "компотик",
				"и так далее", "и прочее", "и тому подобное"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := &Document{
				Name:   tt.fields.Name,
				Text:   tt.fields.Text,
				Tokens: tt.fields.Tokens,
			}
			if doc.ApplyRules(tt.args.rules) != nil {
				t.Errorf("error")
			}

			if !reflect.DeepEqual(doc.Tokens, tt.want) {
				t.Errorf("ApplyRules() = %v, want %v", doc.Tokens, tt.want)
			}
		})
	}
}

func TestDocument_Tokenize(t *testing.T) {
	type fields struct {
		Name   string
		Text   string
		Tokens []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "usual test",
			fields: fields{
				Name:   "",
				Text:   "здравствуйте, как дела? что делаете?    какая-то собака!!!",
				Tokens: nil,
			},
			want: []string{"здравствуйте", "как", "дела", "что", "делаете", "какая-то", "собака"},
		},
		{
			name: "1 token",
			fields: fields{
				Name:   "",
				Text:   "здравствуйте!",
				Tokens: nil,
			},
			want: []string{"здравствуйте"},
		},
		{
			name: "no tokens",
			fields: fields{
				Name:   "",
				Text:   "",
				Tokens: nil,
			},
			want: nil,
		},
		{
			name: "no tokens only sym",
			fields: fields{
				Name:   "",
				Text:   "!!!  .  - ?",
				Tokens: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := &Document{
				Name:   tt.fields.Name,
				Text:   tt.fields.Text,
				Tokens: tt.fields.Tokens,
			}
			doc.Tokenize()
			if tt.want == nil && doc.Tokens != nil {
				t.Errorf("Tokenize() = %v, want %v", doc.Tokens, tt.want)
			} else {
				for i, tok := range tt.want {
					if tok != doc.Tokens[i] {
						t.Errorf("Tokenize() = %v, want %v", doc.Tokens[i], tok)
					}
				}
			}
		})
	}
}
