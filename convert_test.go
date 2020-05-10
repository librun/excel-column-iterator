package iterator

import (
	"testing"
)

type testValue struct {
	ColumnNumber int
	ColumnName   string
}

var (
	testValues = []testValue{
		{0, ""},
		{1, "A"},
		{2, "B"},
		{3, "C"},
		{4, "D"},
		{5, "E"},
		{6, "F"},
		{7, "G"},
		{8, "H"},
		{9, "I"},
		{10, "J"},
		{11, "K"},
		{12, "L"},
		{13, "M"},
		{14, "N"},
		{15, "O"},
		{16, "P"},
		{17, "Q"},
		{18, "R"},
		{19, "S"},
		{20, "T"},
		{21, "U"},
		{22, "V"},
		{23, "W"},
		{24, "X"},
		{25, "Y"},
		{26, "Z"},
		{27, "AA"},
		{28, "AB"},
		{29, "AC"},
		{30, "AD"},
		{31, "AE"},
		{32, "AF"},
		{33, "AG"},
		{34, "AH"},
		{35, "AI"},
		{36, "AJ"},
		{37, "AK"},
		{38, "AL"},
		{39, "AM"},
		{40, "AN"},
		{41, "AO"},
		{42, "AP"},
		{43, "AQ"},
		{44, "AR"},
		{45, "AS"},
		{46, "AT"},
		{47, "AU"},
		{48, "AV"},
		{49, "AW"},
		{50, "AX"},
		{51, "AY"},
		{52, "AZ"},
		{53, "BA"},
		{54, "BB"},
		{55, "BC"},
		{56, "BD"},
		{57, "BE"},
		{58, "BF"},
		{59, "BG"},
		{60, "BH"},
		{61, "BI"},
		{62, "BJ"},
		{63, "BK"},
		{64, "BL"},
		{65, "BM"},
		{66, "BN"},
		{67, "BO"},
		{68, "BP"},
		{69, "BQ"},
		{70, "BR"},
		{71, "BS"},
		{72, "BT"},
		{73, "BU"},
		{74, "BV"},
		{75, "BW"},
		{76, "BX"},
		{77, "BY"},
		{78, "BZ"},
		{79, "CA"},
		{80, "CB"},
		{81, "CC"},
		{82, "CD"},
		{83, "CE"},
		{84, "CF"},
		{85, "CG"},
		{86, "CH"},
		{87, "CI"},
		{88, "CJ"},
		{89, "CK"},
		{90, "CL"},
		{91, "CM"},
		{92, "CN"},
		{93, "CO"},
		{94, "CP"},
		{95, "CQ"},
		{96, "CR"},
		{97, "CS"},
		{98, "CT"},
		{99, "CU"},
		{100, "CV"},
		{101, "CW"},
		{102, "CX"},
		{103, "CY"},
		{104, "CZ"},
		{105, "DA"},
		{106, "DB"},
		{107, "DC"},
		{108, "DD"},
		{109, "DE"},
		{110, "DF"},
		{111, "DG"},
		{112, "DH"},
		{113, "DI"},
		{114, "DJ"},
		{115, "DK"},
		{116, "DL"},
		{117, "DM"},
		{118, "DN"},
		{119, "DO"},
		{120, "DP"},
		{121, "DQ"},
		{122, "DR"},
		{123, "DS"},
		{124, "DT"},
		{125, "DU"},
		{126, "DV"},
		{127, "DW"},
		{128, "DX"},
		{129, "DY"},
		{130, "DZ"},
		{131, "EA"},
		{132, "EB"},
		{133, "EC"},
		{134, "ED"},
		{135, "EE"},
		{136, "EF"},
		{137, "EG"},
		{138, "EH"},
		{139, "EI"},
		{140, "EJ"},
		{141, "EK"},
		{142, "EL"},
		{143, "EM"},
		{144, "EN"},
		{145, "EO"},
		{146, "EP"},
		{147, "EQ"},
		{148, "ER"},
		{149, "ES"},
		{150, "ET"},
		{151, "EU"},
		{152, "EV"},
		{153, "EW"},
		{154, "EX"},
		{155, "EY"},
		{156, "EZ"},
		{157, "FA"},
		{158, "FB"},
		{159, "FC"},
		{160, "FD"},
		{161, "FE"},
		{162, "FF"},
		{163, "FG"},
		{164, "FH"},
		{165, "FI"},
		{166, "FJ"},
		{167, "FK"},
		{168, "FL"},
		{169, "FM"},
		{170, "FN"},
		{171, "FO"},
		{172, "FP"},
		{173, "FQ"},
		{174, "FR"},
		{175, "FS"},
		{176, "FT"},
		{177, "FU"},
		{178, "FV"},
		{179, "FW"},
		{180, "FX"},
		{181, "FY"},
		{182, "FZ"},
		{183, "GA"},
		{184, "GB"},
		{185, "GC"},
		{186, "GD"},
		{187, "GE"},
		{188, "GF"},
		{189, "GG"},
		{190, "GH"},
		{191, "GI"},
		{192, "GJ"},
		{193, "GK"},
		{194, "GL"},
		{195, "GM"},
		{196, "GN"},
		{197, "GO"},
		{198, "GP"},
		{199, "GQ"},
		{200, "GR"},
		{201, "GS"},
		{202, "GT"},
		{203, "GU"},
		{204, "GV"},
		{205, "GW"},
		{206, "GX"},
		{207, "GY"},
		{208, "GZ"},
		{209, "HA"},
		{210, "HB"},
		{211, "HC"},
		{212, "HD"},
		{213, "HE"},
		{214, "HF"},
		{215, "HG"},
		{216, "HH"},
		{217, "HI"},
		{218, "HJ"},
		{219, "HK"},
		{220, "HL"},
		{221, "HM"},
		{222, "HN"},
		{223, "HO"},
		{224, "HP"},
		{225, "HQ"},
		{226, "HR"},
		{227, "HS"},
		{228, "HT"},
		{229, "HU"},
		{230, "HV"},
		{231, "HW"},
		{232, "HX"},
		{233, "HY"},
		{234, "HZ"},
		{235, "IA"},
		{236, "IB"},
		{237, "IC"},
		{238, "ID"},
		{239, "IE"},
		{240, "IF"},
		{241, "IG"},
		{242, "IH"},
		{243, "II"},
		{244, "IJ"},
		{245, "IK"},
		{246, "IL"},
		{247, "IM"},
		{248, "IN"},
		{249, "IO"},
		{250, "IP"},
		{251, "IQ"},
		{252, "IR"},
		{253, "IS"},
		{254, "IT"},
		{255, "IU"},
		{256, "IV"},
		{257, "IW"},
		{258, "IX"},
		{259, "IY"},
		{260, "IZ"},
		{261, "JA"},
		{262, "JB"},
		{263, "JC"},
		{264, "JD"},
		{265, "JE"},
		{266, "JF"},
		{267, "JG"},
		{268, "JH"},
		{269, "JI"},
		{270, "JJ"},
		{271, "JK"},
		{272, "JL"},
		{273, "JM"},
		{274, "JN"},
		{275, "JO"},
		{276, "JP"},
		{277, "JQ"},
		{278, "JR"},
		{279, "JS"},
		{280, "JT"},
		{281, "JU"},
		{282, "JV"},
		{283, "JW"},
		{284, "JX"},
		{285, "JY"},
		{286, "JZ"},
		{287, "KA"},
		{288, "KB"},
		{289, "KC"},
		{290, "KD"},
		{291, "KE"},
		{292, "KF"},
		{293, "KG"},
		{294, "KH"},
		{295, "KI"},
		{296, "KJ"},
		{297, "KK"},
		{298, "KL"},
		{299, "KM"},
		{300, "KN"},
		{301, "KO"},
		{302, "KP"},
		{303, "KQ"},
		{304, "KR"},
		{305, "KS"},
		{306, "KT"},
		{307, "KU"},
		{308, "KV"},
		{309, "KW"},
		{310, "KX"},
		{311, "KY"},
		{312, "KZ"},
		{313, "LA"},
		{314, "LB"},
		{315, "LC"},
		{316, "LD"},
		{317, "LE"},
		{318, "LF"},
		{319, "LG"},
		{320, "LH"},
		{321, "LI"},
		{322, "LJ"},
		{323, "LK"},
		{324, "LL"},
		{325, "LM"},
		{326, "LN"},
		{327, "LO"},
		{328, "LP"},
		{329, "LQ"},
		{330, "LR"},
		{331, "LS"},
		{332, "LT"},
		{333, "LU"},
		{334, "LV"},
		{335, "LW"},
		{336, "LX"},
		{337, "LY"},
		{338, "LZ"},
		{339, "MA"},
		{340, "MB"},
		{341, "MC"},
		{342, "MD"},
		{343, "ME"},
		{344, "MF"},
		{345, "MG"},
		{346, "MH"},
		{347, "MI"},
		{348, "MJ"},
		{349, "MK"},
		{350, "ML"},
		{351, "MM"},
		{352, "MN"},
		{353, "MO"},
		{354, "MP"},
		{355, "MQ"},
		{356, "MR"},
		{357, "MS"},
		{358, "MT"},
		{359, "MU"},
		{360, "MV"},
		{361, "MW"},
		{362, "MX"},
		{363, "MY"},
		{364, "MZ"},
		{365, "NA"},
		{366, "NB"},
		{367, "NC"},
		{368, "ND"},
		{369, "NE"},
		{370, "NF"},
		{371, "NG"},
		{372, "NH"},
		{373, "NI"},
		{374, "NJ"},
		{375, "NK"},
		{376, "NL"},
		{377, "NM"},
		{378, "NN"},
		{379, "NO"},
		{380, "NP"},
		{381, "NQ"},
		{382, "NR"},
		{383, "NS"},
		{384, "NT"},
		{385, "NU"},
		{386, "NV"},
		{387, "NW"},
		{388, "NX"},
		{389, "NY"},
		{390, "NZ"},
		{391, "OA"},
		{392, "OB"},
		{393, "OC"},
		{394, "OD"},
		{395, "OE"},
		{396, "OF"},
		{397, "OG"},
		{398, "OH"},
		{399, "OI"},
		{400, "OJ"},
		{401, "OK"},
		{402, "OL"},
		{403, "OM"},
		{404, "ON"},
		{405, "OO"},
		{406, "OP"},
		{407, "OQ"},
		{408, "OR"},
		{409, "OS"},
		{410, "OT"},
		{411, "OU"},
		{412, "OV"},
		{413, "OW"},
		{414, "OX"},
		{415, "OY"},
		{416, "OZ"},
		{417, "PA"},
		{418, "PB"},
		{419, "PC"},
		{420, "PD"},
		{421, "PE"},
		{422, "PF"},
		{423, "PG"},
		{424, "PH"},
		{425, "PI"},
		{426, "PJ"},
		{427, "PK"},
		{428, "PL"},
		{429, "PM"},
		{430, "PN"},
		{431, "PO"},
		{432, "PP"},
		{433, "PQ"},
		{434, "PR"},
		{435, "PS"},
		{436, "PT"},
		{437, "PU"},
		{438, "PV"},
		{439, "PW"},
		{440, "PX"},
		{441, "PY"},
		{442, "PZ"},
		{443, "QA"},
		{444, "QB"},
		{445, "QC"},
		{446, "QD"},
		{447, "QE"},
		{448, "QF"},
		{449, "QG"},
		{450, "QH"},
		{451, "QI"},
		{452, "QJ"},
		{453, "QK"},
		{454, "QL"},
		{455, "QM"},
		{456, "QN"},
		{457, "QO"},
		{458, "QP"},
		{459, "QQ"},
		{460, "QR"},
		{461, "QS"},
		{462, "QT"},
		{463, "QU"},
		{464, "QV"},
		{465, "QW"},
		{466, "QX"},
		{467, "QY"},
		{468, "QZ"},
		{469, "RA"},
		{470, "RB"},
		{471, "RC"},
		{472, "RD"},
		{473, "RE"},
		{474, "RF"},
		{475, "RG"},
		{476, "RH"},
		{477, "RI"},
		{478, "RJ"},
		{479, "RK"},
		{480, "RL"},
		{481, "RM"},
		{482, "RN"},
		{483, "RO"},
		{484, "RP"},
		{485, "RQ"},
		{486, "RR"},
		{487, "RS"},
		{488, "RT"},
		{489, "RU"},
		{490, "RV"},
		{491, "RW"},
		{492, "RX"},
		{493, "RY"},
		{494, "RZ"},
		{495, "SA"},
		{496, "SB"},
		{497, "SC"},
		{498, "SD"},
		{499, "SE"},
		{500, "SF"},
		{501, "SG"},
		{502, "SH"},
		{503, "SI"},
		{504, "SJ"},
		{505, "SK"},
		{506, "SL"},
		{507, "SM"},
		{508, "SN"},
		{509, "SO"},
		{510, "SP"},
		{511, "SQ"},
		{512, "SR"},
		{513, "SS"},
		{514, "ST"},
		{515, "SU"},
		{516, "SV"},
		{517, "SW"},
		{518, "SX"},
		{519, "SY"},
		{520, "SZ"},
		{521, "TA"},
		{522, "TB"},
		{523, "TC"},
		{524, "TD"},
		{525, "TE"},
		{526, "TF"},
		{527, "TG"},
		{528, "TH"},
		{529, "TI"},
		{530, "TJ"},
		{531, "TK"},
		{532, "TL"},
		{533, "TM"},
		{534, "TN"},
		{535, "TO"},
		{536, "TP"},
		{537, "TQ"},
		{538, "TR"},
		{539, "TS"},
		{540, "TT"},
		{541, "TU"},
		{542, "TV"},
		{543, "TW"},
		{544, "TX"},
		{545, "TY"},
		{546, "TZ"},
		{547, "UA"},
		{548, "UB"},
		{549, "UC"},
		{550, "UD"},
		{551, "UE"},
		{552, "UF"},
		{553, "UG"},
		{554, "UH"},
		{555, "UI"},
		{556, "UJ"},
		{557, "UK"},
		{558, "UL"},
		{559, "UM"},
		{560, "UN"},
		{561, "UO"},
		{562, "UP"},
		{563, "UQ"},
		{564, "UR"},
		{565, "US"},
		{566, "UT"},
		{567, "UU"},
		{568, "UV"},
		{569, "UW"},
		{570, "UX"},
		{571, "UY"},
		{572, "UZ"},
		{573, "VA"},
		{574, "VB"},
		{575, "VC"},
		{576, "VD"},
		{577, "VE"},
		{578, "VF"},
		{579, "VG"},
		{580, "VH"},
		{581, "VI"},
		{582, "VJ"},
		{583, "VK"},
		{584, "VL"},
		{585, "VM"},
		{586, "VN"},
		{587, "VO"},
		{588, "VP"},
		{589, "VQ"},
		{590, "VR"},
		{591, "VS"},
		{592, "VT"},
		{593, "VU"},
		{594, "VV"},
		{595, "VW"},
		{596, "VX"},
		{597, "VY"},
		{598, "VZ"},
		{599, "WA"},
		{600, "WB"},
		{601, "WC"},
		{602, "WD"},
		{603, "WE"},
		{604, "WF"},
		{605, "WG"},
		{606, "WH"},
		{607, "WI"},
		{608, "WJ"},
		{609, "WK"},
		{610, "WL"},
		{611, "WM"},
		{612, "WN"},
		{613, "WO"},
		{614, "WP"},
		{615, "WQ"},
		{616, "WR"},
		{617, "WS"},
		{618, "WT"},
		{619, "WU"},
		{620, "WV"},
		{621, "WW"},
		{622, "WX"},
		{623, "WY"},
		{624, "WZ"},
		{625, "XA"},
		{626, "XB"},
		{627, "XC"},
		{628, "XD"},
		{629, "XE"},
		{630, "XF"},
		{631, "XG"},
		{632, "XH"},
		{633, "XI"},
		{634, "XJ"},
		{635, "XK"},
		{636, "XL"},
		{637, "XM"},
		{638, "XN"},
		{639, "XO"},
		{640, "XP"},
		{641, "XQ"},
		{642, "XR"},
		{643, "XS"},
		{644, "XT"},
		{645, "XU"},
		{646, "XV"},
		{647, "XW"},
		{648, "XX"},
		{649, "XY"},
		{650, "XZ"},
		{651, "YA"},
		{652, "YB"},
		{653, "YC"},
		{654, "YD"},
		{655, "YE"},
		{656, "YF"},
		{657, "YG"},
		{658, "YH"},
		{659, "YI"},
		{660, "YJ"},
		{661, "YK"},
		{662, "YL"},
		{663, "YM"},
		{664, "YN"},
		{665, "YO"},
		{666, "YP"},
		{667, "YQ"},
		{668, "YR"},
		{669, "YS"},
		{670, "YT"},
		{671, "YU"},
		{672, "YV"},
		{673, "YW"},
		{674, "YX"},
		{675, "YY"},
		{676, "YZ"},
		{677, "ZA"},
		{678, "ZB"},
		{679, "ZC"},
		{680, "ZD"},
		{681, "ZE"},
		{682, "ZF"},
		{683, "ZG"},
		{684, "ZH"},
		{685, "ZI"},
		{686, "ZJ"},
		{687, "ZK"},
		{688, "ZL"},
		{689, "ZM"},
		{690, "ZN"},
		{691, "ZO"},
		{692, "ZP"},
		{693, "ZQ"},
		{694, "ZR"},
		{695, "ZS"},
		{696, "ZT"},
		{697, "ZU"},
		{698, "ZV"},
		{699, "ZW"},
		{700, "ZX"},
		{701, "ZY"},
		{702, "ZZ"},
		{703, "AAA"},
		{704, "AAB"},
		{705, "AAC"},
		{706, "AAD"},
		{707, "AAE"},
		{708, "AAF"},
		{709, "AAG"},
		{710, "AAH"},
		{711, "AAI"},
		{712, "AAJ"},
		{713, "AAK"},
		{714, "AAL"},
		{715, "AAM"},
		{716, "AAN"},
		{717, "AAO"},
		{718, "AAP"},
		{719, "AAQ"},
		{720, "AAR"},
		{721, "AAS"},
		{722, "AAT"},
		{723, "AAU"},
		{724, "AAV"},
		{725, "AAW"},
		{726, "AAX"},
		{727, "AAY"},
		{728, "AAZ"},
		{729, "ABA"},
		{730, "ABB"},
		{731, "ABC"},
		{732, "ABD"},
		{733, "ABE"},
		{734, "ABF"},
		{735, "ABG"},
		{736, "ABH"},
		{737, "ABI"},
		{738, "ABJ"},
		{739, "ABK"},
		{740, "ABL"},
		{741, "ABM"},
		{742, "ABN"},
		{743, "ABO"},
		{744, "ABP"},
		{745, "ABQ"},
		{746, "ABR"},
		{747, "ABS"},
		{748, "ABT"},
		{749, "ABU"},
		{750, "ABV"},
		{751, "ABW"},
		{752, "ABX"},
		{753, "ABY"},
		{754, "ABZ"},
		{755, "ACA"},
		{756, "ACB"},
		{757, "ACC"},
		{758, "ACD"},
		{759, "ACE"},
		{760, "ACF"},
		{761, "ACG"},
		{762, "ACH"},
		{763, "ACI"},
		{764, "ACJ"},
		{765, "ACK"},
		{766, "ACL"},
		{767, "ACM"},
		{768, "ACN"},
		{769, "ACO"},
		{770, "ACP"},
		{771, "ACQ"},
		{772, "ACR"},
		{773, "ACS"},
		{774, "ACT"},
		{775, "ACU"},
		{776, "ACV"},
		{777, "ACW"},
		{778, "ACX"},
		{779, "ACY"},
		{780, "ACZ"},
		{781, "ADA"},
		{782, "ADB"},
		{783, "ADC"},
		{784, "ADD"},
		{785, "ADE"},
		{786, "ADF"},
		{787, "ADG"},
		{788, "ADH"},
		{789, "ADI"},
		{790, "ADJ"},
		{791, "ADK"},
		{792, "ADL"},
		{793, "ADM"},
		{794, "ADN"},
		{795, "ADO"},
		{796, "ADP"},
		{797, "ADQ"},
		{798, "ADR"},
		{799, "ADS"},
		{800, "ADT"},
		{801, "ADU"},
		{802, "ADV"},
		{803, "ADW"},
		{804, "ADX"},
		{805, "ADY"},
		{806, "ADZ"},
		{807, "AEA"},
		{808, "AEB"},
		{809, "AEC"},
		{810, "AED"},
		{811, "AEE"},
		{812, "AEF"},
		{813, "AEG"},
		{814, "AEH"},
		{815, "AEI"},
		{816, "AEJ"},
		{817, "AEK"},
		{818, "AEL"},
		{819, "AEM"},
		{820, "AEN"},
		{821, "AEO"},
		{822, "AEP"},
		{823, "AEQ"},
		{824, "AER"},
		{825, "AES"},
		{826, "AET"},
		{827, "AEU"},
		{828, "AEV"},
		{829, "AEW"},
		{830, "AEX"},
		{831, "AEY"},
		{832, "AEZ"},
		{833, "AFA"},
		{834, "AFB"},
		{835, "AFC"},
		{836, "AFD"},
		{837, "AFE"},
		{838, "AFF"},
		{839, "AFG"},
		{840, "AFH"},
		{841, "AFI"},
		{842, "AFJ"},
		{843, "AFK"},
		{844, "AFL"},
		{845, "AFM"},
		{846, "AFN"},
		{847, "AFO"},
		{848, "AFP"},
		{849, "AFQ"},
		{850, "AFR"},
		{851, "AFS"},
		{852, "AFT"},
		{853, "AFU"},
		{854, "AFV"},
		{855, "AFW"},
		{856, "AFX"},
		{857, "AFY"},
		{858, "AFZ"},
		{859, "AGA"},
		{860, "AGB"},
		{861, "AGC"},
		{862, "AGD"},
		{863, "AGE"},
		{864, "AGF"},
		{865, "AGG"},
		{866, "AGH"},
		{867, "AGI"},
		{868, "AGJ"},
		{869, "AGK"},
		{870, "AGL"},
		{871, "AGM"},
		{872, "AGN"},
		{873, "AGO"},
		{874, "AGP"},
		{875, "AGQ"},
		{876, "AGR"},
		{877, "AGS"},
		{878, "AGT"},
		{879, "AGU"},
		{880, "AGV"},
		{881, "AGW"},
		{882, "AGX"},
		{883, "AGY"},
		{884, "AGZ"},
		{885, "AHA"},
		{886, "AHB"},
		{887, "AHC"},
		{888, "AHD"},
		{889, "AHE"},
		{890, "AHF"},
		{891, "AHG"},
		{892, "AHH"},
		{893, "AHI"},
		{894, "AHJ"},
		{895, "AHK"},
		{896, "AHL"},
		{897, "AHM"},
		{898, "AHN"},
		{899, "AHO"},
		{900, "AHP"},
		{901, "AHQ"},
		{902, "AHR"},
		{903, "AHS"},
		{904, "AHT"},
		{905, "AHU"},
		{906, "AHV"},
		{907, "AHW"},
		{908, "AHX"},
		{909, "AHY"},
		{910, "AHZ"},
		{911, "AIA"},
		{912, "AIB"},
		{913, "AIC"},
		{914, "AID"},
		{915, "AIE"},
		{916, "AIF"},
		{917, "AIG"},
		{918, "AIH"},
		{919, "AII"},
		{920, "AIJ"},
		{921, "AIK"},
		{922, "AIL"},
		{923, "AIM"},
		{924, "AIN"},
		{925, "AIO"},
		{926, "AIP"},
		{927, "AIQ"},
		{928, "AIR"},
		{929, "AIS"},
		{930, "AIT"},
		{931, "AIU"},
		{932, "AIV"},
		{933, "AIW"},
		{934, "AIX"},
		{935, "AIY"},
		{936, "AIZ"},
		{937, "AJA"},
		{938, "AJB"},
		{939, "AJC"},
		{940, "AJD"},
		{941, "AJE"},
		{942, "AJF"},
		{943, "AJG"},
		{944, "AJH"},
		{945, "AJI"},
		{946, "AJJ"},
		{947, "AJK"},
		{948, "AJL"},
		{949, "AJM"},
		{950, "AJN"},
		{951, "AJO"},
		{952, "AJP"},
		{953, "AJQ"},
		{954, "AJR"},
		{955, "AJS"},
		{956, "AJT"},
		{957, "AJU"},
		{958, "AJV"},
		{959, "AJW"},
		{960, "AJX"},
		{961, "AJY"},
		{962, "AJZ"},
		{963, "AKA"},
		{964, "AKB"},
		{965, "AKC"},
		{966, "AKD"},
		{967, "AKE"},
		{968, "AKF"},
		{969, "AKG"},
		{970, "AKH"},
		{971, "AKI"},
		{972, "AKJ"},
		{973, "AKK"},
		{974, "AKL"},
		{975, "AKM"},
		{976, "AKN"},
		{977, "AKO"},
		{978, "AKP"},
		{979, "AKQ"},
		{980, "AKR"},
		{981, "AKS"},
		{982, "AKT"},
		{983, "AKU"},
		{984, "AKV"},
		{985, "AKW"},
		{986, "AKX"},
		{987, "AKY"},
		{988, "AKZ"},
		{989, "ALA"},
		{990, "ALB"},
		{991, "ALC"},
		{992, "ALD"},
		{993, "ALE"},
		{994, "ALF"},
		{995, "ALG"},
		{996, "ALH"},
		{997, "ALI"},
		{998, "ALJ"},
		{999, "ALK"},
		{1000, "ALL"},
	}
)

func TestConvertIntToString(t *testing.T) {
	for _, tv := range testValues {
		v := ConvertIntToString(tv.ColumnNumber)
		if tv.ColumnName != v {
			t.Error("Expected ", tv.ColumnName, ", got ", v)
		}
	}
}

func TestConvertStringToInt(t *testing.T) {
	for _, tv := range testValues {
		v := ConvertStringToInt(tv.ColumnName)
		if tv.ColumnNumber != v {
			t.Error("Expected ", tv.ColumnNumber, ", got ", v)
		}
	}

}
