package main

import (
	"fmt"
	"slices"
)

var a = []int{
	77221,
	61169,
	49546,
	11688,
	15820,
	63235,
	28850,
	80735,
	42616,
	55068,
	30851,
	14907,
	64732,
	16709,
	67643,
	44080,
	96480,
	99450,
	21169,
	58096,
	53828,
	47026,
	90101,
	45773,
	71303,
	45304,
	96528,
	19794,
	74831,
	88912,
	98961,
	15788,
	16350,
	96806,
	99959,
	58887,
	42568,
	29278,
	94028,
	42810,
	97490,
	89051,
	73916,
	97751,
	73608,
	62426,
	29429,
	11070,
	78656,
	31308,
	96166,
	36906,
	99662,
	66402,
	37085,
	80907,
	61650,
	84223,
	83891,
	15973,
	68843,
	98854,
	42426,
	25255,
	44559,
	93909,
	28070,
	24880,
	29338,
	64883,
	66924,
	32984,
	32238,
	38816,
	58770,
	74953,
	70592,
	93111,
	31065,
	67370,
	18314,
	71075,
	21544,
	59137,
	20440,
	14393,
	72130,
	43828,
	23262,
	11923,
	47754,
	96043,
	27162,
	20838,
	75009,
	92099,
	22054,
	43518,
	68189,
	50520,
	51425,
	42998,
	52443,
	37256,
	16648,
	81123,
	46974,
	36328,
	76373,
	82592,
	70600,
	27116,
	28090,
	86348,
	53145,
	30203,
	23347,
	76433,
	28046,
	58920,
	25518,
	12791,
	61967,
	98022,
	24326,
	41095,
	24045,
	31274,
	78021,
	63029,
	75684,
	31258,
	40561,
	20457,
	14783,
	39433,
	34774,
	10455,
	26338,
	55227,
	36963,
	31910,
	94894,
	23166,
	48884,
	73036,
	83519,
	50205,
	85926,
	69194,
	25529,
	10182,
	13638,
	83588,
	75911,
	44692,
	43278,
	48148,
	86205,
	82605,
	52770,
	89068,
	66452,
	50217,
	64914,
	76460,
	43026,
	44179,
	47653,
	78061,
	14475,
	77868,
	99645,
	86267,
	58072,
	96428,
	77501,
	52344,
	49534,
	46925,
	40005,
	17329,
	17585,
	32296,
	81836,
	89435,
	55149,
	56971,
	26342,
	57894,
	49762,
	48949,
	40453,
	67613,
	27110,
	25231,
	92546,
	70074,
	39649,
	93231,
	60008,
	18449,
	68828,
	18908,
	55426,
	62522,
	60923,
	85721,
	35153,
	20053,
	24233,
	36237,
	92077,
	30105,
	37032,
	12745,
	41563,
	24700,
	60226,
	50713,
	29972,
	79176,
	40017,
	90724,
	51608,
	46178,
	26151,
	19569,
	34368,
	18948,
	62091,
	89749,
	68315,
	52396,
	14981,
	36038,
	50048,
	60796,
	69942,
	80976,
	56722,
	49489,
	60768,
	70753,
	15233,
	94697,
	40958,
	91709,
	83586,
	88140,
	56993,
	17980,
	12421,
	30183,
	50452,
	91314,
	69995,
	16588,
	22432,
	86088,
	46361,
	97178,
	37444,
	12275,
	47016,
	36822,
	62082,
	68265,
	87344,
	40642,
	87254,
	62372,
	89026,
	88519,
	53514,
	24686,
	23817,
	26704,
	37908,
	22665,
	75378,
	16797,
	55847,
	55932,
	41470,
	43557,
	45439,
	41230,
	25104,
	43736,
	85938,
	74153,
	73956,
	70419,
	81275,
	30656,
	46218,
	67969,
	61544,
	39376,
	14769,
	10199,
	61064,
	75531,
	17211,
	75184,
	24565,
	83922,
	11644,
	71836,
	53030,
	83437,
	33687,
	92054,
	48250,
	10977,
	51050,
	80898,
	26442,
	39875,
	24976,
	31946,
	95278,
	71678,
	49641,
	12962,
	13598,
	40620,
	49062,
	88449,
	80153,
	36918,
	77902,
	41848,
	26032,
	23830,
	49712,
	95464,
	76971,
	67074,
	17629,
	92019,
	20780,
	68372,
	49710,
	94680,
	28626,
	31808,
	25578,
	12020,
	46008,
	27712,
	17168,
	92923,
	94668,
	59490,
	69085,
	28242,
	18302,
	64047,
	62752,
	28873,
	22206,
	57298,
	26420,
	74055,
	74377,
	81846,
	35184,
	72054,
	78666,
	80545,
	69240,
	76319,
	23617,
	53848,
	65082,
	18802,
	63942,
	80208,
	17018,
	78018,
	35520,
	43479,
	18522,
	41260,
	58041,
	96272,
	73601,
	44788,
	67152,
	57671,
	33841,
	41963,
	41010,
	53937,
	75472,
	22253,
	24277,
	62012,
	27539,
	29180,
	27151,
	28752,
	31602,
	80298,
	94689,
	87987,
	51692,
	93653,
	99694,
	50948,
	87830,
	57579,
	60947,
	18514,
	99194,
	57630,
	49914,
	34613,
	32145,
	39569,
	14118,
	38901,
	47443,
	83812,
	61920,
	88684,
	32305,
	58695,
	20016,
	68955,
	96299,
	83258,
	49424,
	97130,
	34668,
	27567,
	57890,
	58399,
	88934,
	70656,
	91507,
	53671,
	75858,
	14726,
	25416,
	90521,
	37289,
	79415,
	21158,
	53210,
	35192,
	67240,
	42967,
	81598,
	12765,
	23853,
	78439,
	13052,
	56642,
	19028,
	89371,
	89795,
	41542,
	49486,
	26630,
	35019,
	62271,
	16433,
	42466,
	63666,
	56231,
	51132,
	33675,
	94670,
	51236,
	52493,
	45346,
	41620,
	80479,
	46200,
	13789,
	98564,
	44974,
	76047,
	60650,
	55723,
	89426,
	33961,
	32872,
	65322,
	29922,
	74923,
	37065,
	66622,
	68253,
	28664,
	62981,
	83161,
	22849,
	11119,
	22980,
	76171,
	96351,
	45573,
	19850,
	97828,
	23867,
	36353,
	46579,
	73434,
	57169,
	68768,
	59354,
	62181,
	34167,
	28374,
	55863,
	25058,
	76126,
	97283,
	25887,
	43067,
	51616,
	97533,
	50692,
	43173,
	93917,
	84713,
	24130,
	56773,
	15310,
	63375,
	75875,
	22284,
	46294,
	29713,
	95318,
	83264,
	39267,
	61101,
	47882,
	24345,
	50674,
	17296,
	52605,
	40817,
	36207,
	16045,
	91784,
	34583,
	71867,
	74898,
	54609,
	33070,
	18570,
	16766,
	62825,
	39793,
	12038,
	37265,
	69479,
	96101,
	49175,
	70708,
	50136,
	27300,
	75794,
	10277,
	30825,
	35563,
	12505,
	11323,
	89922,
	10927,
	50656,
	47973,
	17156,
	98964,
	50980,
	94733,
	22743,
	90255,
	80006,
	10157,
	70496,
	25569,
	79982,
	34897,
	83357,
	26352,
	67432,
	49155,
	36311,
	14394,
	70928,
	67367,
	48195,
	52001,
	63251,
	21894,
	50774,
	85637,
	79243,
	49512,
	22454,
	84214,
	67121,
	49247,
	79313,
	33168,
	48657,
	59824,
	30088,
	74246,
	78389,
	88418,
	93716,
	74962,
	47309,
	93517,
	23895,
	34548,
	32628,
	55988,
	82578,
	51728,
	19440,
	42922,
	52930,
	91562,
	23559,
	92150,
	78315,
	43601,
	87621,
	78580,
	94778,
	68519,
	35509,
	82288,
	65083,
	86995,
	66618,
	18656,
	98537,
	60842,
	68830,
	99689,
	33740,
	45595,
	78239,
	20185,
	87307,
	83009,
	34685,
	13674,
	82450,
	44812,
	54296,
	73558,
	21422,
	49473,
	66108,
	71169,
	78276,
	82977,
	61842,
	31582,
	39064,
	70204,
	71244,
	31642,
	98413,
	61175,
	37168,
	44093,
	18806,
	91249,
	95812,
	93644,
	30669,
	75404,
	53324,
	80675,
	14897,
	34074,
	16366,
	95656,
	73964,
	90705,
	92869,
	99522,
	35903,
	33033,
	55650,
	29289,
	32302,
	76036,
	74870,
	85035,
	79322,
	50332,
	78294,
	93000,
	36560,
	40041,
	70360,
	19074,
	32829,
	61436,
	89938,
	52708,
	83958,
	88155,
	32730,
	64221,
	60391,
	70047,
	12372,
	46501,
	66552,
	56373,
	95596,
	81196,
	32282,
	49504,
	57642,
	95402,
	81625,
	28527,
	71475,
	94280,
	55623,
	46783,
	36641,
	74788,
	87994,
	80189,
	33263,
	23537,
	25755,
	67413,
	74924,
	81215,
	78206,
	36161,
	54925,
	66305,
	42462,
	99051,
	44954,
	88856,
	63026,
	22651,
	39262,
	86176,
	65395,
	12926,
	81994,
	16933,
	46987,
	37215,
	15340,
	40271,
	23747,
	61156,
	65356,
	72336,
	21864,
	99351,
	24599,
	62281,
	81871,
	14029,
	50141,
	46677,
	10745,
	39344,
	38077,
	13983,
	66663,
	64175,
	62355,
	96823,
	46436,
	53388,
	45987,
	87646,
	75997,
	69968,
	30425,
	83211,
	28725,
	47043,
	34475,
	42302,
	66558,
	48964,
	19546,
	51026,
	89346,
	52993,
	29577,
	85932,
	18030,
	65993,
	74692,
	61588,
	50688,
	93760,
	75094,
	68976,
	46278,
	97836,
	17576,
	48519,
	72647,
	54724,
	45186,
	44529,
	90404,
	63175,
	88381,
	32140,
	43348,
	41733,
	84065,
	20389,
	38749,
	37929,
	63747,
	72249,
	98042,
	22741,
	88124,
	16493,
	36599,
	50436,
	67695,
	89091,
	60708,
	28750,
	80224,
	34580,
	79688,
	14870,
	47719,
	57652,
	65541,
	86435,
	95204,
	43800,
	94063,
	24096,
	36628,
	81443,
	61221,
	52782,
	16740,
	68369,
	28895,
	39222,
	97201,
	81119,
	41083,
	72917,
	46148,
	44495,
	99464,
	67254,
	52750,
	63899,
	14617,
	36962,
	41322,
	22504,
	72577,
	67799,
	68284,
	46011,
	50367,
	45832,
	88775,
	32562,
	43756,
	73444,
	67302,
	22133,
	17768,
	69281,
	87364,
	41209,
	37792,
	66667,
	22195,
	65037,
	28634,
	23901,
	40511,
	47779,
	76489,
	77873,
	87523,
	32876,
	62342,
	84222,
	88205,
	78856,
	97888,
	22233,
	35868,
	80155,
	82730,
	12733,
	55038,
	50682,
	42043,
	54719,
	54474,
	61776,
	55564,
	39404,
	94954,
	76258,
	87689,
	65619,
	41846,
	21990,
	60161,
	44955,
	84178,
	32693,
	19125,
	96731,
	98080,
	72947,
	67285,
	78101,
	43064,
	81350,
	46937,
	97686,
	66339,
	39269,
	61837,
	97647,
	44613,
	73127,
	64048,
	15530,
	63926,
	95853,
	26155,
	82222,
	10035,
	16445,
	67972,
	51780,
	52358,
	83942,
	64899,
	56009,
	42288,
	53349,
	76255,
	53044,
	37184,
	24813,
	62309,
	93607,
	27423,
	26128,
	72691,
	42162,
	78247,
	80209,
	38310,
	59228,
	24897,
	55176,
	33062,
	26521,
	93209,
	89881,
	11418,
	16884,
	31995,
	17858,
	49107,
	50512,
	56890,
	76173,
	51569,
	73919,
	62591,
	53735,
	14454,
	19976,
}
var b = []int{
	93653,
	27995,
	69782,
	41563,
	48282,
	37517,
	68189,
	25255,
	65322,
	99897,
	35509,
	49013,
	92011,
	68830,
	36207,
	23559,
	58330,
	40186,
	63410,
	21671,
	79572,
	44529,
	28656,
	73412,
	85991,
	78294,
	23559,
	17882,
	17211,
	11644,
	80435,
	96101,
	86995,
	19028,
	77868,
	90255,
	25255,
	88912,
	63269,
	21548,
	71747,
	24477,
	31808,
	88912,
	77868,
	13754,
	36177,
	10894,
	85185,
	60500,
	91114,
	85134,
	77868,
	41563,
	14475,
	53324,
	53324,
	63753,
	47827,
	77536,
	40593,
	31274,
	97404,
	67531,
	50141,
	50674,
	32858,
	24813,
	89795,
	77868,
	57646,
	66502,
	39062,
	23559,
	65322,
	91636,
	19028,
	11644,
	23537,
	19028,
	47412,
	74334,
	44529,
	50980,
	80749,
	38495,
	90255,
	68830,
	73660,
	50980,
	14164,
	34362,
	28090,
	50713,
	11121,
	31274,
	38248,
	80106,
	13153,
	12727,
	40674,
	31053,
	33051,
	54436,
	95555,
	17329,
	79835,
	77868,
	23559,
	70046,
	22417,
	89690,
	34969,
	23559,
	50980,
	83301,
	77681,
	26236,
	92869,
	50980,
	28090,
	44529,
	31808,
	53056,
	12263,
	13121,
	82448,
	16530,
	21589,
	19658,
	14475,
	36493,
	48589,
	54974,
	88465,
	23537,
	53324,
	14329,
	74416,
	37003,
	40380,
	68189,
	18220,
	44111,
	46787,
	89795,
	25205,
	93653,
	50141,
	23559,
	11642,
	31274,
	61987,
	19028,
	20768,
	24325,
	94888,
	24394,
	99327,
	58818,
	68189,
	84917,
	83679,
	85099,
	52073,
	78294,
	24410,
	50980,
	90255,
	50980,
	71478,
	37150,
	28090,
	70055,
	68830,
	66585,
	50141,
	93653,
	35509,
	12437,
	87987,
	86995,
	89795,
	59660,
	11644,
	65322,
	59423,
	50141,
	95142,
	56227,
	33007,
	65641,
	53505,
	23537,
	87987,
	63910,
	50713,
	27554,
	93653,
	30520,
	79179,
	47209,
	36207,
	72254,
	44661,
	87987,
	25255,
	68107,
	71814,
	24813,
	22904,
	92731,
	68189,
	50674,
	90255,
	54116,
	53324,
	29145,
	85335,
	64775,
	51790,
	19794,
	96101,
	41345,
	39176,
	50141,
	23489,
	44529,
	86052,
	50801,
	17329,
	38984,
	37514,
	19794,
	14475,
	69571,
	50980,
	41563,
	72877,
	25234,
	90255,
	28090,
	65557,
	85065,
	23522,
	87600,
	41532,
	40116,
	19028,
	24977,
	50980,
	53324,
	97982,
	40884,
	53324,
	90651,
	19028,
	28090,
	98658,
	77868,
	65322,
	75705,
	79320,
	96101,
	16942,
	87327,
	31274,
	64314,
	24603,
	23559,
	38755,
	50713,
	68830,
	96101,
	31808,
	41563,
	68705,
	89031,
	65322,
	35601,
	59200,
	86995,
	68189,
	72688,
	96101,
	17329,
	77868,
	76226,
	92869,
	93653,
	93653,
	54726,
	70350,
	37102,
	64160,
	32684,
	32617,
	31808,
	81806,
	56568,
	60336,
	97703,
	40091,
	30486,
	31808,
	74633,
	61765,
	44529,
	25255,
	14393,
	19028,
	33190,
	54998,
	82083,
	68533,
	34363,
	19028,
	68276,
	68830,
	65514,
	76417,
	23537,
	90255,
	26450,
	28580,
	13674,
	30362,
	21256,
	92869,
	79529,
	94081,
	70556,
	99228,
	99097,
	39378,
	50980,
	78294,
	65322,
	27393,
	22132,
	71246,
	92869,
	96521,
	16880,
	88060,
	52634,
	57101,
	60809,
	28090,
	25255,
	33046,
	36207,
	53033,
	87987,
	65843,
	48858,
	38485,
	14664,
	96101,
	71241,
	73532,
	59278,
	18029,
	85374,
	78294,
	30817,
	45566,
	17058,
	27305,
	89795,
	41563,
	96101,
	51322,
	36831,
	14475,
	93653,
	36888,
	93491,
	23559,
	22876,
	68830,
	51611,
	75130,
	88543,
	96101,
	15010,
	21675,
	68830,
	78294,
	52104,
	89795,
	61399,
	56764,
	53394,
	28090,
	65322,
	76675,
	92068,
	93653,
	14475,
	50141,
	41637,
	35922,
	23559,
	47673,
	50906,
	38297,
	48139,
	61967,
	86995,
	96638,
	11644,
	19794,
	93403,
	72082,
	53324,
	25255,
	93653,
	10526,
	94340,
	16105,
	62764,
	31274,
	19016,
	66925,
	31274,
	50141,
	50713,
	31808,
	50713,
	42720,
	31976,
	99509,
	96101,
	68830,
	14043,
	50936,
	35509,
	65959,
	78860,
	50980,
	66788,
	44529,
	87987,
	71137,
	13674,
	31274,
	50075,
	96101,
	19028,
	25355,
	16168,
	81526,
	23559,
	65310,
	93094,
	36207,
	41643,
	22984,
	50141,
	93653,
	82791,
	36244,
	50799,
	90489,
	72630,
	50763,
	68830,
	78294,
	77868,
	57338,
	39984,
	83564,
	96101,
	74200,
	70445,
	17822,
	64460,
	86995,
	31274,
	88912,
	90255,
	60686,
	76843,
	92717,
	80674,
	89430,
	45722,
	97853,
	98414,
	68189,
	12029,
	36159,
	34697,
	59688,
	23559,
	14475,
	53324,
	25761,
	50674,
	50980,
	76308,
	13794,
	19411,
	28221,
	90255,
	79146,
	90483,
	85271,
	50141,
	87987,
	23559,
	68189,
	41682,
	97617,
	64641,
	53324,
	24803,
	57352,
	33040,
	53005,
	51641,
	29253,
	50713,
	61868,
	16129,
	80951,
	41626,
	76055,
	53465,
	13599,
	50980,
	44529,
	50141,
	93653,
	33135,
	77836,
	14393,
	83346,
	61967,
	47771,
	86995,
	96101,
	70347,
	14475,
	12880,
	11644,
	78294,
	93653,
	96101,
	68830,
	69934,
	28090,
	22671,
	27206,
	26257,
	57394,
	87036,
	21503,
	17504,
	82452,
	12706,
	44529,
	44260,
	50713,
	55412,
	14475,
	96101,
	62368,
	87990,
	87987,
	24149,
	91890,
	14393,
	84455,
	65359,
	71304,
	11644,
	92869,
	96101,
	59345,
	23559,
	26134,
	54273,
	65322,
	92869,
	28090,
	41226,
	92869,
	76485,
	86388,
	79503,
	15174,
	72153,
	86995,
	65333,
	51667,
	23559,
	57003,
	26154,
	21651,
	89795,
	43252,
	36207,
	11644,
	57954,
	15919,
	20805,
	86995,
	78035,
	34840,
	19814,
	92888,
	74905,
	68830,
	90052,
	65322,
	28090,
	62226,
	29324,
	50980,
	14393,
	68175,
	25255,
	84730,
	31274,
	40953,
	16926,
	27805,
	45159,
	41563,
	31274,
	62680,
	26879,
	61783,
	91325,
	17838,
	65350,
	77868,
	83273,
	66389,
	52439,
	19794,
	13010,
	95850,
	70145,
	65322,
	52161,
	10132,
	24077,
	89795,
	53090,
	36207,
	36207,
	69366,
	77868,
	88642,
	28090,
	69909,
	75146,
	92869,
	94990,
	48229,
	10547,
	68189,
	74202,
	40690,
	78294,
	15634,
	95459,
	14475,
	18593,
	87987,
	50490,
	37109,
	50141,
	23559,
	11032,
	40848,
	73483,
	32572,
	68189,
	20850,
	92869,
	50141,
	50216,
	73401,
	29952,
	32630,
	28090,
	58896,
	66916,
	59662,
	78294,
	19794,
	42610,
	50980,
	11753,
	58345,
	57644,
	31274,
	77868,
	25255,
	31274,
	44529,
	53324,
	19028,
	53525,
	89795,
	98749,
	89603,
	71046,
	85369,
	28090,
	53324,
	67097,
	75527,
	23559,
	97432,
	20185,
	73527,
	29274,
	14746,
	51614,
	11782,
	72623,
	60515,
	14393,
	34348,
	87987,
	68830,
	37199,
	31808,
	23005,
	99355,
	29073,
	81559,
	14535,
	78773,
	77868,
	89795,
	23537,
	22264,
	72800,
	73774,
	78294,
	92869,
	50141,
	86995,
	49951,
	28090,
	88912,
	89795,
	28090,
	85778,
	24813,
	14475,
	69982,
	14393,
	78294,
	53324,
	30191,
	19794,
	28090,
	63748,
	86674,
	52856,
	11644,
	56048,
	51278,
	77868,
	13349,
	79046,
	23361,
	53324,
	42297,
	48720,
	43636,
	35965,
	80864,
	50038,
	89795,
	67877,
	19028,
	68830,
	79443,
	44754,
	68830,
	92723,
	11707,
	45545,
	93653,
	44529,
	88109,
	64932,
	89795,
	31274,
	94895,
	92753,
	50713,
	78294,
	89795,
	57408,
	54676,
	38385,
	60320,
	22556,
	19922,
	23559,
	89795,
	69541,
	44529,
	96101,
	86835,
	50980,
	44529,
	32708,
	46071,
	93653,
	34155,
	43448,
	41828,
	86357,
	44529,
	25963,
	23537,
	36423,
	88912,
	98953,
	96101,
	93788,
	35143,
	35340,
	93653,
	59381,
	95783,
	33425,
	44529,
	69408,
	41563,
	86149,
	20022,
	92607,
	68830,
	39113,
	96101,
	34798,
	68189,
	29252,
	33458,
	87987,
	19794,
	81239,
	50980,
	11644,
	36588,
	95478,
	17329,
	57166,
	78294,
	31274,
	45653,
	44529,
	65582,
	77602,
	11644,
	50141,
	25255,
	21269,
	18020,
	90525,
	58381,
	28135,
	55754,
	50713,
	50713,
	73436,
	38709,
	39875,
	68830,
	71156,
	61967,
	31808,
	87987,
	87987,
	26340,
	49887,
	16547,
	92449,
	85451,
	77868,
	39885,
	25255,
	27363,
	77868,
	73746,
	28090,
	37556,
	89795,
	50674,
	95801,
	45253,
	86807,
	66992,
	35509,
	50980,
	50980,
	41563,
	40382,
	20051,
	68189,
	17329,
	81437,
	93653,
	78294,
	19794,
	26258,
	23250,
	74279,
	71512,
	23991,
	25255,
	41563,
	31274,
	14475,
	89795,
	67995,
	59930,
	17435,
	22185,
	53026,
	57523,
	39875,
	87987,
	44529,
	93653,
	44529,
	95746,
	23344,
	77580,
	11644,
	23483,
	11644,
	79381,
	60744,
	50674,
	77868,
	86995,
	93568,
	68830,
	80370,
	87987,
	51860,
	91981,
	28090,
	19794,
	87445,
	93653,
	15558,
	68830,
	31274,
	60096,
	11955,
	96101,
	36590,
	65322,
	50713,
	26351,
	31274,
	19028,
	52490,
	19028,
	45568,
	78294,
	50141,
	60052,
	10613,
	31808,
	93653,
	47607,
	23559,
	37760,
	53324,
	39170,
	68189,
	23559,
	50980,
	78294,
	50980,
	98671,
	96101,
	54526,
	65322,
	65309,
	52845,
	23537,
	16739,
	79935,
	69804,
	46609,
}
var c int
var d int

func main() {
	slices.Sort(a)
	slices.Sort(b)
	for _, value_a := range a {
		for _, value_b := range b {
			if value_a == value_b {
				d = d + 1
			}
			c = c + value_a*d
		}
	}

	fmt.Println("Result:", c)
}
