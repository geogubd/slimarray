package slimarray_test

import (
	"fmt"

	"github.com/openacid/slimarray"
)

func ExampleSlimArray() {

	nums := []uint32{
		0, 16, 32, 48, 64, 79, 95, 111, 126, 142, 158, 174, 190, 206, 222, 236,
		252, 268, 275, 278, 281, 283, 285, 289, 296, 301, 304, 307, 311, 313, 318,
		321, 325, 328, 335, 339, 344, 348, 353, 357, 360, 364, 369, 372, 377, 383,
		387, 393, 399, 404, 407, 410, 415, 418, 420, 422, 426, 430, 434, 439, 444,
		446, 448, 451, 456, 459, 462, 465, 470, 473, 479, 482, 488, 490, 494, 500,
		506, 509, 513, 519, 521, 528, 530, 534, 537, 540, 544, 546, 551, 556, 560,
		566, 568, 572, 574, 576, 580, 585, 588, 592, 594, 600, 603, 606, 608, 610,
		614, 620, 623, 628, 630, 632, 638, 644, 647, 653, 658, 660, 662, 665, 670,
		672, 676, 681, 683, 687, 689, 691, 693, 695, 697, 703, 706, 710, 715, 719,
		722, 726, 731, 735, 737, 741, 748, 750, 753, 757, 763, 766, 768, 775, 777,
		782, 785, 791, 795, 798, 800, 806, 811, 815, 818, 821, 824, 829, 832, 836,
		838, 842, 846, 850, 855, 860, 865, 870, 875, 878, 882, 886, 890, 895, 900,
		906, 910, 913, 916, 921, 925, 929, 932, 937, 940, 942, 944, 946, 952, 954,
		956, 958, 962, 966, 968, 971, 975, 979, 983, 987, 989, 994, 997, 1000,
	}

	a := slimarray.NewU32(nums)

	fmt.Println("last elt is:", a.Get(int32(a.Len()-1)))

	st := a.Stat()
	for _, k := range []string{
		"elt_width",
		"mem_elts",
		"bits/elt"} {
		fmt.Printf("%10s : %d\n", k, st[k])
	}

	// Unordered output:
	// last elt is: 1000
	//  elt_width : 3
	//   mem_elts : 112
	//   bits/elt : 16
}
