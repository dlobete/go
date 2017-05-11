// Copyright © 2015-2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package main

import "github.com/platinasystems/go/internal/sriovs"

// The vfs table is 0 based and is adjusted to 1 based beta and production
// units with VfName
var vfs = [][]sriovs.Vf{
	[]sriovs.Vf{
		// pf0
		sriovs.Eth_0_0 | sriovs.Vlan6,
		sriovs.Eth_1_0 | sriovs.Vlan2,
		sriovs.Eth_2_0 | sriovs.Vlan14,
		sriovs.Eth_3_0 | sriovs.Vlan10,
		sriovs.Eth_4_0 | sriovs.Vlan22,
		sriovs.Eth_5_0 | sriovs.Vlan18,
		sriovs.Eth_6_0 | sriovs.Vlan30,
		sriovs.Eth_7_0 | sriovs.Vlan26,
		sriovs.Eth_8_0 | sriovs.Vlan38,
		sriovs.Eth_9_0 | sriovs.Vlan34,
		sriovs.Eth_10_0 | sriovs.Vlan46,
		sriovs.Eth_11_0 | sriovs.Vlan42,
		sriovs.Eth_12_0 | sriovs.Vlan54,
		sriovs.Eth_13_0 | sriovs.Vlan50,
		sriovs.Eth_14_0 | sriovs.Vlan62,
		sriovs.Eth_15_0 | sriovs.Vlan58,
		sriovs.Eth_16_0 | sriovs.Vlan70,
		sriovs.Eth_17_0 | sriovs.Vlan66,
		sriovs.Eth_18_0 | sriovs.Vlan78,
		sriovs.Eth_19_0 | sriovs.Vlan74,
		sriovs.Eth_20_0 | sriovs.Vlan86,
		sriovs.Eth_21_0 | sriovs.Vlan82,
		sriovs.Eth_22_0 | sriovs.Vlan94,
		sriovs.Eth_23_0 | sriovs.Vlan90,
		sriovs.Eth_24_0 | sriovs.Vlan102,
		sriovs.Eth_25_0 | sriovs.Vlan98,
		sriovs.Eth_26_0 | sriovs.Vlan110,
		sriovs.Eth_27_0 | sriovs.Vlan106,
		sriovs.Eth_28_0 | sriovs.Vlan118,
		sriovs.Eth_29_0 | sriovs.Vlan114,
		sriovs.Eth_30_0 | sriovs.Vlan126,
		sriovs.Eth_31_0 | sriovs.Vlan122,
		sriovs.Eth_0_1 | sriovs.Vlan7,
		sriovs.Eth_1_1 | sriovs.Vlan3,
		sriovs.Eth_2_1 | sriovs.Vlan15,
		sriovs.Eth_3_1 | sriovs.Vlan11,
		sriovs.Eth_4_1 | sriovs.Vlan23,
		sriovs.Eth_5_1 | sriovs.Vlan19,
		sriovs.Eth_6_1 | sriovs.Vlan31,
		sriovs.Eth_7_1 | sriovs.Vlan27,
		sriovs.Eth_8_1 | sriovs.Vlan39,
		sriovs.Eth_9_1 | sriovs.Vlan35,
		sriovs.Eth_10_1 | sriovs.Vlan47,
		sriovs.Eth_11_1 | sriovs.Vlan43,
		sriovs.Eth_12_1 | sriovs.Vlan55,
		sriovs.Eth_13_1 | sriovs.Vlan51,
		sriovs.Eth_14_1 | sriovs.Vlan63,
		sriovs.Eth_15_1 | sriovs.Vlan59,
		sriovs.Eth_16_1 | sriovs.Vlan71,
		sriovs.Eth_17_1 | sriovs.Vlan67,
		sriovs.Eth_18_1 | sriovs.Vlan79,
		sriovs.Eth_19_1 | sriovs.Vlan75,
		sriovs.Eth_20_1 | sriovs.Vlan87,
		sriovs.Eth_21_1 | sriovs.Vlan83,
		sriovs.Eth_22_1 | sriovs.Vlan95,
		sriovs.Eth_23_1 | sriovs.Vlan91,
		sriovs.Eth_24_1 | sriovs.Vlan103,
		sriovs.Eth_25_1 | sriovs.Vlan99,
		sriovs.Eth_26_1 | sriovs.Vlan111,
		sriovs.Eth_27_1 | sriovs.Vlan107,
		sriovs.Eth_28_1 | sriovs.Vlan119,
		sriovs.Eth_29_1 | sriovs.Vlan115,
		sriovs.Eth_30_1 | sriovs.Vlan127,
		// sriovs.Eth_31_1 | sriovs.Vlan123,
		/*FIXME
		}, []sriovs.Vf{
			// pf1
			sriovs.Eth_0_2 | sriovs.Vlan8,
			sriovs.Eth_1_2 | sriovs.Vlan4,
			sriovs.Eth_2_2 | sriovs.Vlan16,
			sriovs.Eth_3_2 | sriovs.Vlan12,
			sriovs.Eth_4_2 | sriovs.Vlan24,
			sriovs.Eth_5_2 | sriovs.Vlan20,
			sriovs.Eth_6_2 | sriovs.Vlan32,
			sriovs.Eth_7_2 | sriovs.Vlan28,
			sriovs.Eth_8_2 | sriovs.Vlan40,
			sriovs.Eth_9_2 | sriovs.Vlan36,
			sriovs.Eth_10_2 | sriovs.Vlan48,
			sriovs.Eth_11_2 | sriovs.Vlan44,
			sriovs.Eth_12_2 | sriovs.Vlan56,
			sriovs.Eth_13_2 | sriovs.Vlan52,
			sriovs.Eth_14_2 | sriovs.Vlan64,
			sriovs.Eth_15_2 | sriovs.Vlan60,
			sriovs.Eth_16_2 | sriovs.Vlan72,
			sriovs.Eth_17_2 | sriovs.Vlan68,
			sriovs.Eth_18_2 | sriovs.Vlan80,
			sriovs.Eth_19_2 | sriovs.Vlan76,
			sriovs.Eth_20_2 | sriovs.Vlan88,
			sriovs.Eth_21_2 | sriovs.Vlan84,
			sriovs.Eth_22_2 | sriovs.Vlan96,
			sriovs.Eth_23_2 | sriovs.Vlan92,
			sriovs.Eth_24_2 | sriovs.Vlan104,
			sriovs.Eth_25_2 | sriovs.Vlan100,
			sriovs.Eth_26_2 | sriovs.Vlan112,
			sriovs.Eth_27_2 | sriovs.Vlan108,
			sriovs.Eth_28_2 | sriovs.Vlan120,
			sriovs.Eth_29_2 | sriovs.Vlan116,
			sriovs.Eth_30_2 | sriovs.Vlan128,
			sriovs.Eth_31_2 | sriovs.Vlan124,
			sriovs.Eth_0_3 | sriovs.Vlan9,
			sriovs.Eth_1_3 | sriovs.Vlan5,
			sriovs.Eth_2_3 | sriovs.Vlan17,
			sriovs.Eth_3_3 | sriovs.Vlan13,
			sriovs.Eth_4_3 | sriovs.Vlan25,
			sriovs.Eth_5_3 | sriovs.Vlan21,
			sriovs.Eth_6_3 | sriovs.Vlan33,
			sriovs.Eth_7_3 | sriovs.Vlan29,
			sriovs.Eth_8_3 | sriovs.Vlan41,
			sriovs.Eth_9_3 | sriovs.Vlan37,
			sriovs.Eth_10_3 | sriovs.Vlan49,
			sriovs.Eth_11_3 | sriovs.Vlan45,
			sriovs.Eth_12_3 | sriovs.Vlan57,
			sriovs.Eth_13_3 | sriovs.Vlan53,
			sriovs.Eth_14_3 | sriovs.Vlan65,
			sriovs.Eth_15_3 | sriovs.Vlan61,
			sriovs.Eth_16_3 | sriovs.Vlan73,
			sriovs.Eth_17_3 | sriovs.Vlan69,
			sriovs.Eth_18_3 | sriovs.Vlan81,
			sriovs.Eth_19_3 | sriovs.Vlan77,
			sriovs.Eth_20_3 | sriovs.Vlan89,
			sriovs.Eth_21_3 | sriovs.Vlan85,
			sriovs.Eth_22_3 | sriovs.Vlan97,
			sriovs.Eth_23_3 | sriovs.Vlan93,
			sriovs.Eth_24_3 | sriovs.Vlan105,
			sriovs.Eth_25_3 | sriovs.Vlan101,
			sriovs.Eth_26_3 | sriovs.Vlan113,
			sriovs.Eth_27_3 | sriovs.Vlan109,
			sriovs.Eth_28_3 | sriovs.Vlan121,
			sriovs.Eth_29_3 | sriovs.Vlan117,
			sriovs.Eth_30_3 | sriovs.Vlan129,
			// sriovs.Eth_31_3 | sriovs.Vlan125,
		*/
	},
}