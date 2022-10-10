package generator_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type TestCase struct {
	Input    string
	Expected []string
}

var _ = Describe("Reverse", func() {
	When("input data is 'Hello World'", func() {
		It("should return 'dlroW olleH'", func() {
			Expect(main.Reverse("Hello World")).To(Equal("dlroW olleH"))
		})
	})

	When("input data is 'hello world'", func() {
		It("should return 'dlrow olleh'", func() {
			Expect(main.Reverse("hello world")).To(Equal("dlrow olleh"))
		})
	})

	When("input data is '123hello'", func() {
		It("should return 'olleh321'", func() {
			Expect(main.Reverse("123hello")).To(Equal("olleh321"))
		})
	})

	When("input data is 'SEMANGAT PAGI'", func() {
		It("should return 'IGAP TAGNAMES'", func() {
			Expect(main.Reverse("SEMANGAT PAGI")).To(Equal("IGAP TAGNAMES"))
		})
	})

	When("input data is '!**#&$&#*!(@#'", func() {
		It("should return '#@(!*#&$&#**!'", func() {
			Expect(main.Reverse("!**#&$&#*!(@#")).To(Equal("#@(!*#&$&#**!"))
		})
	})
})

var _ = Describe("Generate", func() {
	When("data input is contain number or and symbol", func() {
		It("should just reverse data", func() {
			testCases := []TestCase{
				{
					Input:    "11111222222333333",
					Expected: []string{"33333322222211111"},
				},
				{
					Input:    "!@$#&*(%#(*&$(#$&^(1231238193812903",
					Expected: []string{"3092183918321321(^&$#($&*(#%(*&#$@!"},
				},
			}

			for _, test := range testCases {
				actual := main.Generate(test.Input)
				Expect(actual).To(Equal(test.Expected[0]))
			}
		})
	})

	When("data input is contain consonant letter", func() {
		It("should reverse data and change letter upper to lower and vice versa", func() {
			testCases := []TestCase{
				{
					Input:    "hjghgjhjghgjhjhghrggtgklkjlkpmnpmvpnvFWJFWKFJWKFWHKFHWKFWKWFJK",
					Expected: []string{"kjfwkwfkwhfkhwfkwjfkwfjwfVNPVMPNMPKLJKLKGTGGRHGHJHJGHGJHJGHGJH"},
				},
				{
					Input:    "hjghgjhjghgjhjhghrggtgklkjlkpmnpmvpnvhjghgjhjghgjhjhghrggtgklkjlkpmnpmvpnvhjghgjhjghgjhjhghrggtgklkjlkpmnpmvpnvhjWSWDWFFGGHJHGJWDJWHGJWJHDJWGFWJFHWJGJWHFJHWJFGWJFGWFHFWJpmvpnvhjghgjhjghgjhjhghrggtgklkjlkpmnpmvpnvhjghgjhjghgjhjhghrggtgklkjlkpmnpmvpnvhjghgjhjghgjhjhghrggtgklkjlkpmnpmvpnvhjghgjhRFRJKHGNMBBlkjlkpmnpmvpnv",
					Expected: []string{"VNPVMPNMPKLJKLbbmnghkjrfrHJGHGJHVNPVMPNMPKLJKLKGTGGRHGHJHJGHGJHJGHGJHVNPVMPNMPKLJKLKGTGGRHGHJHJGHGJHJGHGJHVNPVMPNMPKLJKLKGTGGRHGHJHJGHGJHJGHGJHVNPVMPjwfhfwgfjwgfjwhjfhwjgjwhfjwfgwjdhjwjghwjdwjghjhggffwdwswJHVNPVMPNMPKLJKLKGTGGRHGHJHJGHGJHJGHGJHVNPVMPNMPKLJKLKGTGGRHGHJHJGHGJHJGHGJHVNPVMPNMPKLJKLKGTGGRHGHJHJGHGJHJGHGJH"},
				},
			}

			for _, test := range testCases {
				actual := main.Generate(test.Input)
				Expect(actual).To(Equal(test.Expected[0]))
			}
		})
	})

	When("data input is contain vowel letter", func() {
		It("should reverse and change upper to lower and vice versa and change to next vowel", func() {
			testCases := []TestCase{
				{
					Input:    "asfjdafiowjewfwjwijwefjwefjeifjioHFGWUEFHWEKJNWJCBWVWHVBoahfbwibwknvbwbweifbeqwhivwknvwfhwjofhewojfnwjfnwjofn",
					Expected: []string{"NFUJWNFJWNFJUWIHFUJWHFWVNKWVOHWQIBFOIWBWBVNKWBOWBFHEUbvhwvwbcjwnjkiwhfiawgfhUOJFOIJFIWJFIWJOWJWFWIJWUOFEDJFSE"},
				},
				{
					Input:    "asfjdafiowjewfwjwijwefjwefjeifjioHFGWUEFHWEKJNWJCBWVWHVBoahfbwibwknvbwbweifbeqwhivwknvwfhwjofhewojfnwjfnwjofnasfjdafiowjewfwjwijwefjwefjeifjioHFGWUEFHWEKJNWJCBWVWHVBoahfbwibwknvbwbweifbeqwhivwknvwfhwjofhewojfnwjfnwjofnasfjdafiowjewfwjwijwefjwefjeifjioHFGWUEFHWEKJNWJCBWVWHVBoahfbwibwknvbwbweifbeqwhivwknvwfhwjofhewojfnwjfnwjofnasfjdafiowjewfwjwijwefjwefjeifjioHFGWUEFHWEKJNWJCBWVWHVBoahfbwibwknvbwbweifbeqwhivwknvwfhwjofhewojfnwjfnwjofnasfjdafiowjewfwjwijwefjwefjeifjioHFGWUEFHWEKJNWJCBWVWHVBoahfbwibwknvbwbweifbeqwhivwknvwfhwjofhewojfnwjfnwjofnasfjdafiowjewfwjwijwefjwefjeifjioHFGWUEFHWEKJNWJCBWVWHVBoahfbwibwknvbwbweifbeqwhivwknvwfhwjofhewojfnwjfnwjofnasfjdafiowjewfwjwijwefjwefjeifjioHFGWUEFHWEKJNWJCBWVWHVBoahfbwibwknvbwbweifbeqwhivwknvwfhwjofhewojfnwjfnwjofnasfjdafiowjewfwjwijwefjwefjeifjioHFGWUEFHWEKJNWJCBWVWHVBoahfbwibwknvbwbweifbeqwhivwknvwfhwjofhewojfnwjfnwjofnasfjdafiowjewfwjwijwefjwefjeifjioHFGWUEFHWEKJNWJCBWVWHVBoahfbwibwknvbwbweifbeqwhivwknvwfhwjofhewojfnwjfnwjofn",
					Expected: []string{"NFUJWNFJWNFJUWIHFUJWHFWVNKWVOHWQIBFOIWBWBVNKWBOWBFHEUbvhwvwbcjwnjkiwhfiawgfhUOJFOIJFIWJFIWJOWJWFWIJWUOFEDJFSENFUJWNFJWNFJUWIHFUJWHFWVNKWVOHWQIBFOIWBWBVNKWBOWBFHEUbvhwvwbcjwnjkiwhfiawgfhUOJFOIJFIWJFIWJOWJWFWIJWUOFEDJFSENFUJWNFJWNFJUWIHFUJWHFWVNKWVOHWQIBFOIWBWBVNKWBOWBFHEUbvhwvwbcjwnjkiwhfiawgfhUOJFOIJFIWJFIWJOWJWFWIJWUOFEDJFSENFUJWNFJWNFJUWIHFUJWHFWVNKWVOHWQIBFOIWBWBVNKWBOWBFHEUbvhwvwbcjwnjkiwhfiawgfhUOJFOIJFIWJFIWJOWJWFWIJWUOFEDJFSENFUJWNFJWNFJUWIHFUJWHFWVNKWVOHWQIBFOIWBWBVNKWBOWBFHEUbvhwvwbcjwnjkiwhfiawgfhUOJFOIJFIWJFIWJOWJWFWIJWUOFEDJFSENFUJWNFJWNFJUWIHFUJWHFWVNKWVOHWQIBFOIWBWBVNKWBOWBFHEUbvhwvwbcjwnjkiwhfiawgfhUOJFOIJFIWJFIWJOWJWFWIJWUOFEDJFSENFUJWNFJWNFJUWIHFUJWHFWVNKWVOHWQIBFOIWBWBVNKWBOWBFHEUbvhwvwbcjwnjkiwhfiawgfhUOJFOIJFIWJFIWJOWJWFWIJWUOFEDJFSENFUJWNFJWNFJUWIHFUJWHFWVNKWVOHWQIBFOIWBWBVNKWBOWBFHEUbvhwvwbcjwnjkiwhfiawgfhUOJFOIJFIWJFIWJOWJWFWIJWUOFEDJFSENFUJWNFJWNFJUWIHFUJWHFWVNKWVOHWQIBFOIWBWBVNKWBOWBFHEUbvhwvwbcjwnjkiwhfiawgfhUOJFOIJFIWJFIWJOWJWFWIJWUOFEDJFSE"},
				},
			}

			for _, test := range testCases {
				actual := main.Generate(test.Input)
				Expect(actual).To(Equal(test.Expected[0]))
			}
		})
	})

	When("data input contain space character", func() {
		It("should reverse and remove the space character", func() {
			testCases := []TestCase{
				{
					Input:    "HELLO WORLDS SIOU348 1231J*(*()",
					Expected: []string{")(*(*j1321843auossdlruwullih"},
				},
				{
					Input:    "HGLLO      SADKADJ    sakjadla    12328&*(&*&())    adjsak  123123",
					Expected: []string{"321321KESJDE))(&*&(*&82321ELDEJKESjdekdesullgh"},
				},
			}

			for _, test := range testCases {
				actual := main.Generate(test.Input)
				Expect(actual).To(Equal(test.Expected[0]))
			}
		})
	})
})

var _ = Describe("CheckPassword", func() {
	When("length of password is less than 7", func() {
		It("should return 'sangat lemah'", func() {
			testCases := []TestCase{
				{
					Input:    "",
					Expected: []string{"sangat lemah"},
				},
				{
					Input:    "123456",
					Expected: []string{"sangat lemah"},
				},
				{
					Input:    "adm1n",
					Expected: []string{"sangat lemah"},
				},
				{
					Input:    "adm1!@",
					Expected: []string{"sangat lemah"},
				},
			}

			for _, test := range testCases {
				actual := main.CheckPassword(test.Input)
				Expect(actual).To(Equal(test.Expected[0]))
			}
		})
	})

	Context("length of password is more than or equal 7", func() {
		When("password contain letter or number or both", func() {
			It("should return 'lemah'", func() {
				testCases := []TestCase{
					{
						Input:    "123213213123",
						Expected: []string{"lemah"},
					},
					{
						Input:    "adminadmin",
						Expected: []string{"lemah"},
					},
					{
						Input:    "adminadmin123",
						Expected: []string{"lemah"},
					},
				}

				for _, test := range testCases {
					actual := main.CheckPassword(test.Input)
					Expect(actual).To(Equal(test.Expected[0]))
				}
			})
		})
		When("password contain letter and symbol or number and symbol, or both", func() {
			It("should return 'sedang'", func() {
				testCases := []TestCase{
					{
						Input:    "2133!#123)$#",
						Expected: []string{"sedang"},
					},
					{
						Input:    "admimi@#(*#",
						Expected: []string{"sedang"},
					},
					{
						Input:    "adm1238*(#@",
						Expected: []string{"sedang"},
					},
				}

				for _, test := range testCases {
					actual := main.CheckPassword(test.Input)
					Expect(actual).To(Equal(test.Expected[0]))
				}
			})
		})
	})

	When("length password more than or equal 14 and contain letter, number and symbol", func() {
		It("should return 'kuat'", func() {
			testCases := []TestCase{
				{
					Input:    "2133!#123)$#1232132193831",
					Expected: []string{"kuat"},
				},
				{
					Input:    "adminmi@#(*#asdminsadni",
					Expected: []string{"kuat"},
				},
				{
					Input:    "admin1238*(#@123i@DJHDJSUI@*(#)NDKJKJDSNJ",
					Expected: []string{"kuat"},
				},
			}

			for _, test := range testCases {
				actual := main.CheckPassword(test.Input)
				Expect(actual).To(Equal(test.Expected[0]))
			}
		})
	})
})

var _ = Describe("PasswordGenerator", func() {
	It("Should generate password", func() {
		testCases := []TestCase{
			{
				Input:    "123456",
				Expected: []string{"654321", "sangat lemah"},
			},
			{
				Input:    "DFDFgg",
				Expected: []string{"GGfdfd", "sangat lemah"},
			},
			{
				Input:    "Semangat Pagi",
				Expected: []string{"OGEpTEGNEMIs", "lemah"},
			},
			{
				Input:    "2312321389213",
				Expected: []string{"3129831232132", "lemah"},
			},
			{
				Input:    "2133!#123)$#",
				Expected: []string{"#$)321#!3312", "sedang"},
			},
			{
				Input:    "adm1238*(#@",
				Expected: []string{"@#(*8321MDE", "sedang"},
			},
			{
				Input:    "2133!#123)$#1232132193831",
				Expected: []string{"1383912312321#$)321#!3312", "kuat"},
			},
			{
				Input:    "admin1238*(#@123i@DJHDJSUI@*(#)NDKJKJDSNJ",
				Expected: []string{"jnsdjkjkdn)#(*@oasjdhjd@O321@#(*8321NOMDE", "kuat"},
			},
		}

		for _, test := range testCases {
			actual, actual2 := main.PasswordGenerator(test.Input)
			Expect(actual).To(Equal(test.Expected[0]))
			Expect(actual2).To(Equal(test.Expected[1]))

		}
	})
})
