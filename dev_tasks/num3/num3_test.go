package num3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		desc string
		app  Data
		data []string
		want []string
	}{
		{
			desc: "normal",
			app:  Data{},
			data: []string{
				"Standing on one's head at job interviews forms a lasting impression.",
				"The chic gangster liked to start the day with a pink scarf.",
				"He kept telling himself that one day it would all somehow make sense.",
				"Cats are good pets, for they are clean and are not noisy.",
				"I am my aunt's sister's daughter.",
			},
			want: []string{
				"Cats are good pets, for they are clean and are not noisy.",
				"He kept telling himself that one day it would all somehow make sense.",
				"I am my aunt's sister's daughter.",
				"Standing on one's head at job interviews forms a lasting impression.",
				"The chic gangster liked to start the day with a pink scarf.",
			},
		},
		{
			desc: "not numeric order",
			app:  Data{},
			data: []string{
				"1",
				"5",
				"13",
				"23",
				"11",
				"21",
				"31",
			},
			want: []string{
				"1",
				"11",
				"13",
				"21",
				"23",
				"31",
				"5",
			},
		},
		{
			desc: "reverse order",
			app: Data{
				isReverse: true,
			},
			data: []string{
				"1",
				"5",
				"13",
				"23",
				"11",
				"21",
				"31",
			},
			want: []string{
				"5",
				"31",
				"23",
				"21",
				"13",
				"11",
				"1",
			},
		},
		{
			desc: "delete duplicate",
			app: Data{
				DeleteDuplicate: true,
			},
			data: []string{
				"1",
				"1",
				"5",
				"13",
				"23",
				"11",
				"11",
				"21",
				"31",
				"31",
				"31",
			},
			want: []string{
				"1",
				"11",
				"13",
				"21",
				"23",
				"31",
				"5",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.app.sort(tC.data)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestSortColumns(t *testing.T) {
	testCases := []struct {
		desc string
		app  Data
		data []string
		want []string
	}{
		{
			desc: "by 2nd column",
			app: Data{
				Column: 2,
			},
			data: []string{
				"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
				"Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque.",
				"Ut enim ad minima veniam, quis nostrum exercitationem ullam.",
				"But I must explain to you how all this mistaken idea of denouncing pleasure and praising pain.",
				"At vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis.",
			},
			want: []string{
				"At vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis.",
				"But I must explain to you how all this mistaken idea of denouncing pleasure and praising pain.",
				"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
				"Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque.",
				"Ut enim ad minima veniam, quis nostrum exercitationem ullam.",
			},
		},
		{
			desc: "by column out of range",
			app: Data{
				Column: 200,
			},
			data: []string{
				"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
				"Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque.",
				"Ut enim ad minima veniam, quis nostrum exercitationem ullam.",
				"But I must explain to you how all this mistaken idea of denouncing pleasure and praising pain.",
				"At vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis.",
			},
			want: []string{
				"At vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis.",
				"But I must explain to you how all this mistaken idea of denouncing pleasure and praising pain.",
				"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
				"Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque.",
				"Ut enim ad minima veniam, quis nostrum exercitationem ullam.",
			},
		},
		{
			desc: "numbers numeric order",
			app: Data{
				Column:    1,
				IsNumeric: true,
			},
			data: []string{
				"5",
				"23",
				"1",
				"21",
				"31",
				"13",
				"11",
			},
			want: []string{
				"1",
				"5",
				"11",
				"13",
				"21",
				"23",
				"31",
			},
		},
		{
			desc: "by 2nd column, in reverse",
			app: Data{
				Column:    2,
				IsReverse: true,
			},
			data: []string{
				"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
				"Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque.",
				"Ut enim ad minima veniam, quis nostrum exercitationem ullam.",
				"But I must explain to you how all this mistaken idea of denouncing pleasure and praising pain.",
				"At vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis.",
			},
			want: []string{
				"At vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis.",
				"But I must explain to you how all this mistaken idea of denouncing pleasure and praising pain.",
				"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
				"Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque.",
				"Ut enim ad minima veniam, quis nostrum exercitationem ullam.",
			},
		},
		{
			desc: "delete duplicate",
			app: Data{
				Column:          1,
				DeleteDuplicate: true,
			},
			data: []string{
				"1",
				"1",
				"5",
				"13",
				"23",
				"11",
				"11",
				"21",
				"31",
				"31",
				"31",
			},
			want: []string{
				"1",
				"11",
				"13",
				"21",
				"23",
				"31",
				"5",
			},
		},
		{
			desc: "numeric sort, but column starts with letter",
			app: Data{
				Column:    1,
				IsNumeric: true,
			},
			data: []string{
				"d1",
				"ad5",
				"asbv13",
				"sfg23",
				"fa11",
				"gh21",
				"31",
			},
			want: []string{
				"31",
				"ad5",
				"asbv13",
				"d1",
				"fa11",
				"gh21",
				"sfg23",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.app.sortColumns(tC.data)
			assert.Equal(t, tC.want, got)
		})
	}
}
