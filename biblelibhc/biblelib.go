package biblelibhc

import (
	"bible-app/lattice"
	"strings"
)

// export Init
func Init(bibleText []byte) {
}

// export QueryBook
func Query(searchQuery string, versesWanted uint32) lattice.QueryResponse {
	var queryResponse lattice.QueryResponse

	if strings.HasPrefix(searchQuery, "bcv:") {
		queryResponse.TextTags = []lattice.TextTag{
			lattice.TextTag{
				StartIndex: 779,
				EndIndex:   0,
				TagValue:   "Book:Psalms/003:007",
				TagType:    4,
			},
			lattice.TextTag{
				StartIndex: 913,
				EndIndex:   0,
				TagValue:   "Book:Psalms/003:008",
				TagType:    4,
			},
			lattice.TextTag{
				StartIndex: 985,
				EndIndex:   0,
				TagValue:   "Book:Psalms/004:001",
				TagType:    4,
			},
			lattice.TextTag{
				StartIndex: 985,
				EndIndex:   0,
				TagValue:   "For the Chief Musician; on stringed instruments. A Psalm by David.",
				TagType:    5,
			},
			lattice.TextTag{
				StartIndex: 1107,
				EndIndex:   0,
				TagValue:   "Book:Psalms/004:002",
				TagType:    4,
			},
			lattice.TextTag{
				StartIndex: 1231,
				EndIndex:   0,
				TagValue:   "Book:Psalms/004:003",
				TagType:    4,
			},
			lattice.TextTag{
				StartIndex: 1334,
				EndIndex:   0,
				TagValue:   "Book:Psalms/004:004",
				TagType:    4,
			},
			lattice.TextTag{
				StartIndex: 1422,
				EndIndex:   0,
				TagValue:   "Book:Psalms/004:005",
				TagType:    4,
			},
			lattice.TextTag{
				StartIndex: 1489,
				EndIndex:   0,
				TagValue:   "Book:Psalms/004:006",
				TagType:    4,
			},
			lattice.TextTag{
				StartIndex: 1498,
				EndIndex:   1525,
				TagValue:   "",
				TagType:    6,
			},
			lattice.TextTag{
				StartIndex: 1578,
				EndIndex:   0,
				TagValue:   "Book:Psalms/004:007",
				TagType:    4,
			},
			lattice.TextTag{
				StartIndex: 1675,
				EndIndex:   0,
				TagValue:   "Book:Psalms/004:008",
				TagType:    4,
			},
			lattice.TextTag{
				StartIndex: 1771,
				EndIndex:   0,
				TagValue:   "Book:Psalms/005:001",
				TagType:    4,
			},
			lattice.TextTag{
				StartIndex: 1771,
				EndIndex:   0,
				TagValue:   "For the Chief Musician, with the flutes. A Psalm by David.",
				TagType:    5,
			},
		}
		queryResponse.Body.WriteString("Now therefore be wise, you kings.  Be instructed, you judges of the earth.  Serve Yahweh with fear, and rejoice with trembling.  Give sincere homage, lest he be angry, and you perish in the way, for his wrath will soon be kindled.  Blessed are all those who take refuge in him.  Yahweh, how my adversaries have increased!  Many are those who rise up against me.  Many there are who say of my soul, \"There is no help for him in God.\"  Selah.  But you, Yahweh, are a shield around me, my glory, and the one who lifts up my head.  I cry to Yahweh with my voice, and he answers me out of his holy hill.  Selah.  I laid myself down and slept.  I awakened; for Yahweh sustains me.  I will not be afraid of tens of thousands of people who have set themselves against me on every side.  Arise, Yahweh!  Save me, my God!  For you have struck all of my enemies on the cheek bone.  You have broken the teeth of the wicked.  Salvation belongs to Yahweh.  Your blessing be on your people.  Selah.  Answer me when I call, God of my righteousness.  Give me relief from my distress.  Have mercy on me, and hear my prayer.  You sons of men, how long shall my glory be turned into dishonor?  Will you love vanity, and seek after falsehood?  Selah.  But know that Yahweh has set apart for himself him who is godly: Yahweh will hear when I call to him.  Stand in awe, and don't sin.  Search your own heart on your bed, and be still.  Selah.  Offer the sacrifices of righteousness.  Put your trust in Yahweh.  Many say, \"Who will show us any good?\"  Yahweh, let the light of your face shine on us.  You have put gladness in my heart, more than when their grain and their new wine are increased.  In peace I will both lay myself down and sleep, for you, Yahweh alone, make me live in safety.  Give ear to my words, Yahweh.  Consider my meditation.")
	} else if strings.HasPrefix(searchQuery, "phrase:") || strings.HasPrefix(searchQuery, "seqphrase:") {
		if strings.HasPrefix(searchQuery, "seqphrase:") {
			queryResponse.Body.WriteString("Psalms/002:002,Daniel/009:025,Daniel/009:026,Matthew/001:001,Matthew/001:017,Matthew/001:018,Matthew/001:019,Matthew/002:004,Matthew/011:002,Matthew/016:016,Matthew/016:020,Matthew/022:041,Matthew/023:008,Matthew/023:010,Matthew/024:005,Matthew/024:023,Matthew/026:064,Matthew/026:067,Matthew/027:018,Matthew/027:022,Mark/001:001,Mark/008:030,Mark/012:035,Mark/013:021,Mark/014:062,Mark/015:032,Luke/002:011,Luke/002:026,Luke/002:027,Luke/003:015,Luke/004:041,Luke/004:042,Luke/009:021,Luke/020:041,Luke/022:066,Luke/023:002,Luke/023:036,Luke/023:039,Luke/024:026,Luke/024:046,John/001:018,John/001:020,John/001:025,John/001:042,John/003:028,John/004:025,John/004:030,John/004:042,John/006:069,John/007:027,John/007:028,John/007:032,John/007:041,John/007:042,John/007:043,John/009:022,John/010:025,John/011:028,John/012:034,John/017:003,John/020:030,Acts/002:030,Acts/002:036,Acts/002:038,Acts/003:007,Acts/003:018,Acts/003:019,Acts/004:008,Acts/004:026,Acts/005:042,Acts/008:005,Acts/008:012,Acts/008:017,Acts/008:038,Acts/009:020,Acts/009:022,Acts/009:034,Acts/010:036,Acts/010:048,Acts/011:017,Acts/015:011,Acts/015:024,Acts/016:019,Acts/016:031,Acts/017:002,Acts/018:005,Acts/018:028,Acts/020:022,Acts/024:024,Acts/026:022,Acts/028:030,Romans/001:001,Romans/001:008,Romans/001:016,Romans/002:013,Romans/003:021,Romans/003:023,Romans/005:001,Romans/005:006,Romans/005:008,Romans/005:011,Romans/005:016,Romans/005:017,Romans/005:020,Romans/006:003,Romans/006:004,Romans/006:008,Romans/006:011,Romans/006:023,Romans/007:004,Romans/007:025,Romans/008:001,Romans/008:002,Romans/008:010,Romans/008:011,Romans/008:012,Romans/008:016,Romans/008:035,Romans/008:036,Romans/008:038,Romans/009:001,Romans/009:003,Romans/010:004,Romans/010:008,Romans/012:004,Romans/013:014,Romans/014:009,Romans/014:011,Romans/014:016,Romans/014:018,Romans/015:001,Romans/015:003,Romans/015:005,Romans/015:007,Romans/015:008,Romans/015:015,Romans/015:017,Romans/015:018,Romans/015:029,Romans/015:030,Romans/016:003,Romans/016:006,Romans/016:007,Romans/016:009,Romans/016:010,Romans/016:017,Romans/016:018,Romans/016:021,Romans/016:024,1 Corinthians/001:001,1 Corinthians/001:004,1 Corinthians/001:009,1 Corinthians/001:011,1 Corinthians/001:012,1 Corinthians/001:013,1 Corinthians/001:017,1 Corinthians/001:022,1 Corinthians/001:030,1 Corinthians/002:002,1 Corinthians/003:001,1 Corinthians/003:011,1 Corinthians/003:023,1 Corinthians/004:010,1 Corinthians/004:015,1 Corinthians/004:016,1 Corinthians/004:017,1 Corinthians/005:004,1 Corinthians/005:008,1 Corinthians/006:015,1 Corinthians/006:016,1 Corinthians/008:005,1 Corinthians/008:011,1 Corinthians/008:012,1 Corinthians/009:002,1 Corinthians/009:013,1 Corinthians/009:019,1 Corinthians/009:020,1 Corinthians/010:005,1 Corinthians/010:016,1 Corinthians/010:017,1 Corinthians/011:001,1 Corinthians/011:003,1 Corinthians/012:012,1 Corinthians/012:027,1 Corinthians/015:012,1 Corinthians/015:013,1 Corinthians/015:014,1 Corinthians/015:015,1 Corinthians/015:016,1 Corinthians/015:017,1 Corinthians/015:018,1 Corinthians/015:019,1 Corinthians/015:020,1 Corinthians/015:022,1 Corinthians/015:023,1 Corinthians/015:031,1 Corinthians/015:057,1 Corinthians/016:022,1 Corinthians/016:023,1 Corinthians/016:024,2 Corinthians/001:001,2 Corinthians/001:003,2 Corinthians/001:005,2 Corinthians/001:019,2 Corinthians/001:021,2 Corinthians/002:011,2 Corinthians/002:012,2 Corinthians/002:014,2 Corinthians/002:015,2 Corinthians/003:001,2 Corinthians/003:002,2 Corinthians/003:004,2 Corinthians/003:014,2 Corinthians/004:003,2 Corinthians/004:005,2 Corinthians/005:010,2 Corinthians/005:014,2 Corinthians/005:017,2 Corinthians/005:018,2 Corinthians/005:019,2 Corinthians/005:020,2 Corinthians/005:021,2 Corinthians/006:015,2 Corinthians/008:009,2 Corinthians/008:024,2 Corinthians/009:012,2 Corinthians/010:001,2 Corinthians/010:003,2 Corinthians/010:015,2 Corinthians/011:003,2 Corinthians/011:004,2 Corinthians/011:010,2 Corinthians/011:023,2 Corinthians/011:031,2 Corinthians/012:002,2 Corinthians/012:010,2 Corinthians/012:020,2 Corinthians/013:002,2 Corinthians/013:006,2 Corinthians/013:014,Galatians/001:001,Galatians/001:003,Galatians/001:006,Galatians/001:008,Galatians/001:011,Galatians/001:012,Galatians/001:022,Galatians/002:004,Galatians/002:015,Galatians/002:017,Galatians/002:020,Galatians/003:001,Galatians/003:002,Galatians/003:013,Galatians/003:014,Galatians/003:017,Galatians/003:018,Galatians/003:022,Galatians/003:024,Galatians/003:026,Galatians/003:027,Galatians/003:028,Galatians/004:007,Galatians/004:014,Galatians/004:019,Galatians/005:001,Galatians/005:002,Galatians/005:004,Galatians/005:006,Galatians/005:024,Galatians/006:002,Galatians/006:012,Galatians/006:014,Galatians/006:015,Galatians/006:018,Ephesians/001:001,Ephesians/001:003,Ephesians/001:015,Ephesians/002:004,Ephesians/002:010,Ephesians/002:011,Ephesians/002:013,Ephesians/002:019,Ephesians/003:001,Ephesians/003:008,Ephesians/003:014,Ephesians/003:020,Ephesians/004:007,Ephesians/004:011,Ephesians/004:020,Ephesians/004:032,Ephesians/005:002,Ephesians/005:005,Ephesians/005:014,Ephesians/005:018,Ephesians/005:023,Ephesians/005:024,Ephesians/005:025,Ephesians/005:032,Ephesians/006:005,Ephesians/006:023,Ephesians/006:024,Philippians/001:001,Philippians/001:006,Philippians/001:008,Philippians/001:009,Philippians/001:013,Philippians/001:015,Philippians/001:016,Philippians/001:019,Philippians/001:020,Philippians/001:021,Philippians/001:023,Philippians/001:025,Philippians/001:027,Philippians/001:029,Philippians/002:001,Philippians/002:005,Philippians/002:009,Philippians/002:014,Philippians/002:021,Philippians/002:029,Philippians/003:003,Philippians/003:007,Philippians/003:008,Philippians/003:012,Philippians/003:014,Philippians/003:018,Philippians/003:020,Philippians/004:007,Philippians/004:013,Philippians/004:019,Philippians/004:021,Philippians/004:023,Colossians/001:001,Colossians/001:003,Colossians/001:004,Colossians/001:024,Colossians/001:027,Colossians/002:001,Colossians/002:005,Colossians/002:006,Colossians/002:008,Colossians/002:009,Colossians/002:020,Colossians/003:001,Colossians/003:003,Colossians/003:004,Colossians/003:009,Colossians/003:012,Colossians/003:016,Colossians/003:023,Colossians/004:002,Colossians/004:012,1 Thessalonians/001:001,1 Thessalonians/001:002,1 Thessalonians/002:005,1 Thessalonians/002:014,1 Thessalonians/002:020,1 Thessalonians/003:001,1 Thessalonians/003:011,1 Thessalonians/004:017,1 Thessalonians/005:009,1 Thessalonians/005:018,1 Thessalonians/005:024,1 Thessalonians/005:028,2 Thessalonians/001:001,2 Thessalonians/001:011,2 Thessalonians/002:001,2 Thessalonians/002:013,2 Thessalonians/002:016,2 Thessalonians/003:005,2 Thessalonians/003:006,2 Thessalonians/003:012,2 Thessalonians/003:018,1 Timothy/001:001,1 Timothy/001:012,1 Timothy/001:014,1 Timothy/001:015,1 Timothy/001:016,1 Timothy/002:005,1 Timothy/003:013,1 Timothy/004:006,1 Timothy/005:011,1 Timothy/005:021,1 Timothy/006:003,1 Timothy/006:013,2 Timothy/001:001,2 Timothy/001:008,2 Timothy/001:013,2 Timothy/002:001,2 Timothy/002:003,2 Timothy/002:008,2 Timothy/002:010,2 Timothy/002:019,2 Timothy/003:012,2 Timothy/003:015,2 Timothy/004:001,2 Timothy/004:022,Titus/001:001,Titus/002:011,Titus/003:004,Philemon/001:001,Philemon/001:004,Philemon/001:008,Philemon/001:023,Philemon/001:025,Hebrews/003:005,Hebrews/003:014,Hebrews/005:005,Hebrews/006:001,Hebrews/009:011,Hebrews/009:013,Hebrews/009:024,Hebrews/009:027,Hebrews/010:010,Hebrews/011:024,Hebrews/013:008,Hebrews/013:020,James/001:001,James/002:001,1 Peter/001:001,1 Peter/001:003,1 Peter/001:006,1 Peter/001:010,1 Peter/001:013,1 Peter/001:017,1 Peter/002:005,1 Peter/002:021,1 Peter/003:015,1 Peter/003:018,1 Peter/003:021,1 Peter/004:001,1 Peter/004:012,1 Peter/004:014,1 Peter/005:001,1 Peter/005:010,2 Peter/001:001,2 Peter/001:002,2 Peter/001:008,2 Peter/001:011,2 Peter/001:013,2 Peter/001:016,2 Peter/002:020,2 Peter/003:018,1 John/001:004,1 John/001:007,1 John/002:002,1 John/002:022,1 John/003:023,1 John/004:002,1 John/005:001,1 John/005:006,1 John/005:020,2 John/001:001,2 John/001:007,2 John/001:009,Jude/001:001,Jude/001:004,Jude/001:017,Jude/001:021,Revelation/001:001,Revelation/001:004,Revelation/001:009,Revelation/011:015,Revelation/012:010,Revelation/020:005,Revelation/020:007,Revelation/022:021")
		} else {
			queryResponse.Body.WriteString("Genesis/024:009,Genesis/026:004,Genesis/050:024,Exodus/006:008,Exodus/033:001,Numbers/032:010,Deuteronomy/001:008,Deuteronomy/006:010,Deuteronomy/009:005,Deuteronomy/029:010,Deuteronomy/030:019,Deuteronomy/034:004")
		}
	} else if strings.HasPrefix(searchQuery, "booklist:") {
		queryResponse.Body.WriteString("Genesis,Exodus,Leviticus,Numbers,Deuteronomy,Joshua,Judges,Ruth,1 Samuel,2 Samuel,1 Kings,2 Kings,1 Chronicles,2 Chronicles,Ezra,Nehemiah,Esther,Job,Psalms,Proverbs,Ecclesiastes,Song of Solomon,Isaiah,Jeremiah,Lamentations,Ezekiel,Daniel,Hosea,Joel,Amos,Obadiah,Jonah,Micah,Nahum,Habakkuk,Zephaniah,Haggai,Zechariah,Malachi,Matthew,Mark,Luke,John,Acts,Romans,1 Corinthians,2 Corinthians,Galatians,Ephesians,Philippians,Colossians,1 Thessalonians,2 Thessalonians,1 Timothy,2 Timothy,Titus,Philemon,Hebrews,James,1 Peter,2 Peter,1 John,2 John,3 John,Jude,Revelation,")
	} else if strings.HasPrefix(searchQuery, "bclist:") {
		queryResponse.Body.WriteString("Psalms/001,Psalms/002,Psalms/003,Psalms/004,Psalms/005,Psalms/006,Psalms/007,Psalms/008,Psalms/009,Psalms/010,Psalms/011,Psalms/012,Psalms/013,Psalms/014,Psalms/015,Psalms/016,Psalms/017,Psalms/018,Psalms/019,Psalms/020,Psalms/021,Psalms/022,Psalms/023,Psalms/024,Psalms/025,Psalms/026,Psalms/027,Psalms/028,Psalms/029,Psalms/030,Psalms/031,Psalms/032,Psalms/033,Psalms/034,Psalms/035,Psalms/036,Psalms/037,Psalms/038,Psalms/039,Psalms/040,Psalms/041,Psalms/042,Psalms/043,Psalms/044,Psalms/045,Psalms/046,Psalms/047,Psalms/048,Psalms/049,Psalms/050,Psalms/051,Psalms/052,Psalms/053,Psalms/054,Psalms/055,Psalms/056,Psalms/057,Psalms/058,Psalms/059,Psalms/060,Psalms/061,Psalms/062,Psalms/063,Psalms/064,Psalms/065,Psalms/066,Psalms/067,Psalms/068,Psalms/069,Psalms/070,Psalms/071,Psalms/072,Psalms/073,Psalms/074,Psalms/075,Psalms/076,Psalms/077,Psalms/078,Psalms/079,Psalms/080,Psalms/081,Psalms/082,Psalms/083,Psalms/084,Psalms/085,Psalms/086,Psalms/087,Psalms/088,Psalms/089,Psalms/090,Psalms/091,Psalms/092,Psalms/093,Psalms/094,Psalms/095,Psalms/096,Psalms/097,Psalms/098,Psalms/099,Psalms/100,Psalms/101,Psalms/102,Psalms/103,Psalms/104,Psalms/105,Psalms/106,Psalms/107,Psalms/108,Psalms/109,Psalms/110,Psalms/111,Psalms/112,Psalms/113,Psalms/114,Psalms/115,Psalms/116,Psalms/117,Psalms/118,Psalms/119,Psalms/120,Psalms/121,Psalms/122,Psalms/123,Psalms/124,Psalms/125,Psalms/126,Psalms/127,Psalms/128,Psalms/129,Psalms/130,Psalms/131,Psalms/132,Psalms/133,Psalms/134,Psalms/135,Psalms/136,Psalms/137,Psalms/138,Psalms/139,Psalms/140,Psalms/141,Psalms/142,Psalms/143,Psalms/144,Psalms/145,Psalms/146,Psalms/147,Psalms/148,Psalms/149,Psalms/150")
	} else if strings.HasPrefix(searchQuery, "bcvlist:") {
		queryResponse.Body.WriteString("Psalms/002:001,Psalms/002:002,Psalms/002:003,Psalms/002:004,Psalms/002:005,Psalms/002:006,Psalms/002:007,Psalms/002:008,Psalms/002:009,Psalms/002:010,Psalms/002:011,Psalms/002:012")
	}
	return queryResponse
}
