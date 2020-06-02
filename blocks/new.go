package blocks

type NewMapping map[uint16]int

type NewBlock struct {
	Name     string
	Fallback *string
	Ids      NewMapping
}

var newBlocks = []NewBlock{
	{
		"minecraft:jungle_button[face=wall,facing=west,powered=true]",
		nil,
		NewMapping{
			393: 5387,
			404: 5388,
			477: 5894,
			573: 5894,
		},
	},
	{
		"minecraft:cobblestone_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 5648,
			573: 5648,
			393: 5144,
			404: 5145,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=11,south=up,west=side]",
		nil,
		NewMapping{
			477: 2588,
			573: 2588,
			393: 2284,
			404: 2285,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 11026,
			573: 11026,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=south,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 9158,
			477: 9158,
		},
	},
	{
		"minecraft:birch_fence[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 7574,
			404: 7575,
			477: 8099,
			573: 8099,
		},
	},
	{
		"minecraft:light_gray_banner[rotation=3]",
		nil,
		NewMapping{
			573: 7492,
			393: 6985,
			404: 6986,
			477: 7492,
		},
	},
	{
		"minecraft:chorus_plant[down=true,east=true,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			573: 8532,
			393: 8007,
			404: 8008,
			477: 8532,
		},
	},
	{
		"minecraft:birch_door[facing=north,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7750,
			404: 7751,
			477: 8275,
			573: 8275,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=13,powered=false]",
		nil,
		NewMapping{
			477: 925,
			573: 925,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=north,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7422,
			404: 7423,
			477: 7947,
			573: 7947,
		},
	},
	{
		"minecraft:gray_banner[rotation=7]",
		nil,
		NewMapping{
			393: 6973,
			404: 6974,
			477: 7480,
			573: 7480,
		},
	},
	{
		"minecraft:bell[attachment=floor,facing=east]",
		nil,
		NewMapping{
			477: 11201,
		},
	},
	{
		"minecraft:cactus[age=6]",
		nil,
		NewMapping{
			4:   1302,
			393: 3431,
			404: 3432,
			477: 3935,
			573: 3935,
		},
	},
	{
		"minecraft:green_wall_banner[facing=east]",
		nil,
		NewMapping{
			393: 7165,
			404: 7166,
			477: 7672,
			573: 7672,
		},
	},
	{
		"minecraft:spruce_door[facing=north,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7691,
			404: 7692,
			477: 8216,
			573: 8216,
		},
	},
	{
		"minecraft:magenta_banner[rotation=10]",
		nil,
		NewMapping{
			393: 6896,
			404: 6897,
			477: 7403,
			573: 7403,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5868,
			404: 5869,
			477: 6375,
			573: 6375,
		},
	},
	{
		"minecraft:jungle_fence[east=true,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7591,
			404: 7592,
			477: 8116,
			573: 8116,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10357,
			573: 10357,
		},
	},
	{
		"minecraft:granite_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9925,
			573: 9925,
		},
	},
	{
		"minecraft:dispenser[facing=south,triggered=true]",
		nil,
		NewMapping{
			4:   379,
			393: 237,
			404: 237,
			477: 237,
			573: 237,
		},
	},
	{
		"minecraft:brewing_stand[has_bottle_0=false,has_bottle_1=true,has_bottle_2=false]",
		nil,
		NewMapping{
			404: 4619,
			477: 5122,
			573: 5122,
			4:   1874,
			393: 4618,
		},
	},
	{
		"minecraft:comparator[facing=east,mode=subtract,powered=true]",
		nil,
		NewMapping{
			4:   2415,
			393: 5649,
			404: 5650,
			477: 6156,
			573: 6156,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=14,south=up,west=side]",
		nil,
		NewMapping{
			477: 2903,
			573: 2903,
			393: 2599,
			404: 2600,
		},
	},
	{
		"minecraft:acacia_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 6885,
			573: 6885,
			393: 6378,
			404: 6379,
		},
	},
	{
		"minecraft:bee_nest[facing=east,honey_level=0]",
		nil,
		NewMapping{
			573: 11305,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=11,powered=true]",
		nil,
		NewMapping{
			477: 970,
			573: 970,
		},
	},
	{
		"minecraft:oak_leaves[distance=2,persistent=false]",
		nil,
		NewMapping{
			4:   296,
			393: 147,
			404: 147,
			477: 147,
			573: 147,
		},
	},
	{
		"minecraft:chorus_plant[down=true,east=false,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			404: 8027,
			477: 8551,
			573: 8551,
			393: 8026,
		},
	},
	{
		"minecraft:light_blue_bed[facing=north,occupied=false,part=foot]",
		nil,
		NewMapping{
			404: 799,
			477: 1099,
			573: 1099,
			393: 799,
		},
	},
	{
		"minecraft:oak_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			404: 1713,
			477: 2016,
			573: 2016,
			393: 1712,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10962,
			573: 10962,
		},
	},
	{
		"minecraft:fire[age=3,east=false,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1254,
			404: 1255,
			477: 1558,
			573: 1558,
		},
	},
	{
		"minecraft:cactus[age=12]",
		nil,
		NewMapping{
			393: 3437,
			404: 3438,
			477: 3941,
			573: 3941,
			4:   1308,
		},
	},
	{
		"minecraft:potatoes[age=5]",
		nil,
		NewMapping{
			4:   2277,
			393: 5300,
			404: 5301,
			477: 5807,
			573: 5807,
		},
	},
	{
		"minecraft:acacia_fence[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 8163,
			393: 7638,
			404: 7639,
			477: 8163,
		},
	},
	{
		"minecraft:spruce_fence[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 8044,
			393: 7519,
			404: 7520,
			477: 8044,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 7751,
			393: 7244,
			404: 7245,
			477: 7751,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6780,
			404: 6781,
			477: 7287,
			573: 7287,
		},
	},
	{
		"minecraft:quartz_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 6226,
			573: 6226,
			393: 5719,
			404: 5720,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=3,south=none,west=side]",
		nil,
		NewMapping{
			393: 2794,
			404: 2795,
			477: 3098,
			573: 3098,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 6585,
			393: 6078,
			404: 6079,
			477: 6585,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 5758,
			393: 5254,
			404: 5255,
			477: 5758,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 6799,
			573: 6799,
			393: 6292,
			404: 6293,
		},
	},
	{
		"minecraft:spruce_button[face=floor,facing=north,powered=false]",
		nil,
		NewMapping{
			393: 5328,
			404: 5329,
			477: 5835,
			573: 5835,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=14,south=none,west=none]",
		nil,
		NewMapping{
			573: 2334,
			393: 2030,
			404: 2031,
			477: 2334,
		},
	},
	{
		"minecraft:acacia_fence[east=false,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 7633,
			477: 8157,
			573: 8157,
			393: 7632,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4542,
			404: 4543,
			477: 5046,
			573: 5046,
		},
	},
	{
		"minecraft:fire[age=12,east=true,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1525,
			404: 1526,
			477: 1829,
			573: 1829,
		},
	},
	{
		"minecraft:stone_button[face=ceiling,facing=west,powered=true]",
		nil,
		NewMapping{
			393: 3411,
			404: 3412,
			477: 3915,
			573: 3915,
		},
	},
	{
		"minecraft:diorite_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10209,
			573: 10209,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=south,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6746,
			404: 6747,
			477: 7253,
			573: 7253,
		},
	},
	{
		"minecraft:stone_button[face=floor,facing=west,powered=true]",
		nil,
		NewMapping{
			393: 3395,
			404: 3396,
			477: 3899,
			573: 3899,
		},
	},
	{
		"minecraft:cyan_banner[rotation=3]",
		nil,
		NewMapping{
			393: 7001,
			404: 7002,
			477: 7508,
			573: 7508,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=0,waterlogged=true]",
		nil,
		NewMapping{
			477: 3539,
			573: 3539,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10077,
			573: 10077,
		},
	},
	{
		"minecraft:soul_sand",
		nil,
		NewMapping{
			4:   1408,
			393: 3494,
			404: 3495,
			477: 3998,
			573: 3998,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=9,south=none,west=up]",
		nil,
		NewMapping{
			393: 2847,
			404: 2848,
			477: 3151,
			573: 3151,
		},
	},
	{
		"minecraft:spruce_button[face=floor,facing=east,powered=false]",
		nil,
		NewMapping{
			393: 5334,
			404: 5335,
			477: 5841,
			573: 5841,
		},
	},
	{
		"minecraft:turtle_egg[eggs=2,hatch=1]",
		nil,
		NewMapping{
			393: 8441,
			404: 8442,
			477: 8966,
			573: 8966,
		},
	},
	{
		"minecraft:acacia_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6337,
			404: 6338,
			477: 6844,
			573: 6844,
		},
	},
	{
		"minecraft:blue_shulker_box[facing=west]",
		nil,
		NewMapping{
			4:   3684,
			393: 8286,
			404: 8287,
			477: 8811,
			573: 8811,
		},
	},
	{
		"minecraft:stone_button[face=floor,facing=east,powered=true]",
		nil,
		NewMapping{
			477: 3901,
			573: 3901,
			393: 3397,
			404: 3398,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=14,south=side,west=none]",
		nil,
		NewMapping{
			393: 3035,
			404: 3036,
			477: 3339,
			573: 3339,
		},
	},
	{
		"minecraft:diorite_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10232,
			573: 10232,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9610,
			573: 9610,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6007,
			404: 6008,
			477: 6514,
			573: 6514,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=false,north=false,powered=true,south=true,west=false]",
		nil,
		NewMapping{
			393: 4812,
			404: 4813,
			477: 5316,
			573: 5316,
		},
	},
	{
		"minecraft:fire[age=13,east=false,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			404: 1575,
			477: 1878,
			573: 1878,
			393: 1574,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=7,powered=false]",
		nil,
		NewMapping{
			477: 863,
			573: 863,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=true,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10793,
			573: 10793,
		},
	},
	{
		"minecraft:oak_button[face=ceiling,facing=north,powered=true]",
		nil,
		NewMapping{
			4:   2296,
			393: 5319,
			404: 5320,
			477: 5826,
			573: 5826,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=10,south=side,west=none]",
		nil,
		NewMapping{
			393: 2855,
			404: 2856,
			477: 3159,
			573: 3159,
		},
	},
	{
		"minecraft:spruce_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4923,
			404: 4924,
			477: 5427,
			573: 5427,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9512,
			573: 9512,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6274,
			404: 6275,
			477: 6781,
			573: 6781,
		},
	},
	{
		"minecraft:light_blue_wall_banner[facing=east]",
		nil,
		NewMapping{
			393: 7125,
			404: 7126,
			477: 7632,
			573: 7632,
		},
	},
	{
		"minecraft:iron_bars[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 4195,
			404: 4196,
			477: 4699,
			573: 4699,
		},
	},
	{
		"minecraft:redstone_wall_torch[facing=north,lit=true]",
		nil,
		NewMapping{
			4:   1220,
			393: 3383,
			404: 3384,
			477: 3887,
			573: 3887,
		},
	},
	{
		"minecraft:light_gray_wool",
		nil,
		NewMapping{
			393: 1091,
			404: 1091,
			477: 1391,
			573: 1391,
			4:   568,
		},
	},
	{
		"minecraft:fire[age=13,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1580,
			404: 1581,
			477: 1884,
			573: 1884,
		},
	},
	{
		"minecraft:iron_door[facing=north,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 3317,
			404: 3318,
			477: 3821,
			573: 3821,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 6282,
			477: 6788,
			573: 6788,
			393: 6281,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=5,south=up,west=none]",
		nil,
		NewMapping{
			393: 2951,
			404: 2952,
			477: 3255,
			573: 3255,
		},
	},
	{
		"minecraft:repeater[delay=4,facing=east,locked=true,powered=true]",
		nil,
		NewMapping{
			393: 3573,
			404: 3574,
			477: 4077,
			573: 4077,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=10,south=side,west=side]",
		nil,
		NewMapping{
			393: 1990,
			404: 1991,
			477: 2294,
			573: 2294,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=8,south=up,west=none]",
		nil,
		NewMapping{
			393: 1970,
			404: 1971,
			477: 2274,
			573: 2274,
		},
	},
	{
		"minecraft:stone_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			393: 7297,
			477: 7804,
			573: 7804,
		},
	},
	{
		"minecraft:jungle_door[facing=north,half=upper,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			477: 8337,
			573: 8337,
			393: 7812,
			404: 7813,
		},
	},
	{
		"minecraft:melon_stem[age=4]",
		nil,
		NewMapping{
			4:   1684,
			393: 4264,
			404: 4265,
			477: 4768,
			573: 4768,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=20,powered=true]",
		nil,
		NewMapping{
			393: 338,
			404: 338,
			477: 338,
			573: 338,
		},
	},
	{
		"minecraft:iron_door[facing=west,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			573: 3845,
			393: 3341,
			404: 3342,
			477: 3845,
		},
	},
	{
		"minecraft:jungle_door[facing=west,half=lower,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7845,
			404: 7846,
			477: 8370,
			573: 8370,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=6,south=none,west=side]",
		nil,
		NewMapping{
			393: 2389,
			404: 2390,
			477: 2693,
			573: 2693,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 7233,
			477: 7739,
			573: 7739,
			393: 7232,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9206,
			573: 9206,
		},
	},
	{
		"minecraft:andesite_stairs[facing=north,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9935,
			573: 9935,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=6,south=up,west=up]",
		nil,
		NewMapping{
			393: 2670,
			404: 2671,
			477: 2974,
			573: 2974,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=19,powered=false]",
		nil,
		NewMapping{
			477: 337,
			573: 337,
			393: 337,
			404: 337,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=false,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 11065,
			573: 11065,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10695,
			573: 10695,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=6,powered=false]",
		nil,
		NewMapping{
			393: 611,
			404: 611,
			477: 611,
			573: 611,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5175,
			404: 5176,
			477: 5679,
			573: 5679,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=16,powered=true]",
		nil,
		NewMapping{
			393: 580,
			404: 580,
			477: 580,
			573: 580,
		},
	},
	{
		"minecraft:stripped_spruce_wood[axis=y]",
		nil,
		NewMapping{
			393: 130,
			404: 130,
			477: 130,
			573: 130,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=8,powered=true]",
		nil,
		NewMapping{
			477: 964,
			573: 964,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=9,powered=true]",
		nil,
		NewMapping{
			477: 766,
			573: 766,
		},
	},
	{
		"minecraft:sugar_cane[age=8]",
		nil,
		NewMapping{
			4:   1336,
			393: 3450,
			404: 3451,
			477: 3954,
			573: 3954,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=west,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 6535,
			404: 6536,
			477: 7042,
			573: 7042,
		},
	},
	{
		"minecraft:fire_coral_wall_fan[facing=west,waterlogged=true]",
		nil,
		NewMapping{
			393: 8532,
			404: 8548,
			477: 9092,
			573: 9092,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=east,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7446,
			404: 7447,
			477: 7971,
			573: 7971,
		},
	},
	{
		"minecraft:iron_bars[east=false,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 4207,
			477: 4710,
			573: 4710,
			393: 4206,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 9758,
			477: 9758,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5912,
			404: 5913,
			477: 6419,
			573: 6419,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=south,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			477: 4446,
			573: 4446,
			393: 3942,
			404: 3943,
		},
	},
	{
		"minecraft:dark_oak_fence[east=true,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7651,
			404: 7652,
			477: 8176,
			573: 8176,
		},
	},
	{
		"minecraft:magenta_glazed_terracotta[facing=south]",
		nil,
		NewMapping{
			393: 8322,
			404: 8323,
			477: 8847,
			573: 8847,
			4:   3792,
		},
	},
	{
		"minecraft:dark_oak_leaves[distance=1,persistent=true]",
		nil,
		NewMapping{
			477: 214,
			573: 214,
			4:   2581,
			393: 214,
			404: 214,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=west,in_wall=true,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7440,
			404: 7441,
			477: 7965,
			573: 7965,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			404: 6083,
			477: 6589,
			573: 6589,
			393: 6082,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5217,
			404: 5218,
			477: 5721,
			573: 5721,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=7,powered=false]",
		nil,
		NewMapping{
			477: 713,
			573: 713,
			393: 713,
			404: 713,
		},
	},
	{
		"minecraft:scaffolding[bottom=true,distance=2,waterlogged=true]",
		nil,
		NewMapping{
			477: 11103,
			573: 11103,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=13,south=none,west=side]",
		nil,
		NewMapping{
			573: 2900,
			393: 2596,
			404: 2597,
			477: 2900,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=5,south=up,west=up]",
		nil,
		NewMapping{
			393: 1797,
			404: 1798,
			477: 2101,
			573: 2101,
		},
	},
	{
		"minecraft:dark_oak_door[facing=west,half=upper,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			404: 7973,
			477: 8497,
			573: 8497,
			393: 7972,
		},
	},
	{
		"minecraft:sign[rotation=1,waterlogged=true]",
		nil,
		NewMapping{
			393: 3077,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=north,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			404: 3602,
			477: 4105,
			573: 4105,
			393: 3601,
		},
	},
	{
		"minecraft:frosted_ice[age=3]",
		nil,
		NewMapping{
			393: 8191,
			404: 8192,
			477: 8716,
			573: 8716,
			4:   3395,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5857,
			404: 5858,
			477: 6364,
			573: 6364,
		},
	},
	{
		"minecraft:oak_fence_gate[facing=west,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			393: 4316,
			404: 4317,
			477: 4820,
			573: 4820,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=10,south=none,west=side]",
		nil,
		NewMapping{
			393: 1849,
			404: 1850,
			477: 2153,
			573: 2153,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=9,south=none,west=side]",
		nil,
		NewMapping{
			393: 1984,
			404: 1985,
			477: 2288,
			573: 2288,
		},
	},
	{
		"minecraft:jungle_door[facing=south,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			4:   3125,
			393: 7834,
			404: 7835,
			477: 8359,
			573: 8359,
		},
	},
	{
		"minecraft:oak_button[face=floor,facing=west,powered=false]",
		nil,
		NewMapping{
			404: 5309,
			477: 5815,
			573: 5815,
			393: 5308,
		},
	},
	{
		"minecraft:fire[age=3,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1260,
			404: 1261,
			477: 1564,
			573: 1564,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=7,south=up,west=up]",
		nil,
		NewMapping{
			477: 2695,
			573: 2695,
			393: 2391,
			404: 2392,
		},
	},
	{
		"minecraft:blue_banner[rotation=15]",
		nil,
		NewMapping{
			393: 7045,
			404: 7046,
			477: 7552,
			573: 7552,
		},
	},
	{
		"minecraft:granite_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9887,
			573: 9887,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4723,
			404: 4724,
			477: 5227,
			573: 5227,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 9445,
			477: 9445,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6569,
			404: 6570,
			477: 7076,
			573: 7076,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 7752,
			573: 7752,
			393: 7245,
			404: 7246,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9716,
			573: 9716,
		},
	},
	{
		"minecraft:granite_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9932,
			573: 9932,
		},
	},
	{
		"minecraft:green_shulker_box[facing=west]",
		nil,
		NewMapping{
			393: 8298,
			404: 8299,
			477: 8823,
			573: 8823,
			4:   3716,
		},
	},
	{
		"minecraft:oak_fence[east=true,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 3465,
			404: 3466,
			477: 3969,
			573: 3969,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=east,half=top,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3770,
			404: 3771,
			477: 4274,
			573: 4274,
		},
	},
	{
		"minecraft:fire[age=14,east=false,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1605,
			404: 1606,
			477: 1909,
			573: 1909,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=3,waterlogged=false]",
		nil,
		NewMapping{
			477: 3546,
			573: 3546,
		},
	},
	{
		"minecraft:hopper[enabled=false,facing=down]",
		nil,
		NewMapping{
			393: 5690,
			404: 5691,
			477: 6197,
			573: 6197,
			4:   2472,
		},
	},
	{
		"minecraft:end_rod[facing=down]",
		nil,
		NewMapping{
			4:   3168,
			393: 8002,
			404: 8003,
			477: 8527,
			573: 8527,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=10,powered=false]",
		nil,
		NewMapping{
			393: 619,
			404: 619,
			477: 619,
			573: 619,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5974,
			404: 5975,
			477: 6481,
			573: 6481,
		},
	},
	{
		"minecraft:spruce_stairs[facing=east,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4947,
			404: 4948,
			477: 5451,
			573: 5451,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5854,
			404: 5855,
			477: 6361,
			573: 6361,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9177,
			573: 9177,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=true,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10502,
			573: 10502,
		},
	},
	{
		"minecraft:birch_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 5034,
			404: 5035,
			477: 5538,
			573: 5538,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=5,south=none,west=side]",
		nil,
		NewMapping{
			404: 2525,
			477: 2828,
			573: 2828,
			393: 2524,
		},
	},
	{
		"minecraft:spruce_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			573: 5428,
			393: 4924,
			404: 4925,
			477: 5428,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 6778,
			573: 6778,
			393: 6271,
			404: 6272,
		},
	},
	{
		"minecraft:dead_bubble_coral[waterlogged=false]",
		nil,
		NewMapping{
			404: 8465,
			477: 8989,
			573: 8989,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=11,south=side,west=side]",
		nil,
		NewMapping{
			477: 3167,
			573: 3167,
			393: 2863,
			404: 2864,
		},
	},
	{
		"minecraft:birch_door[facing=east,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7798,
			404: 7799,
			477: 8323,
			573: 8323,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4485,
			404: 4486,
			477: 4989,
			573: 4989,
		},
	},
	{
		"minecraft:granite_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9908,
			573: 9908,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10476,
			477: 10476,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9331,
			573: 9331,
		},
	},
	{
		"minecraft:wall_torch[facing=west]",
		nil,
		NewMapping{
			393: 1133,
			404: 1134,
			477: 1437,
			573: 1437,
			4:   802,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=18,powered=true]",
		nil,
		NewMapping{
			573: 584,
			393: 584,
			404: 584,
			477: 584,
		},
	},
	{
		"minecraft:sea_pickle[pickles=3,waterlogged=false]",
		nil,
		NewMapping{
			393: 8569,
			404: 8585,
			477: 9109,
			573: 9109,
		},
	},
	{
		"minecraft:iron_door[facing=east,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 3361,
			404: 3362,
			477: 3865,
			573: 3865,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=2,south=side,west=up]",
		nil,
		NewMapping{
			393: 2781,
			404: 2782,
			477: 3085,
			573: 3085,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=false,north=false,powered=true,south=false,west=false]",
		nil,
		NewMapping{
			404: 4783,
			477: 5286,
			573: 5286,
			4:   2125,
			393: 4782,
		},
	},
	{
		"minecraft:redstone_ore[lit=false]",
		nil,
		NewMapping{
			4:   1168,
			393: 3380,
			404: 3381,
			477: 3884,
			573: 3884,
		},
	},
	{
		"minecraft:magenta_shulker_box[facing=down]",
		nil,
		NewMapping{
			4:   3536,
			393: 8234,
			404: 8235,
			477: 8759,
			573: 8759,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 5702,
			393: 5198,
			404: 5199,
			477: 5702,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10657,
			573: 10657,
		},
	},
	{
		"minecraft:fire_coral_wall_fan[facing=east,waterlogged=false]",
		nil,
		NewMapping{
			393: 8535,
			404: 8551,
			477: 9095,
			573: 9095,
		},
	},
	{
		"minecraft:purple_banner[rotation=4]",
		nil,
		NewMapping{
			393: 7018,
			404: 7019,
			477: 7525,
			573: 7525,
		},
	},
	{
		"minecraft:diorite_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10217,
			573: 10217,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=false,north=false,powered=false,south=false,west=false]",
		nil,
		NewMapping{
			4:   2122,
			393: 4850,
			404: 4851,
			477: 5354,
			573: 5354,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=5,south=none,west=none]",
		nil,
		NewMapping{
			4:   885,
			393: 2957,
			404: 2958,
			477: 3261,
			573: 3261,
		},
	},
	{
		"minecraft:chorus_plant[down=true,east=true,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			573: 8539,
			393: 8014,
			404: 8015,
			477: 8539,
		},
	},
	{
		"minecraft:purpur_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 8080,
			404: 8081,
			477: 8605,
			573: 8605,
		},
	},
	{
		"minecraft:light_gray_banner[rotation=12]",
		nil,
		NewMapping{
			404: 6995,
			477: 7501,
			573: 7501,
			393: 6994,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=20,powered=false]",
		nil,
		NewMapping{
			477: 889,
			573: 889,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=13,south=none,west=side]",
		nil,
		NewMapping{
			393: 2164,
			404: 2165,
			477: 2468,
			573: 2468,
		},
	},
	{
		"minecraft:smooth_quartz_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			477: 10297,
			573: 10297,
		},
	},
	{
		"minecraft:scaffolding[bottom=false,distance=0,waterlogged=false]",
		nil,
		NewMapping{
			477: 11116,
			573: 11116,
		},
	},
	{
		"minecraft:dark_oak_sapling[stage=1]",
		nil,
		NewMapping{
			4:   109,
			393: 32,
			404: 32,
			477: 32,
			573: 32,
		},
	},
	{
		"minecraft:piston[extended=true,facing=west]",
		nil,
		NewMapping{
			404: 1050,
			477: 1350,
			573: 1350,
			4:   540,
			393: 1050,
		},
	},
	{
		"minecraft:fire[age=14,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			477: 1905,
			573: 1905,
			393: 1601,
			404: 1602,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 7688,
			573: 7688,
			393: 7181,
			404: 7182,
		},
	},
	{
		"minecraft:quartz_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 5710,
			404: 5711,
			477: 6217,
			573: 6217,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=false,north=true,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10555,
			573: 10555,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=west,half=top,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			404: 6530,
			477: 7036,
			573: 7036,
			4:   2686,
			393: 6529,
		},
	},
	{
		"minecraft:repeater[delay=1,facing=east,locked=false,powered=true]",
		nil,
		NewMapping{
			404: 3528,
			477: 4031,
			573: 4031,
			4:   1507,
			393: 3527,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=14,south=none,west=up]",
		nil,
		NewMapping{
			393: 2748,
			404: 2749,
			477: 3052,
			573: 3052,
		},
	},
	{
		"minecraft:black_bed[facing=east,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 1002,
			404: 1002,
			477: 1302,
			573: 1302,
		},
	},
	{
		"minecraft:birch_button[face=floor,facing=south,powered=true]",
		nil,
		NewMapping{
			477: 5860,
			573: 5860,
			393: 5353,
			404: 5354,
		},
	},
	{
		"minecraft:fire[age=0,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			4:   816,
			393: 1166,
			404: 1167,
			477: 1470,
			573: 1470,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=7,south=side,west=side]",
		nil,
		NewMapping{
			404: 2972,
			477: 3275,
			573: 3275,
			393: 2971,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=0,powered=false]",
		nil,
		NewMapping{
			404: 449,
			477: 449,
			573: 449,
			393: 449,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=false,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			477: 4649,
			573: 4649,
			393: 4145,
			404: 4146,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=false,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10676,
			477: 10676,
		},
	},
	{
		"minecraft:repeater[delay=2,facing=north,locked=true,powered=false]",
		nil,
		NewMapping{
			393: 3530,
			404: 3531,
			477: 4034,
			573: 4034,
		},
	},
	{
		"minecraft:fire[age=15,east=true,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			573: 1919,
			393: 1615,
			404: 1616,
			477: 1919,
		},
	},
	{
		"minecraft:jungle_button[face=ceiling,facing=east,powered=false]",
		nil,
		NewMapping{
			573: 5905,
			393: 5398,
			404: 5399,
			477: 5905,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=east,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3719,
			404: 3720,
			477: 4223,
			573: 4223,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10248,
			573: 10248,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 11048,
			573: 11048,
		},
	},
	{
		"minecraft:stone_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9640,
			573: 9640,
		},
	},
	{
		"minecraft:hopper[enabled=true,facing=down]",
		nil,
		NewMapping{
			573: 6192,
			4:   2464,
			393: 5685,
			404: 5686,
			477: 6192,
		},
	},
	{
		"minecraft:dead_fire_coral_wall_fan[facing=north,waterlogged=true]",
		nil,
		NewMapping{
			573: 9048,
			393: 8488,
			404: 8504,
			477: 9048,
		},
	},
	{
		"minecraft:light_blue_banner[rotation=10]",
		nil,
		NewMapping{
			393: 6912,
			404: 6913,
			477: 7419,
			573: 7419,
		},
	},
	{
		"minecraft:purple_banner[rotation=15]",
		nil,
		NewMapping{
			404: 7030,
			477: 7536,
			573: 7536,
			393: 7029,
		},
	},
	{
		"minecraft:white_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5827,
			404: 5828,
			477: 6334,
			573: 6334,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=true,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			573: 4533,
			393: 4029,
			404: 4030,
			477: 4533,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 6525,
			393: 6018,
			404: 6019,
			477: 6525,
		},
	},
	{
		"minecraft:acacia_door[facing=south,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7897,
			404: 7898,
			477: 8422,
			573: 8422,
		},
	},
	{
		"minecraft:potted_blue_orchid",
		nil,
		NewMapping{
			393: 5275,
			404: 5276,
			477: 5779,
			573: 5779,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=west,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3761,
			404: 3762,
			477: 4265,
			573: 4265,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 6503,
			393: 5996,
			404: 5997,
			477: 6503,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6637,
			404: 6638,
			477: 7144,
			573: 7144,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=13,south=side,west=up]",
		nil,
		NewMapping{
			393: 2160,
			404: 2161,
			477: 2464,
			573: 2464,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9543,
			573: 9543,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10165,
			573: 10165,
		},
	},
	{
		"minecraft:pink_shulker_box[facing=west]",
		nil,
		NewMapping{
			404: 8257,
			477: 8781,
			573: 8781,
			4:   3604,
			393: 8256,
		},
	},
	{
		"minecraft:white_stained_glass",
		nil,
		NewMapping{
			393: 3577,
			404: 3578,
			477: 4081,
			573: 4081,
			4:   1520,
		},
	},
	{
		"minecraft:white_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 6327,
			573: 6327,
			393: 5820,
			404: 5821,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=17,powered=false]",
		nil,
		NewMapping{
			573: 483,
			393: 483,
			404: 483,
			477: 483,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=west,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7470,
			404: 7471,
			477: 7995,
			573: 7995,
		},
	},
	{
		"minecraft:damaged_anvil[facing=west]",
		nil,
		NewMapping{
			393: 5577,
			404: 5578,
			477: 6084,
			573: 6084,
			4:   2329,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=4,south=none,west=side]",
		nil,
		NewMapping{
			393: 2083,
			404: 2084,
			477: 2387,
			573: 2387,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=11,south=side,west=none]",
		nil,
		NewMapping{
			393: 2000,
			404: 2001,
			477: 2304,
			573: 2304,
		},
	},
	{
		"minecraft:oak_sign[rotation=15,waterlogged=false]",
		nil,
		NewMapping{
			404: 3107,
			477: 3410,
			573: 3410,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=13,powered=true]",
		nil,
		NewMapping{
			477: 974,
			573: 974,
		},
	},
	{
		"minecraft:fire[age=7,east=true,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1367,
			404: 1368,
			477: 1671,
			573: 1671,
		},
	},
	{
		"minecraft:dead_tube_coral[waterlogged=true]",
		nil,
		NewMapping{
			404: 8460,
			477: 8984,
			573: 8984,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9149,
			573: 9149,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=6,south=none,west=none]",
		nil,
		NewMapping{
			4:   886,
			393: 2966,
			404: 2967,
			477: 3270,
			573: 3270,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=12,powered=true]",
		nil,
		NewMapping{
			393: 622,
			404: 622,
			477: 622,
			573: 622,
		},
	},
	{
		"minecraft:iron_door[facing=west,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 3337,
			404: 3338,
			477: 3841,
			573: 3841,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=13,south=side,west=up]",
		nil,
		NewMapping{
			393: 3024,
			404: 3025,
			477: 3328,
			573: 3328,
		},
	},
	{
		"minecraft:oak_door[facing=north,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 3121,
			404: 3122,
			477: 3585,
			573: 3585,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=false,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10356,
			573: 10356,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=false,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 6108,
			477: 6614,
			573: 6614,
			4:   2568,
			393: 6107,
		},
	},
	{
		"minecraft:repeater[delay=2,facing=north,locked=true,powered=true]",
		nil,
		NewMapping{
			393: 3529,
			404: 3530,
			477: 4033,
			573: 4033,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6314,
			404: 6315,
			477: 6821,
			573: 6821,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9510,
			573: 9510,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 11039,
			573: 11039,
		},
	},
	{
		"minecraft:wall_sign[facing=east,waterlogged=false]",
		nil,
		NewMapping{
			4:   1093,
			393: 3276,
		},
	},
	{
		"minecraft:dead_brain_coral_fan[waterlogged=false]",
		nil,
		NewMapping{
			404: 8563,
			477: 9007,
			573: 9007,
			393: 8547,
		},
	},
	{
		"minecraft:fire[age=7,east=true,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			573: 1668,
			393: 1364,
			404: 1365,
			477: 1668,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9511,
			573: 9511,
		},
	},
	{
		"minecraft:bee_nest[facing=south,honey_level=3]",
		nil,
		NewMapping{
			573: 11296,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=3,south=up,west=none]",
		nil,
		NewMapping{
			393: 2357,
			404: 2358,
			477: 2661,
			573: 2661,
		},
	},
	{
		"minecraft:white_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5824,
			404: 5825,
			477: 6331,
			573: 6331,
		},
	},
	{
		"minecraft:diorite_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 10173,
			573: 10173,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=6,south=none,west=side]",
		nil,
		NewMapping{
			573: 2837,
			393: 2533,
			404: 2534,
			477: 2837,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4674,
			404: 4675,
			477: 5178,
			573: 5178,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=9,powered=false]",
		nil,
		NewMapping{
			477: 1017,
			573: 1017,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=false,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 4006,
			404: 4007,
			477: 4510,
			573: 4510,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 7223,
			393: 6716,
			404: 6717,
			477: 7223,
		},
	},
	{
		"minecraft:birch_door[facing=north,half=upper,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7741,
			404: 7742,
			477: 8266,
			573: 8266,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 6752,
			404: 6753,
			477: 7259,
			573: 7259,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			404: 6686,
			477: 7192,
			573: 7192,
			393: 6685,
		},
	},
	{
		"minecraft:jungle_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5117,
			404: 5118,
			477: 5621,
			573: 5621,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=false,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 4142,
			404: 4143,
			477: 4646,
			573: 4646,
		},
	},
	{
		"minecraft:yellow_bed[facing=south,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 818,
			404: 818,
			477: 1118,
			573: 1118,
		},
	},
	{
		"minecraft:repeater[delay=1,facing=north,locked=false,powered=true]",
		nil,
		NewMapping{
			404: 3516,
			477: 4019,
			573: 4019,
			4:   1506,
			393: 3515,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=north,half=top,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			477: 7007,
			573: 7007,
			393: 6500,
			404: 6501,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=7,south=side,west=up]",
		nil,
		NewMapping{
			393: 1962,
			404: 1963,
			477: 2266,
			573: 2266,
		},
	},
	{
		"minecraft:fire[age=13,east=true,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 1560,
			404: 1561,
			477: 1864,
			573: 1864,
		},
	},
	{
		"minecraft:spruce_fence[east=false,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 8060,
			393: 7535,
			404: 7536,
			477: 8060,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=12,powered=true]",
		nil,
		NewMapping{
			393: 372,
			404: 372,
			477: 372,
			573: 372,
		},
	},
	{
		"minecraft:black_shulker_box[facing=west]",
		nil,
		NewMapping{
			4:   3748,
			393: 8310,
			404: 8311,
			477: 8835,
			573: 8835,
		},
	},
	{
		"minecraft:jungle_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5101,
			404: 5102,
			477: 5605,
			573: 5605,
		},
	},
	{
		"minecraft:birch_door[facing=east,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7793,
			404: 7794,
			477: 8318,
			573: 8318,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=false,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 8058,
			404: 8059,
			477: 8583,
			573: 8583,
		},
	},
	{
		"minecraft:birch_door[facing=north,half=upper,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			404: 7743,
			477: 8267,
			573: 8267,
			393: 7742,
		},
	},
	{
		"minecraft:dark_oak_button[face=wall,facing=north,powered=false]",
		nil,
		NewMapping{
			393: 5432,
			404: 5433,
			477: 5939,
			573: 5939,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=14,south=none,west=side]",
		nil,
		NewMapping{
			393: 2605,
			404: 2606,
			477: 2909,
			573: 2909,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=13,waterlogged=true]",
		nil,
		NewMapping{
			477: 3565,
			573: 3565,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=16,powered=false]",
		nil,
		NewMapping{
			477: 831,
			573: 831,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9706,
			573: 9706,
		},
	},
	{
		"minecraft:lime_bed[facing=east,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 840,
			404: 840,
			477: 1140,
			573: 1140,
		},
	},
	{
		"minecraft:oak_leaves[distance=5,persistent=true]",
		nil,
		NewMapping{
			393: 152,
			404: 152,
			477: 152,
			573: 152,
		},
	},
	{
		"minecraft:fire[age=13,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			573: 1856,
			393: 1552,
			404: 1553,
			477: 1856,
		},
	},
	{
		"minecraft:fire[age=4,east=true,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			477: 1572,
			573: 1572,
			393: 1268,
			404: 1269,
		},
	},
	{
		"minecraft:grindstone[face=ceiling,facing=north]",
		nil,
		NewMapping{
			477: 11173,
			573: 11173,
		},
	},
	{
		"minecraft:oak_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 1665,
			404: 1666,
			477: 1969,
			573: 1969,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=7,south=side,west=none]",
		nil,
		NewMapping{
			477: 2412,
			573: 2412,
			393: 2108,
			404: 2109,
		},
	},
	{
		"minecraft:quartz_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 5726,
			404: 5727,
			477: 6233,
			573: 6233,
		},
	},
	{
		"minecraft:light_blue_banner[rotation=13]",
		nil,
		NewMapping{
			393: 6915,
			404: 6916,
			477: 7422,
			573: 7422,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=4,powered=true]",
		nil,
		NewMapping{
			393: 556,
			404: 556,
			477: 556,
			573: 556,
		},
	},
	{
		"minecraft:activator_rail[powered=true,shape=ascending_north]",
		nil,
		NewMapping{
			4:   2524,
			393: 5784,
			404: 5785,
			477: 6291,
			573: 6291,
		},
	},
	{
		"minecraft:birch_door[facing=north,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			4:   3107,
			393: 7756,
			404: 7757,
			477: 8281,
			573: 8281,
		},
	},
	{
		"minecraft:dark_oak_door[facing=north,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			573: 8473,
			4:   3155,
			393: 7948,
			404: 7949,
			477: 8473,
		},
	},
	{
		"minecraft:acacia_button[face=floor,facing=east,powered=false]",
		nil,
		NewMapping{
			393: 5406,
			404: 5407,
			477: 5913,
			573: 5913,
		},
	},
	{
		"minecraft:acacia_sign[rotation=5,waterlogged=true]",
		nil,
		NewMapping{
			477: 3485,
			573: 3485,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9573,
			573: 9573,
		},
	},
	{
		"minecraft:nether_brick_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			4:   710,
			393: 7332,
			404: 7333,
			477: 7851,
			573: 7851,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=7,south=up,west=up]",
		nil,
		NewMapping{
			573: 2263,
			393: 1959,
			404: 1960,
			477: 2263,
		},
	},
	{
		"minecraft:light_blue_bed[facing=north,occupied=false,part=head]",
		nil,
		NewMapping{
			404: 798,
			477: 1098,
			573: 1098,
			393: 798,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6111,
			404: 6112,
			477: 6618,
			573: 6618,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=20,powered=true]",
		nil,
		NewMapping{
			573: 938,
			477: 938,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5999,
			404: 6000,
			477: 6506,
			573: 6506,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10696,
			477: 10696,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 8064,
			404: 8065,
			477: 8589,
			573: 8589,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=13,south=up,west=side]",
		nil,
		NewMapping{
			404: 2303,
			477: 2606,
			573: 2606,
			393: 2302,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9142,
			573: 9142,
		},
	},
	{
		"minecraft:purpur_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			4:   3264,
			393: 7352,
			404: 7353,
			477: 7877,
			573: 7877,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6218,
			404: 6219,
			477: 6725,
			573: 6725,
		},
	},
	{
		"minecraft:acacia_fence[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7627,
			404: 7628,
			477: 8152,
			573: 8152,
		},
	},
	{
		"minecraft:mushroom_stem[down=false,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 4165,
			404: 4166,
			477: 4669,
			573: 4669,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 6807,
			393: 6300,
			404: 6301,
			477: 6807,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10166,
			573: 10166,
		},
	},
	{
		"minecraft:acacia_wall_sign[facing=west,waterlogged=false]",
		nil,
		NewMapping{
			477: 3762,
			573: 3762,
		},
	},
	{
		"minecraft:purple_bed[facing=east,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 920,
			404: 920,
			477: 1220,
			573: 1220,
		},
	},
	{
		"minecraft:brick_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4339,
			404: 4340,
			477: 4843,
			573: 4843,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=east,half=top,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3772,
			404: 3773,
			477: 4276,
			573: 4276,
		},
	},
	{
		"minecraft:birch_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 5479,
			4:   2163,
			393: 4975,
			404: 4976,
			477: 5479,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=west,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3828,
			404: 3829,
			477: 4332,
			573: 4332,
		},
	},
	{
		"minecraft:brick_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 4873,
			393: 4369,
			404: 4370,
			477: 4873,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=9,south=none,west=side]",
		nil,
		NewMapping{
			393: 2128,
			404: 2129,
			477: 2432,
			573: 2432,
		},
	},
	{
		"minecraft:repeater[delay=2,facing=west,locked=true,powered=false]",
		nil,
		NewMapping{
			393: 3538,
			404: 3539,
			477: 4042,
			573: 4042,
		},
	},
	{
		"minecraft:spruce_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4962,
			404: 4963,
			477: 5466,
			573: 5466,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=east,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 7239,
			404: 7240,
			477: 7746,
			573: 7746,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10800,
			573: 10800,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6736,
			404: 6737,
			477: 7243,
			573: 7243,
		},
	},
	{
		"minecraft:acacia_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 6369,
			477: 6875,
			573: 6875,
			393: 6368,
		},
	},
	{
		"minecraft:fire[age=9,east=true,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 1434,
			404: 1435,
			477: 1738,
			573: 1738,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10475,
			573: 10475,
		},
	},
	{
		"minecraft:campfire[facing=west,lit=true,signal_fire=false,waterlogged=false]",
		nil,
		NewMapping{
			477: 11235,
			573: 11251,
		},
	},
	{
		"minecraft:jungle_leaves[distance=2,persistent=true]",
		nil,
		NewMapping{
			477: 188,
			573: 188,
			4:   303,
			393: 188,
			404: 188,
		},
	},
	{
		"minecraft:birch_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			404: 5019,
			477: 5522,
			573: 5522,
			393: 5018,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=false,north=true,powered=true,south=true,west=false]",
		nil,
		NewMapping{
			393: 4772,
			404: 4773,
			477: 5276,
			573: 5276,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 6711,
			393: 6204,
			404: 6205,
			477: 6711,
		},
	},
	{
		"minecraft:jungle_leaves[distance=3,persistent=false]",
		nil,
		NewMapping{
			393: 191,
			404: 191,
			477: 191,
			573: 191,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6450,
			404: 6451,
			477: 6957,
			573: 6957,
		},
	},
	{
		"minecraft:brick_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 4850,
			573: 4850,
			393: 4346,
			404: 4347,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=1,waterlogged=true]",
		nil,
		NewMapping{
			477: 3541,
			573: 3541,
		},
	},
	{
		"minecraft:andesite_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9957,
			573: 9957,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=true,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			573: 8560,
			393: 8035,
			404: 8036,
			477: 8560,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 3654,
			573: 3654,
			4:   1079,
			393: 3190,
			404: 3191,
		},
	},
	{
		"minecraft:fire[age=8,east=false,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 1416,
			404: 1417,
			477: 1720,
			573: 1720,
		},
	},
	{
		"minecraft:potted_brown_mushroom",
		nil,
		NewMapping{
			393: 5284,
			404: 5285,
			477: 5791,
			573: 5791,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=0,powered=true]",
		nil,
		NewMapping{
			393: 448,
			404: 448,
			477: 448,
			573: 448,
		},
	},
	{
		"minecraft:spruce_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4910,
			404: 4911,
			477: 5414,
			573: 5414,
		},
	},
	{
		"minecraft:oak_fence_gate[facing=east,in_wall=false,open=false,powered=false]",
		nil,
		NewMapping{
			573: 4835,
			4:   1715,
			393: 4331,
			404: 4332,
			477: 4835,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 3687,
			393: 3223,
			404: 3224,
			477: 3687,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=11,powered=true]",
		nil,
		NewMapping{
			477: 270,
			573: 270,
			393: 270,
			404: 270,
		},
	},
	{
		"minecraft:jungle_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 5109,
			404: 5110,
			477: 5613,
			573: 5613,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=south,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3684,
			404: 3685,
			477: 4188,
			573: 4188,
		},
	},
	{
		"minecraft:fire[age=4,east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			573: 1595,
			393: 1291,
			404: 1292,
			477: 1595,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=20,powered=true]",
		nil,
		NewMapping{
			477: 888,
			573: 888,
		},
	},
	{
		"minecraft:pink_banner[rotation=8]",
		nil,
		NewMapping{
			393: 6958,
			404: 6959,
			477: 7465,
			573: 7465,
		},
	},
	{
		"minecraft:campfire[facing=west,lit=true,signal_fire=true,waterlogged=false]",
		nil,
		NewMapping{
			573: 11249,
			477: 11233,
		},
	},
	{
		"minecraft:heavy_weighted_pressure_plate[power=0]",
		nil,
		NewMapping{
			4:   2368,
			393: 5619,
			404: 5620,
			477: 6126,
			573: 6126,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=1,south=side,west=up]",
		nil,
		NewMapping{
			393: 2772,
			404: 2773,
			477: 3076,
			573: 3076,
		},
	},
	{
		"minecraft:jungle_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 5083,
			404: 5084,
			477: 5587,
			573: 5587,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=16,powered=true]",
		nil,
		NewMapping{
			573: 480,
			393: 480,
			404: 480,
			477: 480,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 5201,
			393: 4697,
			404: 4698,
			477: 5201,
		},
	},
	{
		"minecraft:dark_oak_door[facing=south,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			573: 8484,
			393: 7959,
			404: 7960,
			477: 8484,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=4,south=up,west=up]",
		nil,
		NewMapping{
			477: 2812,
			573: 2812,
			393: 2508,
			404: 2509,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10128,
			573: 10128,
		},
	},
	{
		"minecraft:air",
		nil,
		NewMapping{
			477: 0,
			573: 0,
			4:   0,
			393: 0,
			404: 0,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=west,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			573: 7048,
			4:   2674,
			393: 6541,
			404: 6542,
			477: 7048,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 6428,
			573: 6428,
			393: 5921,
			404: 5922,
		},
	},
	{
		"minecraft:fire[age=12,east=true,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			404: 1522,
			477: 1825,
			573: 1825,
			393: 1521,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=4,south=none,west=up]",
		nil,
		NewMapping{
			393: 1794,
			404: 1795,
			477: 2098,
			573: 2098,
		},
	},
	{
		"minecraft:lever[face=floor,facing=north,powered=true]",
		nil,
		NewMapping{
			404: 3278,
			477: 3781,
			573: 3781,
			4:   1117,
			393: 3277,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=10,south=none,west=side]",
		nil,
		NewMapping{
			404: 2282,
			477: 2585,
			573: 2585,
			393: 2281,
		},
	},
	{
		"minecraft:kelp[age=20]",
		nil,
		NewMapping{
			393: 8429,
			404: 8430,
			477: 8954,
			573: 8954,
		},
	},
	{
		"minecraft:piston_head[facing=south,short=true,type=normal]",
		nil,
		NewMapping{
			393: 1067,
			404: 1067,
			477: 1367,
			573: 1367,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9256,
			573: 9256,
		},
	},
	{
		"minecraft:acacia_sapling[stage=0]",
		nil,
		NewMapping{
			4:   100,
			393: 29,
			404: 29,
			477: 29,
			573: 29,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=false,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			404: 6267,
			477: 6773,
			573: 6773,
			393: 6266,
		},
	},
	{
		"minecraft:dark_oak_button[face=floor,facing=east,powered=true]",
		nil,
		NewMapping{
			404: 5430,
			477: 5936,
			573: 5936,
			393: 5429,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 7216,
			404: 7217,
			477: 7723,
			573: 7723,
		},
	},
	{
		"minecraft:brown_bed[facing=south,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 945,
			404: 945,
			477: 1245,
			573: 1245,
		},
	},
	{
		"minecraft:fire[age=14,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			404: 1585,
			477: 1888,
			573: 1888,
			393: 1584,
		},
	},
	{
		"minecraft:acacia_door[facing=south,half=upper,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7886,
			404: 7887,
			477: 8411,
			573: 8411,
		},
	},
	{
		"minecraft:mushroom_stem[down=false,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			573: 4652,
			393: 4148,
			404: 4149,
			477: 4652,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5870,
			404: 5871,
			477: 6377,
			573: 6377,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=23,powered=true]",
		nil,
		NewMapping{
			393: 744,
			404: 744,
			477: 744,
			573: 744,
		},
	},
	{
		"minecraft:smooth_sandstone_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			477: 10291,
			573: 10291,
		},
	},
	{
		"minecraft:diorite_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 11071,
			573: 11071,
		},
	},
	{
		"minecraft:purpur_stairs[facing=west,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			404: 8128,
			477: 8652,
			573: 8652,
			393: 8127,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6638,
			404: 6639,
			477: 7145,
			573: 7145,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=14,south=up,west=side]",
		nil,
		NewMapping{
			393: 2455,
			404: 2456,
			477: 2759,
			573: 2759,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=south,half=top,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3866,
			404: 3867,
			477: 4370,
			573: 4370,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9608,
			573: 9608,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 4037,
			404: 4038,
			477: 4541,
			573: 4541,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=west,half=top,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			477: 4386,
			573: 4386,
			393: 3882,
			404: 3883,
		},
	},
	{
		"minecraft:bamboo[age=0,leaves=none,stage=0]",
		nil,
		NewMapping{
			573: 9116,
			477: 9116,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=8,powered=true]",
		nil,
		NewMapping{
			404: 714,
			477: 714,
			573: 714,
			393: 714,
		},
	},
	{
		"minecraft:player_head[rotation=7]",
		nil,
		NewMapping{
			393: 5518,
			404: 5519,
			477: 6021,
			573: 6021,
		},
	},
	{
		"minecraft:jungle_wall_sign[facing=north,waterlogged=false]",
		nil,
		NewMapping{
			477: 3766,
			573: 3766,
		},
	},
	{
		"minecraft:heavy_weighted_pressure_plate[power=13]",
		nil,
		NewMapping{
			4:   2381,
			393: 5632,
			404: 5633,
			477: 6139,
			573: 6139,
		},
	},
	{
		"minecraft:green_wall_banner[facing=south]",
		nil,
		NewMapping{
			393: 7163,
			404: 7164,
			477: 7670,
			573: 7670,
		},
	},
	{
		"minecraft:acacia_door[facing=west,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			573: 8432,
			393: 7907,
			404: 7908,
			477: 8432,
		},
	},
	{
		"minecraft:dark_oak_door[facing=north,half=upper,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			477: 8458,
			573: 8458,
			393: 7933,
			404: 7934,
		},
	},
	{
		"minecraft:orange_bed[facing=south,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 768,
			404: 768,
			477: 1068,
			573: 1068,
		},
	},
	{
		"minecraft:red_banner[rotation=0]",
		nil,
		NewMapping{
			573: 7585,
			393: 7078,
			404: 7079,
			477: 7585,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9587,
			573: 9587,
		},
	},
	{
		"minecraft:cyan_banner[rotation=11]",
		nil,
		NewMapping{
			393: 7009,
			404: 7010,
			477: 7516,
			573: 7516,
		},
	},
	{
		"minecraft:purpur_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 8123,
			404: 8124,
			477: 8648,
			573: 8648,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9301,
			573: 9301,
		},
	},
	{
		"minecraft:daylight_detector[inverted=false,power=15]",
		nil,
		NewMapping{
			393: 5682,
			404: 5683,
			477: 6189,
			573: 6189,
			4:   2431,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=6,south=up,west=none]",
		nil,
		NewMapping{
			393: 2528,
			404: 2529,
			477: 2832,
			573: 2832,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=10,south=side,west=side]",
		nil,
		NewMapping{
			404: 2999,
			477: 3302,
			573: 3302,
			393: 2998,
		},
	},
	{
		"minecraft:magenta_banner[rotation=15]",
		nil,
		NewMapping{
			477: 7408,
			573: 7408,
			393: 6901,
			404: 6902,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=1,south=up,west=up]",
		nil,
		NewMapping{
			393: 2049,
			404: 2050,
			477: 2353,
			573: 2353,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9605,
			573: 9605,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9858,
			573: 9858,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=east,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9836,
			573: 9836,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=east,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9842,
			573: 9842,
		},
	},
	{
		"minecraft:white_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 6344,
			573: 6344,
			393: 5837,
			404: 5838,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=false,north=false,powered=false,south=false,west=true]",
		nil,
		NewMapping{
			393: 4849,
			404: 4850,
			477: 5353,
			573: 5353,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5173,
			404: 5174,
			477: 5677,
			573: 5677,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=4,south=none,west=up]",
		nil,
		NewMapping{
			393: 2514,
			404: 2515,
			477: 2818,
			573: 2818,
		},
	},
	{
		"minecraft:sunflower[half=upper]",
		nil,
		NewMapping{
			393: 6842,
			404: 6843,
			477: 7349,
			573: 7349,
		},
	},
	{
		"minecraft:repeating_command_block[conditional=true,facing=north]",
		nil,
		NewMapping{
			4:   3370,
			393: 8164,
			404: 8165,
			477: 8689,
			573: 8689,
		},
	},
	{
		"minecraft:jungle_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 5092,
			404: 5093,
			477: 5596,
			573: 5596,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=14,south=side,west=side]",
		nil,
		NewMapping{
			393: 2890,
			404: 2891,
			477: 3194,
			573: 3194,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=15,south=up,west=side]",
		nil,
		NewMapping{
			393: 2320,
			404: 2321,
			477: 2624,
			573: 2624,
		},
	},
	{
		"minecraft:lectern[facing=west,has_book=true,powered=true]",
		nil,
		NewMapping{
			477: 11185,
			573: 11185,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10472,
			573: 10472,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=19,powered=false]",
		nil,
		NewMapping{
			477: 887,
			573: 887,
		},
	},
	{
		"minecraft:bell[attachment=floor,facing=west,powered=true]",
		nil,
		NewMapping{
			573: 11202,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=false,north=false,powered=false,south=false,west=false]",
		nil,
		NewMapping{
			404: 4787,
			477: 5290,
			573: 5290,
			4:   2126,
			393: 4786,
		},
	},
	{
		"minecraft:skeleton_skull[rotation=0]",
		nil,
		NewMapping{
			573: 5954,
			4:   2304,
			393: 5451,
			404: 5452,
			477: 5954,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=3,powered=true]",
		nil,
		NewMapping{
			404: 404,
			477: 404,
			573: 404,
			393: 404,
		},
	},
	{
		"minecraft:chorus_plant[down=true,east=true,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 8003,
			404: 8004,
			477: 8528,
			573: 8528,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=4,south=up,west=none]",
		nil,
		NewMapping{
			477: 2526,
			573: 2526,
			393: 2222,
			404: 2223,
		},
	},
	{
		"minecraft:fire[age=10,east=true,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 1464,
			404: 1465,
			477: 1768,
			573: 1768,
		},
	},
	{
		"minecraft:jungle_fence[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 7596,
			404: 7597,
			477: 8121,
			573: 8121,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=9,south=up,west=side]",
		nil,
		NewMapping{
			393: 2266,
			404: 2267,
			477: 2570,
			573: 2570,
		},
	},
	{
		"minecraft:yellow_concrete",
		nil,
		NewMapping{
			4:   4020,
			393: 8381,
			404: 8382,
			477: 8906,
			573: 8906,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 5864,
			477: 6370,
			573: 6370,
			393: 5863,
		},
	},
	{
		"minecraft:spruce_button[face=wall,facing=north,powered=true]",
		nil,
		NewMapping{
			393: 5335,
			404: 5336,
			477: 5842,
			573: 5842,
		},
	},
	{
		"minecraft:spruce_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4949,
			404: 4950,
			477: 5453,
			573: 5453,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=12,south=none,west=up]",
		nil,
		NewMapping{
			393: 3018,
			404: 3019,
			477: 3322,
			573: 3322,
		},
	},
	{
		"minecraft:glass_pane[east=false,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 4233,
			404: 4234,
			477: 4737,
			573: 4737,
		},
	},
	{
		"minecraft:oak_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 1677,
			404: 1678,
			477: 1981,
			573: 1981,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=21,powered=true]",
		nil,
		NewMapping{
			404: 640,
			477: 640,
			573: 640,
			393: 640,
		},
	},
	{
		"minecraft:red_wall_banner[facing=north]",
		nil,
		NewMapping{
			573: 7673,
			393: 7166,
			404: 7167,
			477: 7673,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4435,
			404: 4436,
			477: 4939,
			573: 4939,
		},
	},
	{
		"minecraft:fire[age=1,east=false,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1197,
			404: 1198,
			477: 1501,
			573: 1501,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9213,
			573: 9213,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9395,
			573: 9395,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 11027,
			573: 11027,
		},
	},
	{
		"minecraft:birch_sign[rotation=0,waterlogged=false]",
		nil,
		NewMapping{
			477: 3444,
			573: 3444,
		},
	},
	{
		"minecraft:snow[layers=5]",
		nil,
		NewMapping{
			393: 3419,
			404: 3420,
			477: 3923,
			573: 3923,
			4:   1252,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=3,south=none,west=up]",
		nil,
		NewMapping{
			393: 2649,
			404: 2650,
			477: 2953,
			573: 2953,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=10,powered=false]",
		nil,
		NewMapping{
			393: 569,
			404: 569,
			477: 569,
			573: 569,
		},
	},
	{
		"minecraft:quartz_stairs[facing=east,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5763,
			404: 5764,
			477: 6270,
			573: 6270,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9327,
			573: 9327,
		},
	},
	{
		"minecraft:piston[extended=true,facing=north]",
		nil,
		NewMapping{
			393: 1047,
			404: 1047,
			477: 1347,
			573: 1347,
			4:   538,
		},
	},
	{
		"minecraft:fire[age=12,east=true,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1529,
			404: 1530,
			477: 1833,
			573: 1833,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=false,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 4141,
			404: 4142,
			477: 4645,
			573: 4645,
		},
	},
	{
		"minecraft:vine[east=true,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			573: 4777,
			393: 4273,
			404: 4274,
			477: 4777,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9497,
			573: 9497,
		},
	},
	{
		"minecraft:cactus[age=5]",
		nil,
		NewMapping{
			393: 3430,
			404: 3431,
			477: 3934,
			573: 3934,
			4:   1301,
		},
	},
	{
		"minecraft:scaffolding[bottom=true,distance=6,waterlogged=false]",
		nil,
		NewMapping{
			573: 11112,
			477: 11112,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=east,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 10073,
			573: 10073,
		},
	},
	{
		"minecraft:red_nether_bricks",
		nil,
		NewMapping{
			573: 8719,
			4:   3440,
			393: 8194,
			404: 8195,
			477: 8719,
		},
	},
	{
		"minecraft:red_bed[facing=south,occupied=false,part=head]",
		nil,
		NewMapping{
			4:   424,
			393: 978,
			404: 978,
			477: 1278,
			573: 1278,
		},
	},
	{
		"minecraft:fire[age=1,east=true,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			404: 1177,
			477: 1480,
			573: 1480,
			393: 1176,
		},
	},
	{
		"minecraft:acacia_door[facing=north,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			404: 7880,
			477: 8404,
			573: 8404,
			393: 7879,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=11,south=up,west=side]",
		nil,
		NewMapping{
			393: 2428,
			404: 2429,
			477: 2732,
			573: 2732,
		},
	},
	{
		"minecraft:oak_wood[axis=x]",
		nil,
		NewMapping{
			393: 108,
			404: 108,
			477: 108,
			573: 108,
		},
	},
	{
		"minecraft:diorite_wall[east=false,north=true,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 11075,
			573: 11075,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=4,powered=true]",
		nil,
		NewMapping{
			573: 906,
			477: 906,
		},
	},
	{
		"minecraft:repeater[delay=4,facing=south,locked=false,powered=true]",
		nil,
		NewMapping{
			4:   1516,
			393: 3567,
			404: 3568,
			477: 4071,
			573: 4071,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=9,south=up,west=up]",
		nil,
		NewMapping{
			573: 2857,
			393: 2553,
			404: 2554,
			477: 2857,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=5,powered=false]",
		nil,
		NewMapping{
			393: 609,
			404: 609,
			477: 609,
			573: 609,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 3225,
			404: 3226,
			477: 3689,
			573: 3689,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=11,south=none,west=up]",
		nil,
		NewMapping{
			477: 2449,
			573: 2449,
			393: 2145,
			404: 2146,
		},
	},
	{
		"minecraft:red_shulker_box[facing=up]",
		nil,
		NewMapping{
			4:   3729,
			393: 8305,
			404: 8306,
			477: 8830,
			573: 8830,
		},
	},
	{
		"minecraft:dead_tube_coral_wall_fan[facing=south,waterlogged=false]",
		nil,
		NewMapping{
			393: 8467,
			404: 8483,
			477: 9027,
			573: 9027,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6639,
			404: 6640,
			477: 7146,
			573: 7146,
		},
	},
	{
		"minecraft:birch_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 5525,
			393: 5021,
			404: 5022,
			477: 5525,
		},
	},
	{
		"minecraft:granite_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			477: 10302,
			573: 10302,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9752,
			573: 9752,
		},
	},
	{
		"minecraft:diorite_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 10213,
			573: 10213,
		},
	},
	{
		"minecraft:stone_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9645,
			573: 9645,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=15,powered=true]",
		nil,
		NewMapping{
			477: 278,
			573: 278,
			393: 278,
			404: 278,
		},
	},
	{
		"minecraft:oak_door[facing=west,half=upper,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			393: 3146,
			404: 3147,
			477: 3610,
			573: 3610,
		},
	},
	{
		"minecraft:piston[extended=false,facing=north]",
		nil,
		NewMapping{
			4:   530,
			393: 1053,
			404: 1053,
			477: 1353,
			573: 1353,
		},
	},
	{
		"minecraft:spruce_door[facing=south,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			4:   3089,
			393: 7708,
			404: 7709,
			477: 8233,
			573: 8233,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   1077,
			393: 3230,
			404: 3231,
			477: 3694,
			573: 3694,
		},
	},
	{
		"minecraft:wither_skeleton_skull[rotation=15]",
		nil,
		NewMapping{
			393: 5486,
			404: 5487,
			477: 5989,
			573: 5989,
		},
	},
	{
		"minecraft:fire[age=12,east=true,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1533,
			404: 1534,
			477: 1837,
			573: 1837,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=3,south=none,west=up]",
		nil,
		NewMapping{
			477: 2809,
			573: 2809,
			393: 2505,
			404: 2506,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6415,
			404: 6416,
			477: 6922,
			573: 6922,
		},
	},
	{
		"minecraft:black_banner[rotation=3]",
		nil,
		NewMapping{
			393: 7097,
			404: 7098,
			477: 7604,
			573: 7604,
		},
	},
	{
		"minecraft:quartz_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			4:   719,
			393: 7336,
			404: 7337,
			477: 7855,
			573: 7855,
		},
	},
	{
		"minecraft:red_terracotta",
		nil,
		NewMapping{
			573: 6325,
			4:   2558,
			393: 5818,
			404: 5819,
			477: 6325,
		},
	},
	{
		"minecraft:structure_block[mode=save]",
		nil,
		NewMapping{
			404: 8595,
			477: 11252,
			573: 11268,
			4:   4080,
			393: 8578,
		},
	},
	{
		"minecraft:light_gray_banner[rotation=15]",
		nil,
		NewMapping{
			393: 6997,
			404: 6998,
			477: 7504,
			573: 7504,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6008,
			404: 6009,
			477: 6515,
			573: 6515,
		},
	},
	{
		"minecraft:red_bed[facing=east,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 984,
			404: 984,
			477: 1284,
			573: 1284,
			4:   431,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5186,
			404: 5187,
			477: 5690,
			573: 5690,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=15,south=side,west=up]",
		nil,
		NewMapping{
			393: 2898,
			404: 2899,
			477: 3202,
			573: 3202,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10069,
			573: 10069,
		},
	},
	{
		"minecraft:cyan_glazed_terracotta[facing=south]",
		nil,
		NewMapping{
			477: 8875,
			573: 8875,
			4:   3904,
			393: 8350,
			404: 8351,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=17,powered=true]",
		nil,
		NewMapping{
			404: 732,
			477: 732,
			573: 732,
			393: 732,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=2,south=up,west=none]",
		nil,
		NewMapping{
			573: 2796,
			393: 2492,
			404: 2493,
			477: 2796,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10896,
			573: 10896,
		},
	},
	{
		"minecraft:barrel[facing=east,open=true]",
		nil,
		NewMapping{
			477: 11137,
			573: 11137,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=false,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 4003,
			404: 4004,
			477: 4507,
			573: 4507,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=south,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			573: 7986,
			393: 7461,
			404: 7462,
			477: 7986,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=north,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			477: 4293,
			573: 4293,
			393: 3789,
			404: 3790,
		},
	},
	{
		"minecraft:birch_door[facing=north,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			477: 8270,
			573: 8270,
			393: 7745,
			404: 7746,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9870,
			477: 9870,
		},
	},
	{
		"minecraft:magenta_glazed_terracotta[facing=west]",
		nil,
		NewMapping{
			573: 8848,
			4:   3793,
			393: 8323,
			404: 8324,
			477: 8848,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=20,powered=true]",
		nil,
		NewMapping{
			477: 838,
			573: 838,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10503,
			573: 10503,
		},
	},
	{
		"minecraft:cyan_banner[rotation=6]",
		nil,
		NewMapping{
			393: 7004,
			404: 7005,
			477: 7511,
			573: 7511,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			404: 5224,
			477: 5727,
			573: 5727,
			393: 5223,
		},
	},
	{
		"minecraft:acacia_button[face=ceiling,facing=north,powered=true]",
		nil,
		NewMapping{
			393: 5415,
			404: 5416,
			477: 5922,
			573: 5922,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=12,powered=true]",
		nil,
		NewMapping{
			393: 422,
			404: 422,
			477: 422,
			573: 422,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=0,powered=true]",
		nil,
		NewMapping{
			393: 248,
			404: 248,
			477: 248,
			573: 248,
		},
	},
	{
		"minecraft:birch_button[face=wall,facing=east,powered=false]",
		nil,
		NewMapping{
			477: 5873,
			573: 5873,
			393: 5366,
			404: 5367,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=5,south=up,west=up]",
		nil,
		NewMapping{
			404: 2662,
			477: 2965,
			573: 2965,
			393: 2661,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=east,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4713,
			404: 4714,
			477: 5217,
			573: 5217,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=12,powered=false]",
		nil,
		NewMapping{
			573: 973,
			477: 973,
		},
	},
	{
		"minecraft:stone_button[face=ceiling,facing=north,powered=false]",
		nil,
		NewMapping{
			4:   1232,
			393: 3408,
			404: 3409,
			477: 3912,
			573: 3912,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=north,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			573: 4169,
			393: 3665,
			404: 3666,
			477: 4169,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=4,south=up,west=none]",
		nil,
		NewMapping{
			393: 1934,
			404: 1935,
			477: 2238,
			573: 2238,
		},
	},
	{
		"minecraft:fire[age=5,east=false,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1325,
			404: 1326,
			477: 1629,
			573: 1629,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9799,
			573: 9799,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 6124,
			477: 6630,
			573: 6630,
			393: 6123,
		},
	},
	{
		"minecraft:iron_door[facing=north,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 3315,
			404: 3316,
			477: 3819,
			573: 3819,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=10,south=none,west=none]",
		nil,
		NewMapping{
			393: 3002,
			404: 3003,
			477: 3306,
			573: 3306,
			4:   890,
		},
	},
	{
		"minecraft:rail[shape=south_east]",
		nil,
		NewMapping{
			393: 3185,
			404: 3186,
			477: 3649,
			573: 3649,
			4:   1062,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=north,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			477: 4365,
			573: 4365,
			393: 3861,
			404: 3862,
		},
	},
	{
		"minecraft:fire[age=6,east=true,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			404: 1338,
			477: 1641,
			573: 1641,
			393: 1337,
		},
	},
	{
		"minecraft:fire[age=0,east=false,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1161,
			404: 1162,
			477: 1465,
			573: 1465,
		},
	},
	{
		"minecraft:purple_terracotta",
		nil,
		NewMapping{
			4:   2554,
			393: 5814,
			404: 5815,
			477: 6321,
			573: 6321,
		},
	},
	{
		"minecraft:fire[age=11,east=false,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1503,
			404: 1504,
			477: 1807,
			573: 1807,
		},
	},
	{
		"minecraft:blue_banner[rotation=11]",
		nil,
		NewMapping{
			477: 7548,
			573: 7548,
			393: 7041,
			404: 7042,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=false,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 6709,
			393: 6202,
			404: 6203,
			477: 6709,
		},
	},
	{
		"minecraft:polished_diorite_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			477: 10271,
			573: 10271,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10080,
			573: 10080,
		},
	},
	{
		"minecraft:ender_chest[facing=north,waterlogged=true]",
		nil,
		NewMapping{
			393: 4731,
			404: 4732,
			477: 5235,
			573: 5235,
		},
	},
	{
		"minecraft:jungle_leaves[distance=7,persistent=false]",
		nil,
		NewMapping{
			393: 199,
			404: 199,
			477: 199,
			573: 199,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=6,south=none,west=side]",
		nil,
		NewMapping{
			404: 2246,
			477: 2549,
			573: 2549,
			393: 2245,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=12,south=up,west=none]",
		nil,
		NewMapping{
			573: 2454,
			393: 2150,
			404: 2151,
			477: 2454,
		},
	},
	{
		"minecraft:fire[age=15,east=true,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1623,
			404: 1624,
			477: 1927,
			573: 1927,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=3,south=side,west=up]",
		nil,
		NewMapping{
			393: 1782,
			404: 1783,
			477: 2086,
			573: 2086,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10575,
			573: 10575,
		},
	},
	{
		"minecraft:lava[level=14]",
		nil,
		NewMapping{
			393: 64,
			404: 64,
			477: 64,
			573: 64,
			4:   174,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=east,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3705,
			404: 3706,
			477: 4209,
			573: 4209,
		},
	},
	{
		"minecraft:mushroom_stem[down=false,east=true,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			573: 4661,
			393: 4157,
			404: 4158,
			477: 4661,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5209,
			404: 5210,
			477: 5713,
			573: 5713,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 6672,
			404: 6673,
			477: 7179,
			573: 7179,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=west,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 6462,
			404: 6463,
			477: 6969,
			573: 6969,
		},
	},
	{
		"minecraft:seagrass",
		nil,
		NewMapping{
			393: 1044,
			404: 1044,
			477: 1344,
			573: 1344,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=15,waterlogged=false]",
		nil,
		NewMapping{
			477: 3570,
			573: 3570,
		},
	},
	{
		"minecraft:birch_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 5003,
			404: 5004,
			477: 5507,
			573: 5507,
		},
	},
	{
		"minecraft:dead_tube_coral_block",
		nil,
		NewMapping{
			393: 8449,
			404: 8450,
			477: 8974,
			573: 8974,
		},
	},
	{
		"minecraft:jungle_sign[rotation=12,waterlogged=true]",
		nil,
		NewMapping{
			477: 3531,
			573: 3531,
		},
	},
	{
		"minecraft:light_gray_glazed_terracotta[facing=west]",
		nil,
		NewMapping{
			573: 8872,
			4:   3889,
			393: 8347,
			404: 8348,
			477: 8872,
		},
	},
	{
		"minecraft:spruce_door[facing=south,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7699,
			404: 7700,
			477: 8224,
			573: 8224,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=8,south=side,west=up]",
		nil,
		NewMapping{
			573: 3139,
			393: 2835,
			404: 2836,
			477: 3139,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=8,south=none,west=none]",
		nil,
		NewMapping{
			404: 2841,
			477: 3144,
			573: 3144,
			393: 2840,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=false,north=true,powered=false,south=true,west=false]",
		nil,
		NewMapping{
			477: 5312,
			573: 5312,
			393: 4808,
			404: 4809,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=north,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9215,
			573: 9215,
		},
	},
	{
		"minecraft:heavy_weighted_pressure_plate[power=14]",
		nil,
		NewMapping{
			393: 5633,
			404: 5634,
			477: 6140,
			573: 6140,
			4:   2382,
		},
	},
	{
		"minecraft:yellow_wall_banner[facing=west]",
		nil,
		NewMapping{
			393: 7128,
			404: 7129,
			477: 7635,
			573: 7635,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 6405,
			393: 5898,
			404: 5899,
			477: 6405,
		},
	},
	{
		"minecraft:birch_button[face=ceiling,facing=west,powered=true]",
		nil,
		NewMapping{
			404: 5372,
			477: 5878,
			573: 5878,
			393: 5371,
		},
	},
	{
		"minecraft:acacia_door[facing=north,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			477: 8406,
			573: 8406,
			393: 7881,
			404: 7882,
		},
	},
	{
		"minecraft:light_gray_bed[facing=west,occupied=true,part=foot]",
		nil,
		NewMapping{
			477: 1185,
			573: 1185,
			393: 885,
			404: 885,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9214,
			573: 9214,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 10120,
			477: 10120,
		},
	},
	{
		"minecraft:blue_bed[facing=west,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 933,
			404: 933,
			477: 1233,
			573: 1233,
		},
	},
	{
		"minecraft:light_blue_bed[facing=east,occupied=false,part=foot]",
		nil,
		NewMapping{
			477: 1111,
			573: 1111,
			393: 811,
			404: 811,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=east,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3641,
			404: 3642,
			477: 4145,
			573: 4145,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=west,in_wall=true,open=false,powered=true]",
		nil,
		NewMapping{
			573: 7900,
			393: 7375,
			404: 7376,
			477: 7900,
		},
	},
	{
		"minecraft:fire[age=15,east=true,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 1626,
			404: 1627,
			477: 1930,
			573: 1930,
		},
	},
	{
		"minecraft:birch_fence[east=true,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 7553,
			404: 7554,
			477: 8078,
			573: 8078,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=east,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			573: 4211,
			393: 3707,
			404: 3708,
			477: 4211,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6159,
			404: 6160,
			477: 6666,
			573: 6666,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9359,
			573: 9359,
		},
	},
	{
		"minecraft:purple_shulker_box[facing=south]",
		nil,
		NewMapping{
			573: 8804,
			4:   3667,
			393: 8279,
			404: 8280,
			477: 8804,
		},
	},
	{
		"minecraft:dispenser[facing=west,triggered=true]",
		nil,
		NewMapping{
			573: 239,
			4:   380,
			393: 239,
			404: 239,
			477: 239,
		},
	},
	{
		"minecraft:potted_jungle_sapling",
		nil,
		NewMapping{
			393: 5269,
			404: 5270,
			477: 5773,
			573: 5773,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=15,powered=false]",
		nil,
		NewMapping{
			477: 379,
			573: 379,
			393: 379,
			404: 379,
		},
	},
	{
		"minecraft:mushroom_stem[down=false,east=false,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			573: 4677,
			393: 4173,
			404: 4174,
			477: 4677,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=19,powered=false]",
		nil,
		NewMapping{
			477: 937,
			573: 937,
		},
	},
	{
		"minecraft:bamboo[age=1,leaves=none,stage=1]",
		nil,
		NewMapping{
			573: 9123,
			477: 9123,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=19,powered=true]",
		nil,
		NewMapping{
			477: 1036,
			573: 1036,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=10,south=none,west=side]",
		nil,
		NewMapping{
			393: 2713,
			404: 2714,
			477: 3017,
			573: 3017,
		},
	},
	{
		"minecraft:gray_carpet",
		nil,
		NewMapping{
			4:   2743,
			393: 6830,
			404: 6831,
			477: 7337,
			573: 7337,
		},
	},
	{
		"minecraft:water[level=15]",
		nil,
		NewMapping{
			404: 49,
			477: 49,
			573: 49,
			4:   143,
			393: 49,
		},
	},
	{
		"minecraft:sea_pickle[pickles=2,waterlogged=false]",
		nil,
		NewMapping{
			393: 8567,
			404: 8583,
			477: 9107,
			573: 9107,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4727,
			404: 4728,
			477: 5231,
			573: 5231,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=7,powered=true]",
		nil,
		NewMapping{
			393: 612,
			404: 612,
			477: 612,
			573: 612,
		},
	},
	{
		"minecraft:birch_door[facing=south,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7759,
			404: 7760,
			477: 8284,
			573: 8284,
		},
	},
	{
		"minecraft:dark_oak_door[facing=west,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			573: 8496,
			393: 7971,
			404: 7972,
			477: 8496,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6143,
			404: 6144,
			477: 6650,
			573: 6650,
		},
	},
	{
		"minecraft:fire[age=2,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1230,
			404: 1231,
			477: 1534,
			573: 1534,
			4:   818,
		},
	},
	{
		"minecraft:birch_leaves[distance=6,persistent=false]",
		nil,
		NewMapping{
			477: 183,
			573: 183,
			393: 183,
			404: 183,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=14,south=none,west=up]",
		nil,
		NewMapping{
			393: 2604,
			404: 2605,
			477: 2908,
			573: 2908,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=10,powered=false]",
		nil,
		NewMapping{
			393: 319,
			404: 319,
			477: 319,
			573: 319,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10705,
			573: 10705,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5948,
			404: 5949,
			477: 6455,
			573: 6455,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10670,
			573: 10670,
		},
	},
	{
		"minecraft:bell[attachment=double_wall,facing=east,powered=false]",
		nil,
		NewMapping{
			573: 11229,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   1073,
			393: 3240,
			404: 3241,
			477: 3704,
			573: 3704,
		},
	},
	{
		"minecraft:command_block[conditional=false,facing=down]",
		nil,
		NewMapping{
			477: 5639,
			573: 5639,
			4:   2192,
			393: 5135,
			404: 5136,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=12,south=none,west=none]",
		nil,
		NewMapping{
			573: 3180,
			393: 2876,
			404: 2877,
			477: 3180,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4456,
			404: 4457,
			477: 4960,
			573: 4960,
		},
	},
	{
		"minecraft:oak_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 1670,
			404: 1671,
			477: 1974,
			573: 1974,
		},
	},
	{
		"minecraft:white_banner[rotation=11]",
		nil,
		NewMapping{
			4:   2827,
			393: 6865,
			404: 6866,
			477: 7372,
			573: 7372,
		},
	},
	{
		"minecraft:purpur_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			404: 8142,
			477: 8666,
			573: 8666,
			393: 8141,
		},
	},
	{
		"minecraft:light_blue_bed[facing=north,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 797,
			404: 797,
			477: 1097,
			573: 1097,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=19,powered=false]",
		nil,
		NewMapping{
			393: 287,
			404: 287,
			477: 287,
			573: 287,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=3,south=side,west=side]",
		nil,
		NewMapping{
			393: 2503,
			404: 2504,
			477: 2807,
			573: 2807,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9817,
			573: 9817,
		},
	},
	{
		"minecraft:bell[attachment=single_wall,facing=south]",
		nil,
		NewMapping{
			477: 11207,
		},
	},
	{
		"minecraft:acacia_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6391,
			404: 6392,
			477: 6898,
			573: 6898,
		},
	},
	{
		"minecraft:oak_door[facing=east,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			573: 3629,
			393: 3165,
			404: 3166,
			477: 3629,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6288,
			404: 6289,
			477: 6795,
			573: 6795,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=22,powered=true]",
		nil,
		NewMapping{
			393: 742,
			404: 742,
			477: 742,
			573: 742,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 11035,
			477: 11035,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10678,
			573: 10678,
		},
	},
	{
		"minecraft:lever[face=floor,facing=north,powered=false]",
		nil,
		NewMapping{
			573: 3782,
			4:   1109,
			393: 3278,
			404: 3279,
			477: 3782,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=6,south=up,west=side]",
		nil,
		NewMapping{
			393: 2815,
			404: 2816,
			477: 3119,
			573: 3119,
		},
	},
	{
		"minecraft:acacia_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6336,
			404: 6337,
			477: 6843,
			573: 6843,
		},
	},
	{
		"minecraft:lever[face=ceiling,facing=south,powered=false]",
		nil,
		NewMapping{
			393: 3296,
			404: 3297,
			477: 3800,
			573: 3800,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=2,south=up,west=up]",
		nil,
		NewMapping{
			393: 2202,
			404: 2203,
			477: 2506,
			573: 2506,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=east,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3965,
			404: 3966,
			477: 4469,
			573: 4469,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9792,
			573: 9792,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9786,
			477: 9786,
		},
	},
	{
		"minecraft:fire[age=14,east=false,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1609,
			404: 1610,
			477: 1913,
			573: 1913,
		},
	},
	{
		"minecraft:oak_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			404: 1723,
			477: 2026,
			573: 2026,
			393: 1722,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 4922,
			393: 4418,
			404: 4419,
			477: 4922,
		},
	},
	{
		"minecraft:granite_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 9929,
			477: 9929,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=false,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10970,
			573: 10970,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9210,
			573: 9210,
		},
	},
	{
		"minecraft:spruce_leaves[distance=5,persistent=true]",
		nil,
		NewMapping{
			393: 166,
			404: 166,
			477: 166,
			573: 166,
		},
	},
	{
		"minecraft:jungle_door[facing=north,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7810,
			404: 7811,
			477: 8335,
			573: 8335,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=east,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			573: 4221,
			393: 3717,
			404: 3718,
			477: 4221,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=5,powered=true]",
		nil,
		NewMapping{
			393: 558,
			404: 558,
			477: 558,
			573: 558,
		},
	},
	{
		"minecraft:black_bed[facing=south,occupied=false,part=head]",
		nil,
		NewMapping{
			573: 1294,
			393: 994,
			404: 994,
			477: 1294,
		},
	},
	{
		"minecraft:jungle_log[axis=y]",
		nil,
		NewMapping{
			477: 82,
			573: 82,
			4:   275,
			393: 82,
			404: 82,
		},
	},
	{
		"minecraft:acacia_fence[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 7613,
			404: 7614,
			477: 8138,
			573: 8138,
		},
	},
	{
		"minecraft:spruce_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4959,
			404: 4960,
			477: 5463,
			573: 5463,
		},
	},
	{
		"minecraft:brewing_stand[has_bottle_0=true,has_bottle_1=false,has_bottle_2=true]",
		nil,
		NewMapping{
			4:   1877,
			393: 4615,
			404: 4616,
			477: 5119,
			573: 5119,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=9,south=side,west=up]",
		nil,
		NewMapping{
			393: 1836,
			404: 1837,
			477: 2140,
			573: 2140,
		},
	},
	{
		"minecraft:fire[age=9,east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			477: 1755,
			573: 1755,
			393: 1451,
			404: 1452,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10817,
			573: 10817,
		},
	},
	{
		"minecraft:brick_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 4880,
			393: 4376,
			404: 4377,
			477: 4880,
		},
	},
	{
		"minecraft:jungle_stairs[facing=east,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 5610,
			393: 5106,
			404: 5107,
			477: 5610,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10943,
			573: 10943,
		},
	},
	{
		"minecraft:granite_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9904,
			573: 9904,
		},
	},
	{
		"minecraft:brick_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4379,
			404: 4380,
			477: 4883,
			573: 4883,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=11,south=side,west=side]",
		nil,
		NewMapping{
			393: 1999,
			404: 2000,
			477: 2303,
			573: 2303,
		},
	},
	{
		"minecraft:oak_sign[rotation=6,waterlogged=false]",
		nil,
		NewMapping{
			404: 3089,
			477: 3392,
			573: 3392,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9530,
			477: 9530,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9438,
			573: 9438,
		},
	},
	{
		"minecraft:spruce_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			404: 4939,
			477: 5442,
			573: 5442,
			393: 4938,
		},
	},
	{
		"minecraft:purpur_stairs[facing=east,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 8137,
			404: 8138,
			477: 8662,
			573: 8662,
		},
	},
	{
		"minecraft:grindstone[face=floor,facing=north]",
		nil,
		NewMapping{
			477: 11165,
			573: 11165,
		},
	},
	{
		"minecraft:chest[facing=west,type=left,waterlogged=false]",
		nil,
		NewMapping{
			404: 1744,
			477: 2047,
			573: 2047,
			393: 1743,
		},
	},
	{
		"minecraft:jungle_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5077,
			404: 5078,
			477: 5581,
			573: 5581,
		},
	},
	{
		"minecraft:iron_door[facing=west,half=lower,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 3343,
			404: 3344,
			477: 3847,
			573: 3847,
		},
	},
	{
		"minecraft:lime_banner[rotation=13]",
		nil,
		NewMapping{
			393: 6947,
			404: 6948,
			477: 7454,
			573: 7454,
		},
	},
	{
		"minecraft:acacia_fence[east=false,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 8160,
			393: 7635,
			404: 7636,
			477: 8160,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9531,
			573: 9531,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=true,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10886,
			573: 10886,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=22,powered=true]",
		nil,
		NewMapping{
			393: 392,
			404: 392,
			477: 392,
			573: 392,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10367,
			573: 10367,
		},
	},
	{
		"minecraft:diorite_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 11083,
			573: 11083,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=east,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9596,
			573: 9596,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 6557,
			393: 6050,
			404: 6051,
			477: 6557,
		},
	},
	{
		"minecraft:repeater[delay=3,facing=west,locked=true,powered=true]",
		nil,
		NewMapping{
			573: 4057,
			393: 3553,
			404: 3554,
			477: 4057,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			573: 7218,
			393: 6711,
			404: 6712,
			477: 7218,
		},
	},
	{
		"minecraft:dead_horn_coral_fan[waterlogged=true]",
		nil,
		NewMapping{
			573: 9012,
			393: 8552,
			404: 8568,
			477: 9012,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6421,
			404: 6422,
			477: 6928,
			573: 6928,
		},
	},
	{
		"minecraft:furnace[facing=north,lit=false]",
		nil,
		NewMapping{
			573: 3372,
			4:   978,
			393: 3068,
			404: 3069,
			477: 3372,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=6,powered=true]",
		nil,
		NewMapping{
			393: 610,
			404: 610,
			477: 610,
			573: 610,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=1,south=side,west=up]",
		nil,
		NewMapping{
			393: 2628,
			404: 2629,
			477: 2932,
			573: 2932,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=south,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6666,
			404: 6667,
			477: 7173,
			573: 7173,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=11,powered=true]",
		nil,
		NewMapping{
			477: 1020,
			573: 1020,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9482,
			573: 9482,
		},
	},
	{
		"minecraft:lime_glazed_terracotta[facing=west]",
		nil,
		NewMapping{
			573: 8860,
			4:   3841,
			393: 8335,
			404: 8336,
			477: 8860,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=2,south=up,west=none]",
		nil,
		NewMapping{
			573: 2508,
			393: 2204,
			404: 2205,
			477: 2508,
		},
	},
	{
		"minecraft:acacia_door[facing=south,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			404: 7896,
			477: 8420,
			573: 8420,
			393: 7895,
		},
	},
	{
		"minecraft:repeater[delay=2,facing=east,locked=true,powered=false]",
		nil,
		NewMapping{
			393: 3542,
			404: 3543,
			477: 4046,
			573: 4046,
		},
	},
	{
		"minecraft:brick_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4371,
			404: 4372,
			477: 4875,
			573: 4875,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=north,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			573: 4234,
			393: 3730,
			404: 3731,
			477: 4234,
		},
	},
	{
		"minecraft:diorite_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10199,
			573: 10199,
		},
	},
	{
		"minecraft:campfire[facing=west,lit=false,signal_fire=false,waterlogged=false]",
		nil,
		NewMapping{
			573: 11255,
			477: 11239,
		},
	},
	{
		"minecraft:red_bed[facing=south,occupied=false,part=foot]",
		nil,
		NewMapping{
			4:   416,
			393: 979,
			404: 979,
			477: 1279,
			573: 1279,
		},
	},
	{
		"minecraft:acacia_stairs[facing=south,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6357,
			404: 6358,
			477: 6864,
			573: 6864,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=north,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			573: 4353,
			393: 3849,
			404: 3850,
			477: 4353,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=9,powered=true]",
		nil,
		NewMapping{
			573: 916,
			477: 916,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9396,
			573: 9396,
		},
	},
	{
		"minecraft:spruce_log[axis=y]",
		nil,
		NewMapping{
			4:   273,
			393: 76,
			404: 76,
			477: 76,
			573: 76,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=north,in_wall=true,open=false,powered=true]",
		nil,
		NewMapping{
			573: 7980,
			393: 7455,
			404: 7456,
			477: 7980,
		},
	},
	{
		"minecraft:dark_oak_fence[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 8185,
			393: 7660,
			404: 7661,
			477: 8185,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=5,powered=false]",
		nil,
		NewMapping{
			573: 809,
			477: 809,
		},
	},
	{
		"minecraft:grindstone[face=ceiling,facing=south]",
		nil,
		NewMapping{
			477: 11174,
			573: 11174,
		},
	},
	{
		"minecraft:stone_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9622,
			573: 9622,
		},
	},
	{
		"minecraft:campfire[facing=east,lit=true,signal_fire=true,waterlogged=false]",
		nil,
		NewMapping{
			477: 11241,
			573: 11257,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9534,
			573: 9534,
		},
	},
	{
		"minecraft:oak_fence[east=false,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			4:   1360,
			393: 3491,
			404: 3492,
			477: 3995,
			573: 3995,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			477: 4509,
			573: 4509,
			393: 4005,
			404: 4006,
		},
	},
	{
		"minecraft:birch_door[facing=north,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7747,
			404: 7748,
			477: 8272,
			573: 8272,
		},
	},
	{
		"minecraft:oak_fence[east=false,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 3983,
			393: 3479,
			404: 3480,
			477: 3983,
		},
	},
	{
		"minecraft:oak_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 1667,
			404: 1668,
			477: 1971,
			573: 1971,
		},
	},
	{
		"minecraft:oak_door[facing=east,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			404: 3168,
			477: 3631,
			573: 3631,
			393: 3167,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10408,
			573: 10408,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=north,in_wall=false,open=true,powered=true]",
		nil,
		NewMapping{
			4:   2974,
			393: 7425,
			404: 7426,
			477: 7950,
			573: 7950,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=true,north=false,powered=true,south=false,west=true]",
		nil,
		NewMapping{
			393: 4829,
			404: 4830,
			477: 5333,
			573: 5333,
		},
	},
	{
		"minecraft:acacia_fence[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 7633,
			404: 7634,
			477: 8158,
			573: 8158,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=8,south=side,west=up]",
		nil,
		NewMapping{
			393: 2547,
			404: 2548,
			477: 2851,
			573: 2851,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=20,powered=true]",
		nil,
		NewMapping{
			404: 288,
			477: 288,
			573: 288,
			393: 288,
		},
	},
	{
		"minecraft:daylight_detector[inverted=false,power=13]",
		nil,
		NewMapping{
			393: 5680,
			404: 5681,
			477: 6187,
			573: 6187,
			4:   2429,
		},
	},
	{
		"minecraft:light_blue_banner[rotation=4]",
		nil,
		NewMapping{
			477: 7413,
			573: 7413,
			393: 6906,
			404: 6907,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=south,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3865,
			404: 3866,
			477: 4369,
			573: 4369,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 6176,
			477: 6682,
			573: 6682,
			393: 6175,
		},
	},
	{
		"minecraft:fire[age=15,east=true,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1622,
			404: 1623,
			477: 1926,
			573: 1926,
		},
	},
	{
		"minecraft:cyan_shulker_box[facing=west]",
		nil,
		NewMapping{
			4:   3652,
			393: 8274,
			404: 8275,
			477: 8799,
			573: 8799,
		},
	},
	{
		"minecraft:potted_red_mushroom",
		nil,
		NewMapping{
			393: 5283,
			404: 5284,
			477: 5790,
			573: 5790,
		},
	},
	{
		"minecraft:acacia_fence[east=true,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 8146,
			573: 8146,
			393: 7621,
			404: 7622,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10334,
			573: 10334,
		},
	},
	{
		"minecraft:birch_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 5500,
			573: 5500,
			393: 4996,
			404: 4997,
		},
	},
	{
		"minecraft:oak_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 1684,
			477: 1987,
			573: 1987,
			393: 1683,
		},
	},
	{
		"minecraft:fire[age=13,east=true,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1566,
			404: 1567,
			477: 1870,
			573: 1870,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10146,
			573: 10146,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=true,north=false,powered=true,south=false,west=true]",
		nil,
		NewMapping{
			404: 4798,
			477: 5301,
			573: 5301,
			393: 4797,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=0,south=up,west=none]",
		nil,
		NewMapping{
			477: 2634,
			573: 2634,
			393: 2330,
			404: 2331,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=false,north=true,powered=true,south=true,west=false]",
		nil,
		NewMapping{
			477: 5340,
			573: 5340,
			393: 4836,
			404: 4837,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=1,powered=false]",
		nil,
		NewMapping{
			573: 951,
			477: 951,
		},
	},
	{
		"minecraft:stone_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9670,
			477: 9670,
		},
	},
	{
		"minecraft:oak_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 1682,
			477: 1985,
			573: 1985,
			393: 1681,
		},
	},
	{
		"minecraft:stone_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9667,
			573: 9667,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9390,
			573: 9390,
		},
	},
	{
		"minecraft:melon",
		nil,
		NewMapping{
			4:   1648,
			393: 4243,
			404: 4244,
			477: 4747,
			573: 4747,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=west,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3885,
			404: 3886,
			477: 4389,
			573: 4389,
		},
	},
	{
		"minecraft:oak_door[facing=north,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			404: 3110,
			477: 3573,
			573: 3573,
			393: 3109,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=15,south=none,west=none]",
		nil,
		NewMapping{
			573: 2343,
			393: 2039,
			404: 2040,
			477: 2343,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=false,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 6798,
			573: 6798,
			393: 6291,
			404: 6292,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9203,
			573: 9203,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4533,
			404: 4534,
			477: 5037,
			573: 5037,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6327,
			404: 6328,
			477: 6834,
			573: 6834,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=north,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			477: 4236,
			573: 4236,
			393: 3732,
			404: 3733,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=9,south=none,west=up]",
		nil,
		NewMapping{
			393: 2271,
			404: 2272,
			477: 2575,
			573: 2575,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=2,south=none,west=up]",
		nil,
		NewMapping{
			477: 2944,
			573: 2944,
			393: 2640,
			404: 2641,
		},
	},
	{
		"minecraft:gray_shulker_box[facing=east]",
		nil,
		NewMapping{
			404: 8261,
			477: 8785,
			573: 8785,
			4:   3621,
			393: 8260,
		},
	},
	{
		"minecraft:stone_button[face=floor,facing=north,powered=true]",
		nil,
		NewMapping{
			4:   1245,
			393: 3391,
			404: 3392,
			477: 3895,
			573: 3895,
		},
	},
	{
		"minecraft:brick_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			404: 4391,
			477: 4894,
			573: 4894,
			393: 4390,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=4,south=side,west=up]",
		nil,
		NewMapping{
			477: 3103,
			573: 3103,
			393: 2799,
			404: 2800,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9738,
			573: 9738,
		},
	},
	{
		"minecraft:iron_door[facing=north,half=lower,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 3311,
			404: 3312,
			477: 3815,
			573: 3815,
		},
	},
	{
		"minecraft:quartz_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 5736,
			404: 5737,
			477: 6243,
			573: 6243,
		},
	},
	{
		"minecraft:birch_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			393: 7271,
			404: 7272,
			477: 7778,
			573: 7778,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9391,
			573: 9391,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10931,
			573: 10931,
		},
	},
	{
		"minecraft:smoker[facing=north,lit=true]",
		nil,
		NewMapping{
			477: 11147,
			573: 11147,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=8,south=side,west=none]",
		nil,
		NewMapping{
			393: 2837,
			404: 2838,
			477: 3141,
			573: 3141,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5862,
			404: 5863,
			477: 6369,
			573: 6369,
		},
	},
	{
		"minecraft:orange_bed[facing=west,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 774,
			404: 774,
			477: 1074,
			573: 1074,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=15,south=side,west=none]",
		nil,
		NewMapping{
			404: 1893,
			477: 2196,
			573: 2196,
			393: 1892,
		},
	},
	{
		"minecraft:dark_oak_leaves[distance=4,persistent=false]",
		nil,
		NewMapping{
			573: 221,
			393: 221,
			404: 221,
			477: 221,
		},
	},
	{
		"minecraft:sugar_cane[age=15]",
		nil,
		NewMapping{
			4:   1343,
			393: 3457,
			404: 3458,
			477: 3961,
			573: 3961,
		},
	},
	{
		"minecraft:oak_fence[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 3467,
			404: 3468,
			477: 3971,
			573: 3971,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 3721,
			573: 3721,
			393: 3257,
			404: 3258,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=9,south=side,west=none]",
		nil,
		NewMapping{
			573: 2718,
			393: 2414,
			404: 2415,
			477: 2718,
		},
	},
	{
		"minecraft:spruce_sign[rotation=12,waterlogged=false]",
		nil,
		NewMapping{
			477: 3436,
			573: 3436,
		},
	},
	{
		"minecraft:pumpkin_stem[age=4]",
		nil,
		NewMapping{
			404: 4257,
			477: 4760,
			573: 4760,
			4:   1668,
			393: 4256,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=13,south=side,west=none]",
		nil,
		NewMapping{
			393: 3026,
			404: 3027,
			477: 3330,
			573: 3330,
		},
	},
	{
		"minecraft:dark_prismarine_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			393: 6815,
			404: 6816,
			477: 7322,
			573: 7322,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 6608,
			573: 6608,
			393: 6101,
			404: 6102,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9606,
			573: 9606,
		},
	},
	{
		"minecraft:cornflower",
		nil,
		NewMapping{
			477: 1421,
			573: 1421,
		},
	},
	{
		"minecraft:jungle_sign[rotation=11,waterlogged=true]",
		nil,
		NewMapping{
			477: 3529,
			573: 3529,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=south,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			573: 4378,
			393: 3874,
			404: 3875,
			477: 4378,
		},
	},
	{
		"minecraft:acacia_door[facing=north,half=upper,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7870,
			404: 7871,
			477: 8395,
			573: 8395,
		},
	},
	{
		"minecraft:player_head[rotation=12]",
		nil,
		NewMapping{
			393: 5523,
			404: 5524,
			477: 6026,
			573: 6026,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=13,south=up,west=up]",
		nil,
		NewMapping{
			477: 3181,
			573: 3181,
			393: 2877,
			404: 2878,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=true,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10538,
			573: 10538,
		},
	},
	{
		"minecraft:magenta_concrete",
		nil,
		NewMapping{
			393: 8379,
			404: 8380,
			477: 8904,
			573: 8904,
			4:   4018,
		},
	},
	{
		"minecraft:cobblestone_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5137,
			404: 5138,
			477: 5641,
			573: 5641,
		},
	},
	{
		"minecraft:fire[age=6,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 1328,
			404: 1329,
			477: 1632,
			573: 1632,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5923,
			404: 5924,
			477: 6430,
			573: 6430,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9335,
			573: 9335,
		},
	},
	{
		"minecraft:dark_oak_fence_gate[facing=south,in_wall=false,open=true,powered=true]",
		nil,
		NewMapping{
			4:   2988,
			393: 7497,
			404: 7498,
			477: 8022,
			573: 8022,
		},
	},
	{
		"minecraft:gray_banner[rotation=1]",
		nil,
		NewMapping{
			477: 7474,
			573: 7474,
			393: 6967,
			404: 6968,
		},
	},
	{
		"minecraft:purpur_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 8109,
			404: 8110,
			477: 8634,
			573: 8634,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10735,
			573: 10735,
		},
	},
	{
		"minecraft:andesite_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9979,
			573: 9979,
		},
	},
	{
		"minecraft:tube_coral",
		nil,
		NewMapping{
			393: 8459,
		},
	},
	{
		"minecraft:brick_stairs[facing=east,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4401,
			404: 4402,
			477: 4905,
			573: 4905,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=7,south=none,west=side]",
		nil,
		NewMapping{
			393: 2110,
			404: 2111,
			477: 2414,
			573: 2414,
		},
	},
	{
		"minecraft:bubble_coral_wall_fan[facing=north,waterlogged=false]",
		nil,
		NewMapping{
			477: 9081,
			573: 9081,
			393: 8521,
			404: 8537,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=true,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 3998,
			404: 3999,
			477: 4502,
			573: 4502,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=true,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 4089,
			404: 4090,
			477: 4593,
			573: 4593,
		},
	},
	{
		"minecraft:mushroom_stem[down=false,east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 4175,
			404: 4176,
			477: 4679,
			573: 4679,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=true,north=true,powered=true,south=false,west=true]",
		nil,
		NewMapping{
			393: 4853,
			404: 4854,
			477: 5357,
			573: 5357,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9574,
			573: 9574,
		},
	},
	{
		"minecraft:stone_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9692,
			573: 9692,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=21,powered=true]",
		nil,
		NewMapping{
			393: 290,
			404: 290,
			477: 290,
			573: 290,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=13,south=none,west=none]",
		nil,
		NewMapping{
			393: 2453,
			404: 2454,
			477: 2757,
			573: 2757,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 5992,
			477: 6498,
			573: 6498,
			393: 5991,
		},
	},
	{
		"minecraft:sugar_cane[age=0]",
		nil,
		NewMapping{
			393: 3442,
			404: 3443,
			477: 3946,
			573: 3946,
			4:   1328,
		},
	},
	{
		"minecraft:fire[age=5,east=false,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			477: 1618,
			573: 1618,
			393: 1314,
			404: 1315,
		},
	},
	{
		"minecraft:pink_bed[facing=south,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 850,
			404: 850,
			477: 1150,
			573: 1150,
		},
	},
	{
		"minecraft:acacia_door[facing=east,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7927,
			404: 7928,
			477: 8452,
			573: 8452,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=2,powered=false]",
		nil,
		NewMapping{
			393: 403,
			404: 403,
			477: 403,
			573: 403,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10731,
			573: 10731,
		},
	},
	{
		"minecraft:red_glazed_terracotta[facing=west]",
		nil,
		NewMapping{
			4:   3985,
			393: 8371,
			404: 8372,
			477: 8896,
			573: 8896,
		},
	},
	{
		"minecraft:acacia_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 6380,
			477: 6886,
			573: 6886,
			393: 6379,
		},
	},
	{
		"minecraft:birch_fence[east=true,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 7559,
			477: 8083,
			573: 8083,
			393: 7558,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=1,powered=true]",
		nil,
		NewMapping{
			477: 750,
			573: 750,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=1,south=up,west=none]",
		nil,
		NewMapping{
			393: 2339,
			404: 2340,
			477: 2643,
			573: 2643,
		},
	},
	{
		"minecraft:dark_oak_door[facing=west,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7979,
			404: 7980,
			477: 8504,
			573: 8504,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=west,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3895,
			404: 3896,
			477: 4399,
			573: 4399,
		},
	},
	{
		"minecraft:orange_stained_glass",
		nil,
		NewMapping{
			4:   1521,
			393: 3578,
			404: 3579,
			477: 4082,
			573: 4082,
		},
	},
	{
		"minecraft:white_glazed_terracotta[facing=south]",
		nil,
		NewMapping{
			573: 8839,
			4:   3760,
			393: 8314,
			404: 8315,
			477: 8839,
		},
	},
	{
		"minecraft:powered_rail[powered=false,shape=north_south]",
		nil,
		NewMapping{
			404: 1010,
			477: 1310,
			573: 1310,
			4:   432,
			393: 1010,
		},
	},
	{
		"minecraft:dark_oak_button[face=ceiling,facing=south,powered=true]",
		nil,
		NewMapping{
			393: 5441,
			404: 5442,
			477: 5948,
			573: 5948,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			477: 4584,
			573: 4584,
			393: 4080,
			404: 4081,
		},
	},
	{
		"minecraft:gray_bed[facing=west,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 868,
			404: 868,
			477: 1168,
			573: 1168,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10863,
			573: 10863,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=west,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			573: 4144,
			4:   1538,
			393: 3640,
			404: 3641,
			477: 4144,
		},
	},
	{
		"minecraft:quartz_stairs[facing=west,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 5750,
			477: 6256,
			573: 6256,
			393: 5749,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=14,powered=false]",
		nil,
		NewMapping{
			573: 577,
			393: 577,
			404: 577,
			477: 577,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=east,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			404: 3783,
			477: 4286,
			573: 4286,
			393: 3782,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6316,
			404: 6317,
			477: 6823,
			573: 6823,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10780,
			573: 10780,
		},
	},
	{
		"minecraft:dragon_head[rotation=4]",
		nil,
		NewMapping{
			404: 5556,
			477: 6058,
			573: 6058,
			393: 5555,
		},
	},
	{
		"minecraft:dark_oak_fence_gate[facing=east,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7509,
			404: 7510,
			477: 8034,
			573: 8034,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 3263,
			404: 3264,
			477: 3727,
			573: 3727,
		},
	},
	{
		"minecraft:quartz_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 6251,
			573: 6251,
			393: 5744,
			404: 5745,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=west,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			477: 4333,
			573: 4333,
			393: 3829,
			404: 3830,
		},
	},
	{
		"minecraft:moving_piston[facing=south,type=normal]",
		nil,
		NewMapping{
			4:   579,
			393: 1103,
			404: 1103,
			477: 1403,
			573: 1403,
		},
	},
	{
		"minecraft:birch_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4976,
			404: 4977,
			477: 5480,
			573: 5480,
		},
	},
	{
		"minecraft:oak_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 1653,
			404: 1654,
			477: 1957,
			573: 1957,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=5,south=up,west=side]",
		nil,
		NewMapping{
			573: 2534,
			393: 2230,
			404: 2231,
			477: 2534,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=24,powered=true]",
		nil,
		NewMapping{
			393: 596,
			404: 596,
			477: 596,
			573: 596,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6466,
			404: 6467,
			477: 6973,
			573: 6973,
		},
	},
	{
		"minecraft:dead_brain_coral[waterlogged=true]",
		nil,
		NewMapping{
			477: 8986,
			573: 8986,
			404: 8462,
		},
	},
	{
		"minecraft:granite_stairs[facing=south,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9881,
			573: 9881,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=true,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10922,
			573: 10922,
		},
	},
	{
		"minecraft:nether_bricks",
		nil,
		NewMapping{
			4:   1792,
			393: 4495,
			404: 4496,
			477: 4999,
			573: 4999,
		},
	},
	{
		"minecraft:spruce_sapling[stage=1]",
		nil,
		NewMapping{
			404: 24,
			477: 24,
			573: 24,
			4:   105,
			393: 24,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=false,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 6581,
			573: 6581,
			393: 6074,
			404: 6075,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=19,powered=false]",
		nil,
		NewMapping{
			393: 637,
			404: 637,
			477: 637,
			573: 637,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 10661,
			477: 10661,
		},
	},
	{
		"minecraft:stone_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9661,
			573: 9661,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   2886,
			393: 7198,
			404: 7199,
			477: 7705,
			573: 7705,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=west,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3633,
			404: 3634,
			477: 4137,
			573: 4137,
		},
	},
	{
		"minecraft:fire[age=14,east=true,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			404: 1591,
			477: 1894,
			573: 1894,
			393: 1590,
		},
	},
	{
		"minecraft:birch_button[face=wall,facing=south,powered=true]",
		nil,
		NewMapping{
			393: 5361,
			404: 5362,
			477: 5868,
			573: 5868,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=15,south=side,west=side]",
		nil,
		NewMapping{
			404: 2324,
			477: 2627,
			573: 2627,
			393: 2323,
		},
	},
	{
		"minecraft:acacia_door[facing=east,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			404: 7932,
			477: 8456,
			573: 8456,
			393: 7931,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=3,south=side,west=side]",
		nil,
		NewMapping{
			573: 3239,
			393: 2935,
			404: 2936,
			477: 3239,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9820,
			573: 9820,
		},
	},
	{
		"minecraft:smooth_stone",
		nil,
		NewMapping{
			4:   696,
			393: 7353,
			404: 7354,
			477: 7878,
			573: 7878,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=7,south=none,west=side]",
		nil,
		NewMapping{
			393: 2686,
			404: 2687,
			477: 2990,
			573: 2990,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=3,south=up,west=side]",
		nil,
		NewMapping{
			393: 2212,
			404: 2213,
			477: 2516,
			573: 2516,
		},
	},
	{
		"minecraft:birch_door[facing=east,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			477: 8319,
			573: 8319,
			393: 7794,
			404: 7795,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=1,south=side,west=none]",
		nil,
		NewMapping{
			477: 2502,
			573: 2502,
			393: 2198,
			404: 2199,
		},
	},
	{
		"minecraft:dark_oak_button[face=wall,facing=south,powered=true]",
		nil,
		NewMapping{
			393: 5433,
			404: 5434,
			477: 5940,
			573: 5940,
		},
	},
	{
		"minecraft:fire[age=3,east=true,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			404: 1247,
			477: 1550,
			573: 1550,
			393: 1246,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10096,
			573: 10096,
		},
	},
	{
		"minecraft:mossy_stone_brick_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			477: 10265,
			573: 10265,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5964,
			404: 5965,
			477: 6471,
			573: 6471,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 10244,
			573: 10244,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10668,
			573: 10668,
		},
	},
	{
		"minecraft:lava[level=6]",
		nil,
		NewMapping{
			4:   166,
			393: 56,
			404: 56,
			477: 56,
			573: 56,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=true,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 4130,
			404: 4131,
			477: 4634,
			573: 4634,
		},
	},
	{
		"minecraft:jungle_button[face=floor,facing=north,powered=false]",
		nil,
		NewMapping{
			393: 5376,
			404: 5377,
			477: 5883,
			573: 5883,
		},
	},
	{
		"minecraft:purple_banner[rotation=12]",
		nil,
		NewMapping{
			393: 7026,
			404: 7027,
			477: 7533,
			573: 7533,
		},
	},
	{
		"minecraft:fire[age=0,east=true,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			477: 1439,
			573: 1439,
			393: 1135,
			404: 1136,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=23,powered=true]",
		nil,
		NewMapping{
			477: 1044,
			573: 1044,
		},
	},
	{
		"minecraft:bell[attachment=ceiling,facing=south]",
		nil,
		NewMapping{
			477: 11203,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=true,north=true,powered=true,south=false,west=true]",
		nil,
		NewMapping{
			393: 4757,
			404: 4758,
			477: 5261,
			573: 5261,
		},
	},
	{
		"minecraft:fire[age=14,east=false,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1606,
			404: 1607,
			477: 1910,
			573: 1910,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10464,
			573: 10464,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=4,south=side,west=side]",
		nil,
		NewMapping{
			477: 3104,
			573: 3104,
			393: 2800,
			404: 2801,
		},
	},
	{
		"minecraft:brown_bed[facing=north,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 942,
			404: 942,
			477: 1242,
			573: 1242,
		},
	},
	{
		"minecraft:yellow_banner[rotation=4]",
		nil,
		NewMapping{
			393: 6922,
			404: 6923,
			477: 7429,
			573: 7429,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=4,south=side,west=none]",
		nil,
		NewMapping{
			573: 2961,
			393: 2657,
			404: 2658,
			477: 2961,
		},
	},
	{
		"minecraft:fire[age=8,east=false,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			477: 1719,
			573: 1719,
			393: 1415,
			404: 1416,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=true,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10627,
			573: 10627,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10790,
			573: 10790,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10030,
			573: 10030,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9444,
			573: 9444,
		},
	},
	{
		"minecraft:fire[age=14,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1612,
			404: 1613,
			477: 1916,
			573: 1916,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=15,south=none,west=none]",
		nil,
		NewMapping{
			477: 3351,
			573: 3351,
			4:   895,
			393: 3047,
			404: 3048,
		},
	},
	{
		"minecraft:stone_button[face=floor,facing=west,powered=false]",
		nil,
		NewMapping{
			404: 3397,
			477: 3900,
			573: 3900,
			393: 3396,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=9,south=up,west=up]",
		nil,
		NewMapping{
			404: 1834,
			477: 2137,
			573: 2137,
			393: 1833,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=5,south=side,west=none]",
		nil,
		NewMapping{
			573: 2970,
			393: 2666,
			404: 2667,
			477: 2970,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=13,south=up,west=up]",
		nil,
		NewMapping{
			393: 2013,
			404: 2014,
			477: 2317,
			573: 2317,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=15,south=none,west=up]",
		nil,
		NewMapping{
			393: 3045,
			404: 3046,
			477: 3349,
			573: 3349,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=3,powered=false]",
		nil,
		NewMapping{
			573: 905,
			477: 905,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10031,
			573: 10031,
		},
	},
	{
		"minecraft:powered_rail[powered=false,shape=ascending_west]",
		nil,
		NewMapping{
			4:   435,
			393: 1013,
			404: 1013,
			477: 1313,
			573: 1313,
		},
	},
	{
		"minecraft:birch_fence[east=true,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7555,
			404: 7556,
			477: 8080,
			573: 8080,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=13,south=side,west=side]",
		nil,
		NewMapping{
			393: 3025,
			404: 3026,
			477: 3329,
			573: 3329,
		},
	},
	{
		"minecraft:acacia_leaves[distance=4,persistent=false]",
		nil,
		NewMapping{
			393: 207,
			404: 207,
			477: 207,
			573: 207,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=15,powered=false]",
		nil,
		NewMapping{
			393: 279,
			404: 279,
			477: 279,
			573: 279,
		},
	},
	{
		"minecraft:jungle_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 5082,
			404: 5083,
			477: 5586,
			573: 5586,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 11062,
			573: 11062,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9155,
			573: 9155,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=8,south=side,west=up]",
		nil,
		NewMapping{
			393: 2259,
			404: 2260,
			477: 2563,
			573: 2563,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=1,south=none,west=up]",
		nil,
		NewMapping{
			393: 2055,
			404: 2056,
			477: 2359,
			573: 2359,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=north,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3672,
			404: 3673,
			477: 4176,
			573: 4176,
		},
	},
	{
		"minecraft:andesite_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9986,
			573: 9986,
		},
	},
	{
		"minecraft:diorite_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 11086,
			573: 11086,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=12,powered=false]",
		nil,
		NewMapping{
			393: 523,
			404: 523,
			477: 523,
			573: 523,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=2,south=side,west=side]",
		nil,
		NewMapping{
			393: 2062,
			404: 2063,
			477: 2366,
			573: 2366,
		},
	},
	{
		"minecraft:magenta_banner[rotation=11]",
		nil,
		NewMapping{
			393: 6897,
			404: 6898,
			477: 7404,
			573: 7404,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10749,
			573: 10749,
		},
	},
	{
		"minecraft:composter[level=6]",
		nil,
		NewMapping{
			477: 11268,
			573: 11284,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=1,powered=true]",
		nil,
		NewMapping{
			393: 700,
			404: 700,
			477: 700,
			573: 700,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6135,
			404: 6136,
			477: 6642,
			573: 6642,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10386,
			573: 10386,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 5161,
			393: 4657,
			404: 4658,
			477: 5161,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9821,
			573: 9821,
		},
	},
	{
		"minecraft:jungle_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5097,
			404: 5098,
			477: 5601,
			573: 5601,
		},
	},
	{
		"minecraft:green_banner[rotation=3]",
		nil,
		NewMapping{
			393: 7065,
			404: 7066,
			477: 7572,
			573: 7572,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10850,
			573: 10850,
		},
	},
	{
		"minecraft:spruce_stairs[facing=east,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4946,
			404: 4947,
			477: 5450,
			573: 5450,
		},
	},
	{
		"minecraft:jungle_fence[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 8117,
			393: 7592,
			404: 7593,
			477: 8117,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=10,powered=true]",
		nil,
		NewMapping{
			404: 718,
			477: 718,
			573: 718,
			393: 718,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=15,south=side,west=none]",
		nil,
		NewMapping{
			573: 3060,
			393: 2756,
			404: 2757,
			477: 3060,
		},
	},
	{
		"minecraft:light_gray_bed[facing=west,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 886,
			404: 886,
			477: 1186,
			573: 1186,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9466,
			477: 9466,
		},
	},
	{
		"minecraft:wheat[age=4]",
		nil,
		NewMapping{
			4:   948,
			393: 3055,
			404: 3056,
			477: 3359,
			573: 3359,
		},
	},
	{
		"minecraft:oak_fence_gate[facing=east,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			393: 4325,
			404: 4326,
			477: 4829,
			573: 4829,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6764,
			404: 6765,
			477: 7271,
			573: 7271,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=13,powered=false]",
		nil,
		NewMapping{
			477: 375,
			573: 375,
			393: 375,
			404: 375,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 10154,
			573: 10154,
		},
	},
	{
		"minecraft:birch_door[facing=east,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			404: 7803,
			477: 8327,
			573: 8327,
			4:   3108,
			393: 7802,
		},
	},
	{
		"minecraft:birch_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 5543,
			393: 5039,
			404: 5040,
			477: 5543,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=true,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10474,
			573: 10474,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10404,
			477: 10404,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=east,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10162,
			573: 10162,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10807,
			573: 10807,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9708,
			573: 9708,
		},
	},
	{
		"minecraft:iron_bars[east=false,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 4709,
			573: 4709,
			393: 4205,
			404: 4206,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=west,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			573: 4463,
			393: 3959,
			404: 3960,
			477: 4463,
		},
	},
	{
		"minecraft:sign[rotation=2,waterlogged=true]",
		nil,
		NewMapping{
			393: 3079,
		},
	},
	{
		"minecraft:brown_wall_banner[facing=east]",
		nil,
		NewMapping{
			393: 7161,
			404: 7162,
			477: 7668,
			573: 7668,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9269,
			573: 9269,
		},
	},
	{
		"minecraft:brick_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4365,
			404: 4366,
			477: 4869,
			573: 4869,
		},
	},
	{
		"minecraft:birch_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4990,
			404: 4991,
			477: 5494,
			573: 5494,
		},
	},
	{
		"minecraft:birch_sign[rotation=2,waterlogged=true]",
		nil,
		NewMapping{
			477: 3447,
			573: 3447,
		},
	},
	{
		"minecraft:oak_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 1702,
			404: 1703,
			477: 2006,
			573: 2006,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=14,south=side,west=up]",
		nil,
		NewMapping{
			393: 2601,
			404: 2602,
			477: 2905,
			573: 2905,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=8,south=none,west=side]",
		nil,
		NewMapping{
			477: 3143,
			573: 3143,
			393: 2839,
			404: 2840,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=false,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 4073,
			404: 4074,
			477: 4577,
			573: 4577,
		},
	},
	{
		"minecraft:iron_bars[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 4193,
			404: 4194,
			477: 4697,
			573: 4697,
		},
	},
	{
		"minecraft:granite_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9879,
			573: 9879,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=12,south=up,west=none]",
		nil,
		NewMapping{
			393: 2294,
			404: 2295,
			477: 2598,
			573: 2598,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=east,in_wall=true,open=false,powered=true]",
		nil,
		NewMapping{
			404: 7416,
			477: 7940,
			573: 7940,
			393: 7415,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6663,
			404: 6664,
			477: 7170,
			573: 7170,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9828,
			573: 9828,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9427,
			573: 9427,
		},
	},
	{
		"minecraft:dispenser[facing=west,triggered=false]",
		nil,
		NewMapping{
			4:   372,
			393: 240,
			404: 240,
			477: 240,
			573: 240,
		},
	},
	{
		"minecraft:fire[age=6,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1339,
			404: 1340,
			477: 1643,
			573: 1643,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9744,
			573: 9744,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=15,south=none,west=side]",
		nil,
		NewMapping{
			393: 2758,
			404: 2759,
			477: 3062,
			573: 3062,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=true,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10822,
			573: 10822,
		},
	},
	{
		"minecraft:scaffolding[bottom=false,distance=6,waterlogged=false]",
		nil,
		NewMapping{
			477: 11128,
			573: 11128,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=east,in_wall=false,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7451,
			404: 7452,
			477: 7976,
			573: 7976,
			4:   2971,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			477: 4551,
			573: 4551,
			4:   1588,
			393: 4047,
			404: 4048,
		},
	},
	{
		"minecraft:sugar_cane[age=3]",
		nil,
		NewMapping{
			573: 3949,
			4:   1331,
			393: 3445,
			404: 3446,
			477: 3949,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=12,powered=false]",
		nil,
		NewMapping{
			393: 273,
			404: 273,
			477: 273,
			573: 273,
		},
	},
	{
		"minecraft:spruce_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4942,
			404: 4943,
			477: 5446,
			573: 5446,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=14,powered=false]",
		nil,
		NewMapping{
			393: 677,
			404: 677,
			477: 677,
			573: 677,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10642,
			573: 10642,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9565,
			573: 9565,
		},
	},
	{
		"minecraft:oak_log[axis=y]",
		nil,
		NewMapping{
			393: 73,
			404: 73,
			477: 73,
			573: 73,
			4:   272,
		},
	},
	{
		"minecraft:sugar_cane[age=10]",
		nil,
		NewMapping{
			4:   1338,
			393: 3452,
			404: 3453,
			477: 3956,
			573: 3956,
		},
	},
	{
		"minecraft:jungle_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 5625,
			393: 5121,
			404: 5122,
			477: 5625,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=false,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 4038,
			404: 4039,
			477: 4542,
			573: 4542,
		},
	},
	{
		"minecraft:iron_door[facing=east,half=lower,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			573: 3863,
			393: 3359,
			404: 3360,
			477: 3863,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=20,powered=false]",
		nil,
		NewMapping{
			477: 389,
			573: 389,
			393: 389,
			404: 389,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6734,
			404: 6735,
			477: 7241,
			573: 7241,
		},
	},
	{
		"minecraft:turtle_egg[eggs=2,hatch=2]",
		nil,
		NewMapping{
			477: 8967,
			573: 8967,
			393: 8442,
			404: 8443,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9223,
			573: 9223,
		},
	},
	{
		"minecraft:oak_door[facing=north,half=upper,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			573: 3571,
			393: 3107,
			404: 3108,
			477: 3571,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=east,in_wall=true,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7383,
			404: 7384,
			477: 7908,
			573: 7908,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10623,
			573: 10623,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10675,
			573: 10675,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=7,south=up,west=none]",
		nil,
		NewMapping{
			393: 2825,
			404: 2826,
			477: 3129,
			573: 3129,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=4,south=side,west=side]",
		nil,
		NewMapping{
			393: 1936,
			404: 1937,
			477: 2240,
			573: 2240,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=1,south=up,west=side]",
		nil,
		NewMapping{
			393: 2050,
			404: 2051,
			477: 2354,
			573: 2354,
		},
	},
	{
		"minecraft:sugar_cane[age=4]",
		nil,
		NewMapping{
			4:   1332,
			393: 3446,
			404: 3447,
			477: 3950,
			573: 3950,
		},
	},
	{
		"minecraft:repeater[delay=2,facing=south,locked=false,powered=true]",
		nil,
		NewMapping{
			4:   1508,
			393: 3535,
			404: 3536,
			477: 4039,
			573: 4039,
		},
	},
	{
		"minecraft:mushroom_stem[down=false,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 4178,
			404: 4179,
			477: 4682,
			573: 4682,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 6761,
			404: 6762,
			477: 7268,
			573: 7268,
		},
	},
	{
		"minecraft:fire[age=5,east=false,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1321,
			404: 1322,
			477: 1625,
			573: 1625,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=true,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			404: 4058,
			477: 4561,
			573: 4561,
			393: 4057,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=false,north=true,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10569,
			573: 10569,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9175,
			573: 9175,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=17,powered=false]",
		nil,
		NewMapping{
			573: 933,
			477: 933,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10463,
			573: 10463,
		},
	},
	{
		"minecraft:stone_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9665,
			573: 9665,
		},
	},
	{
		"minecraft:jigsaw[facing=down]",
		nil,
		NewMapping{
			477: 11261,
			573: 11277,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=north,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			477: 7883,
			573: 7883,
			393: 7358,
			404: 7359,
		},
	},
	{
		"minecraft:birch_door[facing=south,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7766,
			404: 7767,
			477: 8291,
			573: 8291,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4715,
			404: 4716,
			477: 5219,
			573: 5219,
		},
	},
	{
		"minecraft:fire[age=15,east=false,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1631,
			404: 1632,
			477: 1935,
			573: 1935,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=4,south=side,west=none]",
		nil,
		NewMapping{
			393: 1793,
			404: 1794,
			477: 2097,
			573: 2097,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9529,
			573: 9529,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=12,powered=false]",
		nil,
		NewMapping{
			477: 873,
			573: 873,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10791,
			573: 10791,
		},
	},
	{
		"minecraft:acacia_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 6870,
			4:   2610,
			393: 6363,
			404: 6364,
			477: 6870,
		},
	},
	{
		"minecraft:blue_banner[rotation=12]",
		nil,
		NewMapping{
			393: 7042,
			404: 7043,
			477: 7549,
			573: 7549,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=16,powered=true]",
		nil,
		NewMapping{
			393: 330,
			404: 330,
			477: 330,
			573: 330,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6091,
			404: 6092,
			477: 6598,
			573: 6598,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=false,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 4009,
			404: 4010,
			477: 4513,
			573: 4513,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 11043,
			477: 11043,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10720,
			573: 10720,
		},
	},
	{
		"minecraft:sign[rotation=2,waterlogged=false]",
		nil,
		NewMapping{
			4:   1010,
			393: 3080,
		},
	},
	{
		"minecraft:redstone_torch[lit=false]",
		nil,
		NewMapping{
			404: 3383,
			477: 3886,
			573: 3886,
			4:   1205,
			393: 3382,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 6558,
			393: 6051,
			404: 6052,
			477: 6558,
		},
	},
	{
		"minecraft:skeleton_skull[rotation=7]",
		nil,
		NewMapping{
			393: 5458,
			404: 5459,
			477: 5961,
			573: 5961,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=north,half=top,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3594,
			404: 3595,
			477: 4098,
			573: 4098,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=0,south=none,west=side]",
		nil,
		NewMapping{
			404: 2192,
			477: 2495,
			573: 2495,
			393: 2191,
		},
	},
	{
		"minecraft:dead_tube_coral_wall_fan[facing=east,waterlogged=true]",
		nil,
		NewMapping{
			393: 8470,
			404: 8486,
			477: 9030,
			573: 9030,
		},
	},
	{
		"minecraft:player_head[rotation=3]",
		nil,
		NewMapping{
			404: 5515,
			477: 6017,
			573: 6017,
			393: 5514,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10803,
			573: 10803,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=1,south=side,west=side]",
		nil,
		NewMapping{
			393: 2773,
			404: 2774,
			477: 3077,
			573: 3077,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4470,
			404: 4471,
			477: 4974,
			573: 4974,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=10,south=up,west=side]",
		nil,
		NewMapping{
			404: 2132,
			477: 2435,
			573: 2435,
			393: 2131,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=5,south=side,west=side]",
		nil,
		NewMapping{
			573: 2537,
			393: 2233,
			404: 2234,
			477: 2537,
		},
	},
	{
		"minecraft:birch_door[facing=east,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7799,
			404: 7800,
			477: 8324,
			573: 8324,
		},
	},
	{
		"minecraft:white_banner[rotation=10]",
		nil,
		NewMapping{
			404: 6865,
			477: 7371,
			573: 7371,
			4:   2826,
			393: 6864,
		},
	},
	{
		"minecraft:chipped_anvil[facing=south]",
		nil,
		NewMapping{
			573: 6079,
			4:   2324,
			393: 5572,
			404: 5573,
			477: 6079,
		},
	},
	{
		"minecraft:magenta_bed[facing=north,occupied=true,part=foot]",
		nil,
		NewMapping{
			404: 781,
			477: 1081,
			573: 1081,
			393: 781,
		},
	},
	{
		"minecraft:yellow_wall_banner[facing=east]",
		nil,
		NewMapping{
			393: 7129,
			404: 7130,
			477: 7636,
			573: 7636,
		},
	},
	{
		"minecraft:cyan_bed[facing=west,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 900,
			404: 900,
			477: 1200,
			573: 1200,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=6,south=up,west=side]",
		nil,
		NewMapping{
			477: 2255,
			573: 2255,
			393: 1951,
			404: 1952,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9848,
			573: 9848,
		},
	},
	{
		"minecraft:stone_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 9628,
			477: 9628,
		},
	},
	{
		"minecraft:orange_shulker_box[facing=west]",
		nil,
		NewMapping{
			4:   3524,
			393: 8226,
			404: 8227,
			477: 8751,
			573: 8751,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=11,powered=false]",
		nil,
		NewMapping{
			573: 371,
			393: 371,
			404: 371,
			477: 371,
		},
	},
	{
		"minecraft:acacia_door[facing=north,half=lower,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7877,
			404: 7878,
			477: 8402,
			573: 8402,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=4,south=side,west=side]",
		nil,
		NewMapping{
			393: 2224,
			404: 2225,
			477: 2528,
			573: 2528,
		},
	},
	{
		"minecraft:acacia_button[face=wall,facing=north,powered=true]",
		nil,
		NewMapping{
			393: 5407,
			404: 5408,
			477: 5914,
			573: 5914,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10693,
			573: 10693,
		},
	},
	{
		"minecraft:granite_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9910,
			573: 9910,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=false,north=false,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10581,
			573: 10581,
		},
	},
	{
		"minecraft:trapped_chest[facing=west,type=single,waterlogged=false]",
		nil,
		NewMapping{
			4:   2340,
			393: 5592,
			404: 5593,
			477: 6099,
			573: 6099,
		},
	},
	{
		"minecraft:fire[age=4,east=true,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1265,
			404: 1266,
			477: 1569,
			573: 1569,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=south,in_wall=true,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7464,
			404: 7465,
			477: 7989,
			573: 7989,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 6836,
			393: 6329,
			404: 6330,
			477: 6836,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=1,south=up,west=up]",
		nil,
		NewMapping{
			404: 2482,
			477: 2785,
			573: 2785,
			393: 2481,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=10,powered=true]",
		nil,
		NewMapping{
			393: 318,
			404: 318,
			477: 318,
			573: 318,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5220,
			404: 5221,
			477: 5724,
			573: 5724,
		},
	},
	{
		"minecraft:iron_door[facing=west,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 3344,
			404: 3345,
			477: 3848,
			573: 3848,
		},
	},
	{
		"minecraft:sugar_cane[age=9]",
		nil,
		NewMapping{
			573: 3955,
			4:   1337,
			393: 3451,
			404: 3452,
			477: 3955,
		},
	},
	{
		"minecraft:tube_coral_wall_fan[facing=south,waterlogged=false]",
		nil,
		NewMapping{
			393: 8507,
			404: 8523,
			477: 9067,
			573: 9067,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4666,
			404: 4667,
			477: 5170,
			573: 5170,
		},
	},
	{
		"minecraft:dragon_head[rotation=15]",
		nil,
		NewMapping{
			393: 5566,
			404: 5567,
			477: 6069,
			573: 6069,
		},
	},
	{
		"minecraft:repeater[delay=4,facing=north,locked=true,powered=false]",
		nil,
		NewMapping{
			393: 3562,
			404: 3563,
			477: 4066,
			573: 4066,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=0,powered=false]",
		nil,
		NewMapping{
			393: 599,
			404: 599,
			477: 599,
			573: 599,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10239,
			573: 10239,
		},
	},
	{
		"minecraft:acacia_door[facing=west,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7910,
			404: 7911,
			477: 8435,
			573: 8435,
		},
	},
	{
		"minecraft:farmland[moisture=1]",
		nil,
		NewMapping{
			477: 3364,
			573: 3364,
			4:   961,
			393: 3060,
			404: 3061,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=false,north=false,powered=true,south=false,west=false]",
		nil,
		NewMapping{
			477: 5350,
			573: 5350,
			4:   2121,
			393: 4846,
			404: 4847,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=1,south=side,west=side]",
		nil,
		NewMapping{
			393: 2629,
			404: 2630,
			477: 2933,
			573: 2933,
		},
	},
	{
		"minecraft:fire[age=0,east=false,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			573: 1458,
			393: 1154,
			404: 1155,
			477: 1458,
		},
	},
	{
		"minecraft:iron_door[facing=south,half=upper,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 3322,
			404: 3323,
			477: 3826,
			573: 3826,
		},
	},
	{
		"minecraft:light_blue_banner[rotation=6]",
		nil,
		NewMapping{
			393: 6908,
			404: 6909,
			477: 7415,
			573: 7415,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=3,south=side,west=up]",
		nil,
		NewMapping{
			393: 2646,
			404: 2647,
			477: 2950,
			573: 2950,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=11,south=up,west=none]",
		nil,
		NewMapping{
			393: 1997,
			404: 1998,
			477: 2301,
			573: 2301,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4716,
			404: 4717,
			477: 5220,
			573: 5220,
		},
	},
	{
		"minecraft:shulker_box[facing=south]",
		nil,
		NewMapping{
			393: 8213,
			404: 8214,
			477: 8738,
			573: 8738,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9474,
			573: 9474,
		},
	},
	{
		"minecraft:iron_door[facing=east,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			4:   1140,
			393: 3364,
			404: 3365,
			477: 3868,
			573: 3868,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6165,
			404: 6166,
			477: 6672,
			573: 6672,
		},
	},
	{
		"minecraft:fire[age=15,east=true,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			573: 1922,
			393: 1618,
			404: 1619,
			477: 1922,
		},
	},
	{
		"minecraft:jungle_button[face=wall,facing=south,powered=true]",
		nil,
		NewMapping{
			573: 5892,
			393: 5385,
			404: 5386,
			477: 5892,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=true,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 6559,
			393: 6052,
			404: 6053,
			477: 6559,
		},
	},
	{
		"minecraft:birch_sign[rotation=9,waterlogged=false]",
		nil,
		NewMapping{
			477: 3462,
			573: 3462,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4693,
			404: 4694,
			477: 5197,
			573: 5197,
		},
	},
	{
		"minecraft:green_wall_banner[facing=west]",
		nil,
		NewMapping{
			477: 7671,
			573: 7671,
			393: 7164,
			404: 7165,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			404: 6763,
			477: 7269,
			573: 7269,
			393: 6762,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=false,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10557,
			573: 10557,
		},
	},
	{
		"minecraft:diorite_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 11080,
			573: 11080,
		},
	},
	{
		"minecraft:birch_button[face=ceiling,facing=south,powered=false]",
		nil,
		NewMapping{
			393: 5370,
			404: 5371,
			477: 5877,
			573: 5877,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 6651,
			393: 6144,
			404: 6145,
			477: 6651,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=south,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3620,
			404: 3621,
			477: 4124,
			573: 4124,
			4:   1541,
		},
	},
	{
		"minecraft:gray_wool",
		nil,
		NewMapping{
			477: 1390,
			573: 1390,
			4:   567,
			393: 1090,
			404: 1090,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=west,in_wall=true,open=false,powered=false]",
		nil,
		NewMapping{
			477: 7901,
			573: 7901,
			393: 7376,
			404: 7377,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=9,powered=true]",
		nil,
		NewMapping{
			477: 666,
			573: 666,
			393: 666,
			404: 666,
		},
	},
	{
		"minecraft:spruce_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4936,
			404: 4937,
			477: 5440,
			573: 5440,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=south,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			477: 7017,
			573: 7017,
			393: 6510,
			404: 6511,
		},
	},
	{
		"minecraft:acacia_sign[rotation=0,waterlogged=true]",
		nil,
		NewMapping{
			477: 3475,
			573: 3475,
		},
	},
	{
		"minecraft:dead_horn_coral_wall_fan[facing=east,waterlogged=false]",
		nil,
		NewMapping{
			573: 9063,
			393: 8503,
			404: 8519,
			477: 9063,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=west,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			573: 4208,
			393: 3704,
			404: 3705,
			477: 4208,
		},
	},
	{
		"minecraft:acacia_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6366,
			404: 6367,
			477: 6873,
			573: 6873,
		},
	},
	{
		"minecraft:jungle_button[face=wall,facing=west,powered=false]",
		nil,
		NewMapping{
			573: 5895,
			393: 5388,
			404: 5389,
			477: 5895,
		},
	},
	{
		"minecraft:oak_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			393: 7257,
			404: 7258,
			477: 7764,
			573: 7764,
		},
	},
	{
		"minecraft:chorus_flower[age=4]",
		nil,
		NewMapping{
			4:   3204,
			393: 8071,
			404: 8072,
			477: 8596,
			573: 8596,
		},
	},
	{
		"minecraft:wall_torch[facing=south]",
		nil,
		NewMapping{
			573: 1436,
			4:   803,
			393: 1132,
			404: 1133,
			477: 1436,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=true,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5180,
			404: 5181,
			477: 5684,
			573: 5684,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=0,south=up,west=none]",
		nil,
		NewMapping{
			404: 1755,
			477: 2058,
			573: 2058,
			393: 1754,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9259,
			573: 9259,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 10114,
			573: 10114,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=0,south=none,west=up]",
		nil,
		NewMapping{
			393: 2334,
			404: 2335,
			477: 2638,
			573: 2638,
		},
	},
	{
		"minecraft:purpur_stairs[facing=east,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 8146,
			404: 8147,
			477: 8671,
			573: 8671,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=3,south=none,west=up]",
		nil,
		NewMapping{
			393: 2361,
			404: 2362,
			477: 2665,
			573: 2665,
		},
	},
	{
		"minecraft:dead_brain_coral_block",
		nil,
		NewMapping{
			573: 8975,
			393: 8450,
			404: 8451,
			477: 8975,
		},
	},
	{
		"minecraft:andesite_stairs[facing=south,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9953,
			573: 9953,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=5,south=none,west=side]",
		nil,
		NewMapping{
			404: 2669,
			477: 2972,
			573: 2972,
			393: 2668,
		},
	},
	{
		"minecraft:oak_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 1666,
			404: 1667,
			477: 1970,
			573: 1970,
		},
	},
	{
		"minecraft:fire[age=7,east=false,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 1384,
			404: 1385,
			477: 1688,
			573: 1688,
		},
	},
	{
		"minecraft:anvil[facing=north]",
		nil,
		NewMapping{
			573: 6074,
			4:   2322,
			393: 5567,
			404: 5568,
			477: 6074,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=south,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6589,
			404: 6590,
			477: 7096,
			573: 7096,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10820,
			477: 10820,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=8,south=up,west=up]",
		nil,
		NewMapping{
			393: 1968,
			404: 1969,
			477: 2272,
			573: 2272,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=9,waterlogged=true]",
		nil,
		NewMapping{
			477: 3557,
			573: 3557,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=13,waterlogged=false]",
		nil,
		NewMapping{
			477: 3566,
			573: 3566,
		},
	},
	{
		"minecraft:oak_door[facing=north,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			477: 3586,
			573: 3586,
			4:   1027,
			393: 3122,
			404: 3123,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=24,powered=false]",
		nil,
		NewMapping{
			393: 447,
			404: 447,
			477: 447,
			573: 447,
		},
	},
	{
		"minecraft:fire[age=4,east=true,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			477: 1571,
			573: 1571,
			393: 1267,
			404: 1268,
		},
	},
	{
		"minecraft:iron_door[facing=west,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			404: 3347,
			477: 3850,
			573: 3850,
			393: 3346,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=west,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			573: 4261,
			393: 3757,
			404: 3758,
			477: 4261,
		},
	},
	{
		"minecraft:blast_furnace[facing=south,lit=false]",
		nil,
		NewMapping{
			477: 11158,
			573: 11158,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9815,
			573: 9815,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=east,half=top,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			404: 3968,
			477: 4471,
			573: 4471,
			393: 3967,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=4,south=up,west=none]",
		nil,
		NewMapping{
			393: 2942,
			404: 2943,
			477: 3246,
			573: 3246,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6251,
			404: 6252,
			477: 6758,
			573: 6758,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=17,powered=true]",
		nil,
		NewMapping{
			393: 582,
			404: 582,
			477: 582,
			573: 582,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=east,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7381,
			404: 7382,
			477: 7906,
			573: 7906,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6595,
			404: 6596,
			477: 7102,
			573: 7102,
		},
	},
	{
		"minecraft:purpur_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 8129,
			404: 8130,
			477: 8654,
			573: 8654,
		},
	},
	{
		"minecraft:red_shulker_box[facing=east]",
		nil,
		NewMapping{
			393: 8302,
			404: 8303,
			477: 8827,
			573: 8827,
			4:   3733,
		},
	},
	{
		"minecraft:lime_glazed_terracotta[facing=north]",
		nil,
		NewMapping{
			4:   3842,
			393: 8333,
			404: 8334,
			477: 8858,
			573: 8858,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 4143,
			404: 4144,
			477: 4647,
			573: 4647,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=2,south=up,west=none]",
		nil,
		NewMapping{
			404: 2781,
			477: 3084,
			573: 3084,
			393: 2780,
		},
	},
	{
		"minecraft:yellow_bed[facing=south,occupied=true,part=head]",
		nil,
		NewMapping{
			573: 1116,
			393: 816,
			404: 816,
			477: 1116,
		},
	},
	{
		"minecraft:repeater[delay=3,facing=south,locked=true,powered=true]",
		nil,
		NewMapping{
			393: 3549,
			404: 3550,
			477: 4053,
			573: 4053,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5981,
			404: 5982,
			477: 6488,
			573: 6488,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=12,powered=true]",
		nil,
		NewMapping{
			477: 872,
			573: 872,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=2,south=side,west=up]",
		nil,
		NewMapping{
			573: 2509,
			393: 2205,
			404: 2206,
			477: 2509,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=11,south=none,west=up]",
		nil,
		NewMapping{
			393: 2865,
			404: 2866,
			477: 3169,
			573: 3169,
		},
	},
	{
		"minecraft:spruce_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 5432,
			393: 4928,
			404: 4929,
			477: 5432,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=north,in_wall=true,open=false,powered=false]",
		nil,
		NewMapping{
			477: 7981,
			573: 7981,
			393: 7456,
			404: 7457,
		},
	},
	{
		"minecraft:fire[age=14,east=true,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 1594,
			404: 1595,
			477: 1898,
			573: 1898,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9194,
			573: 9194,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9220,
			573: 9220,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=5,south=up,west=side]",
		nil,
		NewMapping{
			393: 2806,
			404: 2807,
			477: 3110,
			573: 3110,
		},
	},
	{
		"minecraft:birch_fence[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 7560,
			404: 7561,
			477: 8085,
			573: 8085,
		},
	},
	{
		"minecraft:diorite_wall[east=false,north=false,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 11098,
			573: 11098,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9417,
			573: 9417,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=14,south=none,west=side]",
		nil,
		NewMapping{
			393: 2173,
			404: 2174,
			477: 2477,
			573: 2477,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=12,powered=true]",
		nil,
		NewMapping{
			477: 822,
			573: 822,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=false,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 11060,
			477: 11060,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			4:   3184,
			393: 8066,
			404: 8067,
			477: 8591,
			573: 8591,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=21,powered=true]",
		nil,
		NewMapping{
			393: 390,
			404: 390,
			477: 390,
			573: 390,
		},
	},
	{
		"minecraft:fire[age=5,east=true,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1303,
			404: 1304,
			477: 1607,
			573: 1607,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=0,south=up,west=up]",
		nil,
		NewMapping{
			393: 2040,
			404: 2041,
			477: 2344,
			573: 2344,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=9,south=side,west=side]",
		nil,
		NewMapping{
			393: 1981,
			404: 1982,
			477: 2285,
			573: 2285,
		},
	},
	{
		"minecraft:white_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5821,
			404: 5822,
			477: 6328,
			573: 6328,
		},
	},
	{
		"minecraft:jungle_sign[rotation=1,waterlogged=false]",
		nil,
		NewMapping{
			477: 3510,
			573: 3510,
		},
	},
	{
		"minecraft:birch_wall_sign[facing=east,waterlogged=true]",
		nil,
		NewMapping{
			477: 3755,
			573: 3755,
		},
	},
	{
		"minecraft:heavy_weighted_pressure_plate[power=15]",
		nil,
		NewMapping{
			404: 5635,
			477: 6141,
			573: 6141,
			4:   2383,
			393: 5634,
		},
	},
	{
		"minecraft:acacia_door[facing=south,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			477: 8424,
			573: 8424,
			393: 7899,
			404: 7900,
		},
	},
	{
		"minecraft:stripped_acacia_log[axis=z]",
		nil,
		NewMapping{
			477: 101,
			573: 101,
			393: 101,
			404: 101,
		},
	},
	{
		"minecraft:iron_door[facing=south,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 3330,
			404: 3331,
			477: 3834,
			573: 3834,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 3201,
			404: 3202,
			477: 3665,
			573: 3665,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9365,
			573: 9365,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10663,
			573: 10663,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=1,powered=false]",
		nil,
		NewMapping{
			477: 901,
			573: 901,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 5698,
			573: 5698,
			393: 5194,
			404: 5195,
		},
	},
	{
		"minecraft:acacia_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			573: 6849,
			393: 6342,
			404: 6343,
			477: 6849,
		},
	},
	{
		"minecraft:iron_door[facing=east,half=upper,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 3352,
			404: 3353,
			477: 3856,
			573: 3856,
		},
	},
	{
		"minecraft:acacia_sign[rotation=8,waterlogged=false]",
		nil,
		NewMapping{
			477: 3492,
			573: 3492,
		},
	},
	{
		"minecraft:yellow_stained_glass",
		nil,
		NewMapping{
			4:   1524,
			393: 3581,
			404: 3582,
			477: 4085,
			573: 4085,
		},
	},
	{
		"minecraft:red_bed[facing=south,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 977,
			404: 977,
			477: 1277,
			573: 1277,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=9,south=none,west=none]",
		nil,
		NewMapping{
			573: 2145,
			393: 1841,
			404: 1842,
			477: 2145,
		},
	},
	{
		"minecraft:player_head[rotation=9]",
		nil,
		NewMapping{
			477: 6023,
			573: 6023,
			393: 5520,
			404: 5521,
		},
	},
	{
		"minecraft:polished_diorite_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			477: 10273,
			573: 10273,
		},
	},
	{
		"minecraft:chorus_plant[down=true,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 8021,
			404: 8022,
			477: 8546,
			573: 8546,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=7,powered=false]",
		nil,
		NewMapping{
			393: 613,
			404: 613,
			477: 613,
			573: 613,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=true,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10370,
			573: 10370,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=9,powered=true]",
		nil,
		NewMapping{
			573: 1016,
			477: 1016,
		},
	},
	{
		"minecraft:jungle_fence[east=true,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 7590,
			404: 7591,
			477: 8115,
			573: 8115,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=false,north=true,powered=true,south=true,west=false]",
		nil,
		NewMapping{
			393: 4804,
			404: 4805,
			477: 5308,
			573: 5308,
		},
	},
	{
		"minecraft:bubble_coral_wall_fan[facing=south,waterlogged=false]",
		nil,
		NewMapping{
			477: 9083,
			573: 9083,
			393: 8523,
			404: 8539,
		},
	},
	{
		"minecraft:trapped_chest[facing=west,type=left,waterlogged=true]",
		nil,
		NewMapping{
			393: 5593,
			404: 5594,
			477: 6100,
			573: 6100,
		},
	},
	{
		"minecraft:granite_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9918,
			573: 9918,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=false,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6226,
			404: 6227,
			477: 6733,
			573: 6733,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=east,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			404: 7479,
			477: 8003,
			573: 8003,
			393: 7478,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=10,south=none,west=side]",
		nil,
		NewMapping{
			393: 2137,
			404: 2138,
			477: 2441,
			573: 2441,
		},
	},
	{
		"minecraft:jungle_fence[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 8113,
			573: 8113,
			393: 7588,
			404: 7589,
		},
	},
	{
		"minecraft:glass_pane[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 4740,
			393: 4236,
			404: 4237,
			477: 4740,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=west,in_wall=false,open=true,powered=false]",
		nil,
		NewMapping{
			4:   2933,
			393: 7378,
			404: 7379,
			477: 7903,
			573: 7903,
		},
	},
	{
		"minecraft:birch_door[facing=south,half=upper,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7760,
			404: 7761,
			477: 8285,
			573: 8285,
		},
	},
	{
		"minecraft:birch_button[face=ceiling,facing=south,powered=true]",
		nil,
		NewMapping{
			393: 5369,
			404: 5370,
			477: 5876,
			573: 5876,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6282,
			404: 6283,
			477: 6789,
			573: 6789,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=10,south=side,west=up]",
		nil,
		NewMapping{
			573: 2437,
			393: 2133,
			404: 2134,
			477: 2437,
		},
	},
	{
		"minecraft:granite_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			573: 10303,
			477: 10303,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10625,
			573: 10625,
		},
	},
	{
		"minecraft:oak_door[facing=east,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			4:   1035,
			393: 3161,
			404: 3162,
			477: 3625,
			573: 3625,
		},
	},
	{
		"minecraft:dropper[facing=east,triggered=true]",
		nil,
		NewMapping{
			4:   2541,
			393: 5794,
			404: 5795,
			477: 6301,
			573: 6301,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=west,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3758,
			404: 3759,
			477: 4262,
			573: 4262,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 6786,
			393: 6279,
			404: 6280,
			477: 6786,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=south,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 6514,
			404: 6515,
			477: 7021,
			573: 7021,
		},
	},
	{
		"minecraft:fire[age=13,east=true,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			477: 1868,
			573: 1868,
			393: 1564,
			404: 1565,
		},
	},
	{
		"minecraft:lime_glazed_terracotta[facing=east]",
		nil,
		NewMapping{
			573: 8861,
			4:   3843,
			393: 8336,
			404: 8337,
			477: 8861,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=5,south=up,west=side]",
		nil,
		NewMapping{
			393: 2518,
			404: 2519,
			477: 2822,
			573: 2822,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=8,powered=false]",
		nil,
		NewMapping{
			573: 265,
			393: 265,
			404: 265,
			477: 265,
		},
	},
	{
		"minecraft:fire[age=8,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			573: 1724,
			393: 1420,
			404: 1421,
			477: 1724,
		},
	},
	{
		"minecraft:spruce_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4907,
			404: 4908,
			477: 5411,
			573: 5411,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			404: 7178,
			477: 7684,
			573: 7684,
			393: 7177,
		},
	},
	{
		"minecraft:red_bed[facing=north,occupied=true,part=foot]",
		nil,
		NewMapping{
			404: 973,
			477: 1273,
			573: 1273,
			393: 973,
		},
	},
	{
		"minecraft:acacia_stairs[facing=west,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 6888,
			573: 6888,
			393: 6381,
			404: 6382,
		},
	},
	{
		"minecraft:spruce_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4898,
			404: 4899,
			477: 5402,
			573: 5402,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=true,north=true,powered=false,south=false,west=false]",
		nil,
		NewMapping{
			404: 4827,
			477: 5330,
			573: 5330,
			393: 4826,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=6,south=side,west=side]",
		nil,
		NewMapping{
			573: 3122,
			393: 2818,
			404: 2819,
			477: 3122,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=10,powered=false]",
		nil,
		NewMapping{
			393: 369,
			404: 369,
			477: 369,
			573: 369,
		},
	},
	{
		"minecraft:potted_cornflower",
		nil,
		NewMapping{
			477: 5787,
			573: 5787,
		},
	},
	{
		"minecraft:hopper[enabled=true,facing=north]",
		nil,
		NewMapping{
			4:   2466,
			393: 5686,
			404: 5687,
			477: 6193,
			573: 6193,
		},
	},
	{
		"minecraft:fire[age=2,east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1227,
			404: 1228,
			477: 1531,
			573: 1531,
		},
	},
	{
		"minecraft:jungle_leaves[distance=4,persistent=true]",
		nil,
		NewMapping{
			393: 192,
			404: 192,
			477: 192,
			573: 192,
		},
	},
	{
		"minecraft:granite_stairs[facing=east,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9915,
			573: 9915,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=5,powered=false]",
		nil,
		NewMapping{
			477: 1009,
			573: 1009,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 4950,
			573: 4950,
			393: 4446,
			404: 4447,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 6420,
			477: 6926,
			573: 6926,
			393: 6419,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=1,powered=false]",
		nil,
		NewMapping{
			393: 601,
			404: 601,
			477: 601,
			573: 601,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5954,
			404: 5955,
			477: 6461,
			573: 6461,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=6,south=none,west=side]",
		nil,
		NewMapping{
			393: 2965,
			404: 2966,
			477: 3269,
			573: 3269,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 7229,
			573: 7229,
			393: 6722,
			404: 6723,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=23,powered=false]",
		nil,
		NewMapping{
			477: 395,
			573: 395,
			393: 395,
			404: 395,
		},
	},
	{
		"minecraft:magenta_stained_glass",
		nil,
		NewMapping{
			4:   1522,
			393: 3579,
			404: 3580,
			477: 4083,
			573: 4083,
		},
	},
	{
		"minecraft:spruce_door[facing=east,half=upper,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			4:   3097,
			393: 7732,
			404: 7733,
			477: 8257,
			573: 8257,
		},
	},
	{
		"minecraft:birch_leaves[distance=2,persistent=true]",
		nil,
		NewMapping{
			477: 174,
			573: 174,
			4:   302,
			393: 174,
			404: 174,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=north,half=top,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3792,
			404: 3793,
			477: 4296,
			573: 4296,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=east,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3653,
			404: 3654,
			477: 4157,
			573: 4157,
		},
	},
	{
		"minecraft:sign[rotation=12,waterlogged=true]",
		nil,
		NewMapping{
			393: 3099,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 6965,
			573: 6965,
			393: 6458,
			404: 6459,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=7,south=side,west=side]",
		nil,
		NewMapping{
			477: 2555,
			573: 2555,
			393: 2251,
			404: 2252,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9443,
			573: 9443,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6179,
			404: 6180,
			477: 6686,
			573: 6686,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=9,south=side,west=side]",
		nil,
		NewMapping{
			573: 3005,
			393: 2701,
			404: 2702,
			477: 3005,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=south,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 10113,
			573: 10113,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 10579,
			477: 10579,
		},
	},
	{
		"minecraft:fire[age=11,east=false,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1508,
			404: 1509,
			477: 1812,
			573: 1812,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=north,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3661,
			404: 3662,
			477: 4165,
			573: 4165,
		},
	},
	{
		"minecraft:dark_oak_fence[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 7645,
			404: 7646,
			477: 8170,
			573: 8170,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=2,powered=true]",
		nil,
		NewMapping{
			393: 602,
			404: 602,
			477: 602,
			573: 602,
		},
	},
	{
		"minecraft:andesite_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9950,
			477: 9950,
		},
	},
	{
		"minecraft:mushroom_stem[down=false,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 4159,
			404: 4160,
			477: 4663,
			573: 4663,
		},
	},
	{
		"minecraft:glass_pane[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 4731,
			393: 4227,
			404: 4228,
			477: 4731,
		},
	},
	{
		"minecraft:white_banner[rotation=12]",
		nil,
		NewMapping{
			404: 6867,
			477: 7373,
			573: 7373,
			4:   2828,
			393: 6866,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=west,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3891,
			404: 3892,
			477: 4395,
			573: 4395,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=6,south=none,west=none]",
		nil,
		NewMapping{
			393: 2534,
			404: 2535,
			477: 2838,
			573: 2838,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=10,south=side,west=none]",
		nil,
		NewMapping{
			404: 3000,
			477: 3303,
			573: 3303,
			393: 2999,
		},
	},
	{
		"minecraft:dark_oak_button[face=floor,facing=west,powered=false]",
		nil,
		NewMapping{
			393: 5428,
			404: 5429,
			477: 5935,
			573: 5935,
		},
	},
	{
		"minecraft:birch_fence[east=false,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7575,
			404: 7576,
			477: 8100,
			573: 8100,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=false,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 6829,
			573: 6829,
			393: 6322,
			404: 6323,
		},
	},
	{
		"minecraft:acacia_fence[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 7628,
			404: 7629,
			477: 8153,
			573: 8153,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4717,
			404: 4718,
			477: 5221,
			573: 5221,
		},
	},
	{
		"minecraft:magenta_bed[facing=east,occupied=true,part=foot]",
		nil,
		NewMapping{
			573: 1093,
			393: 793,
			404: 793,
			477: 1093,
		},
	},
	{
		"minecraft:gray_bed[facing=north,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 862,
			404: 862,
			477: 1162,
			573: 1162,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 6699,
			477: 7205,
			573: 7205,
			393: 6698,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=3,south=up,west=up]",
		nil,
		NewMapping{
			477: 2083,
			573: 2083,
			393: 1779,
			404: 1780,
		},
	},
	{
		"minecraft:dropper[facing=south,triggered=true]",
		nil,
		NewMapping{
			4:   2539,
			393: 5796,
			404: 5797,
			477: 6303,
			573: 6303,
		},
	},
	{
		"minecraft:fire[age=10,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			573: 1790,
			4:   826,
			393: 1486,
			404: 1487,
			477: 1790,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=true,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 4001,
			404: 4002,
			477: 4505,
			573: 4505,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=14,south=side,west=none]",
		nil,
		NewMapping{
			477: 3195,
			573: 3195,
			393: 2891,
			404: 2892,
		},
	},
	{
		"minecraft:wither_skeleton_wall_skull[facing=west]",
		nil,
		NewMapping{
			393: 5469,
			404: 5470,
			477: 5992,
			573: 5992,
		},
	},
	{
		"minecraft:stone_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9625,
			573: 9625,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=false,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			4:   2573,
			393: 6267,
			404: 6268,
			477: 6774,
			573: 6774,
		},
	},
	{
		"minecraft:purple_bed[facing=west,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 919,
			404: 919,
			477: 1219,
			573: 1219,
		},
	},
	{
		"minecraft:brick_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4404,
			404: 4405,
			477: 4908,
			573: 4908,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=6,powered=true]",
		nil,
		NewMapping{
			573: 660,
			393: 660,
			404: 660,
			477: 660,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=17,powered=false]",
		nil,
		NewMapping{
			573: 383,
			393: 383,
			404: 383,
			477: 383,
		},
	},
	{
		"minecraft:oak_door[facing=south,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 3127,
			404: 3128,
			477: 3591,
			573: 3591,
		},
	},
	{
		"minecraft:iron_bars[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 4194,
			404: 4195,
			477: 4698,
			573: 4698,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=false,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			404: 4132,
			477: 4635,
			573: 4635,
			393: 4131,
		},
	},
	{
		"minecraft:end_portal_frame[eye=true,facing=north]",
		nil,
		NewMapping{
			393: 4626,
			404: 4627,
			477: 5130,
			573: 5130,
			4:   1926,
		},
	},
	{
		"minecraft:purple_wool",
		nil,
		NewMapping{
			404: 1093,
			477: 1393,
			573: 1393,
			4:   570,
			393: 1093,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=false,north=false,powered=true,south=true,west=true]",
		nil,
		NewMapping{
			477: 5283,
			573: 5283,
			393: 4779,
			404: 4780,
		},
	},
	{
		"minecraft:fire[age=13,east=false,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			477: 1885,
			573: 1885,
			393: 1581,
			404: 1582,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=7,south=up,west=none]",
		nil,
		NewMapping{
			393: 2681,
			404: 2682,
			477: 2985,
			573: 2985,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=5,south=up,west=up]",
		nil,
		NewMapping{
			573: 2821,
			393: 2517,
			404: 2518,
			477: 2821,
		},
	},
	{
		"minecraft:white_bed[facing=south,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 753,
			404: 753,
			477: 1053,
			573: 1053,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 7183,
			404: 7184,
			477: 7690,
			573: 7690,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=true,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 10818,
			477: 10818,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=false,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10646,
			573: 10646,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=2,south=none,west=up]",
		nil,
		NewMapping{
			477: 2080,
			573: 2080,
			393: 1776,
			404: 1777,
		},
	},
	{
		"minecraft:fire[age=8,east=false,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			477: 1718,
			573: 1718,
			393: 1414,
			404: 1415,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=9,south=none,west=up]",
		nil,
		NewMapping{
			393: 2559,
			404: 2560,
			477: 2863,
			573: 2863,
		},
	},
	{
		"minecraft:oak_door[facing=south,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 3132,
			404: 3133,
			477: 3596,
			573: 3596,
		},
	},
	{
		"minecraft:cut_sandstone_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			573: 7822,
			477: 7822,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10091,
			573: 10091,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=true,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			404: 8043,
			477: 8567,
			573: 8567,
			393: 8042,
		},
	},
	{
		"minecraft:green_bed[facing=south,occupied=true,part=foot]",
		nil,
		NewMapping{
			404: 961,
			477: 1261,
			573: 1261,
			393: 961,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=6,powered=false]",
		nil,
		NewMapping{
			404: 361,
			477: 361,
			573: 361,
			393: 361,
		},
	},
	{
		"minecraft:lectern[facing=west,has_book=false,powered=true]",
		nil,
		NewMapping{
			573: 11187,
			477: 11187,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10242,
			573: 10242,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10635,
			573: 10635,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10348,
			477: 10348,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9147,
			573: 9147,
		},
	},
	{
		"minecraft:stone_button[face=wall,facing=south,powered=false]",
		nil,
		NewMapping{
			4:   1235,
			393: 3402,
			404: 3403,
			477: 3906,
			573: 3906,
		},
	},
	{
		"minecraft:end_gateway",
		nil,
		NewMapping{
			573: 8688,
			4:   3344,
			393: 8163,
			404: 8164,
			477: 8688,
		},
	},
	{
		"minecraft:white_glazed_terracotta[facing=east]",
		nil,
		NewMapping{
			404: 8317,
			477: 8841,
			573: 8841,
			4:   3763,
			393: 8316,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=south,half=top,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			404: 6514,
			477: 7020,
			573: 7020,
			4:   2685,
			393: 6513,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=6,powered=false]",
		nil,
		NewMapping{
			477: 861,
			573: 861,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=false,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10458,
			573: 10458,
		},
	},
	{
		"minecraft:fire[age=7,east=true,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1374,
			404: 1375,
			477: 1678,
			573: 1678,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10717,
			573: 10717,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=2,powered=false]",
		nil,
		NewMapping{
			477: 1003,
			573: 1003,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=12,south=side,west=up]",
		nil,
		NewMapping{
			393: 2007,
			404: 2008,
			477: 2311,
			573: 2311,
		},
	},
	{
		"minecraft:oak_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 1715,
			404: 1716,
			477: 2019,
			573: 2019,
		},
	},
	{
		"minecraft:jungle_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 5571,
			393: 5067,
			404: 5068,
			477: 5571,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=7,south=side,west=side]",
		nil,
		NewMapping{
			573: 2843,
			393: 2539,
			404: 2540,
			477: 2843,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=15,south=up,west=up]",
		nil,
		NewMapping{
			393: 2031,
			404: 2032,
			477: 2335,
			573: 2335,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9409,
			573: 9409,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=true,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10683,
			573: 10683,
		},
	},
	{
		"minecraft:potatoes[age=7]",
		nil,
		NewMapping{
			477: 5809,
			573: 5809,
			4:   2279,
			393: 5302,
			404: 5303,
		},
	},
	{
		"minecraft:jungle_leaves[distance=6,persistent=true]",
		nil,
		NewMapping{
			393: 196,
			404: 196,
			477: 196,
			573: 196,
		},
	},
	{
		"minecraft:spruce_fence[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 7547,
			477: 8071,
			573: 8071,
			393: 7546,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=11,south=none,west=side]",
		nil,
		NewMapping{
			393: 2578,
			404: 2579,
			477: 2882,
			573: 2882,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=2,powered=false]",
		nil,
		NewMapping{
			477: 853,
			573: 853,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=true,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			4:   1598,
			393: 3987,
			404: 3988,
			477: 4491,
			573: 4491,
		},
	},
	{
		"minecraft:lime_bed[facing=south,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 832,
			404: 832,
			477: 1132,
			573: 1132,
		},
	},
	{
		"minecraft:dark_oak_door[facing=north,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			573: 8464,
			393: 7939,
			404: 7940,
			477: 8464,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10831,
			573: 10831,
		},
	},
	{
		"minecraft:brick_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			404: 4351,
			477: 4854,
			573: 4854,
			393: 4350,
		},
	},
	{
		"minecraft:andesite_stairs[facing=south,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9961,
			573: 9961,
		},
	},
	{
		"minecraft:smooth_red_sandstone_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			477: 10264,
			573: 10264,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6243,
			404: 6244,
			477: 6750,
			573: 6750,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5219,
			404: 5220,
			477: 5723,
			573: 5723,
		},
	},
	{
		"minecraft:jungle_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 5080,
			477: 5583,
			573: 5583,
			393: 5079,
		},
	},
	{
		"minecraft:cyan_banner[rotation=12]",
		nil,
		NewMapping{
			404: 7011,
			477: 7517,
			573: 7517,
			393: 7010,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=5,south=none,west=none]",
		nil,
		NewMapping{
			393: 2525,
			404: 2526,
			477: 2829,
			573: 2829,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9490,
			573: 9490,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 7141,
			573: 7141,
			393: 6634,
			404: 6635,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9453,
			573: 9453,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=24,powered=false]",
		nil,
		NewMapping{
			477: 797,
			573: 797,
		},
	},
	{
		"minecraft:brick_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   1735,
			393: 4333,
			404: 4334,
			477: 4837,
			573: 4837,
		},
	},
	{
		"minecraft:spruce_button[face=ceiling,facing=south,powered=false]",
		nil,
		NewMapping{
			393: 5346,
			404: 5347,
			477: 5853,
			573: 5853,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9537,
			573: 9537,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=15,powered=false]",
		nil,
		NewMapping{
			477: 879,
			573: 879,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9402,
			573: 9402,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 10089,
			477: 10089,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10127,
			573: 10127,
		},
	},
	{
		"minecraft:bamboo_sapling",
		nil,
		NewMapping{
			477: 9115,
			573: 9115,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=5,south=none,west=side]",
		nil,
		NewMapping{
			404: 2957,
			477: 3260,
			573: 3260,
			393: 2956,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5246,
			404: 5247,
			477: 5750,
			573: 5750,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=false,north=false,powered=true,south=false,west=true]",
		nil,
		NewMapping{
			573: 5349,
			393: 4845,
			404: 4846,
			477: 5349,
		},
	},
	{
		"minecraft:cyan_wall_banner[facing=east]",
		nil,
		NewMapping{
			393: 7149,
			404: 7150,
			477: 7656,
			573: 7656,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=6,powered=true]",
		nil,
		NewMapping{
			477: 1010,
			573: 1010,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=15,south=none,west=side]",
		nil,
		NewMapping{
			393: 2326,
			404: 2327,
			477: 2630,
			573: 2630,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=3,south=side,west=none]",
		nil,
		NewMapping{
			477: 2664,
			573: 2664,
			393: 2360,
			404: 2361,
		},
	},
	{
		"minecraft:honey_block",
		nil,
		NewMapping{
			573: 11335,
		},
	},
	{
		"minecraft:comparator[facing=north,mode=compare,powered=false]",
		nil,
		NewMapping{
			393: 5636,
			404: 5637,
			477: 6143,
			573: 6143,
			4:   2402,
		},
	},
	{
		"minecraft:pumpkin_stem[age=7]",
		nil,
		NewMapping{
			393: 4259,
			404: 4260,
			477: 4763,
			573: 4763,
			4:   1671,
		},
	},
	{
		"minecraft:quartz_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 6211,
			573: 6211,
			393: 5704,
			404: 5705,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=north,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3795,
			404: 3796,
			477: 4299,
			573: 4299,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=14,south=up,west=none]",
		nil,
		NewMapping{
			393: 2888,
			404: 2889,
			477: 3192,
			573: 3192,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9741,
			573: 9741,
		},
	},
	{
		"minecraft:dropper[facing=west,triggered=true]",
		nil,
		NewMapping{
			4:   2540,
			393: 5798,
			404: 5799,
			477: 6305,
			573: 6305,
		},
	},
	{
		"minecraft:lime_bed[facing=north,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 831,
			404: 831,
			477: 1131,
			573: 1131,
		},
	},
	{
		"minecraft:jungle_door[facing=west,half=upper,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7844,
			404: 7845,
			477: 8369,
			573: 8369,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=16,powered=false]",
		nil,
		NewMapping{
			393: 731,
			404: 731,
			477: 731,
			573: 731,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=3,south=up,west=none]",
		nil,
		NewMapping{
			393: 2501,
			404: 2502,
			477: 2805,
			573: 2805,
		},
	},
	{
		"minecraft:fire[age=13,east=false,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1577,
			404: 1578,
			477: 1881,
			573: 1881,
		},
	},
	{
		"minecraft:yellow_bed[facing=west,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 822,
			404: 822,
			477: 1122,
			573: 1122,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9400,
			573: 9400,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9426,
			573: 9426,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 10131,
			477: 10131,
		},
	},
	{
		"minecraft:chorus_flower[age=3]",
		nil,
		NewMapping{
			4:   3203,
			393: 8070,
			404: 8071,
			477: 8595,
			573: 8595,
		},
	},
	{
		"minecraft:wheat[age=5]",
		nil,
		NewMapping{
			393: 3056,
			404: 3057,
			477: 3360,
			573: 3360,
			4:   949,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 6994,
			393: 6487,
			404: 6488,
			477: 6994,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10396,
			477: 10396,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10382,
			573: 10382,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			4:   1605,
			393: 4112,
			404: 4113,
			477: 4616,
			573: 4616,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=9,south=side,west=none]",
		nil,
		NewMapping{
			393: 2270,
			404: 2271,
			477: 2574,
			573: 2574,
		},
	},
	{
		"minecraft:iron_bars[east=false,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			404: 4202,
			477: 4705,
			573: 4705,
			393: 4201,
		},
	},
	{
		"minecraft:andesite_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10011,
			573: 10011,
		},
	},
	{
		"minecraft:pink_tulip",
		nil,
		NewMapping{
			4:   615,
			393: 1119,
			404: 1119,
			477: 1419,
			573: 1419,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=7,south=side,west=none]",
		nil,
		NewMapping{
			573: 2988,
			393: 2684,
			404: 2685,
			477: 2988,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=west,half=top,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3952,
			404: 3953,
			477: 4456,
			573: 4456,
		},
	},
	{
		"minecraft:fire[age=0,east=false,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			477: 1469,
			573: 1469,
			393: 1165,
			404: 1166,
		},
	},
	{
		"minecraft:andesite_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 10004,
			573: 10004,
		},
	},
	{
		"minecraft:brain_coral[waterlogged=true]",
		nil,
		NewMapping{
			404: 8472,
			477: 8996,
			573: 8996,
		},
	},
	{
		"minecraft:birch_sign[rotation=12,waterlogged=true]",
		nil,
		NewMapping{
			477: 3467,
			573: 3467,
		},
	},
	{
		"minecraft:brewing_stand[has_bottle_0=true,has_bottle_1=true,has_bottle_2=true]",
		nil,
		NewMapping{
			393: 4613,
			404: 4614,
			477: 5117,
			573: 5117,
			4:   1879,
		},
	},
	{
		"minecraft:observer[facing=up,powered=true]",
		nil,
		NewMapping{
			4:   3497,
			393: 8207,
			404: 8208,
			477: 8732,
			573: 8732,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6735,
			404: 6736,
			477: 7242,
			573: 7242,
		},
	},
	{
		"minecraft:spruce_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4958,
			404: 4959,
			477: 5462,
			573: 5462,
		},
	},
	{
		"minecraft:green_bed[facing=south,occupied=false,part=head]",
		nil,
		NewMapping{
			573: 1262,
			393: 962,
			404: 962,
			477: 1262,
		},
	},
	{
		"minecraft:pink_shulker_box[facing=up]",
		nil,
		NewMapping{
			393: 8257,
			404: 8258,
			477: 8782,
			573: 8782,
			4:   3601,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=2,south=side,west=side]",
		nil,
		NewMapping{
			573: 2078,
			393: 1774,
			404: 1775,
			477: 2078,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=false,north=false,powered=false,south=true,west=false]",
		nil,
		NewMapping{
			573: 5320,
			393: 4816,
			404: 4817,
			477: 5320,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=0,south=side,west=side]",
		nil,
		NewMapping{
			393: 2764,
			404: 2765,
			477: 3068,
			573: 3068,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9712,
			573: 9712,
		},
	},
	{
		"minecraft:dropper[facing=up,triggered=true]",
		nil,
		NewMapping{
			4:   2537,
			393: 5800,
			404: 5801,
			477: 6307,
			573: 6307,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5258,
			404: 5259,
			477: 5762,
			573: 5762,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5931,
			404: 5932,
			477: 6438,
			573: 6438,
		},
	},
	{
		"minecraft:nether_brick_fence[east=true,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 4502,
			404: 4503,
			477: 5006,
			573: 5006,
		},
	},
	{
		"minecraft:vine[east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			404: 4298,
			477: 4801,
			573: 4801,
			393: 4297,
		},
	},
	{
		"minecraft:fire[age=7,east=true,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			477: 1663,
			573: 1663,
			393: 1359,
			404: 1360,
		},
	},
	{
		"minecraft:bell[attachment=floor,facing=north]",
		nil,
		NewMapping{
			477: 11198,
		},
	},
	{
		"minecraft:hopper[enabled=true,facing=west]",
		nil,
		NewMapping{
			4:   2468,
			393: 5688,
			404: 5689,
			477: 6195,
			573: 6195,
		},
	},
	{
		"minecraft:cobweb",
		nil,
		NewMapping{
			404: 1040,
			477: 1340,
			573: 1340,
			4:   480,
			393: 1040,
		},
	},
	{
		"minecraft:fire_coral_wall_fan[facing=north,waterlogged=true]",
		nil,
		NewMapping{
			393: 8528,
			404: 8544,
			477: 9088,
			573: 9088,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			404: 4661,
			477: 5164,
			573: 5164,
			393: 4660,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=11,south=up,west=none]",
		nil,
		NewMapping{
			573: 3021,
			393: 2717,
			404: 2718,
			477: 3021,
		},
	},
	{
		"minecraft:blue_bed[facing=north,occupied=true,part=foot]",
		nil,
		NewMapping{
			573: 1225,
			393: 925,
			404: 925,
			477: 1225,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=15,south=side,west=none]",
		nil,
		NewMapping{
			393: 2324,
			404: 2325,
			477: 2628,
			573: 2628,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10504,
			573: 10504,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10832,
			573: 10832,
		},
	},
	{
		"minecraft:damaged_anvil[facing=north]",
		nil,
		NewMapping{
			477: 6082,
			573: 6082,
			4:   2330,
			393: 5575,
			404: 5576,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=false,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 4136,
			404: 4137,
			477: 4640,
			573: 4640,
		},
	},
	{
		"minecraft:fire[age=2,east=false,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1215,
			404: 1216,
			477: 1519,
			573: 1519,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=false,north=true,powered=false,south=false,west=false]",
		nil,
		NewMapping{
			393: 4842,
			404: 4843,
			477: 5346,
			573: 5346,
		},
	},
	{
		"minecraft:oak_sign[rotation=0,waterlogged=false]",
		nil,
		NewMapping{
			573: 3380,
			404: 3077,
			477: 3380,
		},
	},
	{
		"minecraft:birch_door[facing=north,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7751,
			404: 7752,
			477: 8276,
			573: 8276,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 7132,
			573: 7132,
			393: 6625,
			404: 6626,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 6925,
			393: 6418,
			404: 6419,
			477: 6925,
		},
	},
	{
		"minecraft:cyan_banner[rotation=0]",
		nil,
		NewMapping{
			404: 6999,
			477: 7505,
			573: 7505,
			393: 6998,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=7,south=none,west=up]",
		nil,
		NewMapping{
			404: 2254,
			477: 2557,
			573: 2557,
			393: 2253,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=14,powered=true]",
		nil,
		NewMapping{
			573: 376,
			393: 376,
			404: 376,
			477: 376,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=8,powered=false]",
		nil,
		NewMapping{
			477: 965,
			573: 965,
		},
	},
	{
		"minecraft:pink_banner[rotation=7]",
		nil,
		NewMapping{
			477: 7464,
			573: 7464,
			393: 6957,
			404: 6958,
		},
	},
	{
		"minecraft:zombie_head[rotation=12]",
		nil,
		NewMapping{
			393: 5503,
			404: 5504,
			477: 6006,
			573: 6006,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=north,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			404: 3861,
			477: 4364,
			573: 4364,
			393: 3860,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=5,south=side,west=side]",
		nil,
		NewMapping{
			573: 3113,
			393: 2809,
			404: 2810,
			477: 3113,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=south,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9401,
			573: 9401,
		},
	},
	{
		"minecraft:purpur_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 8645,
			573: 8645,
			393: 8120,
			404: 8121,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=12,powered=false]",
		nil,
		NewMapping{
			393: 723,
			404: 723,
			477: 723,
			573: 723,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9383,
			573: 9383,
		},
	},
	{
		"minecraft:fire[age=0,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			404: 1137,
			477: 1440,
			573: 1440,
			393: 1136,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			404: 4607,
			477: 5110,
			573: 5110,
			393: 4606,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 6372,
			573: 6372,
			393: 5865,
			404: 5866,
		},
	},
	{
		"minecraft:jungle_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			393: 7277,
			404: 7278,
			477: 7784,
			573: 7784,
		},
	},
	{
		"minecraft:dark_oak_door[facing=north,half=lower,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7941,
			404: 7942,
			477: 8466,
			573: 8466,
		},
	},
	{
		"minecraft:fire[age=7,east=true,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 1370,
			404: 1371,
			477: 1674,
			573: 1674,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 4992,
			573: 4992,
			393: 4488,
			404: 4489,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=west,half=top,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3760,
			404: 3761,
			477: 4264,
			573: 4264,
		},
	},
	{
		"minecraft:birch_sign[rotation=1,waterlogged=false]",
		nil,
		NewMapping{
			477: 3446,
			573: 3446,
		},
	},
	{
		"minecraft:daylight_detector[inverted=false,power=14]",
		nil,
		NewMapping{
			393: 5681,
			404: 5682,
			477: 6188,
			573: 6188,
			4:   2430,
		},
	},
	{
		"minecraft:end_portal_frame[eye=false,facing=north]",
		nil,
		NewMapping{
			4:   1922,
			393: 4630,
			404: 4631,
			477: 5134,
			573: 5134,
		},
	},
	{
		"minecraft:brick_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 4858,
			573: 4858,
			393: 4354,
			404: 4355,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=13,south=up,west=up]",
		nil,
		NewMapping{
			393: 2589,
			404: 2590,
			477: 2893,
			573: 2893,
		},
	},
	{
		"minecraft:purpur_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			393: 7349,
			404: 7350,
			477: 7874,
			573: 7874,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=10,south=side,west=none]",
		nil,
		NewMapping{
			404: 1848,
			477: 2151,
			573: 2151,
			393: 1847,
		},
	},
	{
		"minecraft:pink_bed[facing=north,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 844,
			404: 844,
			477: 1144,
			573: 1144,
		},
	},
	{
		"minecraft:gray_bed[facing=north,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 861,
			404: 861,
			477: 1161,
			573: 1161,
		},
	},
	{
		"minecraft:oak_door[facing=south,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			573: 3598,
			393: 3134,
			404: 3135,
			477: 3598,
		},
	},
	{
		"minecraft:birch_button[face=ceiling,facing=north,powered=true]",
		nil,
		NewMapping{
			393: 5367,
			404: 5368,
			477: 5874,
			573: 5874,
		},
	},
	{
		"minecraft:birch_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4988,
			404: 4989,
			477: 5492,
			573: 5492,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9462,
			573: 9462,
		},
	},
	{
		"minecraft:cyan_banner[rotation=4]",
		nil,
		NewMapping{
			393: 7002,
			404: 7003,
			477: 7509,
			573: 7509,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=west,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9263,
			573: 9263,
		},
	},
	{
		"minecraft:lectern[facing=north,has_book=true,powered=true]",
		nil,
		NewMapping{
			477: 11177,
			573: 11177,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=20,powered=true]",
		nil,
		NewMapping{
			477: 1038,
			573: 1038,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=10,powered=true]",
		nil,
		NewMapping{
			477: 968,
			573: 968,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=east,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			404: 3836,
			477: 4339,
			573: 4339,
			393: 3835,
		},
	},
	{
		"minecraft:fire[age=12,east=true,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			404: 1524,
			477: 1827,
			573: 1827,
			393: 1523,
		},
	},
	{
		"minecraft:zombie_head[rotation=2]",
		nil,
		NewMapping{
			393: 5493,
			404: 5494,
			477: 5996,
			573: 5996,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=7,south=up,west=up]",
		nil,
		NewMapping{
			393: 2103,
			404: 2104,
			477: 2407,
			573: 2407,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10544,
			573: 10544,
		},
	},
	{
		"minecraft:diorite_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10207,
			573: 10207,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10250,
			573: 10250,
		},
	},
	{
		"minecraft:fire[age=6,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			4:   822,
			393: 1358,
			404: 1359,
			477: 1662,
			573: 1662,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 6628,
			393: 6121,
			404: 6122,
			477: 6628,
		},
	},
	{
		"minecraft:blue_ice",
		nil,
		NewMapping{
			393: 8572,
			404: 8588,
			477: 9112,
			573: 9112,
		},
	},
	{
		"minecraft:cobblestone_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 5670,
			393: 5166,
			404: 5167,
			477: 5670,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10805,
			573: 10805,
		},
	},
	{
		"minecraft:acacia_door[facing=west,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7913,
			404: 7914,
			477: 8438,
			573: 8438,
		},
	},
	{
		"minecraft:turtle_egg[eggs=4,hatch=2]",
		nil,
		NewMapping{
			393: 8448,
			404: 8449,
			477: 8973,
			573: 8973,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10252,
			573: 10252,
		},
	},
	{
		"minecraft:beehive[facing=south,honey_level=1]",
		nil,
		NewMapping{
			573: 11318,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9317,
			573: 9317,
		},
	},
	{
		"minecraft:bell[attachment=ceiling,facing=north]",
		nil,
		NewMapping{
			477: 11202,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=true,north=true,powered=false,south=true,west=true]",
		nil,
		NewMapping{
			393: 4823,
			404: 4824,
			477: 5327,
			573: 5327,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=11,south=up,west=none]",
		nil,
		NewMapping{
			393: 2861,
			404: 2862,
			477: 3165,
			573: 3165,
		},
	},
	{
		"minecraft:birch_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			573: 5478,
			393: 4974,
			404: 4975,
			477: 5478,
		},
	},
	{
		"minecraft:jungle_door[facing=south,half=upper,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			477: 8349,
			573: 8349,
			393: 7824,
			404: 7825,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6653,
			404: 6654,
			477: 7160,
			573: 7160,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6093,
			404: 6094,
			477: 6600,
			573: 6600,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=12,south=side,west=up]",
		nil,
		NewMapping{
			404: 3016,
			477: 3319,
			573: 3319,
			393: 3015,
		},
	},
	{
		"minecraft:cactus[age=3]",
		nil,
		NewMapping{
			4:   1299,
			393: 3428,
			404: 3429,
			477: 3932,
			573: 3932,
		},
	},
	{
		"minecraft:chest[facing=north,type=single,waterlogged=false]",
		nil,
		NewMapping{
			404: 1730,
			477: 2033,
			573: 2033,
			4:   866,
			393: 1729,
		},
	},
	{
		"minecraft:white_glazed_terracotta[facing=west]",
		nil,
		NewMapping{
			4:   3761,
			393: 8315,
			404: 8316,
			477: 8840,
			573: 8840,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5942,
			404: 5943,
			477: 6449,
			573: 6449,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 6420,
			393: 5913,
			404: 5914,
			477: 6420,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=true,north=true,powered=false,south=true,west=true]",
		nil,
		NewMapping{
			393: 4791,
			404: 4792,
			477: 5295,
			573: 5295,
		},
	},
	{
		"minecraft:mushroom_stem[down=false,east=true,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 4162,
			404: 4163,
			477: 4666,
			573: 4666,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=6,south=side,west=up]",
		nil,
		NewMapping{
			477: 2977,
			573: 2977,
			393: 2673,
			404: 2674,
		},
	},
	{
		"minecraft:water[level=5]",
		nil,
		NewMapping{
			393: 39,
			404: 39,
			477: 39,
			573: 39,
			4:   133,
		},
	},
	{
		"minecraft:lava[level=11]",
		nil,
		NewMapping{
			4:   171,
			393: 61,
			404: 61,
			477: 61,
			573: 61,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 4599,
			404: 4600,
			477: 5103,
			573: 5103,
			4:   1824,
		},
	},
	{
		"minecraft:purple_bed[facing=north,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 910,
			404: 910,
			477: 1210,
			573: 1210,
		},
	},
	{
		"minecraft:oak_fence[east=true,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 3976,
			393: 3472,
			404: 3473,
			477: 3976,
		},
	},
	{
		"minecraft:brick_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 4332,
			404: 4333,
			477: 4836,
			573: 4836,
		},
	},
	{
		"minecraft:fire[age=11,east=true,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			404: 1496,
			477: 1799,
			573: 1799,
			393: 1495,
		},
	},
	{
		"minecraft:granite_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9912,
			573: 9912,
		},
	},
	{
		"minecraft:granite_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9897,
			573: 9897,
		},
	},
	{
		"minecraft:tall_seagrass[half=upper]",
		nil,
		NewMapping{
			393: 1045,
			404: 1045,
			477: 1345,
			573: 1345,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5188,
			404: 5189,
			477: 5692,
			573: 5692,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6729,
			404: 6730,
			477: 7236,
			573: 7236,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10344,
			573: 10344,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=west,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10142,
			573: 10142,
		},
	},
	{
		"minecraft:fire[age=11,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1518,
			404: 1519,
			477: 1822,
			573: 1822,
			4:   827,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=5,powered=false]",
		nil,
		NewMapping{
			573: 459,
			393: 459,
			404: 459,
			477: 459,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10915,
			573: 10915,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=true,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10819,
			573: 10819,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=south,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 10123,
			573: 10123,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=west,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9823,
			573: 9823,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9538,
			573: 9538,
		},
	},
	{
		"minecraft:dark_oak_door[facing=north,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			4:   3159,
			393: 7946,
			404: 7947,
			477: 8471,
			573: 8471,
		},
	},
	{
		"minecraft:red_stained_glass",
		nil,
		NewMapping{
			393: 3591,
			404: 3592,
			477: 4095,
			573: 4095,
			4:   1534,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 6772,
			404: 6773,
			477: 7279,
			573: 7279,
		},
	},
	{
		"minecraft:birch_door[facing=south,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7769,
			404: 7770,
			477: 8294,
			573: 8294,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=21,powered=false]",
		nil,
		NewMapping{
			393: 391,
			404: 391,
			477: 391,
			573: 391,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=6,south=up,west=none]",
		nil,
		NewMapping{
			393: 2672,
			404: 2673,
			477: 2976,
			573: 2976,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=3,powered=false]",
		nil,
		NewMapping{
			393: 605,
			404: 605,
			477: 605,
			573: 605,
		},
	},
	{
		"minecraft:player_head[rotation=8]",
		nil,
		NewMapping{
			404: 5520,
			477: 6022,
			573: 6022,
			393: 5519,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=1,powered=true]",
		nil,
		NewMapping{
			393: 400,
			404: 400,
			477: 400,
			573: 400,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 10899,
			477: 10899,
		},
	},
	{
		"minecraft:lever[face=wall,facing=north,powered=true]",
		nil,
		NewMapping{
			573: 3789,
			4:   1116,
			393: 3285,
			404: 3286,
			477: 3789,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=13,south=side,west=up]",
		nil,
		NewMapping{
			393: 2880,
			404: 2881,
			477: 3184,
			573: 3184,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=false,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 4041,
			404: 4042,
			477: 4545,
			573: 4545,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=11,south=side,west=none]",
		nil,
		NewMapping{
			477: 3168,
			573: 3168,
			393: 2864,
			404: 2865,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10711,
			573: 10711,
		},
	},
	{
		"minecraft:acacia_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   2611,
			393: 6343,
			404: 6344,
			477: 6850,
			573: 6850,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=15,south=side,west=side]",
		nil,
		NewMapping{
			393: 2755,
			404: 2756,
			477: 3059,
			573: 3059,
		},
	},
	{
		"minecraft:red_sandstone_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			393: 7341,
			404: 7342,
			477: 7860,
			573: 7860,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=2,south=side,west=up]",
		nil,
		NewMapping{
			393: 2493,
			404: 2494,
			477: 2797,
			573: 2797,
		},
	},
	{
		"minecraft:composter[level=2]",
		nil,
		NewMapping{
			477: 11264,
			573: 11280,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=22,powered=false]",
		nil,
		NewMapping{
			393: 543,
			404: 543,
			477: 543,
			573: 543,
		},
	},
	{
		"minecraft:brown_banner[rotation=8]",
		nil,
		NewMapping{
			393: 7054,
			404: 7055,
			477: 7561,
			573: 7561,
		},
	},
	{
		"minecraft:lime_wall_banner[facing=east]",
		nil,
		NewMapping{
			477: 7640,
			573: 7640,
			393: 7133,
			404: 7134,
		},
	},
	{
		"minecraft:diorite_stairs[facing=south,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 10193,
			573: 10193,
		},
	},
	{
		"minecraft:kelp[age=8]",
		nil,
		NewMapping{
			477: 8942,
			573: 8942,
			393: 8417,
			404: 8418,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=14,south=up,west=none]",
		nil,
		NewMapping{
			393: 2024,
			404: 2025,
			477: 2328,
			573: 2328,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=false,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5263,
			404: 5264,
			477: 5767,
			573: 5767,
		},
	},
	{
		"minecraft:oak_button[face=ceiling,facing=north,powered=false]",
		nil,
		NewMapping{
			4:   2288,
			393: 5320,
			404: 5321,
			477: 5827,
			573: 5827,
		},
	},
	{
		"minecraft:sign[rotation=15,waterlogged=false]",
		nil,
		NewMapping{
			4:   1023,
			393: 3106,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=11,south=up,west=up]",
		nil,
		NewMapping{
			404: 1852,
			477: 2155,
			573: 2155,
			393: 1851,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 3211,
			404: 3212,
			477: 3675,
			573: 3675,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=west,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3830,
			404: 3831,
			477: 4334,
			573: 4334,
		},
	},
	{
		"minecraft:oak_sign[rotation=12,waterlogged=false]",
		nil,
		NewMapping{
			404: 3101,
			477: 3404,
			573: 3404,
		},
	},
	{
		"minecraft:dark_oak_fence[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 7674,
			404: 7675,
			477: 8199,
			573: 8199,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=north,half=top,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			573: 4290,
			393: 3786,
			404: 3787,
			477: 4290,
		},
	},
	{
		"minecraft:red_banner[rotation=5]",
		nil,
		NewMapping{
			393: 7083,
			404: 7084,
			477: 7590,
			573: 7590,
		},
	},
	{
		"minecraft:granite_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9928,
			573: 9928,
		},
	},
	{
		"minecraft:acacia_sign[rotation=7,waterlogged=false]",
		nil,
		NewMapping{
			477: 3490,
			573: 3490,
		},
	},
	{
		"minecraft:blue_shulker_box[facing=up]",
		nil,
		NewMapping{
			404: 8288,
			477: 8812,
			573: 8812,
			4:   3681,
			393: 8287,
		},
	},
	{
		"minecraft:acacia_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6400,
			404: 6401,
			477: 6907,
			573: 6907,
		},
	},
	{
		"minecraft:birch_door[facing=south,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			404: 7762,
			477: 8286,
			573: 8286,
			393: 7761,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=true,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 10345,
			477: 10345,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 10037,
			477: 10037,
		},
	},
	{
		"minecraft:diorite_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			477: 10330,
			573: 10330,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9802,
			573: 9802,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10784,
			573: 10784,
		},
	},
	{
		"minecraft:repeater[delay=1,facing=west,locked=true,powered=false]",
		nil,
		NewMapping{
			393: 3522,
			404: 3523,
			477: 4026,
			573: 4026,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 4031,
			404: 4032,
			477: 4535,
			573: 4535,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 6974,
			573: 6974,
			393: 6467,
			404: 6468,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=7,south=side,west=up]",
		nil,
		NewMapping{
			573: 2410,
			393: 2106,
			404: 2107,
			477: 2410,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 11044,
			573: 11044,
		},
	},
	{
		"minecraft:dead_tube_coral_wall_fan[facing=north,waterlogged=true]",
		nil,
		NewMapping{
			404: 8480,
			477: 9024,
			573: 9024,
			393: 8464,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=8,south=none,west=up]",
		nil,
		NewMapping{
			404: 2839,
			477: 3142,
			573: 3142,
			393: 2838,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10449,
			573: 10449,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			4:   1604,
			393: 4111,
			404: 4112,
			477: 4615,
			573: 4615,
		},
	},
	{
		"minecraft:repeater[delay=1,facing=north,locked=false,powered=false]",
		nil,
		NewMapping{
			393: 3516,
			404: 3517,
			477: 4020,
			573: 4020,
			4:   1490,
		},
	},
	{
		"minecraft:oak_fence_gate[facing=north,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			393: 4301,
			404: 4302,
			477: 4805,
			573: 4805,
		},
	},
	{
		"minecraft:spruce_door[facing=west,half=lower,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7717,
			404: 7718,
			477: 8242,
			573: 8242,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=true,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 6784,
			573: 6784,
			393: 6277,
			404: 6278,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=false,north=false,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10585,
			573: 10585,
		},
	},
	{
		"minecraft:sponge",
		nil,
		NewMapping{
			4:   304,
			393: 228,
			404: 228,
			477: 228,
			573: 228,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9544,
			573: 9544,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=0,south=none,west=none]",
		nil,
		NewMapping{
			393: 2768,
			404: 2769,
			477: 3072,
			573: 3072,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=6,south=up,west=side]",
		nil,
		NewMapping{
			477: 3263,
			573: 3263,
			393: 2959,
			404: 2960,
		},
	},
	{
		"minecraft:scaffolding[bottom=false,distance=5,waterlogged=true]",
		nil,
		NewMapping{
			477: 11125,
			573: 11125,
		},
	},
	{
		"minecraft:diorite_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10181,
			573: 10181,
		},
	},
	{
		"minecraft:birch_leaves[distance=4,persistent=false]",
		nil,
		NewMapping{
			393: 179,
			404: 179,
			477: 179,
			573: 179,
		},
	},
	{
		"minecraft:fire[age=13,east=false,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			477: 1875,
			573: 1875,
			393: 1571,
			404: 1572,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 6954,
			393: 6447,
			404: 6448,
			477: 6954,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=south,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			404: 3812,
			477: 4315,
			573: 4315,
			393: 3811,
		},
	},
	{
		"minecraft:dark_prismarine_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			393: 6816,
			404: 6817,
			477: 7323,
			573: 7323,
		},
	},
	{
		"minecraft:bubble_coral_wall_fan[facing=west,waterlogged=false]",
		nil,
		NewMapping{
			477: 9085,
			573: 9085,
			393: 8525,
			404: 8541,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=14,south=side,west=up]",
		nil,
		NewMapping{
			573: 2329,
			393: 2025,
			404: 2026,
			477: 2329,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 5716,
			573: 5716,
			393: 5212,
			404: 5213,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=false,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10873,
			573: 10873,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 11040,
			573: 11040,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 11051,
			573: 11051,
		},
	},
	{
		"minecraft:repeater[delay=1,facing=north,locked=true,powered=true]",
		nil,
		NewMapping{
			393: 3513,
			404: 3514,
			477: 4017,
			573: 4017,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 4079,
			404: 4080,
			477: 4583,
			573: 4583,
		},
	},
	{
		"minecraft:jungle_button[face=floor,facing=west,powered=false]",
		nil,
		NewMapping{
			573: 5887,
			393: 5380,
			404: 5381,
			477: 5887,
		},
	},
	{
		"minecraft:cobblestone_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5157,
			404: 5158,
			477: 5661,
			573: 5661,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=12,powered=true]",
		nil,
		NewMapping{
			573: 322,
			393: 322,
			404: 322,
			477: 322,
		},
	},
	{
		"minecraft:iron_door[facing=east,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			4:   1136,
			393: 3366,
			404: 3367,
			477: 3870,
			573: 3870,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=14,powered=false]",
		nil,
		NewMapping{
			477: 627,
			573: 627,
			393: 627,
			404: 627,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 8036,
			404: 8037,
			477: 8561,
			573: 8561,
		},
	},
	{
		"minecraft:oak_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 1675,
			404: 1676,
			477: 1979,
			573: 1979,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=23,powered=true]",
		nil,
		NewMapping{
			477: 994,
			573: 994,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9424,
			573: 9424,
		},
	},
	{
		"minecraft:observer[facing=south,powered=true]",
		nil,
		NewMapping{
			393: 8203,
			404: 8204,
			477: 8728,
			573: 8728,
			4:   3499,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 5693,
			393: 5189,
			404: 5190,
			477: 5693,
		},
	},
	{
		"minecraft:acacia_door[facing=west,half=upper,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7901,
			404: 7902,
			477: 8426,
			573: 8426,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 6561,
			404: 6562,
			477: 7068,
			573: 7068,
		},
	},
	{
		"minecraft:stone_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 9639,
			477: 9639,
		},
	},
	{
		"minecraft:yellow_banner[rotation=14]",
		nil,
		NewMapping{
			477: 7439,
			573: 7439,
			393: 6932,
			404: 6933,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=0,south=up,west=none]",
		nil,
		NewMapping{
			404: 2619,
			477: 2922,
			573: 2922,
			393: 2618,
		},
	},
	{
		"minecraft:jungle_pressure_plate[powered=false]",
		nil,
		NewMapping{
			393: 3374,
			404: 3375,
			477: 3878,
			573: 3878,
		},
	},
	{
		"minecraft:cactus[age=10]",
		nil,
		NewMapping{
			404: 3436,
			477: 3939,
			573: 3939,
			4:   1306,
			393: 3435,
		},
	},
	{
		"minecraft:wither_skeleton_skull[rotation=7]",
		nil,
		NewMapping{
			393: 5478,
			404: 5479,
			477: 5981,
			573: 5981,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=14,south=side,west=side]",
		nil,
		NewMapping{
			477: 3338,
			573: 3338,
			393: 3034,
			404: 3035,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=12,south=up,west=none]",
		nil,
		NewMapping{
			393: 2006,
			404: 2007,
			477: 2310,
			573: 2310,
		},
	},
	{
		"minecraft:jungle_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 5050,
			404: 5051,
			477: 5554,
			573: 5554,
		},
	},
	{
		"minecraft:kelp[age=18]",
		nil,
		NewMapping{
			573: 8952,
			393: 8427,
			404: 8428,
			477: 8952,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 9157,
			477: 9157,
		},
	},
	{
		"minecraft:diorite_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10228,
			573: 10228,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10335,
			573: 10335,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=23,powered=false]",
		nil,
		NewMapping{
			393: 495,
			404: 495,
			477: 495,
			573: 495,
		},
	},
	{
		"minecraft:piston_head[facing=north,short=true,type=sticky]",
		nil,
		NewMapping{
			393: 1060,
			404: 1060,
			477: 1360,
			573: 1360,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9384,
			573: 9384,
		},
	},
	{
		"minecraft:spruce_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4919,
			404: 4920,
			477: 5423,
			573: 5423,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=1,south=none,west=side]",
		nil,
		NewMapping{
			573: 2360,
			393: 2056,
			404: 2057,
			477: 2360,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=east,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			477: 4219,
			573: 4219,
			393: 3715,
			404: 3716,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=10,south=side,west=side]",
		nil,
		NewMapping{
			477: 3158,
			573: 3158,
			393: 2854,
			404: 2855,
		},
	},
	{
		"minecraft:fire[age=11,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1499,
			404: 1500,
			477: 1803,
			573: 1803,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 6466,
			393: 5959,
			404: 5960,
			477: 6466,
		},
	},
	{
		"minecraft:void_air",
		nil,
		NewMapping{
			393: 8574,
			404: 8591,
			477: 9129,
			573: 9129,
		},
	},
	{
		"minecraft:ender_chest[facing=east,waterlogged=true]",
		nil,
		NewMapping{
			393: 4737,
			404: 4738,
			477: 5241,
			573: 5241,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=16,powered=true]",
		nil,
		NewMapping{
			477: 830,
			573: 830,
		},
	},
	{
		"minecraft:jigsaw[facing=east]",
		nil,
		NewMapping{
			477: 11257,
			573: 11273,
		},
	},
	{
		"minecraft:tall_grass[half=lower]",
		nil,
		NewMapping{
			4:   2802,
			393: 6851,
			404: 6852,
			477: 7358,
			573: 7358,
		},
	},
	{
		"minecraft:farmland[moisture=4]",
		nil,
		NewMapping{
			573: 3367,
			4:   964,
			393: 3063,
			404: 3064,
			477: 3367,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=13,south=up,west=none]",
		nil,
		NewMapping{
			393: 2735,
			404: 2736,
			477: 3039,
			573: 3039,
		},
	},
	{
		"minecraft:quartz_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 5715,
			404: 5716,
			477: 6222,
			573: 6222,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9354,
			573: 9354,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=0,south=up,west=up]",
		nil,
		NewMapping{
			393: 2328,
			404: 2329,
			477: 2632,
			573: 2632,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=13,south=side,west=none]",
		nil,
		NewMapping{
			404: 2163,
			477: 2466,
			573: 2466,
			393: 2162,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9797,
			573: 9797,
		},
	},
	{
		"minecraft:dispenser[facing=east,triggered=false]",
		nil,
		NewMapping{
			4:   373,
			393: 236,
			404: 236,
			477: 236,
			573: 236,
		},
	},
	{
		"minecraft:repeater[delay=2,facing=east,locked=false,powered=true]",
		nil,
		NewMapping{
			4:   1511,
			393: 3543,
			404: 3544,
			477: 4047,
			573: 4047,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=west,half=top,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			573: 4135,
			393: 3631,
			404: 3632,
			477: 4135,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=true,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 8040,
			404: 8041,
			477: 8565,
			573: 8565,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=true,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			404: 8045,
			477: 8569,
			573: 8569,
			393: 8044,
		},
	},
	{
		"minecraft:birch_door[facing=north,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			477: 8268,
			573: 8268,
			393: 7743,
			404: 7744,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=false,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6170,
			404: 6171,
			477: 6677,
			573: 6677,
		},
	},
	{
		"minecraft:cocoa[age=0,facing=south]",
		nil,
		NewMapping{
			393: 4639,
			404: 4640,
			477: 5143,
			573: 5143,
			4:   2032,
		},
	},
	{
		"minecraft:farmland[moisture=5]",
		nil,
		NewMapping{
			404: 3065,
			477: 3368,
			573: 3368,
			4:   965,
			393: 3064,
		},
	},
	{
		"minecraft:sticky_piston[extended=true,facing=up]",
		nil,
		NewMapping{
			4:   473,
			393: 1032,
			404: 1032,
			477: 1332,
			573: 1332,
		},
	},
	{
		"minecraft:yellow_bed[facing=west,occupied=false,part=foot]",
		nil,
		NewMapping{
			404: 823,
			477: 1123,
			573: 1123,
			393: 823,
		},
	},
	{
		"minecraft:birch_stairs[facing=east,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 5033,
			404: 5034,
			477: 5537,
			573: 5537,
		},
	},
	{
		"minecraft:light_blue_bed[facing=west,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 806,
			404: 806,
			477: 1106,
			573: 1106,
		},
	},
	{
		"minecraft:dark_oak_fence_gate[facing=south,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7494,
			404: 7495,
			477: 8019,
			573: 8019,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6210,
			404: 6211,
			477: 6717,
			573: 6717,
		},
	},
	{
		"minecraft:clay",
		nil,
		NewMapping{
			393: 3441,
			404: 3442,
			477: 3945,
			573: 3945,
			4:   1312,
		},
	},
	{
		"minecraft:powered_rail[powered=true,shape=north_south]",
		nil,
		NewMapping{
			573: 1304,
			4:   440,
			393: 1004,
			404: 1004,
			477: 1304,
		},
	},
	{
		"minecraft:white_banner[rotation=0]",
		nil,
		NewMapping{
			4:   2816,
			393: 6854,
			404: 6855,
			477: 7361,
			573: 7361,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=6,powered=false]",
		nil,
		NewMapping{
			477: 261,
			573: 261,
			393: 261,
			404: 261,
		},
	},
	{
		"minecraft:dark_oak_door[facing=east,half=lower,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7989,
			404: 7990,
			477: 8514,
			573: 8514,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=true,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5240,
			404: 5241,
			477: 5744,
			573: 5744,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=0,powered=true]",
		nil,
		NewMapping{
			477: 898,
			573: 898,
		},
	},
	{
		"minecraft:polished_andesite_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			477: 10323,
			573: 10323,
		},
	},
	{
		"minecraft:smooth_quartz_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			477: 10298,
			573: 10298,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=14,powered=true]",
		nil,
		NewMapping{
			573: 726,
			393: 726,
			404: 726,
			477: 726,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=11,south=side,west=up]",
		nil,
		NewMapping{
			393: 2862,
			404: 2863,
			477: 3166,
			573: 3166,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=6,waterlogged=false]",
		nil,
		NewMapping{
			477: 3552,
			573: 3552,
		},
	},
	{
		"minecraft:cake[bites=1]",
		nil,
		NewMapping{
			4:   1473,
			393: 3507,
			404: 3508,
			477: 4011,
			573: 4011,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   1827,
			393: 4539,
			404: 4540,
			477: 5043,
			573: 5043,
		},
	},
	{
		"minecraft:spruce_fence[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 8049,
			393: 7524,
			404: 7525,
			477: 8049,
		},
	},
	{
		"minecraft:gray_wall_banner[facing=west]",
		nil,
		NewMapping{
			393: 7140,
			404: 7141,
			477: 7647,
			573: 7647,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=8,south=up,west=side]",
		nil,
		NewMapping{
			393: 2257,
			404: 2258,
			477: 2561,
			573: 2561,
		},
	},
	{
		"minecraft:tripwire_hook[attached=false,facing=west,powered=true]",
		nil,
		NewMapping{
			4:   2105,
			393: 4751,
			404: 4752,
			477: 5255,
			573: 5255,
		},
	},
	{
		"minecraft:chiseled_sandstone",
		nil,
		NewMapping{
			4:   385,
			393: 246,
			404: 246,
			477: 246,
			573: 246,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10757,
			573: 10757,
		},
	},
	{
		"minecraft:birch_sign[rotation=6,waterlogged=true]",
		nil,
		NewMapping{
			477: 3455,
			573: 3455,
		},
	},
	{
		"minecraft:pumpkin_stem[age=5]",
		nil,
		NewMapping{
			477: 4761,
			573: 4761,
			4:   1669,
			393: 4257,
			404: 4258,
		},
	},
	{
		"minecraft:piston[extended=false,facing=up]",
		nil,
		NewMapping{
			477: 1357,
			573: 1357,
			4:   529,
			393: 1057,
			404: 1057,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=3,south=none,west=none]",
		nil,
		NewMapping{
			393: 2075,
			404: 2076,
			477: 2379,
			573: 2379,
		},
	},
	{
		"minecraft:chorus_plant[down=true,east=true,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			573: 8533,
			393: 8008,
			404: 8009,
			477: 8533,
		},
	},
	{
		"minecraft:green_bed[facing=south,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 963,
			404: 963,
			477: 1263,
			573: 1263,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9699,
			573: 9699,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9266,
			477: 9266,
		},
	},
	{
		"minecraft:purple_banner[rotation=5]",
		nil,
		NewMapping{
			404: 7020,
			477: 7526,
			573: 7526,
			393: 7019,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 6472,
			573: 6472,
			393: 5965,
			404: 5966,
		},
	},
	{
		"minecraft:lever[face=ceiling,facing=south,powered=true]",
		nil,
		NewMapping{
			477: 3799,
			573: 3799,
			393: 3295,
			404: 3296,
		},
	},
	{
		"minecraft:fire[age=0,east=false,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			573: 1460,
			393: 1156,
			404: 1157,
			477: 1460,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=10,powered=true]",
		nil,
		NewMapping{
			477: 818,
			573: 818,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9489,
			573: 9489,
		},
	},
	{
		"minecraft:water[level=1]",
		nil,
		NewMapping{
			4:   145,
			393: 35,
			404: 35,
			477: 35,
			573: 35,
		},
	},
	{
		"minecraft:orange_banner[rotation=2]",
		nil,
		NewMapping{
			393: 6872,
			404: 6873,
			477: 7379,
			573: 7379,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=8,powered=false]",
		nil,
		NewMapping{
			393: 615,
			404: 615,
			477: 615,
			573: 615,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=true,north=false,powered=false,south=false,west=true]",
		nil,
		NewMapping{
			573: 5369,
			393: 4865,
			404: 4866,
			477: 5369,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6048,
			404: 6049,
			477: 6555,
			573: 6555,
		},
	},
	{
		"minecraft:acacia_door[facing=north,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7875,
			404: 7876,
			477: 8400,
			573: 8400,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=2,south=none,west=none]",
		nil,
		NewMapping{
			404: 1779,
			477: 2082,
			573: 2082,
			393: 1778,
		},
	},
	{
		"minecraft:water[level=2]",
		nil,
		NewMapping{
			393: 36,
			404: 36,
			477: 36,
			573: 36,
			4:   130,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=11,south=up,west=up]",
		nil,
		NewMapping{
			404: 2860,
			477: 3163,
			573: 3163,
			393: 2859,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=9,south=side,west=side]",
		nil,
		NewMapping{
			393: 2557,
			404: 2558,
			477: 2861,
			573: 2861,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=7,south=side,west=up]",
		nil,
		NewMapping{
			573: 2698,
			393: 2394,
			404: 2395,
			477: 2698,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6757,
			404: 6758,
			477: 7264,
			573: 7264,
		},
	},
	{
		"minecraft:lava[level=12]",
		nil,
		NewMapping{
			477: 62,
			573: 62,
			4:   188,
			393: 62,
			404: 62,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=13,south=none,west=side]",
		nil,
		NewMapping{
			477: 2180,
			573: 2180,
			393: 1876,
			404: 1877,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4420,
			404: 4421,
			477: 4924,
			573: 4924,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9864,
			573: 9864,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9480,
			573: 9480,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 6705,
			573: 6705,
			393: 6198,
			404: 6199,
		},
	},
	{
		"minecraft:dark_oak_button[face=floor,facing=north,powered=true]",
		nil,
		NewMapping{
			393: 5423,
			404: 5424,
			477: 5930,
			573: 5930,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			404: 6696,
			477: 7202,
			573: 7202,
			393: 6695,
		},
	},
	{
		"minecraft:acacia_sign[rotation=6,waterlogged=false]",
		nil,
		NewMapping{
			477: 3488,
			573: 3488,
		},
	},
	{
		"minecraft:spruce_sign[rotation=4,waterlogged=false]",
		nil,
		NewMapping{
			477: 3420,
			573: 3420,
		},
	},
	{
		"minecraft:orange_banner[rotation=11]",
		nil,
		NewMapping{
			393: 6881,
			404: 6882,
			477: 7388,
			573: 7388,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=west,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7469,
			404: 7470,
			477: 7994,
			573: 7994,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=15,south=none,west=up]",
		nil,
		NewMapping{
			393: 2181,
			404: 2182,
			477: 2485,
			573: 2485,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6109,
			404: 6110,
			477: 6616,
			573: 6616,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=8,south=none,west=up]",
		nil,
		NewMapping{
			477: 2566,
			573: 2566,
			393: 2262,
			404: 2263,
		},
	},
	{
		"minecraft:acacia_sign[rotation=4,waterlogged=true]",
		nil,
		NewMapping{
			477: 3483,
			573: 3483,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9493,
			573: 9493,
		},
	},
	{
		"minecraft:birch_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 5012,
			404: 5013,
			477: 5516,
			573: 5516,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 3988,
			404: 3989,
			477: 4492,
			573: 4492,
		},
	},
	{
		"minecraft:skeleton_skull[rotation=9]",
		nil,
		NewMapping{
			477: 5963,
			573: 5963,
			393: 5460,
			404: 5461,
		},
	},
	{
		"minecraft:spruce_sign[rotation=13,waterlogged=true]",
		nil,
		NewMapping{
			477: 3437,
			573: 3437,
		},
	},
	{
		"minecraft:spruce_sign[rotation=7,waterlogged=true]",
		nil,
		NewMapping{
			477: 3425,
			573: 3425,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=true,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10729,
			573: 10729,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9764,
			573: 9764,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9420,
			573: 9420,
		},
	},
	{
		"minecraft:spruce_leaves[distance=1,persistent=false]",
		nil,
		NewMapping{
			404: 159,
			477: 159,
			573: 159,
			4:   289,
			393: 159,
		},
	},
	{
		"minecraft:poppy",
		nil,
		NewMapping{
			477: 1412,
			573: 1412,
			4:   608,
			393: 1112,
			404: 1112,
		},
	},
	{
		"minecraft:acacia_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			4:   2004,
			393: 7286,
			404: 7287,
			477: 7793,
			573: 7793,
		},
	},
	{
		"minecraft:red_banner[rotation=4]",
		nil,
		NewMapping{
			393: 7082,
			404: 7083,
			477: 7589,
			573: 7589,
		},
	},
	{
		"minecraft:fire[age=15,east=true,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1620,
			404: 1621,
			477: 1924,
			573: 1924,
		},
	},
	{
		"minecraft:purple_bed[facing=west,occupied=true,part=head]",
		nil,
		NewMapping{
			573: 1216,
			393: 916,
			404: 916,
			477: 1216,
		},
	},
	{
		"minecraft:dead_bubble_coral_wall_fan[facing=south,waterlogged=false]",
		nil,
		NewMapping{
			393: 8483,
			404: 8499,
			477: 9043,
			573: 9043,
		},
	},
	{
		"minecraft:gray_banner[rotation=8]",
		nil,
		NewMapping{
			477: 7481,
			573: 7481,
			393: 6974,
			404: 6975,
		},
	},
	{
		"minecraft:fire[age=15,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			4:   831,
			393: 1646,
			404: 1647,
			477: 1950,
			573: 1950,
		},
	},
	{
		"minecraft:dark_oak_fence_gate[facing=north,in_wall=false,open=false,powered=false]",
		nil,
		NewMapping{
			4:   2978,
			393: 7492,
			404: 7493,
			477: 8017,
			573: 8017,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=2,south=up,west=up]",
		nil,
		NewMapping{
			393: 2058,
			404: 2059,
			477: 2362,
			573: 2362,
		},
	},
	{
		"minecraft:purpur_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 8616,
			573: 8616,
			393: 8091,
			404: 8092,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 7713,
			573: 7713,
			393: 7206,
			404: 7207,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=8,south=none,west=up]",
		nil,
		NewMapping{
			573: 2854,
			393: 2550,
			404: 2551,
			477: 2854,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 4490,
			477: 4993,
			573: 4993,
			393: 4489,
		},
	},
	{
		"minecraft:magenta_bed[facing=west,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 789,
			404: 789,
			477: 1089,
			573: 1089,
		},
	},
	{
		"minecraft:purple_bed[facing=west,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 918,
			404: 918,
			477: 1218,
			573: 1218,
		},
	},
	{
		"minecraft:dead_brain_coral_wall_fan[facing=east,waterlogged=true]",
		nil,
		NewMapping{
			404: 8494,
			477: 9038,
			573: 9038,
			393: 8478,
		},
	},
	{
		"minecraft:powered_rail[powered=true,shape=ascending_west]",
		nil,
		NewMapping{
			477: 1307,
			573: 1307,
			4:   443,
			393: 1007,
			404: 1007,
		},
	},
	{
		"minecraft:acacia_door[facing=east,half=upper,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			573: 8445,
			4:   3144,
			393: 7920,
			404: 7921,
			477: 8445,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 6490,
			477: 6996,
			573: 6996,
			393: 6489,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 6955,
			393: 6448,
			404: 6449,
			477: 6955,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=true,north=true,powered=true,south=true,west=true]",
		nil,
		NewMapping{
			393: 4787,
			404: 4788,
			477: 5291,
			573: 5291,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=east,half=top,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			477: 4408,
			573: 4408,
			393: 3904,
			404: 3905,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9862,
			573: 9862,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 10233,
			573: 10233,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10501,
			573: 10501,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=false,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6066,
			404: 6067,
			477: 6573,
			573: 6573,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 11024,
			573: 11024,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9465,
			573: 9465,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=east,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 7297,
			573: 7297,
			393: 6790,
			404: 6791,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=north,in_wall=false,open=true,powered=true]",
		nil,
		NewMapping{
			404: 7362,
			477: 7886,
			573: 7886,
			4:   2942,
			393: 7361,
		},
	},
	{
		"minecraft:jungle_button[face=ceiling,facing=south,powered=true]",
		nil,
		NewMapping{
			393: 5393,
			404: 5394,
			477: 5900,
			573: 5900,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=10,south=none,west=up]",
		nil,
		NewMapping{
			393: 1848,
			404: 1849,
			477: 2152,
			573: 2152,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=13,south=up,west=none]",
		nil,
		NewMapping{
			393: 2015,
			404: 2016,
			477: 2319,
			573: 2319,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 3236,
			404: 3237,
			477: 3700,
			573: 3700,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=5,powered=false]",
		nil,
		NewMapping{
			573: 759,
			477: 759,
		},
	},
	{
		"minecraft:acacia_log[axis=x]",
		nil,
		NewMapping{
			393: 84,
			404: 84,
			477: 84,
			573: 84,
			4:   2596,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=3,south=up,west=side]",
		nil,
		NewMapping{
			393: 2788,
			404: 2789,
			477: 3092,
			573: 3092,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 4666,
			477: 5169,
			573: 5169,
			393: 4665,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=false,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 4100,
			404: 4101,
			477: 4604,
			573: 4604,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=11,powered=false]",
		nil,
		NewMapping{
			477: 821,
			573: 821,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=false,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			4:   1585,
			393: 4039,
			404: 4040,
			477: 4543,
			573: 4543,
		},
	},
	{
		"minecraft:oak_button[face=wall,facing=south,powered=true]",
		nil,
		NewMapping{
			4:   2299,
			393: 5313,
			404: 5314,
			477: 5820,
			573: 5820,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=0,south=side,west=side]",
		nil,
		NewMapping{
			573: 2204,
			393: 1900,
			404: 1901,
			477: 2204,
		},
	},
	{
		"minecraft:andesite_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9967,
			573: 9967,
		},
	},
	{
		"minecraft:brain_coral_fan[waterlogged=true]",
		nil,
		NewMapping{
			404: 8572,
			477: 9016,
			573: 9016,
			393: 8556,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=north,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6563,
			404: 6564,
			477: 7070,
			573: 7070,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=0,powered=true]",
		nil,
		NewMapping{
			477: 598,
			573: 598,
			393: 598,
			404: 598,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=15,south=none,west=side]",
		nil,
		NewMapping{
			404: 2183,
			477: 2486,
			573: 2486,
			393: 2182,
		},
	},
	{
		"minecraft:spruce_sign[rotation=0,waterlogged=false]",
		nil,
		NewMapping{
			477: 3412,
			573: 3412,
		},
	},
	{
		"minecraft:chain_command_block[conditional=false,facing=north]",
		nil,
		NewMapping{
			573: 8707,
			4:   3378,
			393: 8182,
			404: 8183,
			477: 8707,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5878,
			404: 5879,
			477: 6385,
			573: 6385,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 6441,
			573: 6441,
			393: 5934,
			404: 5935,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=0,south=up,west=up]",
		nil,
		NewMapping{
			573: 2776,
			393: 2472,
			404: 2473,
			477: 2776,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9378,
			573: 9378,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=true,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 4034,
			404: 4035,
			477: 4538,
			573: 4538,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=11,south=side,west=up]",
		nil,
		NewMapping{
			393: 1998,
			404: 1999,
			477: 2302,
			573: 2302,
		},
	},
	{
		"minecraft:andesite_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9952,
			573: 9952,
		},
	},
	{
		"minecraft:dark_oak_door[facing=west,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			4:   3154,
			393: 7980,
			404: 7981,
			477: 8505,
			573: 8505,
		},
	},
	{
		"minecraft:dark_oak_fence[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 7669,
			404: 7670,
			477: 8194,
			573: 8194,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10830,
			573: 10830,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=3,powered=true]",
		nil,
		NewMapping{
			477: 854,
			573: 854,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=east,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9193,
			573: 9193,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=5,powered=false]",
		nil,
		NewMapping{
			393: 559,
			404: 559,
			477: 559,
			573: 559,
		},
	},
	{
		"minecraft:acacia_door[facing=east,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			477: 8454,
			573: 8454,
			393: 7929,
			404: 7930,
		},
	},
	{
		"minecraft:quartz_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 5708,
			404: 5709,
			477: 6215,
			573: 6215,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=9,south=up,west=up]",
		nil,
		NewMapping{
			573: 3001,
			393: 2697,
			404: 2698,
			477: 3001,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=north,half=top,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 6495,
			404: 6496,
			477: 7002,
			573: 7002,
		},
	},
	{
		"minecraft:quartz_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 5714,
			404: 5715,
			477: 6221,
			573: 6221,
		},
	},
	{
		"minecraft:spruce_pressure_plate[powered=false]",
		nil,
		NewMapping{
			393: 3370,
			404: 3371,
			477: 3874,
			573: 3874,
		},
	},
	{
		"minecraft:oak_fence_gate[facing=east,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			393: 4324,
			404: 4325,
			477: 4828,
			573: 4828,
		},
	},
	{
		"minecraft:andesite_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9994,
			573: 9994,
		},
	},
	{
		"minecraft:beehive[facing=south,honey_level=4]",
		nil,
		NewMapping{
			573: 11321,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=20,powered=true]",
		nil,
		NewMapping{
			393: 438,
			404: 438,
			477: 438,
			573: 438,
		},
	},
	{
		"minecraft:purple_wall_banner[facing=east]",
		nil,
		NewMapping{
			393: 7153,
			404: 7154,
			477: 7660,
			573: 7660,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=15,south=side,west=none]",
		nil,
		NewMapping{
			393: 2180,
			404: 2181,
			477: 2484,
			573: 2484,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=13,powered=false]",
		nil,
		NewMapping{
			477: 875,
			573: 875,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=14,powered=false]",
		nil,
		NewMapping{
			573: 877,
			477: 877,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10395,
			573: 10395,
		},
	},
	{
		"minecraft:jungle_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			404: 7279,
			477: 7785,
			573: 7785,
			4:   2019,
			393: 7278,
		},
	},
	{
		"minecraft:light_blue_bed[facing=west,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 805,
			404: 805,
			477: 1105,
			573: 1105,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=18,powered=false]",
		nil,
		NewMapping{
			393: 285,
			404: 285,
			477: 285,
			573: 285,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=6,south=none,west=none]",
		nil,
		NewMapping{
			393: 2390,
			404: 2391,
			477: 2694,
			573: 2694,
		},
	},
	{
		"minecraft:purple_banner[rotation=0]",
		nil,
		NewMapping{
			393: 7014,
			404: 7015,
			477: 7521,
			573: 7521,
		},
	},
	{
		"minecraft:acacia_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 6880,
			4:   2613,
			393: 6373,
			404: 6374,
			477: 6880,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=7,powered=true]",
		nil,
		NewMapping{
			393: 362,
			404: 362,
			477: 362,
			573: 362,
		},
	},
	{
		"minecraft:granite_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9923,
			573: 9923,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4478,
			404: 4479,
			477: 4982,
			573: 4982,
		},
	},
	{
		"minecraft:end_rod[facing=south]",
		nil,
		NewMapping{
			573: 8524,
			4:   3171,
			393: 7999,
			404: 8000,
			477: 8524,
		},
	},
	{
		"minecraft:lime_shulker_box[facing=east]",
		nil,
		NewMapping{
			4:   3589,
			393: 8248,
			404: 8249,
			477: 8773,
			573: 8773,
		},
	},
	{
		"minecraft:birch_door[facing=south,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			477: 8293,
			573: 8293,
			393: 7768,
			404: 7769,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10088,
			573: 10088,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10829,
			573: 10829,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=north,in_wall=false,open=false,powered=false]",
		nil,
		NewMapping{
			4:   2994,
			393: 7460,
			404: 7461,
			477: 7985,
			573: 7985,
		},
	},
	{
		"minecraft:gray_glazed_terracotta[facing=east]",
		nil,
		NewMapping{
			477: 8869,
			573: 8869,
			4:   3875,
			393: 8344,
			404: 8345,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=true,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 8050,
			404: 8051,
			477: 8575,
			573: 8575,
		},
	},
	{
		"minecraft:spruce_fence[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 7534,
			404: 7535,
			477: 8059,
			573: 8059,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=false,north=true,powered=true,south=false,west=false]",
		nil,
		NewMapping{
			573: 5278,
			393: 4774,
			404: 4775,
			477: 5278,
		},
	},
	{
		"minecraft:red_wall_banner[facing=east]",
		nil,
		NewMapping{
			393: 7169,
			404: 7170,
			477: 7676,
			573: 7676,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5902,
			404: 5903,
			477: 6409,
			573: 6409,
		},
	},
	{
		"minecraft:granite_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9901,
			573: 9901,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9388,
			573: 9388,
		},
	},
	{
		"minecraft:player_head[rotation=14]",
		nil,
		NewMapping{
			393: 5525,
			404: 5526,
			477: 6028,
			573: 6028,
		},
	},
	{
		"minecraft:fire[age=12,east=false,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 1538,
			404: 1539,
			477: 1842,
			573: 1842,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=0,powered=false]",
		nil,
		NewMapping{
			404: 349,
			477: 349,
			573: 349,
			393: 349,
		},
	},
	{
		"minecraft:iron_bars[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 4708,
			573: 4708,
			393: 4204,
			404: 4205,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10961,
			573: 10961,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=east,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3784,
			404: 3785,
			477: 4288,
			573: 4288,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10621,
			573: 10621,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6738,
			404: 6739,
			477: 7245,
			573: 7245,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=false,north=false,powered=true,south=true,west=false]",
		nil,
		NewMapping{
			393: 4876,
			404: 4877,
			477: 5380,
			573: 5380,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=3,south=up,west=none]",
		nil,
		NewMapping{
			573: 2229,
			393: 1925,
			404: 1926,
			477: 2229,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 6442,
			573: 6442,
			393: 5935,
			404: 5936,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5210,
			404: 5211,
			477: 5714,
			573: 5714,
		},
	},
	{
		"minecraft:end_stone",
		nil,
		NewMapping{
			4:   1936,
			393: 4634,
			404: 4635,
			477: 5138,
			573: 5138,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=false,north=false,powered=false,south=true,west=true]",
		nil,
		NewMapping{
			393: 4847,
			404: 4848,
			477: 5351,
			573: 5351,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=3,south=none,west=none]",
		nil,
		NewMapping{
			393: 2795,
			404: 2796,
			477: 3099,
			573: 3099,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=1,south=none,west=none]",
		nil,
		NewMapping{
			393: 2201,
			404: 2202,
			477: 2505,
			573: 2505,
		},
	},
	{
		"minecraft:acacia_sign[rotation=15,waterlogged=true]",
		nil,
		NewMapping{
			477: 3505,
			573: 3505,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6765,
			404: 6766,
			477: 7272,
			573: 7272,
		},
	},
	{
		"minecraft:vine[east=false,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 4285,
			404: 4286,
			477: 4789,
			573: 4789,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6633,
			404: 6634,
			477: 7140,
			573: 7140,
		},
	},
	{
		"minecraft:brown_concrete",
		nil,
		NewMapping{
			573: 8914,
			4:   4028,
			393: 8389,
			404: 8390,
			477: 8914,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6236,
			404: 6237,
			477: 6743,
			573: 6743,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 7190,
			404: 7191,
			477: 7697,
			573: 7697,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 7241,
			404: 7242,
			477: 7748,
			573: 7748,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=4,powered=true]",
		nil,
		NewMapping{
			573: 656,
			393: 656,
			404: 656,
			477: 656,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5995,
			404: 5996,
			477: 6502,
			573: 6502,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=12,powered=false]",
		nil,
		NewMapping{
			477: 923,
			573: 923,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=3,waterlogged=true]",
		nil,
		NewMapping{
			477: 3545,
			573: 3545,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10732,
			573: 10732,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=north,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			477: 4240,
			573: 4240,
			393: 3736,
			404: 3737,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=west,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			404: 7439,
			477: 7963,
			573: 7963,
			393: 7438,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=16,powered=true]",
		nil,
		NewMapping{
			393: 380,
			404: 380,
			477: 380,
			573: 380,
		},
	},
	{
		"minecraft:smoker[facing=south,lit=false]",
		nil,
		NewMapping{
			477: 11150,
			573: 11150,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6597,
			404: 6598,
			477: 7104,
			573: 7104,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=4,powered=true]",
		nil,
		NewMapping{
			477: 1006,
			573: 1006,
		},
	},
	{
		"minecraft:nether_brick_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			4:   718,
			393: 7330,
			404: 7331,
			477: 7849,
			573: 7849,
		},
	},
	{
		"minecraft:light_weighted_pressure_plate[power=10]",
		nil,
		NewMapping{
			404: 5614,
			477: 6120,
			573: 6120,
			4:   2362,
			393: 5613,
		},
	},
	{
		"minecraft:command_block[conditional=true,facing=south]",
		nil,
		NewMapping{
			4:   2203,
			393: 5126,
			404: 5127,
			477: 5630,
			573: 5630,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=13,powered=false]",
		nil,
		NewMapping{
			393: 475,
			404: 475,
			477: 475,
			573: 475,
		},
	},
	{
		"minecraft:fire[age=2,east=true,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 1202,
			404: 1203,
			477: 1506,
			573: 1506,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=east,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3651,
			404: 3652,
			477: 4155,
			573: 4155,
		},
	},
	{
		"minecraft:birch_fence[east=true,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 7557,
			404: 7558,
			477: 8082,
			573: 8082,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=7,south=side,west=none]",
		nil,
		NewMapping{
			393: 2828,
			404: 2829,
			477: 3132,
			573: 3132,
		},
	},
	{
		"minecraft:wall_sign[facing=north,waterlogged=false]",
		nil,
		NewMapping{
			4:   1090,
			393: 3270,
		},
	},
	{
		"minecraft:black_banner[rotation=15]",
		nil,
		NewMapping{
			393: 7109,
			404: 7110,
			477: 7616,
			573: 7616,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=14,south=side,west=none]",
		nil,
		NewMapping{
			393: 1883,
			404: 1884,
			477: 2187,
			573: 2187,
		},
	},
	{
		"minecraft:acacia_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 6389,
			477: 6895,
			573: 6895,
			393: 6388,
		},
	},
	{
		"minecraft:birch_fence[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			404: 7550,
			477: 8074,
			573: 8074,
			393: 7549,
		},
	},
	{
		"minecraft:moving_piston[facing=down,type=normal]",
		nil,
		NewMapping{
			4:   576,
			393: 1109,
			404: 1109,
			477: 1409,
			573: 1409,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=10,powered=false]",
		nil,
		NewMapping{
			477: 419,
			573: 419,
			393: 419,
			404: 419,
		},
	},
	{
		"minecraft:campfire[facing=west,lit=false,signal_fire=true,waterlogged=true]",
		nil,
		NewMapping{
			477: 11236,
			573: 11252,
		},
	},
	{
		"minecraft:hopper[enabled=false,facing=east]",
		nil,
		NewMapping{
			4:   2477,
			393: 5694,
			404: 5695,
			477: 6201,
			573: 6201,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=11,south=side,west=side]",
		nil,
		NewMapping{
			393: 3007,
			404: 3008,
			477: 3311,
			573: 3311,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6017,
			404: 6018,
			477: 6524,
			573: 6524,
		},
	},
	{
		"minecraft:oak_stairs[facing=east,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 1717,
			404: 1718,
			477: 2021,
			573: 2021,
		},
	},
	{
		"minecraft:birch_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   2166,
			393: 4985,
			404: 4986,
			477: 5489,
			573: 5489,
		},
	},
	{
		"minecraft:jungle_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 5124,
			477: 5627,
			573: 5627,
			393: 5123,
		},
	},
	{
		"minecraft:lime_wall_banner[facing=south]",
		nil,
		NewMapping{
			404: 7132,
			477: 7638,
			573: 7638,
			393: 7131,
		},
	},
	{
		"minecraft:kelp[age=0]",
		nil,
		NewMapping{
			573: 8934,
			393: 8409,
			404: 8410,
			477: 8934,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=15,south=up,west=none]",
		nil,
		NewMapping{
			393: 2753,
			404: 2754,
			477: 3057,
			573: 3057,
		},
	},
	{
		"minecraft:iron_bars[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 4696,
			393: 4192,
			404: 4193,
			477: 4696,
		},
	},
	{
		"minecraft:jungle_sign[rotation=10,waterlogged=false]",
		nil,
		NewMapping{
			477: 3528,
			573: 3528,
		},
	},
	{
		"minecraft:lever[face=wall,facing=north,powered=false]",
		nil,
		NewMapping{
			573: 3790,
			4:   1108,
			393: 3286,
			404: 3287,
			477: 3790,
		},
	},
	{
		"minecraft:dark_oak_door[facing=east,half=upper,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7982,
			404: 7983,
			477: 8507,
			573: 8507,
		},
	},
	{
		"minecraft:iron_door[facing=south,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			573: 3825,
			393: 3321,
			404: 3322,
			477: 3825,
		},
	},
	{
		"minecraft:oak_door[facing=west,half=upper,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 3140,
			404: 3141,
			477: 3604,
			573: 3604,
		},
	},
	{
		"minecraft:fire_coral_wall_fan[facing=south,waterlogged=false]",
		nil,
		NewMapping{
			393: 8531,
			404: 8547,
			477: 9091,
			573: 9091,
		},
	},
	{
		"minecraft:repeater[delay=3,facing=north,locked=true,powered=false]",
		nil,
		NewMapping{
			404: 3547,
			477: 4050,
			573: 4050,
			393: 3546,
		},
	},
	{
		"minecraft:acacia_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 6340,
			477: 6846,
			573: 6846,
			393: 6339,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6303,
			404: 6304,
			477: 6810,
			573: 6810,
		},
	},
	{
		"minecraft:dead_fire_coral_wall_fan[facing=north,waterlogged=false]",
		nil,
		NewMapping{
			393: 8489,
			404: 8505,
			477: 9049,
			573: 9049,
		},
	},
	{
		"minecraft:acacia_wall_sign[facing=east,waterlogged=false]",
		nil,
		NewMapping{
			477: 3764,
			573: 3764,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6325,
			404: 6326,
			477: 6832,
			573: 6832,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=5,powered=false]",
		nil,
		NewMapping{
			477: 859,
			573: 859,
		},
	},
	{
		"minecraft:light_blue_banner[rotation=5]",
		nil,
		NewMapping{
			393: 6907,
			404: 6908,
			477: 7414,
			573: 7414,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9392,
			573: 9392,
		},
	},
	{
		"minecraft:repeater[delay=2,facing=north,locked=false,powered=false]",
		nil,
		NewMapping{
			393: 3532,
			404: 3533,
			477: 4036,
			573: 4036,
			4:   1494,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=7,south=none,west=side]",
		nil,
		NewMapping{
			393: 2398,
			404: 2399,
			477: 2702,
			573: 2702,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=north,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4530,
			404: 4531,
			477: 5034,
			573: 5034,
		},
	},
	{
		"minecraft:fire[age=13,east=false,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			477: 1877,
			573: 1877,
			393: 1573,
			404: 1574,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=2,powered=true]",
		nil,
		NewMapping{
			393: 502,
			404: 502,
			477: 502,
			573: 502,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=south,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			573: 4316,
			393: 3812,
			404: 3813,
			477: 4316,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6214,
			404: 6215,
			477: 6721,
			573: 6721,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=15,south=side,west=side]",
		nil,
		NewMapping{
			393: 2179,
			404: 2180,
			477: 2483,
			573: 2483,
		},
	},
	{
		"minecraft:observer[facing=west,powered=true]",
		nil,
		NewMapping{
			573: 8730,
			4:   3500,
			393: 8205,
			404: 8206,
			477: 8730,
		},
	},
	{
		"minecraft:rose_bush[half=upper]",
		nil,
		NewMapping{
			393: 6846,
			404: 6847,
			477: 7353,
			573: 7353,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=south,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			477: 7923,
			573: 7923,
			393: 7398,
			404: 7399,
		},
	},
	{
		"minecraft:iron_bars[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 4704,
			573: 4704,
			393: 4200,
			404: 4201,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=north,in_wall=true,open=false,powered=false]",
		nil,
		NewMapping{
			573: 7885,
			393: 7360,
			404: 7361,
			477: 7885,
		},
	},
	{
		"minecraft:fire[age=7,east=true,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1372,
			404: 1373,
			477: 1676,
			573: 1676,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=12,south=side,west=up]",
		nil,
		NewMapping{
			393: 2295,
			404: 2296,
			477: 2599,
			573: 2599,
		},
	},
	{
		"minecraft:brick_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 4879,
			573: 4879,
			393: 4375,
			404: 4376,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=9,south=side,west=side]",
		nil,
		NewMapping{
			573: 2141,
			393: 1837,
			404: 1838,
			477: 2141,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=5,south=up,west=up]",
		nil,
		NewMapping{
			477: 3109,
			573: 3109,
			393: 2805,
			404: 2806,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6025,
			404: 6026,
			477: 6532,
			573: 6532,
		},
	},
	{
		"minecraft:light_blue_shulker_box[facing=north]",
		nil,
		NewMapping{
			4:   3554,
			393: 8235,
			404: 8236,
			477: 8760,
			573: 8760,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   2883,
			393: 7188,
			404: 7189,
			477: 7695,
			573: 7695,
		},
	},
	{
		"minecraft:dead_bubble_coral_wall_fan[facing=east,waterlogged=false]",
		nil,
		NewMapping{
			573: 9047,
			393: 8487,
			404: 8503,
			477: 9047,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6759,
			404: 6760,
			477: 7266,
			573: 7266,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10707,
			573: 10707,
		},
	},
	{
		"minecraft:end_stone_brick_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			477: 10286,
			573: 10286,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			477: 4522,
			573: 4522,
			393: 4018,
			404: 4019,
		},
	},
	{
		"minecraft:magenta_banner[rotation=5]",
		nil,
		NewMapping{
			404: 6892,
			477: 7398,
			573: 7398,
			393: 6891,
		},
	},
	{
		"minecraft:birch_door[facing=west,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7784,
			404: 7785,
			477: 8309,
			573: 8309,
		},
	},
	{
		"minecraft:spruce_fence[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 7528,
			404: 7529,
			477: 8053,
			573: 8053,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=south,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9153,
			573: 9153,
		},
	},
	{
		"minecraft:oak_fence[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 3984,
			573: 3984,
			393: 3480,
			404: 3481,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=true,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5957,
			404: 5958,
			477: 6464,
			573: 6464,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4695,
			404: 4696,
			477: 5199,
			573: 5199,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=false,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6035,
			404: 6036,
			477: 6542,
			573: 6542,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10851,
			573: 10851,
		},
	},
	{
		"minecraft:chiseled_red_sandstone",
		nil,
		NewMapping{
			477: 7682,
			573: 7682,
			4:   2865,
			393: 7175,
			404: 7176,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 7089,
			573: 7089,
			393: 6582,
			404: 6583,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=14,south=up,west=up]",
		nil,
		NewMapping{
			393: 2742,
			404: 2743,
			477: 3046,
			573: 3046,
		},
	},
	{
		"minecraft:fire[age=9,east=true,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 1426,
			404: 1427,
			477: 1730,
			573: 1730,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10039,
			573: 10039,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=13,south=side,west=none]",
		nil,
		NewMapping{
			404: 2883,
			477: 3186,
			573: 3186,
			393: 2882,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=1,south=up,west=side]",
		nil,
		NewMapping{
			573: 3074,
			393: 2770,
			404: 2771,
			477: 3074,
		},
	},
	{
		"minecraft:bell[attachment=double_wall,facing=north]",
		nil,
		NewMapping{
			477: 11210,
		},
	},
	{
		"minecraft:andesite_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9972,
			573: 9972,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=10,south=none,west=up]",
		nil,
		NewMapping{
			393: 2136,
			404: 2137,
			477: 2440,
			573: 2440,
		},
	},
	{
		"minecraft:gray_bed[facing=south,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 864,
			404: 864,
			477: 1164,
			573: 1164,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=false,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6194,
			404: 6195,
			477: 6701,
			573: 6701,
		},
	},
	{
		"minecraft:pink_bed[facing=south,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 849,
			404: 849,
			477: 1149,
			573: 1149,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=true,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10492,
			573: 10492,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9133,
			573: 9133,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9430,
			573: 9430,
		},
	},
	{
		"minecraft:gray_concrete_powder",
		nil,
		NewMapping{
			4:   4039,
			393: 8400,
			404: 8401,
			477: 8925,
			573: 8925,
		},
	},
	{
		"minecraft:cobblestone_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			4:   691,
			393: 7316,
			404: 7317,
			477: 7835,
			573: 7835,
		},
	},
	{
		"minecraft:acacia_fence[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 7629,
			404: 7630,
			477: 8154,
			573: 8154,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=6,south=side,west=none]",
		nil,
		NewMapping{
			393: 2243,
			404: 2244,
			477: 2547,
			573: 2547,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=5,south=up,west=side]",
		nil,
		NewMapping{
			393: 2086,
			404: 2087,
			477: 2390,
			573: 2390,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=5,powered=false]",
		nil,
		NewMapping{
			573: 959,
			477: 959,
		},
	},
	{
		"minecraft:jungle_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5111,
			404: 5112,
			477: 5615,
			573: 5615,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=7,south=none,west=none]",
		nil,
		NewMapping{
			393: 1967,
			404: 1968,
			477: 2271,
			573: 2271,
		},
	},
	{
		"minecraft:dark_oak_door[facing=north,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			477: 8463,
			573: 8463,
			393: 7938,
			404: 7939,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=2,powered=false]",
		nil,
		NewMapping{
			573: 903,
			477: 903,
		},
	},
	{
		"minecraft:beehive[facing=west,honey_level=1]",
		nil,
		NewMapping{
			573: 11324,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10641,
			573: 10641,
		},
	},
	{
		"minecraft:diorite_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			477: 10327,
			573: 10327,
		},
	},
	{
		"minecraft:diorite_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10190,
			573: 10190,
		},
	},
	{
		"minecraft:purpur_block",
		nil,
		NewMapping{
			573: 8598,
			4:   3216,
			393: 8073,
			404: 8074,
			477: 8598,
		},
	},
	{
		"minecraft:repeater[delay=2,facing=east,locked=true,powered=true]",
		nil,
		NewMapping{
			393: 3541,
			404: 3542,
			477: 4045,
			573: 4045,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=14,powered=true]",
		nil,
		NewMapping{
			573: 276,
			393: 276,
			404: 276,
			477: 276,
		},
	},
	{
		"minecraft:fire[age=6,east=true,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1331,
			404: 1332,
			477: 1635,
			573: 1635,
		},
	},
	{
		"minecraft:jungle_button[face=wall,facing=north,powered=false]",
		nil,
		NewMapping{
			393: 5384,
			404: 5385,
			477: 5891,
			573: 5891,
		},
	},
	{
		"minecraft:stripped_acacia_log[axis=x]",
		nil,
		NewMapping{
			393: 99,
			404: 99,
			477: 99,
			573: 99,
		},
	},
	{
		"minecraft:chorus_plant[down=true,east=true,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 8013,
			404: 8014,
			477: 8538,
			573: 8538,
		},
	},
	{
		"minecraft:dark_oak_door[facing=north,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			477: 8462,
			573: 8462,
			393: 7937,
			404: 7938,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=south,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6440,
			404: 6441,
			477: 6947,
			573: 6947,
		},
	},
	{
		"minecraft:stone_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 9654,
			477: 9654,
		},
	},
	{
		"minecraft:blast_furnace[facing=west,lit=false]",
		nil,
		NewMapping{
			477: 11160,
			573: 11160,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=south,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9313,
			573: 9313,
		},
	},
	{
		"minecraft:repeater[delay=1,facing=west,locked=false,powered=false]",
		nil,
		NewMapping{
			573: 4028,
			4:   1489,
			393: 3524,
			404: 3525,
			477: 4028,
		},
	},
	{
		"minecraft:water[level=10]",
		nil,
		NewMapping{
			4:   138,
			393: 44,
			404: 44,
			477: 44,
			573: 44,
		},
	},
	{
		"minecraft:jungle_log[axis=z]",
		nil,
		NewMapping{
			573: 83,
			4:   283,
			393: 83,
			404: 83,
			477: 83,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6229,
			404: 6230,
			477: 6736,
			573: 6736,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=2,south=side,west=none]",
		nil,
		NewMapping{
			393: 2783,
			404: 2784,
			477: 3087,
			573: 3087,
		},
	},
	{
		"minecraft:orange_bed[facing=south,occupied=false,part=foot]",
		nil,
		NewMapping{
			404: 771,
			477: 1071,
			573: 1071,
			393: 771,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=true,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 5894,
			477: 6400,
			573: 6400,
			393: 5893,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=north,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			404: 3922,
			477: 4425,
			573: 4425,
			393: 3921,
		},
	},
	{
		"minecraft:lectern[facing=south,has_book=false,powered=false]",
		nil,
		NewMapping{
			477: 11184,
			573: 11184,
		},
	},
	{
		"minecraft:loom[facing=west]",
		nil,
		NewMapping{
			477: 11133,
			573: 11133,
		},
	},
	{
		"minecraft:white_stained_glass_pane[east=false,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 6358,
			573: 6358,
			4:   2560,
			393: 5851,
			404: 5852,
		},
	},
	{
		"minecraft:jungle_leaves[distance=2,persistent=false]",
		nil,
		NewMapping{
			573: 189,
			4:   299,
			393: 189,
			404: 189,
			477: 189,
		},
	},
	{
		"minecraft:allium",
		nil,
		NewMapping{
			4:   610,
			393: 1114,
			404: 1114,
			477: 1414,
			573: 1414,
		},
	},
	{
		"minecraft:jungle_door[facing=north,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7816,
			404: 7817,
			477: 8341,
			573: 8341,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=13,south=none,west=none]",
		nil,
		NewMapping{
			404: 2022,
			477: 2325,
			573: 2325,
			393: 2021,
		},
	},
	{
		"minecraft:skeleton_skull[rotation=10]",
		nil,
		NewMapping{
			404: 5462,
			477: 5964,
			573: 5964,
			393: 5461,
		},
	},
	{
		"minecraft:kelp[age=19]",
		nil,
		NewMapping{
			393: 8428,
			404: 8429,
			477: 8953,
			573: 8953,
		},
	},
	{
		"minecraft:blue_bed[facing=south,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 928,
			404: 928,
			477: 1228,
			573: 1228,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=24,powered=true]",
		nil,
		NewMapping{
			393: 496,
			404: 496,
			477: 496,
			573: 496,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6677,
			404: 6678,
			477: 7184,
			573: 7184,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 6825,
			393: 6318,
			404: 6319,
			477: 6825,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4454,
			404: 4455,
			477: 4958,
			573: 4958,
		},
	},
	{
		"minecraft:oak_door[facing=north,half=upper,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			573: 3574,
			393: 3110,
			404: 3111,
			477: 3574,
		},
	},
	{
		"minecraft:acacia_sign[rotation=8,waterlogged=true]",
		nil,
		NewMapping{
			573: 3491,
			477: 3491,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=south,in_wall=true,open=false,powered=false]",
		nil,
		NewMapping{
			404: 7369,
			477: 7893,
			573: 7893,
			393: 7368,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			404: 7188,
			477: 7694,
			573: 7694,
			393: 7187,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 7182,
			404: 7183,
			477: 7689,
			573: 7689,
		},
	},
	{
		"minecraft:fire[age=6,east=true,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			404: 1330,
			477: 1633,
			573: 1633,
			393: 1329,
		},
	},
	{
		"minecraft:birch_door[facing=west,half=upper,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			477: 8301,
			573: 8301,
			393: 7776,
			404: 7777,
		},
	},
	{
		"minecraft:oak_button[face=floor,facing=east,powered=false]",
		nil,
		NewMapping{
			404: 5311,
			477: 5817,
			573: 5817,
			393: 5310,
		},
	},
	{
		"minecraft:green_banner[rotation=9]",
		nil,
		NewMapping{
			573: 7578,
			393: 7071,
			404: 7072,
			477: 7578,
		},
	},
	{
		"minecraft:yellow_shulker_box[facing=north]",
		nil,
		NewMapping{
			4:   3570,
			393: 8241,
			404: 8242,
			477: 8766,
			573: 8766,
		},
	},
	{
		"minecraft:magenta_shulker_box[facing=west]",
		nil,
		NewMapping{
			573: 8757,
			4:   3540,
			393: 8232,
			404: 8233,
			477: 8757,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 6556,
			573: 6556,
			393: 6049,
			404: 6050,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=14,south=side,west=up]",
		nil,
		NewMapping{
			404: 2458,
			477: 2761,
			573: 2761,
			393: 2457,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=21,powered=true]",
		nil,
		NewMapping{
			573: 790,
			477: 790,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10600,
			573: 10600,
		},
	},
	{
		"minecraft:daylight_detector[inverted=false,power=1]",
		nil,
		NewMapping{
			4:   2417,
			393: 5668,
			404: 5669,
			477: 6175,
			573: 6175,
		},
	},
	{
		"minecraft:magenta_banner[rotation=9]",
		nil,
		NewMapping{
			393: 6895,
			404: 6896,
			477: 7402,
			573: 7402,
		},
	},
	{
		"minecraft:purpur_stairs[facing=north,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 8079,
			404: 8080,
			477: 8604,
			573: 8604,
		},
	},
	{
		"minecraft:light_gray_banner[rotation=9]",
		nil,
		NewMapping{
			573: 7498,
			393: 6991,
			404: 6992,
			477: 7498,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6424,
			404: 6425,
			477: 6931,
			573: 6931,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=true,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			4:   1587,
			393: 4024,
			404: 4025,
			477: 4528,
			573: 4528,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6604,
			404: 6605,
			477: 7111,
			573: 7111,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 4920,
			393: 4416,
			404: 4417,
			477: 4920,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=true,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10698,
			573: 10698,
		},
	},
	{
		"minecraft:smooth_quartz_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			477: 10296,
			573: 10296,
		},
	},
	{
		"minecraft:repeating_command_block[conditional=true,facing=up]",
		nil,
		NewMapping{
			4:   3369,
			393: 8168,
			404: 8169,
			477: 8693,
			573: 8693,
		},
	},
	{
		"minecraft:powered_rail[powered=true,shape=ascending_north]",
		nil,
		NewMapping{
			4:   444,
			393: 1008,
			404: 1008,
			477: 1308,
			573: 1308,
		},
	},
	{
		"minecraft:chorus_plant[down=true,east=true,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			573: 8535,
			393: 8010,
			404: 8011,
			477: 8535,
		},
	},
	{
		"minecraft:fire[age=13,east=true,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1555,
			404: 1556,
			477: 1859,
			573: 1859,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=0,south=up,west=side]",
		nil,
		NewMapping{
			393: 2761,
			404: 2762,
			477: 3065,
			573: 3065,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=true,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			477: 4600,
			573: 4600,
			4:   1606,
			393: 4096,
			404: 4097,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=0,south=side,west=side]",
		nil,
		NewMapping{
			393: 2044,
			404: 2045,
			477: 2348,
			573: 2348,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=12,south=up,west=up]",
		nil,
		NewMapping{
			393: 1860,
			404: 1861,
			477: 2164,
			573: 2164,
		},
	},
	{
		"minecraft:diorite_stairs[facing=south,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 10203,
			573: 10203,
		},
	},
	{
		"minecraft:oak_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			477: 7768,
			573: 7768,
			393: 7261,
			404: 7262,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=23,powered=true]",
		nil,
		NewMapping{
			393: 294,
			404: 294,
			477: 294,
			573: 294,
		},
	},
	{
		"minecraft:chorus_plant[down=true,east=false,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			573: 8548,
			393: 8023,
			404: 8024,
			477: 8548,
		},
	},
	{
		"minecraft:birch_sign[rotation=6,waterlogged=false]",
		nil,
		NewMapping{
			477: 3456,
			573: 3456,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=4,waterlogged=false]",
		nil,
		NewMapping{
			477: 3548,
			573: 3548,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=2,powered=false]",
		nil,
		NewMapping{
			477: 953,
			573: 953,
		},
	},
	{
		"minecraft:magenta_shulker_box[facing=south]",
		nil,
		NewMapping{
			393: 8231,
			404: 8232,
			477: 8756,
			573: 8756,
			4:   3539,
		},
	},
	{
		"minecraft:jungle_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5061,
			404: 5062,
			477: 5565,
			573: 5565,
		},
	},
	{
		"minecraft:oak_leaves[distance=3,persistent=true]",
		nil,
		NewMapping{
			573: 148,
			393: 148,
			404: 148,
			477: 148,
		},
	},
	{
		"minecraft:cobblestone_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5158,
			404: 5159,
			477: 5662,
			573: 5662,
		},
	},
	{
		"minecraft:jungle_sign[rotation=3,waterlogged=false]",
		nil,
		NewMapping{
			477: 3514,
			573: 3514,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 3714,
			4:   1076,
			393: 3250,
			404: 3251,
			477: 3714,
		},
	},
	{
		"minecraft:white_concrete",
		nil,
		NewMapping{
			4:   4016,
			393: 8377,
			404: 8378,
			477: 8902,
			573: 8902,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10610,
			573: 10610,
		},
	},
	{
		"minecraft:andesite_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9940,
			573: 9940,
		},
	},
	{
		"minecraft:stonecutter[facing=north]",
		nil,
		NewMapping{
			477: 11194,
			573: 11194,
		},
	},
	{
		"minecraft:comparator[facing=north,mode=compare,powered=true]",
		nil,
		NewMapping{
			404: 5636,
			477: 6142,
			573: 6142,
			4:   2410,
			393: 5635,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=12,south=up,west=side]",
		nil,
		NewMapping{
			393: 2437,
			404: 2438,
			477: 2741,
			573: 2741,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=11,south=none,west=none]",
		nil,
		NewMapping{
			404: 2436,
			477: 2739,
			573: 2739,
			393: 2435,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=0,south=none,west=side]",
		nil,
		NewMapping{
			573: 2063,
			393: 1759,
			404: 1760,
			477: 2063,
		},
	},
	{
		"minecraft:dark_oak_fence_gate[facing=east,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			477: 8035,
			573: 8035,
			393: 7510,
			404: 7511,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=9,south=side,west=side]",
		nil,
		NewMapping{
			573: 2429,
			393: 2125,
			404: 2126,
			477: 2429,
		},
	},
	{
		"minecraft:creeper_head[rotation=15]",
		nil,
		NewMapping{
			404: 5547,
			477: 6049,
			573: 6049,
			393: 5546,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 7220,
			404: 7221,
			477: 7727,
			573: 7727,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 9601,
			477: 9601,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9192,
			573: 9192,
		},
	},
	{
		"minecraft:brown_glazed_terracotta[facing=east]",
		nil,
		NewMapping{
			573: 8889,
			4:   3955,
			393: 8364,
			404: 8365,
			477: 8889,
		},
	},
	{
		"minecraft:dead_fire_coral_wall_fan[facing=west,waterlogged=true]",
		nil,
		NewMapping{
			393: 8492,
			404: 8508,
			477: 9052,
			573: 9052,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 6975,
			393: 6468,
			404: 6469,
			477: 6975,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=false,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 4077,
			404: 4078,
			477: 4581,
			573: 4581,
		},
	},
	{
		"minecraft:birch_door[facing=south,half=upper,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			477: 8283,
			573: 8283,
			393: 7758,
			404: 7759,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9784,
			573: 9784,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10451,
			573: 10451,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 7139,
			393: 6632,
			404: 6633,
			477: 7139,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=6,south=up,west=up]",
		nil,
		NewMapping{
			573: 2254,
			393: 1950,
			404: 1951,
			477: 2254,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=false,north=true,powered=false,south=false,west=true]",
		nil,
		NewMapping{
			393: 4841,
			404: 4842,
			477: 5345,
			573: 5345,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9425,
			573: 9425,
		},
	},
	{
		"minecraft:andesite_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 9975,
			477: 9975,
		},
	},
	{
		"minecraft:dark_oak_door[facing=south,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7964,
			404: 7965,
			477: 8489,
			573: 8489,
			4:   3153,
		},
	},
	{
		"minecraft:rail[shape=south_west]",
		nil,
		NewMapping{
			4:   1063,
			393: 3186,
			404: 3187,
			477: 3650,
			573: 3650,
		},
	},
	{
		"minecraft:lime_bed[facing=north,occupied=true,part=foot]",
		nil,
		NewMapping{
			573: 1129,
			393: 829,
			404: 829,
			477: 1129,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 3222,
			477: 3685,
			573: 3685,
			393: 3221,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=8,powered=false]",
		nil,
		NewMapping{
			477: 1015,
			573: 1015,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=7,waterlogged=true]",
		nil,
		NewMapping{
			477: 3553,
			573: 3553,
		},
	},
	{
		"minecraft:andesite_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9946,
			573: 9946,
		},
	},
	{
		"minecraft:beehive[facing=south,honey_level=0]",
		nil,
		NewMapping{
			573: 11317,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=12,south=up,west=up]",
		nil,
		NewMapping{
			393: 2292,
			404: 2293,
			477: 2596,
			573: 2596,
		},
	},
	{
		"minecraft:grass_block[snowy=true]",
		nil,
		NewMapping{
			477: 8,
			573: 8,
			393: 8,
			404: 8,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=1,powered=true]",
		nil,
		NewMapping{
			573: 500,
			393: 500,
			404: 500,
			477: 500,
		},
	},
	{
		"minecraft:white_glazed_terracotta[facing=north]",
		nil,
		NewMapping{
			4:   3762,
			393: 8313,
			404: 8314,
			477: 8838,
			573: 8838,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 6490,
			393: 5983,
			404: 5984,
			477: 6490,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 6465,
			477: 6971,
			573: 6971,
			393: 6464,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			404: 6313,
			477: 6819,
			573: 6819,
			393: 6312,
		},
	},
	{
		"minecraft:sandstone_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			393: 7299,
			404: 7300,
			477: 7812,
			573: 7812,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			404: 6153,
			477: 6659,
			573: 6659,
			393: 6152,
		},
	},
	{
		"minecraft:diorite_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10220,
			573: 10220,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=6,south=side,west=side]",
		nil,
		NewMapping{
			393: 2674,
			404: 2675,
			477: 2978,
			573: 2978,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=5,south=none,west=up]",
		nil,
		NewMapping{
			573: 2107,
			393: 1803,
			404: 1804,
			477: 2107,
		},
	},
	{
		"minecraft:brown_bed[facing=west,occupied=true,part=head]",
		nil,
		NewMapping{
			573: 1248,
			393: 948,
			404: 948,
			477: 1248,
		},
	},
	{
		"minecraft:quartz_stairs[facing=east,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 5762,
			404: 5763,
			477: 6269,
			573: 6269,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=true,north=true,powered=false,south=true,west=true]",
		nil,
		NewMapping{
			393: 4855,
			404: 4856,
			477: 5359,
			573: 5359,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6154,
			404: 6155,
			477: 6661,
			573: 6661,
		},
	},
	{
		"minecraft:creeper_head[rotation=0]",
		nil,
		NewMapping{
			393: 5531,
			404: 5532,
			477: 6034,
			573: 6034,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=3,south=none,west=none]",
		nil,
		NewMapping{
			393: 2219,
			404: 2220,
			477: 2523,
			573: 2523,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10884,
			573: 10884,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 4469,
			477: 4972,
			573: 4972,
			393: 4468,
		},
	},
	{
		"minecraft:andesite_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 9954,
			477: 9954,
		},
	},
	{
		"minecraft:stone_stairs[facing=south,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			573: 9633,
			477: 9633,
		},
	},
	{
		"minecraft:vine[east=true,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			4:   1705,
			393: 4279,
			404: 4280,
			477: 4783,
			573: 4783,
		},
	},
	{
		"minecraft:brick_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4378,
			404: 4379,
			477: 4882,
			573: 4882,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4561,
			404: 4562,
			477: 5065,
			573: 5065,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=true,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 4056,
			404: 4057,
			477: 4560,
			573: 4560,
		},
	},
	{
		"minecraft:oak_wood[axis=z]",
		nil,
		NewMapping{
			573: 110,
			393: 110,
			404: 110,
			477: 110,
		},
	},
	{
		"minecraft:furnace[facing=west,lit=false]",
		nil,
		NewMapping{
			573: 3376,
			4:   980,
			393: 3072,
			404: 3073,
			477: 3376,
		},
	},
	{
		"minecraft:piston_head[facing=south,short=false,type=normal]",
		nil,
		NewMapping{
			4:   547,
			393: 1069,
			404: 1069,
			477: 1369,
			573: 1369,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=17,powered=true]",
		nil,
		NewMapping{
			393: 482,
			404: 482,
			477: 482,
			573: 482,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6694,
			404: 6695,
			477: 7201,
			573: 7201,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9287,
			573: 9287,
		},
	},
	{
		"minecraft:vine[east=false,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			477: 4795,
			573: 4795,
			4:   1700,
			393: 4291,
			404: 4292,
		},
	},
	{
		"minecraft:player_head[rotation=6]",
		nil,
		NewMapping{
			393: 5517,
			404: 5518,
			477: 6020,
			573: 6020,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=7,south=none,west=up]",
		nil,
		NewMapping{
			573: 3277,
			393: 2973,
			404: 2974,
			477: 3277,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=false,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10841,
			573: 10841,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=true,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10953,
			573: 10953,
		},
	},
	{
		"minecraft:birch_sign[rotation=0,waterlogged=true]",
		nil,
		NewMapping{
			573: 3443,
			477: 3443,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=12,south=none,west=none]",
		nil,
		NewMapping{
			393: 3020,
			404: 3021,
			477: 3324,
			573: 3324,
			4:   892,
		},
	},
	{
		"minecraft:fire[age=11,east=false,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 1504,
			404: 1505,
			477: 1808,
			573: 1808,
		},
	},
	{
		"minecraft:dark_oak_door[facing=north,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			477: 8472,
			573: 8472,
			393: 7947,
			404: 7948,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 6934,
			393: 6427,
			404: 6428,
			477: 6934,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 11028,
			573: 11028,
		},
	},
	{
		"minecraft:beehive[facing=north,honey_level=3]",
		nil,
		NewMapping{
			573: 11314,
		},
	},
	{
		"minecraft:jungle_door[facing=west,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			573: 8375,
			4:   3126,
			393: 7850,
			404: 7851,
			477: 8375,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=east,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3710,
			404: 3711,
			477: 4214,
			573: 4214,
		},
	},
	{
		"minecraft:cyan_bed[facing=east,occupied=true,part=head]",
		nil,
		NewMapping{
			573: 1204,
			393: 904,
			404: 904,
			477: 1204,
		},
	},
	{
		"minecraft:glass_pane[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 4224,
			404: 4225,
			477: 4728,
			573: 4728,
		},
	},
	{
		"minecraft:stone_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9688,
			573: 9688,
		},
	},
	{
		"minecraft:fire[age=14,east=true,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1587,
			404: 1588,
			477: 1891,
			573: 1891,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10664,
			573: 10664,
		},
	},
	{
		"minecraft:cut_sandstone_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			477: 7821,
			573: 7821,
		},
	},
	{
		"minecraft:purpur_pillar[axis=x]",
		nil,
		NewMapping{
			393: 8074,
			404: 8075,
			477: 8599,
			573: 8599,
			4:   3236,
		},
	},
	{
		"minecraft:attached_melon_stem[facing=west]",
		nil,
		NewMapping{
			393: 4250,
			404: 4251,
			477: 4754,
			573: 4754,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 4949,
			573: 4949,
			393: 4445,
			404: 4446,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=6,south=up,west=none]",
		nil,
		NewMapping{
			393: 2384,
			404: 2385,
			477: 2688,
			573: 2688,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=1,south=none,west=up]",
		nil,
		NewMapping{
			573: 2791,
			393: 2487,
			404: 2488,
			477: 2791,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=south,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3869,
			404: 3870,
			477: 4373,
			573: 4373,
		},
	},
	{
		"minecraft:lime_shulker_box[facing=up]",
		nil,
		NewMapping{
			4:   3585,
			393: 8251,
			404: 8252,
			477: 8776,
			573: 8776,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 6735,
			573: 6735,
			393: 6228,
			404: 6229,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=11,south=side,west=up]",
		nil,
		NewMapping{
			393: 2718,
			404: 2719,
			477: 3022,
			573: 3022,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=12,powered=false]",
		nil,
		NewMapping{
			393: 623,
			404: 623,
			477: 623,
			573: 623,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 7251,
			404: 7252,
			477: 7758,
			573: 7758,
		},
	},
	{
		"minecraft:spruce_button[face=ceiling,facing=west,powered=false]",
		nil,
		NewMapping{
			393: 5348,
			404: 5349,
			477: 5855,
			573: 5855,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=15,south=up,west=none]",
		nil,
		NewMapping{
			393: 2897,
			404: 2898,
			477: 3201,
			573: 3201,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5950,
			404: 5951,
			477: 6457,
			573: 6457,
		},
	},
	{
		"minecraft:fire[age=9,east=true,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1438,
			404: 1439,
			477: 1742,
			573: 1742,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=15,south=side,west=up]",
		nil,
		NewMapping{
			573: 2338,
			393: 2034,
			404: 2035,
			477: 2338,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=north,half=top,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			573: 4418,
			393: 3914,
			404: 3915,
			477: 4418,
		},
	},
	{
		"minecraft:magma_block",
		nil,
		NewMapping{
			4:   3408,
			393: 8192,
			404: 8193,
			477: 8717,
			573: 8717,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 3226,
			404: 3227,
			477: 3690,
			573: 3690,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=12,powered=false]",
		nil,
		NewMapping{
			573: 323,
			393: 323,
			404: 323,
			477: 323,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=north,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3734,
			404: 3735,
			477: 4238,
			573: 4238,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6724,
			404: 6725,
			477: 7231,
			573: 7231,
		},
	},
	{
		"minecraft:scaffolding[bottom=false,distance=7,waterlogged=false]",
		nil,
		NewMapping{
			477: 11130,
			573: 11130,
		},
	},
	{
		"minecraft:bee_nest[facing=north,honey_level=3]",
		nil,
		NewMapping{
			573: 11290,
		},
	},
	{
		"minecraft:daylight_detector[inverted=false,power=10]",
		nil,
		NewMapping{
			4:   2426,
			393: 5677,
			404: 5678,
			477: 6184,
			573: 6184,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   2624,
			393: 6483,
			404: 6484,
			477: 6990,
			573: 6990,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=23,powered=false]",
		nil,
		NewMapping{
			393: 295,
			404: 295,
			477: 295,
			573: 295,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10771,
			573: 10771,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9226,
			573: 9226,
		},
	},
	{
		"minecraft:brick_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4398,
			404: 4399,
			477: 4902,
			573: 4902,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=true,north=false,powered=false,south=true,west=true]",
		nil,
		NewMapping{
			393: 4799,
			404: 4800,
			477: 5303,
			573: 5303,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			573: 9863,
			477: 9863,
		},
	},
	{
		"minecraft:vine[east=true,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			4:   1711,
			393: 4270,
			404: 4271,
			477: 4774,
			573: 4774,
		},
	},
	{
		"minecraft:cactus[age=11]",
		nil,
		NewMapping{
			573: 3940,
			4:   1307,
			393: 3436,
			404: 3437,
			477: 3940,
		},
	},
	{
		"minecraft:brown_glazed_terracotta[facing=west]",
		nil,
		NewMapping{
			4:   3953,
			393: 8363,
			404: 8364,
			477: 8888,
			573: 8888,
		},
	},
	{
		"minecraft:birch_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 5468,
			573: 5468,
			393: 4964,
			404: 4965,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=8,south=none,west=none]",
		nil,
		NewMapping{
			393: 2120,
			404: 2121,
			477: 2424,
			573: 2424,
		},
	},
	{
		"minecraft:lectern[facing=east,has_book=false,powered=false]",
		nil,
		NewMapping{
			477: 11192,
			573: 11192,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=20,powered=false]",
		nil,
		NewMapping{
			477: 939,
			573: 939,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10674,
			573: 10674,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9589,
			573: 9589,
		},
	},
	{
		"minecraft:spruce_sign[rotation=6,waterlogged=true]",
		nil,
		NewMapping{
			477: 3423,
			573: 3423,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9385,
			573: 9385,
		},
	},
	{
		"minecraft:cyan_glazed_terracotta[facing=north]",
		nil,
		NewMapping{
			477: 8874,
			573: 8874,
			4:   3906,
			393: 8349,
			404: 8350,
		},
	},
	{
		"minecraft:purpur_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 8096,
			404: 8097,
			477: 8621,
			573: 8621,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4570,
			404: 4571,
			477: 5074,
			573: 5074,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=south,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7429,
			404: 7430,
			477: 7954,
			573: 7954,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9751,
			573: 9751,
		},
	},
	{
		"minecraft:birch_sign[rotation=14,waterlogged=false]",
		nil,
		NewMapping{
			477: 3472,
			573: 3472,
		},
	},
	{
		"minecraft:sticky_piston[extended=false,facing=down]",
		nil,
		NewMapping{
			404: 1039,
			477: 1339,
			573: 1339,
			4:   464,
			393: 1039,
		},
	},
	{
		"minecraft:spruce_button[face=wall,facing=east,powered=false]",
		nil,
		NewMapping{
			393: 5342,
			404: 5343,
			477: 5849,
			573: 5849,
		},
	},
	{
		"minecraft:green_banner[rotation=2]",
		nil,
		NewMapping{
			393: 7064,
			404: 7065,
			477: 7571,
			573: 7571,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=22,powered=true]",
		nil,
		NewMapping{
			404: 642,
			477: 642,
			573: 642,
			393: 642,
		},
	},
	{
		"minecraft:birch_button[face=ceiling,facing=west,powered=false]",
		nil,
		NewMapping{
			393: 5372,
			404: 5373,
			477: 5879,
			573: 5879,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=2,south=up,west=side]",
		nil,
		NewMapping{
			393: 2923,
			404: 2924,
			477: 3227,
			573: 3227,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9720,
			573: 9720,
		},
	},
	{
		"minecraft:end_stone_brick_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			477: 10283,
			573: 10283,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=4,powered=false]",
		nil,
		NewMapping{
			477: 357,
			573: 357,
			393: 357,
			404: 357,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6275,
			404: 6276,
			477: 6782,
			573: 6782,
		},
	},
	{
		"minecraft:dark_oak_door[facing=south,half=lower,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7957,
			404: 7958,
			477: 8482,
			573: 8482,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6070,
			404: 6071,
			477: 6577,
			573: 6577,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 5056,
			573: 5056,
			393: 4552,
			404: 4553,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9736,
			573: 9736,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=true,north=false,powered=false,south=true,west=true]",
		nil,
		NewMapping{
			393: 4831,
			404: 4832,
			477: 5335,
			573: 5335,
		},
	},
	{
		"minecraft:oak_button[face=floor,facing=west,powered=true]",
		nil,
		NewMapping{
			404: 5308,
			477: 5814,
			573: 5814,
			393: 5307,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10354,
			573: 10354,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=7,powered=true]",
		nil,
		NewMapping{
			477: 962,
			573: 962,
		},
	},
	{
		"minecraft:dark_oak_door[facing=east,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7986,
			404: 7987,
			477: 8511,
			573: 8511,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=3,south=up,west=up]",
		nil,
		NewMapping{
			393: 2643,
			404: 2644,
			477: 2947,
			573: 2947,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=13,south=none,west=up]",
		nil,
		NewMapping{
			477: 3331,
			573: 3331,
			393: 3027,
			404: 3028,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=true,north=true,powered=true,south=false,west=true]",
		nil,
		NewMapping{
			393: 4821,
			404: 4822,
			477: 5325,
			573: 5325,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6614,
			404: 6615,
			477: 7121,
			573: 7121,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=18,powered=false]",
		nil,
		NewMapping{
			393: 385,
			404: 385,
			477: 385,
			573: 385,
		},
	},
	{
		"minecraft:spruce_door[facing=south,half=upper,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7693,
			404: 7694,
			477: 8218,
			573: 8218,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=10,south=none,west=none]",
		nil,
		NewMapping{
			393: 2858,
			404: 2859,
			477: 3162,
			573: 3162,
		},
	},
	{
		"minecraft:birch_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4971,
			404: 4972,
			477: 5475,
			573: 5475,
		},
	},
	{
		"minecraft:campfire[facing=north,lit=false,signal_fire=true,waterlogged=true]",
		nil,
		NewMapping{
			477: 11220,
			573: 11236,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9826,
			573: 9826,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=0,powered=true]",
		nil,
		NewMapping{
			477: 748,
			573: 748,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=12,powered=true]",
		nil,
		NewMapping{
			477: 922,
			573: 922,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 3218,
			404: 3219,
			477: 3682,
			573: 3682,
		},
	},
	{
		"minecraft:acacia_door[facing=east,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7926,
			404: 7927,
			477: 8451,
			573: 8451,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=9,south=side,west=up]",
		nil,
		NewMapping{
			404: 2269,
			477: 2572,
			573: 2572,
			393: 2268,
		},
	},
	{
		"minecraft:jungle_planks",
		nil,
		NewMapping{
			573: 18,
			4:   83,
			393: 18,
			404: 18,
			477: 18,
		},
	},
	{
		"minecraft:observer[facing=north,powered=false]",
		nil,
		NewMapping{
			573: 8725,
			4:   3490,
			393: 8200,
			404: 8201,
			477: 8725,
		},
	},
	{
		"minecraft:jungle_door[facing=north,half=upper,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7805,
			404: 7806,
			477: 8330,
			573: 8330,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=north,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 6496,
			404: 6497,
			477: 7003,
			573: 7003,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=20,powered=true]",
		nil,
		NewMapping{
			393: 738,
			404: 738,
			477: 738,
			573: 738,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=15,powered=false]",
		nil,
		NewMapping{
			393: 729,
			404: 729,
			477: 729,
			573: 729,
		},
	},
	{
		"minecraft:smoker[facing=south,lit=true]",
		nil,
		NewMapping{
			477: 11149,
			573: 11149,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9778,
			573: 9778,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10343,
			573: 10343,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=true,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 8048,
			404: 8049,
			477: 8573,
			573: 8573,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=true,north=true,powered=true,south=false,west=true]",
		nil,
		NewMapping{
			393: 4789,
			404: 4790,
			477: 5293,
			573: 5293,
		},
	},
	{
		"minecraft:fire[age=12,east=true,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1519,
			404: 1520,
			477: 1823,
			573: 1823,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 10839,
			477: 10839,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=south,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9721,
			573: 9721,
		},
	},
	{
		"minecraft:fire[age=5,east=true,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1300,
			404: 1301,
			477: 1604,
			573: 1604,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9146,
			477: 9146,
		},
	},
	{
		"minecraft:stone_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9634,
			573: 9634,
		},
	},
	{
		"minecraft:wheat[age=7]",
		nil,
		NewMapping{
			4:   951,
			393: 3058,
			404: 3059,
			477: 3362,
			573: 3362,
		},
	},
	{
		"minecraft:spruce_door[facing=west,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			4:   3094,
			393: 7722,
			404: 7723,
			477: 8247,
			573: 8247,
		},
	},
	{
		"minecraft:brown_bed[facing=north,occupied=false,part=foot]",
		nil,
		NewMapping{
			477: 1243,
			573: 1243,
			393: 943,
			404: 943,
		},
	},
	{
		"minecraft:oak_door[facing=west,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			404: 3146,
			477: 3609,
			573: 3609,
			393: 3145,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=8,south=side,west=none]",
		nil,
		NewMapping{
			393: 1973,
			404: 1974,
			477: 2277,
			573: 2277,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10444,
			573: 10444,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9703,
			573: 9703,
		},
	},
	{
		"minecraft:fire[age=1,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1179,
			404: 1180,
			477: 1483,
			573: 1483,
		},
	},
	{
		"minecraft:light_gray_banner[rotation=7]",
		nil,
		NewMapping{
			393: 6989,
			404: 6990,
			477: 7496,
			573: 7496,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4550,
			404: 4551,
			477: 5054,
			573: 5054,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=8,south=none,west=none]",
		nil,
		NewMapping{
			393: 2264,
			404: 2265,
			477: 2568,
			573: 2568,
		},
	},
	{
		"minecraft:red_banner[rotation=11]",
		nil,
		NewMapping{
			573: 7596,
			393: 7089,
			404: 7090,
			477: 7596,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9540,
			573: 9540,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 5168,
			573: 5168,
			393: 4664,
			404: 4665,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6268,
			404: 6269,
			477: 6775,
			573: 6775,
		},
	},
	{
		"minecraft:purple_bed[facing=south,occupied=true,part=head]",
		nil,
		NewMapping{
			573: 1212,
			393: 912,
			404: 912,
			477: 1212,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 6456,
			477: 6962,
			573: 6962,
			393: 6455,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9134,
			573: 9134,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=4,south=none,west=none]",
		nil,
		NewMapping{
			393: 2660,
			404: 2661,
			477: 2964,
			573: 2964,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=13,south=side,west=side]",
		nil,
		NewMapping{
			393: 2161,
			404: 2162,
			477: 2465,
			573: 2465,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   1829,
			393: 4569,
			404: 4570,
			477: 5073,
			573: 5073,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=true,north=false,powered=true,south=false,west=false]",
		nil,
		NewMapping{
			404: 4831,
			477: 5334,
			573: 5334,
			393: 4830,
		},
	},
	{
		"minecraft:repeater[delay=1,facing=east,locked=true,powered=false]",
		nil,
		NewMapping{
			477: 4030,
			573: 4030,
			393: 3526,
			404: 3527,
		},
	},
	{
		"minecraft:oak_fence[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 3981,
			573: 3981,
			393: 3477,
			404: 3478,
		},
	},
	{
		"minecraft:diorite_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10195,
			573: 10195,
		},
	},
	{
		"minecraft:stripped_spruce_log[axis=x]",
		nil,
		NewMapping{
			477: 90,
			573: 90,
			393: 90,
			404: 90,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4722,
			404: 4723,
			477: 5226,
			573: 5226,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 3662,
			393: 3198,
			404: 3199,
			477: 3662,
		},
	},
	{
		"minecraft:oak_sign[rotation=9,waterlogged=true]",
		nil,
		NewMapping{
			404: 3094,
			477: 3397,
			573: 3397,
		},
	},
	{
		"minecraft:scaffolding[bottom=false,distance=3,waterlogged=false]",
		nil,
		NewMapping{
			477: 11122,
			573: 11122,
		},
	},
	{
		"minecraft:spruce_fence[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 7541,
			404: 7542,
			477: 8066,
			573: 8066,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10341,
			573: 10341,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=10,powered=true]",
		nil,
		NewMapping{
			477: 1018,
			573: 1018,
		},
	},
	{
		"minecraft:light_gray_glazed_terracotta[facing=east]",
		nil,
		NewMapping{
			4:   3891,
			393: 8348,
			404: 8349,
			477: 8873,
			573: 8873,
		},
	},
	{
		"minecraft:pink_banner[rotation=1]",
		nil,
		NewMapping{
			393: 6951,
			404: 6952,
			477: 7458,
			573: 7458,
		},
	},
	{
		"minecraft:green_banner[rotation=15]",
		nil,
		NewMapping{
			404: 7078,
			477: 7584,
			573: 7584,
			393: 7077,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=6,south=none,west=side]",
		nil,
		NewMapping{
			404: 1814,
			477: 2117,
			573: 2117,
			393: 1813,
		},
	},
	{
		"minecraft:oak_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 1662,
			404: 1663,
			477: 1966,
			573: 1966,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=20,powered=true]",
		nil,
		NewMapping{
			477: 988,
			573: 988,
		},
	},
	{
		"minecraft:cyan_wall_banner[facing=west]",
		nil,
		NewMapping{
			393: 7148,
			404: 7149,
			477: 7655,
			573: 7655,
		},
	},
	{
		"minecraft:water[level=8]",
		nil,
		NewMapping{
			393: 42,
			404: 42,
			477: 42,
			573: 42,
			4:   136,
		},
	},
	{
		"minecraft:melon_stem[age=3]",
		nil,
		NewMapping{
			393: 4263,
			404: 4264,
			477: 4767,
			573: 4767,
			4:   1683,
		},
	},
	{
		"minecraft:quartz_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			404: 5742,
			477: 6248,
			573: 6248,
			4:   2501,
			393: 5741,
		},
	},
	{
		"minecraft:birch_stairs[facing=south,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			404: 4993,
			477: 5496,
			573: 5496,
			393: 4992,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=south,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3933,
			404: 3934,
			477: 4437,
			573: 4437,
		},
	},
	{
		"minecraft:purpur_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 8099,
			404: 8100,
			477: 8624,
			573: 8624,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=3,south=side,west=none]",
		nil,
		NewMapping{
			393: 2216,
			404: 2217,
			477: 2520,
			573: 2520,
		},
	},
	{
		"minecraft:magenta_bed[facing=north,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 782,
			404: 782,
			477: 1082,
			573: 1082,
		},
	},
	{
		"minecraft:quartz_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			393: 7339,
			404: 7340,
			477: 7858,
			573: 7858,
		},
	},
	{
		"minecraft:birch_sign[rotation=3,waterlogged=true]",
		nil,
		NewMapping{
			573: 3449,
			477: 3449,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=11,powered=false]",
		nil,
		NewMapping{
			477: 921,
			573: 921,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9472,
			573: 9472,
		},
	},
	{
		"minecraft:jungle_door[facing=south,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7825,
			404: 7826,
			477: 8350,
			573: 8350,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=14,south=side,west=up]",
		nil,
		NewMapping{
			404: 2746,
			477: 3049,
			573: 3049,
			393: 2745,
		},
	},
	{
		"minecraft:acacia_stairs[facing=east,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 6901,
			573: 6901,
			393: 6394,
			404: 6395,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10787,
			573: 10787,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=4,powered=true]",
		nil,
		NewMapping{
			573: 956,
			477: 956,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=5,south=side,west=up]",
		nil,
		NewMapping{
			393: 2808,
			404: 2809,
			477: 3112,
			573: 3112,
		},
	},
	{
		"minecraft:cobblestone_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			393: 7313,
			404: 7314,
			477: 7832,
			573: 7832,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9137,
			573: 9137,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=13,south=up,west=up]",
		nil,
		NewMapping{
			393: 2301,
			404: 2302,
			477: 2605,
			573: 2605,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=south,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3875,
			404: 3876,
			477: 4379,
			573: 4379,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=2,south=side,west=none]",
		nil,
		NewMapping{
			393: 2927,
			404: 2928,
			477: 3231,
			573: 3231,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=6,south=up,west=up]",
		nil,
		NewMapping{
			393: 2382,
			404: 2383,
			477: 2686,
			573: 2686,
		},
	},
	{
		"minecraft:oak_wall_sign[facing=west,waterlogged=false]",
		nil,
		NewMapping{
			404: 3275,
			477: 3738,
			573: 3738,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 5040,
			393: 4536,
			404: 4537,
			477: 5040,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=north,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3731,
			404: 3732,
			477: 4235,
			573: 4235,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 6392,
			393: 5885,
			404: 5886,
			477: 6392,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=15,powered=false]",
		nil,
		NewMapping{
			477: 979,
			573: 979,
		},
	},
	{
		"minecraft:lever[face=wall,facing=south,powered=false]",
		nil,
		NewMapping{
			573: 3792,
			4:   1107,
			393: 3288,
			404: 3289,
			477: 3792,
		},
	},
	{
		"minecraft:detector_rail[powered=true,shape=east_west]",
		nil,
		NewMapping{
			573: 1317,
			4:   457,
			393: 1017,
			404: 1017,
			477: 1317,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=west,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3765,
			404: 3766,
			477: 4269,
			573: 4269,
		},
	},
	{
		"minecraft:diorite_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 10230,
			477: 10230,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9386,
			573: 9386,
		},
	},
	{
		"minecraft:polished_granite_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			477: 10254,
			573: 10254,
		},
	},
	{
		"minecraft:brick_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4387,
			404: 4388,
			477: 4891,
			573: 4891,
		},
	},
	{
		"minecraft:cut_sandstone_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			477: 7819,
			573: 7819,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=14,south=up,west=none]",
		nil,
		NewMapping{
			573: 3048,
			393: 2744,
			404: 2745,
			477: 3048,
		},
	},
	{
		"minecraft:zombie_wall_head[facing=west]",
		nil,
		NewMapping{
			404: 5490,
			477: 6012,
			573: 6012,
			393: 5489,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 9309,
			477: 9309,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=true,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10620,
			573: 10620,
		},
	},
	{
		"minecraft:lime_banner[rotation=11]",
		nil,
		NewMapping{
			393: 6945,
			404: 6946,
			477: 7452,
			573: 7452,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=3,south=up,west=none]",
		nil,
		NewMapping{
			393: 2213,
			404: 2214,
			477: 2517,
			573: 2517,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=11,south=none,west=up]",
		nil,
		NewMapping{
			393: 2721,
			404: 2722,
			477: 3025,
			573: 3025,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=4,south=none,west=side]",
		nil,
		NewMapping{
			393: 2515,
			404: 2516,
			477: 2819,
			573: 2819,
		},
	},
	{
		"minecraft:fire[age=7,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1388,
			404: 1389,
			477: 1692,
			573: 1692,
		},
	},
	{
		"minecraft:pink_wall_banner[facing=north]",
		nil,
		NewMapping{
			393: 7134,
			404: 7135,
			477: 7641,
			573: 7641,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=6,powered=false]",
		nil,
		NewMapping{
			573: 711,
			393: 711,
			404: 711,
			477: 711,
		},
	},
	{
		"minecraft:fire[age=8,east=false,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			573: 1714,
			393: 1410,
			404: 1411,
			477: 1714,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=14,south=up,west=side]",
		nil,
		NewMapping{
			393: 2167,
			404: 2168,
			477: 2471,
			573: 2471,
		},
	},
	{
		"minecraft:stone_stairs[facing=south,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9643,
			573: 9643,
		},
	},
	{
		"minecraft:sticky_piston[extended=true,facing=south]",
		nil,
		NewMapping{
			404: 1030,
			477: 1330,
			573: 1330,
			4:   475,
			393: 1030,
		},
	},
	{
		"minecraft:oak_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			4:   2024,
			393: 7258,
			404: 7259,
			477: 7765,
			573: 7765,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 4959,
			573: 4959,
			393: 4455,
			404: 4456,
		},
	},
	{
		"minecraft:oak_door[facing=east,half=upper,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 3155,
			404: 3156,
			477: 3619,
			573: 3619,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=13,powered=false]",
		nil,
		NewMapping{
			393: 325,
			404: 325,
			477: 325,
			573: 325,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9545,
			573: 9545,
		},
	},
	{
		"minecraft:white_bed[facing=east,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 760,
			404: 760,
			477: 1060,
			573: 1060,
		},
	},
	{
		"minecraft:black_bed[facing=west,occupied=true,part=foot]",
		nil,
		NewMapping{
			573: 1297,
			393: 997,
			404: 997,
			477: 1297,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=west,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3894,
			404: 3895,
			477: 4398,
			573: 4398,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 6571,
			404: 6572,
			477: 7078,
			573: 7078,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10017,
			573: 10017,
		},
	},
	{
		"minecraft:repeater[delay=2,facing=south,locked=true,powered=false]",
		nil,
		NewMapping{
			393: 3534,
			404: 3535,
			477: 4038,
			573: 4038,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9319,
			573: 9319,
		},
	},
	{
		"minecraft:dark_oak_door[facing=east,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7996,
			404: 7997,
			477: 8521,
			573: 8521,
			4:   3152,
		},
	},
	{
		"minecraft:iron_bars[east=false,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 4207,
			404: 4208,
			477: 4711,
			573: 4711,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=false,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6131,
			404: 6132,
			477: 6638,
			573: 6638,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=true,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6309,
			404: 6310,
			477: 6816,
			573: 6816,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=7,south=none,west=none]",
		nil,
		NewMapping{
			477: 2415,
			573: 2415,
			393: 2111,
			404: 2112,
		},
	},
	{
		"minecraft:jungle_pressure_plate[powered=true]",
		nil,
		NewMapping{
			393: 3373,
			404: 3374,
			477: 3877,
			573: 3877,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=5,south=up,west=none]",
		nil,
		NewMapping{
			393: 2231,
			404: 2232,
			477: 2535,
			573: 2535,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=13,powered=false]",
		nil,
		NewMapping{
			573: 525,
			393: 525,
			404: 525,
			477: 525,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=4,south=none,west=none]",
		nil,
		NewMapping{
			4:   884,
			393: 2948,
			404: 2949,
			477: 3252,
			573: 3252,
		},
	},
	{
		"minecraft:obsidian",
		nil,
		NewMapping{
			573: 1433,
			4:   784,
			393: 1129,
			404: 1130,
			477: 1433,
		},
	},
	{
		"minecraft:repeater[delay=3,facing=east,locked=false,powered=true]",
		nil,
		NewMapping{
			573: 4063,
			4:   1515,
			393: 3559,
			404: 3560,
			477: 4063,
		},
	},
	{
		"minecraft:sign[rotation=0,waterlogged=false]",
		nil,
		NewMapping{
			4:   1008,
			393: 3076,
		},
	},
	{
		"minecraft:spruce_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4902,
			404: 4903,
			477: 5406,
			573: 5406,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4438,
			404: 4439,
			477: 4942,
			573: 4942,
		},
	},
	{
		"minecraft:birch_fence[east=false,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 7576,
			404: 7577,
			477: 8101,
			573: 8101,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4535,
			404: 4536,
			477: 5039,
			573: 5039,
		},
	},
	{
		"minecraft:diorite_wall[east=false,north=true,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 11082,
			573: 11082,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=1,south=up,west=none]",
		nil,
		NewMapping{
			393: 2051,
			404: 2052,
			477: 2355,
			573: 2355,
		},
	},
	{
		"minecraft:creeper_head[rotation=7]",
		nil,
		NewMapping{
			477: 6041,
			573: 6041,
			393: 5538,
			404: 5539,
		},
	},
	{
		"minecraft:spruce_door[facing=south,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			573: 8222,
			393: 7697,
			404: 7698,
			477: 8222,
		},
	},
	{
		"minecraft:sign[rotation=8,waterlogged=true]",
		nil,
		NewMapping{
			393: 3091,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10429,
			573: 10429,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=2,south=none,west=side]",
		nil,
		NewMapping{
			393: 2353,
			404: 2354,
			477: 2657,
			573: 2657,
		},
	},
	{
		"minecraft:player_wall_head[facing=west]",
		nil,
		NewMapping{
			477: 6032,
			573: 6032,
			393: 5509,
			404: 5510,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9267,
			573: 9267,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=12,powered=false]",
		nil,
		NewMapping{
			477: 373,
			573: 373,
			393: 373,
			404: 373,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=14,south=none,west=none]",
		nil,
		NewMapping{
			477: 3054,
			573: 3054,
			393: 2750,
			404: 2751,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 10783,
			477: 10783,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=east,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 9602,
			477: 9602,
		},
	},
	{
		"minecraft:lime_concrete",
		nil,
		NewMapping{
			573: 8907,
			4:   4021,
			393: 8382,
			404: 8383,
			477: 8907,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=true,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			477: 4622,
			573: 4622,
			393: 4118,
			404: 4119,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=east,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			573: 7907,
			393: 7382,
			404: 7383,
			477: 7907,
		},
	},
	{
		"minecraft:kelp[age=9]",
		nil,
		NewMapping{
			404: 8419,
			477: 8943,
			573: 8943,
			393: 8418,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=22,powered=true]",
		nil,
		NewMapping{
			393: 592,
			404: 592,
			477: 592,
			573: 592,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=11,south=up,west=side]",
		nil,
		NewMapping{
			477: 2876,
			573: 2876,
			393: 2572,
			404: 2573,
		},
	},
	{
		"minecraft:pink_bed[facing=west,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 854,
			404: 854,
			477: 1154,
			573: 1154,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10151,
			573: 10151,
		},
	},
	{
		"minecraft:anvil[facing=west]",
		nil,
		NewMapping{
			4:   2321,
			393: 5569,
			404: 5570,
			477: 6076,
			573: 6076,
		},
	},
	{
		"minecraft:red_bed[facing=north,occupied=false,part=head]",
		nil,
		NewMapping{
			4:   426,
			393: 974,
			404: 974,
			477: 1274,
			573: 1274,
		},
	},
	{
		"minecraft:oak_door[facing=south,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 3137,
			404: 3138,
			477: 3601,
			573: 3601,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=0,south=side,west=side]",
		nil,
		NewMapping{
			393: 2188,
			404: 2189,
			477: 2492,
			573: 2492,
		},
	},
	{
		"minecraft:magenta_banner[rotation=6]",
		nil,
		NewMapping{
			393: 6892,
			404: 6893,
			477: 7399,
			573: 7399,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10150,
			573: 10150,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=8,south=up,west=up]",
		nil,
		NewMapping{
			393: 2976,
			404: 2977,
			477: 3280,
			573: 3280,
		},
	},
	{
		"minecraft:purpur_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 8618,
			393: 8093,
			404: 8094,
			477: 8618,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10098,
			573: 10098,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9533,
			573: 9533,
		},
	},
	{
		"minecraft:bamboo[age=1,leaves=small,stage=0]",
		nil,
		NewMapping{
			477: 9124,
			573: 9124,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 3724,
			573: 3724,
			4:   1072,
			393: 3260,
			404: 3261,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=11,south=none,west=up]",
		nil,
		NewMapping{
			393: 3009,
			404: 3010,
			477: 3313,
			573: 3313,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 9404,
			477: 9404,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10598,
			573: 10598,
		},
	},
	{
		"minecraft:diorite_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			477: 10326,
			573: 10326,
		},
	},
	{
		"minecraft:iron_door[facing=north,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 3312,
			404: 3313,
			477: 3816,
			573: 3816,
		},
	},
	{
		"minecraft:pink_banner[rotation=10]",
		nil,
		NewMapping{
			393: 6960,
			404: 6961,
			477: 7467,
			573: 7467,
		},
	},
	{
		"minecraft:petrified_oak_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			393: 7309,
			404: 7310,
			477: 7828,
			573: 7828,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 6761,
			477: 7267,
			573: 7267,
			393: 6760,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10638,
			573: 10638,
		},
	},
	{
		"minecraft:purpur_pillar[axis=y]",
		nil,
		NewMapping{
			4:   3232,
			393: 8075,
			404: 8076,
			477: 8600,
			573: 8600,
		},
	},
	{
		"minecraft:jungle_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 5098,
			404: 5099,
			477: 5602,
			573: 5602,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9526,
			477: 9526,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 11009,
			573: 11009,
		},
	},
	{
		"minecraft:oak_door[facing=east,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 3169,
			404: 3170,
			477: 3633,
			573: 3633,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			573: 4605,
			393: 4101,
			404: 4102,
			477: 4605,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10138,
			573: 10138,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10614,
			573: 10614,
		},
	},
	{
		"minecraft:attached_melon_stem[facing=east]",
		nil,
		NewMapping{
			393: 4251,
			404: 4252,
			477: 4755,
			573: 4755,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=false,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6130,
			404: 6131,
			477: 6637,
			573: 6637,
		},
	},
	{
		"minecraft:fire[age=7,east=false,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			477: 1686,
			573: 1686,
			393: 1382,
			404: 1383,
		},
	},
	{
		"minecraft:quartz_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 6219,
			573: 6219,
			393: 5712,
			404: 5713,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=15,south=up,west=none]",
		nil,
		NewMapping{
			404: 3042,
			477: 3345,
			573: 3345,
			393: 3041,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5257,
			404: 5258,
			477: 5761,
			573: 5761,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=7,south=up,west=up]",
		nil,
		NewMapping{
			393: 2535,
			404: 2536,
			477: 2839,
			573: 2839,
		},
	},
	{
		"minecraft:potted_orange_tulip",
		nil,
		NewMapping{
			393: 5279,
			404: 5280,
			477: 5783,
			573: 5783,
		},
	},
	{
		"minecraft:birch_stairs[facing=west,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 5014,
			404: 5015,
			477: 5518,
			573: 5518,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 7756,
			573: 7756,
			393: 7249,
			404: 7250,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=5,south=side,west=up]",
		nil,
		NewMapping{
			573: 2968,
			393: 2664,
			404: 2665,
			477: 2968,
		},
	},
	{
		"minecraft:brick_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 4870,
			393: 4366,
			404: 4367,
			477: 4870,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=0,powered=false]",
		nil,
		NewMapping{
			477: 799,
			573: 799,
		},
	},
	{
		"minecraft:bamboo[age=0,leaves=large,stage=1]",
		nil,
		NewMapping{
			573: 9121,
			477: 9121,
		},
	},
	{
		"minecraft:lime_shulker_box[facing=south]",
		nil,
		NewMapping{
			393: 8249,
			404: 8250,
			477: 8774,
			573: 8774,
			4:   3587,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6221,
			404: 6222,
			477: 6728,
			573: 6728,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=east,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3783,
			404: 3784,
			477: 4287,
			573: 4287,
		},
	},
	{
		"minecraft:acacia_door[facing=north,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7880,
			404: 7881,
			477: 8405,
			573: 8405,
		},
	},
	{
		"minecraft:shulker_box[facing=north]",
		nil,
		NewMapping{
			573: 8736,
			393: 8211,
			404: 8212,
			477: 8736,
		},
	},
	{
		"minecraft:jungle_sign[rotation=14,waterlogged=false]",
		nil,
		NewMapping{
			477: 3536,
			573: 3536,
		},
	},
	{
		"minecraft:hay_block[axis=x]",
		nil,
		NewMapping{
			4:   2724,
			393: 6820,
			404: 6821,
			477: 7327,
			573: 7327,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=north,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4652,
			404: 4653,
			477: 5156,
			573: 5156,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=south,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 3219,
			404: 3220,
			477: 3683,
			573: 3683,
		},
	},
	{
		"minecraft:fire[age=10,east=true,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 1466,
			404: 1467,
			477: 1770,
			573: 1770,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10736,
			573: 10736,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=13,south=up,west=none]",
		nil,
		NewMapping{
			573: 3183,
			393: 2879,
			404: 2880,
			477: 3183,
		},
	},
	{
		"minecraft:end_stone_brick_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			573: 10288,
			477: 10288,
		},
	},
	{
		"minecraft:pink_shulker_box[facing=east]",
		nil,
		NewMapping{
			4:   3605,
			393: 8254,
			404: 8255,
			477: 8779,
			573: 8779,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=5,south=side,west=up]",
		nil,
		NewMapping{
			393: 1944,
			404: 1945,
			477: 2248,
			573: 2248,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=true,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 3989,
			404: 3990,
			477: 4493,
			573: 4493,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=10,south=up,west=up]",
		nil,
		NewMapping{
			477: 2578,
			573: 2578,
			393: 2274,
			404: 2275,
		},
	},
	{
		"minecraft:purpur_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 8610,
			573: 8610,
			393: 8085,
			404: 8086,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9364,
			573: 9364,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10897,
			573: 10897,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=0,south=none,west=none]",
		nil,
		NewMapping{
			4:   880,
			393: 2912,
			404: 2913,
			477: 3216,
			573: 3216,
		},
	},
	{
		"minecraft:oak_pressure_plate[powered=false]",
		nil,
		NewMapping{
			573: 3872,
			4:   1152,
			393: 3368,
			404: 3369,
			477: 3872,
		},
	},
	{
		"minecraft:dead_fire_coral_fan[waterlogged=false]",
		nil,
		NewMapping{
			393: 8551,
			404: 8567,
			477: 9011,
			573: 9011,
		},
	},
	{
		"minecraft:spruce_fence[east=true,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7523,
			404: 7524,
			477: 8048,
			573: 8048,
		},
	},
	{
		"minecraft:magenta_banner[rotation=14]",
		nil,
		NewMapping{
			393: 6900,
			404: 6901,
			477: 7407,
			573: 7407,
		},
	},
	{
		"minecraft:snow[layers=4]",
		nil,
		NewMapping{
			477: 3922,
			573: 3922,
			4:   1251,
			393: 3418,
			404: 3419,
		},
	},
	{
		"minecraft:fire[age=4,east=true,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1278,
			404: 1279,
			477: 1582,
			573: 1582,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=10,powered=true]",
		nil,
		NewMapping{
			393: 668,
			404: 668,
			477: 668,
			573: 668,
		},
	},
	{
		"minecraft:fire[age=10,east=true,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1462,
			404: 1463,
			477: 1766,
			573: 1766,
		},
	},
	{
		"minecraft:barrel[facing=east,open=false]",
		nil,
		NewMapping{
			477: 11138,
			573: 11138,
		},
	},
	{
		"minecraft:andesite_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9960,
			573: 9960,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=north,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			404: 3599,
			477: 4102,
			573: 4102,
			393: 3598,
		},
	},
	{
		"minecraft:quartz_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 5755,
			404: 5756,
			477: 6262,
			573: 6262,
		},
	},
	{
		"minecraft:brick_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 4350,
			477: 4853,
			573: 4853,
			393: 4349,
		},
	},
	{
		"minecraft:blue_banner[rotation=13]",
		nil,
		NewMapping{
			393: 7043,
			404: 7044,
			477: 7550,
			573: 7550,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=true,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10921,
			573: 10921,
		},
	},
	{
		"minecraft:dark_oak_door[facing=north,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			404: 7946,
			477: 8470,
			573: 8470,
			393: 7945,
		},
	},
	{
		"minecraft:andesite_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9987,
			573: 9987,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9205,
			573: 9205,
		},
	},
	{
		"minecraft:prismarine",
		nil,
		NewMapping{
			4:   2688,
			393: 6558,
			404: 6559,
			477: 7065,
			573: 7065,
		},
	},
	{
		"minecraft:moving_piston[facing=east,type=sticky]",
		nil,
		NewMapping{
			4:   589,
			393: 1102,
			404: 1102,
			477: 1402,
			573: 1402,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5222,
			404: 5223,
			477: 5726,
			573: 5726,
		},
	},
	{
		"minecraft:stripped_dark_oak_log[axis=y]",
		nil,
		NewMapping{
			573: 103,
			393: 103,
			404: 103,
			477: 103,
		},
	},
	{
		"minecraft:fire[age=2,east=true,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			404: 1202,
			477: 1505,
			573: 1505,
			393: 1201,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9239,
			573: 9239,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=12,south=none,west=none]",
		nil,
		NewMapping{
			393: 2156,
			404: 2157,
			477: 2460,
			573: 2460,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4484,
			404: 4485,
			477: 4988,
			573: 4988,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=0,south=none,west=side]",
		nil,
		NewMapping{
			404: 2480,
			477: 2783,
			573: 2783,
			393: 2479,
		},
	},
	{
		"minecraft:jungle_door[facing=south,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			573: 8355,
			393: 7830,
			404: 7831,
			477: 8355,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 7136,
			573: 7136,
			393: 6629,
			404: 6630,
		},
	},
	{
		"minecraft:fire[age=4,east=false,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1283,
			404: 1284,
			477: 1587,
			573: 1587,
		},
	},
	{
		"minecraft:glass_pane[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 4232,
			404: 4233,
			477: 4736,
			573: 4736,
		},
	},
	{
		"minecraft:spruce_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 4894,
			477: 5397,
			573: 5397,
			393: 4893,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10595,
			573: 10595,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6029,
			404: 6030,
			477: 6536,
			573: 6536,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 5198,
			573: 5198,
			393: 4694,
			404: 4695,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10485,
			573: 10485,
		},
	},
	{
		"minecraft:jungle_door[facing=west,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7842,
			404: 7843,
			477: 8367,
			573: 8367,
		},
	},
	{
		"minecraft:spruce_door[facing=north,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			573: 8204,
			393: 7679,
			404: 7680,
			477: 8204,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=false,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			573: 8584,
			393: 8059,
			404: 8060,
			477: 8584,
		},
	},
	{
		"minecraft:birch_door[facing=north,half=lower,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			573: 8274,
			393: 7749,
			404: 7750,
			477: 8274,
		},
	},
	{
		"minecraft:oak_fence[east=true,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 3461,
			404: 3462,
			477: 3965,
			573: 3965,
		},
	},
	{
		"minecraft:purpur_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 8132,
			404: 8133,
			477: 8657,
			573: 8657,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=south,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3801,
			404: 3802,
			477: 4305,
			573: 4305,
		},
	},
	{
		"minecraft:fire[age=14,east=false,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1607,
			404: 1608,
			477: 1911,
			573: 1911,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10991,
			573: 10991,
		},
	},
	{
		"minecraft:fire[age=7,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			4:   823,
			393: 1390,
			404: 1391,
			477: 1694,
			573: 1694,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=6,south=up,west=none]",
		nil,
		NewMapping{
			393: 1952,
			404: 1953,
			477: 2256,
			573: 2256,
		},
	},
	{
		"minecraft:blue_bed[facing=north,occupied=false,part=head]",
		nil,
		NewMapping{
			573: 1226,
			393: 926,
			404: 926,
			477: 1226,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=0,south=none,west=up]",
		nil,
		NewMapping{
			477: 2926,
			573: 2926,
			393: 2622,
			404: 2623,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5204,
			404: 5205,
			477: 5708,
			573: 5708,
		},
	},
	{
		"minecraft:cyan_banner[rotation=8]",
		nil,
		NewMapping{
			393: 7006,
			404: 7007,
			477: 7513,
			573: 7513,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			404: 7229,
			477: 7735,
			573: 7735,
			4:   2881,
			393: 7228,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6640,
			404: 6641,
			477: 7147,
			573: 7147,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 6972,
			573: 6972,
			393: 6465,
			404: 6466,
		},
	},
	{
		"minecraft:birch_door[facing=south,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7763,
			404: 7764,
			477: 8288,
			573: 8288,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9800,
			477: 9800,
		},
	},
	{
		"minecraft:fire[age=7,east=false,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 1376,
			404: 1377,
			477: 1680,
			573: 1680,
		},
	},
	{
		"minecraft:jungle_sign[rotation=13,waterlogged=false]",
		nil,
		NewMapping{
			477: 3534,
			573: 3534,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=north,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3785,
			404: 3786,
			477: 4289,
			573: 4289,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 4660,
			477: 5163,
			573: 5163,
			393: 4659,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=west,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3637,
			404: 3638,
			477: 4141,
			573: 4141,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6019,
			404: 6020,
			477: 6526,
			573: 6526,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=2,south=none,west=up]",
		nil,
		NewMapping{
			477: 3088,
			573: 3088,
			393: 2784,
			404: 2785,
		},
	},
	{
		"minecraft:tube_coral_wall_fan[facing=west,waterlogged=false]",
		nil,
		NewMapping{
			393: 8509,
			404: 8525,
			477: 9069,
			573: 9069,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 9532,
			477: 9532,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9789,
			573: 9789,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=6,south=side,west=up]",
		nil,
		NewMapping{
			477: 2257,
			573: 2257,
			393: 1953,
			404: 1954,
		},
	},
	{
		"minecraft:acacia_leaves[distance=5,persistent=true]",
		nil,
		NewMapping{
			477: 208,
			573: 208,
			393: 208,
			404: 208,
		},
	},
	{
		"minecraft:cyan_bed[facing=north,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 893,
			404: 893,
			477: 1193,
			573: 1193,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=4,south=up,west=up]",
		nil,
		NewMapping{
			393: 2796,
			404: 2797,
			477: 3100,
			573: 3100,
		},
	},
	{
		"minecraft:fire[age=6,east=false,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			573: 1652,
			393: 1348,
			404: 1349,
			477: 1652,
		},
	},
	{
		"minecraft:purpur_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 8086,
			404: 8087,
			477: 8611,
			573: 8611,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=north,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			573: 7010,
			393: 6503,
			404: 6504,
			477: 7010,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=false,north=true,powered=true,south=false,west=true]",
		nil,
		NewMapping{
			393: 4773,
			404: 4774,
			477: 5277,
			573: 5277,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=3,south=side,west=side]",
		nil,
		NewMapping{
			393: 2647,
			404: 2648,
			477: 2951,
			573: 2951,
		},
	},
	{
		"minecraft:diorite_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10191,
			573: 10191,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9432,
			573: 9432,
		},
	},
	{
		"minecraft:bell[attachment=double_wall,facing=south,powered=false]",
		nil,
		NewMapping{
			573: 11225,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=east,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			4:   2675,
			393: 6557,
			404: 6558,
			477: 7064,
			573: 7064,
		},
	},
	{
		"minecraft:yellow_banner[rotation=9]",
		nil,
		NewMapping{
			573: 7434,
			393: 6927,
			404: 6928,
			477: 7434,
		},
	},
	{
		"minecraft:horn_coral",
		nil,
		NewMapping{
			393: 8463,
		},
	},
	{
		"minecraft:cyan_bed[facing=south,occupied=false,part=head]",
		nil,
		NewMapping{
			477: 1198,
			573: 1198,
			393: 898,
			404: 898,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4586,
			404: 4587,
			477: 5090,
			573: 5090,
		},
	},
	{
		"minecraft:oak_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 1718,
			404: 1719,
			477: 2022,
			573: 2022,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=south,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9318,
			573: 9318,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=21,powered=false]",
		nil,
		NewMapping{
			477: 891,
			573: 891,
		},
	},
	{
		"minecraft:purple_concrete",
		nil,
		NewMapping{
			4:   4026,
			393: 8387,
			404: 8388,
			477: 8912,
			573: 8912,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=north,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 6504,
			404: 6505,
			477: 7011,
			573: 7011,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6420,
			404: 6421,
			477: 6927,
			573: 6927,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=true,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 4055,
			404: 4056,
			477: 4559,
			573: 4559,
		},
	},
	{
		"minecraft:white_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 6346,
			573: 6346,
			393: 5839,
			404: 5840,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9325,
			573: 9325,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 10353,
			477: 10353,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=0,powered=false]",
		nil,
		NewMapping{
			477: 749,
			573: 749,
		},
	},
	{
		"minecraft:lever[face=wall,facing=west,powered=true]",
		nil,
		NewMapping{
			4:   1114,
			393: 3289,
			404: 3290,
			477: 3793,
			573: 3793,
		},
	},
	{
		"minecraft:spruce_button[face=wall,facing=north,powered=false]",
		nil,
		NewMapping{
			393: 5336,
			404: 5337,
			477: 5843,
			573: 5843,
		},
	},
	{
		"minecraft:light_gray_wall_banner[facing=east]",
		nil,
		NewMapping{
			393: 7145,
			404: 7146,
			477: 7652,
			573: 7652,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=13,powered=true]",
		nil,
		NewMapping{
			393: 724,
			404: 724,
			477: 724,
			573: 724,
		},
	},
	{
		"minecraft:fire[age=3,east=false,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			477: 1565,
			573: 1565,
			393: 1261,
			404: 1262,
		},
	},
	{
		"minecraft:fire[age=8,east=true,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 1400,
			404: 1401,
			477: 1704,
			573: 1704,
		},
	},
	{
		"minecraft:oak_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 1690,
			404: 1691,
			477: 1994,
			573: 1994,
		},
	},
	{
		"minecraft:dark_oak_fence[east=true,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 8178,
			393: 7653,
			404: 7654,
			477: 8178,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10624,
			477: 10624,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=true,north=false,powered=false,south=true,west=true]",
		nil,
		NewMapping{
			573: 5367,
			393: 4863,
			404: 4864,
			477: 5367,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=east,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6710,
			404: 6711,
			477: 7217,
			573: 7217,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4673,
			404: 4674,
			477: 5177,
			573: 5177,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=false,north=false,powered=false,south=false,west=true]",
		nil,
		NewMapping{
			393: 4817,
			404: 4818,
			477: 5321,
			573: 5321,
		},
	},
	{
		"minecraft:acacia_fence[east=true,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 8147,
			573: 8147,
			393: 7622,
			404: 7623,
		},
	},
	{
		"minecraft:jungle_door[facing=north,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7807,
			404: 7808,
			477: 8332,
			573: 8332,
		},
	},
	{
		"minecraft:comparator[facing=south,mode=subtract,powered=true]",
		nil,
		NewMapping{
			4:   2412,
			393: 5641,
			404: 5642,
			477: 6148,
			573: 6148,
		},
	},
	{
		"minecraft:tripwire_hook[attached=false,facing=west,powered=false]",
		nil,
		NewMapping{
			573: 5256,
			4:   2097,
			393: 4752,
			404: 4753,
			477: 5256,
		},
	},
	{
		"minecraft:spruce_door[facing=west,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7714,
			404: 7715,
			477: 8239,
			573: 8239,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=west,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			477: 4325,
			573: 4325,
			393: 3821,
			404: 3822,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=west,half=top,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3626,
			404: 3627,
			477: 4130,
			573: 4130,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=6,south=side,west=none]",
		nil,
		NewMapping{
			393: 2963,
			404: 2964,
			477: 3267,
			573: 3267,
		},
	},
	{
		"minecraft:scaffolding[bottom=false,distance=2,waterlogged=false]",
		nil,
		NewMapping{
			477: 11120,
			573: 11120,
		},
	},
	{
		"minecraft:quartz_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5737,
			404: 5738,
			477: 6244,
			573: 6244,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10403,
			573: 10403,
		},
	},
	{
		"minecraft:dark_oak_wood[axis=y]",
		nil,
		NewMapping{
			4:   2605,
			393: 124,
			404: 124,
			477: 124,
			573: 124,
		},
	},
	{
		"minecraft:cocoa[age=0,facing=west]",
		nil,
		NewMapping{
			4:   2033,
			393: 4640,
			404: 4641,
			477: 5144,
			573: 5144,
		},
	},
	{
		"minecraft:oak_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 1658,
			404: 1659,
			477: 1962,
			573: 1962,
		},
	},
	{
		"minecraft:fire[age=14,east=false,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			477: 1908,
			573: 1908,
			393: 1604,
			404: 1605,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5976,
			404: 5977,
			477: 6483,
			573: 6483,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			404: 4719,
			477: 5222,
			573: 5222,
			393: 4718,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=east,half=top,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3708,
			404: 3709,
			477: 4212,
			573: 4212,
		},
	},
	{
		"minecraft:pink_bed[facing=east,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 857,
			404: 857,
			477: 1157,
			573: 1157,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 3242,
			404: 3243,
			477: 3706,
			573: 3706,
		},
	},
	{
		"minecraft:spruce_door[facing=south,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			573: 8227,
			393: 7702,
			404: 7703,
			477: 8227,
		},
	},
	{
		"minecraft:light_blue_banner[rotation=8]",
		nil,
		NewMapping{
			404: 6911,
			477: 7417,
			573: 7417,
			393: 6910,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=10,south=side,west=up]",
		nil,
		NewMapping{
			573: 3157,
			393: 2853,
			404: 2854,
			477: 3157,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9231,
			573: 9231,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=0,south=up,west=none]",
		nil,
		NewMapping{
			477: 3210,
			573: 3210,
			393: 2906,
			404: 2907,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			404: 5214,
			477: 5717,
			573: 5717,
			393: 5213,
		},
	},
	{
		"minecraft:purple_banner[rotation=3]",
		nil,
		NewMapping{
			393: 7017,
			404: 7018,
			477: 7524,
			573: 7524,
		},
	},
	{
		"minecraft:ladder[facing=north,waterlogged=true]",
		nil,
		NewMapping{
			393: 3171,
			404: 3172,
			477: 3635,
			573: 3635,
		},
	},
	{
		"minecraft:turtle_egg[eggs=3,hatch=2]",
		nil,
		NewMapping{
			393: 8445,
			404: 8446,
			477: 8970,
			573: 8970,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 11041,
			573: 11041,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=6,waterlogged=true]",
		nil,
		NewMapping{
			573: 3551,
			477: 3551,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10480,
			573: 10480,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=true,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 5217,
			477: 5720,
			573: 5720,
			393: 5216,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 11021,
			573: 11021,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9460,
			573: 9460,
		},
	},
	{
		"minecraft:detector_rail[powered=false,shape=ascending_north]",
		nil,
		NewMapping{
			393: 1026,
			404: 1026,
			477: 1326,
			573: 1326,
			4:   452,
		},
	},
	{
		"minecraft:dark_oak_fence_gate[facing=north,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7485,
			404: 7486,
			477: 8010,
			573: 8010,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=11,south=none,west=none]",
		nil,
		NewMapping{
			404: 2292,
			477: 2595,
			573: 2595,
			393: 2291,
		},
	},
	{
		"minecraft:birch_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5031,
			404: 5032,
			477: 5535,
			573: 5535,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=11,powered=true]",
		nil,
		NewMapping{
			573: 320,
			393: 320,
			404: 320,
			477: 320,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10632,
			573: 10632,
		},
	},
	{
		"minecraft:dark_prismarine",
		nil,
		NewMapping{
			4:   2690,
			393: 6560,
			404: 6561,
			477: 7067,
			573: 7067,
		},
	},
	{
		"minecraft:petrified_oak_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			4:   706,
			393: 7308,
			404: 7309,
			477: 7827,
			573: 7827,
		},
	},
	{
		"minecraft:stripped_birch_log[axis=z]",
		nil,
		NewMapping{
			393: 95,
			404: 95,
			477: 95,
			573: 95,
		},
	},
	{
		"minecraft:birch_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 5532,
			573: 5532,
			393: 5028,
			404: 5029,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9845,
			573: 9845,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=0,south=side,west=none]",
		nil,
		NewMapping{
			393: 2621,
			404: 2622,
			477: 2925,
			573: 2925,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 11045,
			573: 11045,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=true,north=true,powered=false,south=true,west=false]",
		nil,
		NewMapping{
			393: 4824,
			404: 4825,
			477: 5328,
			573: 5328,
		},
	},
	{
		"minecraft:black_banner[rotation=4]",
		nil,
		NewMapping{
			404: 7099,
			477: 7605,
			573: 7605,
			393: 7098,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=11,powered=false]",
		nil,
		NewMapping{
			573: 521,
			393: 521,
			404: 521,
			477: 521,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4725,
			404: 4726,
			477: 5229,
			573: 5229,
		},
	},
	{
		"minecraft:fire[age=2,east=true,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1203,
			404: 1204,
			477: 1507,
			573: 1507,
		},
	},
	{
		"minecraft:jungle_door[facing=south,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7836,
			404: 7837,
			477: 8361,
			573: 8361,
			4:   3121,
		},
	},
	{
		"minecraft:skeleton_skull[rotation=5]",
		nil,
		NewMapping{
			404: 5457,
			477: 5959,
			573: 5959,
			393: 5456,
		},
	},
	{
		"minecraft:green_bed[facing=west,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 964,
			404: 964,
			477: 1264,
			573: 1264,
		},
	},
	{
		"minecraft:scaffolding[bottom=true,distance=3,waterlogged=true]",
		nil,
		NewMapping{
			477: 11105,
			573: 11105,
		},
	},
	{
		"minecraft:lectern[facing=west,has_book=false,powered=false]",
		nil,
		NewMapping{
			477: 11188,
			573: 11188,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9416,
			477: 9416,
		},
	},
	{
		"minecraft:stripped_birch_wood[axis=x]",
		nil,
		NewMapping{
			393: 132,
			404: 132,
			477: 132,
			573: 132,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=9,south=side,west=up]",
		nil,
		NewMapping{
			573: 2716,
			393: 2412,
			404: 2413,
			477: 2716,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=8,powered=false]",
		nil,
		NewMapping{
			393: 665,
			404: 665,
			477: 665,
			573: 665,
		},
	},
	{
		"minecraft:chorus_plant[down=true,east=false,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 8030,
			404: 8031,
			477: 8555,
			573: 8555,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 10958,
			477: 10958,
		},
	},
	{
		"minecraft:daylight_detector[inverted=true,power=12]",
		nil,
		NewMapping{
			4:   2860,
			393: 5663,
			404: 5664,
			477: 6170,
			573: 6170,
		},
	},
	{
		"minecraft:gray_shulker_box[facing=south]",
		nil,
		NewMapping{
			4:   3619,
			393: 8261,
			404: 8262,
			477: 8786,
			573: 8786,
		},
	},
	{
		"minecraft:purpur_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   3253,
			393: 8118,
			404: 8119,
			477: 8643,
			573: 8643,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=1,powered=false]",
		nil,
		NewMapping{
			393: 501,
			404: 501,
			477: 501,
			573: 501,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=21,powered=false]",
		nil,
		NewMapping{
			477: 841,
			573: 841,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=14,south=none,west=none]",
		nil,
		NewMapping{
			4:   894,
			393: 3038,
			404: 3039,
			477: 3342,
			573: 3342,
		},
	},
	{
		"minecraft:vine[east=false,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			477: 4797,
			573: 4797,
			393: 4293,
			404: 4294,
		},
	},
	{
		"minecraft:purpur_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 8639,
			573: 8639,
			393: 8114,
			404: 8115,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=true,north=true,powered=false,south=true,west=false]",
		nil,
		NewMapping{
			393: 4792,
			404: 4793,
			477: 5296,
			573: 5296,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=south,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9398,
			573: 9398,
		},
	},
	{
		"minecraft:fire[age=1,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			4:   817,
			393: 1198,
			404: 1199,
			477: 1502,
			573: 1502,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=6,south=none,west=none]",
		nil,
		NewMapping{
			393: 2246,
			404: 2247,
			477: 2550,
			573: 2550,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=7,south=up,west=up]",
		nil,
		NewMapping{
			573: 2119,
			393: 1815,
			404: 1816,
			477: 2119,
		},
	},
	{
		"minecraft:birch_fence[east=false,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 8097,
			573: 8097,
			393: 7572,
			404: 7573,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=12,south=up,west=up]",
		nil,
		NewMapping{
			404: 2005,
			477: 2308,
			573: 2308,
			393: 2004,
		},
	},
	{
		"minecraft:trapped_chest[facing=north,type=right,waterlogged=true]",
		nil,
		NewMapping{
			404: 5584,
			477: 6090,
			573: 6090,
			393: 5583,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=south,in_wall=true,open=false,powered=true]",
		nil,
		NewMapping{
			404: 7368,
			477: 7892,
			573: 7892,
			393: 7367,
		},
	},
	{
		"minecraft:diorite_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			477: 10329,
			573: 10329,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 6970,
			573: 6970,
			4:   2625,
			393: 6463,
			404: 6464,
		},
	},
	{
		"minecraft:orange_bed[facing=north,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 767,
			404: 767,
			477: 1067,
			573: 1067,
		},
	},
	{
		"minecraft:purple_banner[rotation=6]",
		nil,
		NewMapping{
			573: 7527,
			393: 7020,
			404: 7021,
			477: 7527,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=7,powered=true]",
		nil,
		NewMapping{
			393: 662,
			404: 662,
			477: 662,
			573: 662,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=3,south=side,west=up]",
		nil,
		NewMapping{
			393: 2070,
			404: 2071,
			477: 2374,
			573: 2374,
		},
	},
	{
		"minecraft:cracked_stone_bricks",
		nil,
		NewMapping{
			4:   1570,
			393: 3985,
			404: 3986,
			477: 4483,
			573: 4483,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=3,south=none,west=none]",
		nil,
		NewMapping{
			393: 2507,
			404: 2508,
			477: 2811,
			573: 2811,
		},
	},
	{
		"minecraft:fire[age=8,east=true,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1393,
			404: 1394,
			477: 1697,
			573: 1697,
		},
	},
	{
		"minecraft:birch_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 5471,
			393: 4967,
			404: 4968,
			477: 5471,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=south,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 6432,
			404: 6433,
			477: 6939,
			573: 6939,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=11,south=none,west=none]",
		nil,
		NewMapping{
			393: 2723,
			404: 2724,
			477: 3027,
			573: 3027,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 6682,
			404: 6683,
			477: 7189,
			573: 7189,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10149,
			573: 10149,
		},
	},
	{
		"minecraft:oak_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 2023,
			4:   848,
			393: 1719,
			404: 1720,
			477: 2023,
		},
	},
	{
		"minecraft:spruce_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   2144,
			393: 4955,
			404: 4956,
			477: 5459,
			573: 5459,
		},
	},
	{
		"minecraft:cyan_shulker_box[facing=east]",
		nil,
		NewMapping{
			404: 8273,
			477: 8797,
			573: 8797,
			4:   3653,
			393: 8272,
		},
	},
	{
		"minecraft:black_banner[rotation=1]",
		nil,
		NewMapping{
			393: 7095,
			404: 7096,
			477: 7602,
			573: 7602,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 6997,
			573: 6997,
			393: 6490,
			404: 6491,
		},
	},
	{
		"minecraft:polished_diorite_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			477: 10275,
			573: 10275,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=south,half=top,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3680,
			404: 3681,
			477: 4184,
			573: 4184,
		},
	},
	{
		"minecraft:mushroom_stem[down=false,east=true,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 4153,
			404: 4154,
			477: 4657,
			573: 4657,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=south,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 10033,
			573: 10033,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10461,
			573: 10461,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10995,
			573: 10995,
		},
	},
	{
		"minecraft:iron_bars[east=false,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 4197,
			404: 4198,
			477: 4701,
			573: 4701,
		},
	},
	{
		"minecraft:jungle_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 5582,
			573: 5582,
			393: 5078,
			404: 5079,
		},
	},
	{
		"minecraft:nether_brick_fence[east=false,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 4522,
			404: 4523,
			477: 5026,
			573: 5026,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=7,south=none,west=up]",
		nil,
		NewMapping{
			393: 1965,
			404: 1966,
			477: 2269,
			573: 2269,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=2,south=up,west=side]",
		nil,
		NewMapping{
			404: 1772,
			477: 2075,
			573: 2075,
			393: 1771,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=1,south=up,west=up]",
		nil,
		NewMapping{
			393: 2769,
			404: 2770,
			477: 3073,
			573: 3073,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=false,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			573: 8576,
			393: 8051,
			404: 8052,
			477: 8576,
		},
	},
	{
		"minecraft:spruce_door[facing=south,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7698,
			404: 7699,
			477: 8223,
			573: 8223,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6574,
			404: 6575,
			477: 7081,
			573: 7081,
		},
	},
	{
		"minecraft:black_bed[facing=south,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 993,
			404: 993,
			477: 1293,
			573: 1293,
		},
	},
	{
		"minecraft:spruce_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			4:   2001,
			393: 7268,
			404: 7269,
			477: 7775,
			573: 7775,
		},
	},
	{
		"minecraft:dark_oak_sapling[stage=0]",
		nil,
		NewMapping{
			4:   101,
			393: 31,
			404: 31,
			477: 31,
			573: 31,
		},
	},
	{
		"minecraft:red_bed[facing=north,occupied=false,part=foot]",
		nil,
		NewMapping{
			477: 1275,
			573: 1275,
			4:   418,
			393: 975,
			404: 975,
		},
	},
	{
		"minecraft:acacia_leaves[distance=1,persistent=false]",
		nil,
		NewMapping{
			4:   2576,
			393: 201,
			404: 201,
			477: 201,
			573: 201,
		},
	},
	{
		"minecraft:pink_bed[facing=west,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 855,
			404: 855,
			477: 1155,
			573: 1155,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=3,south=none,west=side]",
		nil,
		NewMapping{
			393: 2938,
			404: 2939,
			477: 3242,
			573: 3242,
		},
	},
	{
		"minecraft:dead_brain_coral[waterlogged=false]",
		nil,
		NewMapping{
			404: 8463,
			477: 8987,
			573: 8987,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=10,powered=false]",
		nil,
		NewMapping{
			477: 769,
			573: 769,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=13,powered=true]",
		nil,
		NewMapping{
			477: 324,
			573: 324,
			393: 324,
			404: 324,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=5,south=up,west=up]",
		nil,
		NewMapping{
			393: 1941,
			404: 1942,
			477: 2245,
			573: 2245,
		},
	},
	{
		"minecraft:dark_oak_fence[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 7666,
			404: 7667,
			477: 8191,
			573: 8191,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10418,
			573: 10418,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=false,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6331,
			404: 6332,
			477: 6838,
			573: 6838,
			4:   2575,
		},
	},
	{
		"minecraft:jukebox[has_record=false]",
		nil,
		NewMapping{
			4:   1344,
			393: 3459,
			404: 3460,
			477: 3963,
			573: 3963,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=22,powered=false]",
		nil,
		NewMapping{
			393: 393,
			404: 393,
			477: 393,
			573: 393,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6438,
			404: 6439,
			477: 6945,
			573: 6945,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=east,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			477: 4337,
			573: 4337,
			393: 3833,
			404: 3834,
		},
	},
	{
		"minecraft:birch_door[facing=east,half=upper,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7790,
			404: 7791,
			477: 8315,
			573: 8315,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=10,south=up,west=none]",
		nil,
		NewMapping{
			573: 2724,
			393: 2420,
			404: 2421,
			477: 2724,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 10013,
			573: 10013,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 11005,
			573: 11005,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 11061,
			573: 11061,
		},
	},
	{
		"minecraft:green_concrete",
		nil,
		NewMapping{
			4:   4029,
			393: 8390,
			404: 8391,
			477: 8915,
			573: 8915,
		},
	},
	{
		"minecraft:pink_banner[rotation=6]",
		nil,
		NewMapping{
			393: 6956,
			404: 6957,
			477: 7463,
			573: 7463,
		},
	},
	{
		"minecraft:oak_fence_gate[facing=east,in_wall=true,open=false,powered=true]",
		nil,
		NewMapping{
			573: 4830,
			393: 4326,
			404: 4327,
			477: 4830,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=10,powered=false]",
		nil,
		NewMapping{
			393: 719,
			404: 719,
			477: 719,
			573: 719,
		},
	},
	{
		"minecraft:fire[age=2,east=false,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			404: 1224,
			477: 1527,
			573: 1527,
			393: 1223,
		},
	},
	{
		"minecraft:mossy_cobblestone",
		nil,
		NewMapping{
			4:   768,
			393: 1128,
			404: 1129,
			477: 1432,
			573: 1432,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 6964,
			573: 6964,
			393: 6457,
			404: 6458,
		},
	},
	{
		"minecraft:spruce_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4932,
			404: 4933,
			477: 5436,
			573: 5436,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=east,in_wall=true,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7479,
			404: 7480,
			477: 8004,
			573: 8004,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 7235,
			393: 6728,
			404: 6729,
			477: 7235,
		},
	},
	{
		"minecraft:gray_banner[rotation=14]",
		nil,
		NewMapping{
			393: 6980,
			404: 6981,
			477: 7487,
			573: 7487,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6086,
			404: 6087,
			477: 6593,
			573: 6593,
		},
	},
	{
		"minecraft:large_fern[half=upper]",
		nil,
		NewMapping{
			393: 6852,
			404: 6853,
			477: 7359,
			573: 7359,
		},
	},
	{
		"minecraft:spruce_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 5415,
			573: 5415,
			393: 4911,
			404: 4912,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=6,south=up,west=up]",
		nil,
		NewMapping{
			477: 2398,
			573: 2398,
			393: 2094,
			404: 2095,
		},
	},
	{
		"minecraft:spruce_door[facing=north,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7681,
			404: 7682,
			477: 8206,
			573: 8206,
		},
	},
	{
		"minecraft:dark_prismarine_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			404: 6818,
			477: 7324,
			573: 7324,
			393: 6817,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=15,south=up,west=side]",
		nil,
		NewMapping{
			393: 1888,
			404: 1889,
			477: 2192,
			573: 2192,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9564,
			573: 9564,
		},
	},
	{
		"minecraft:stone_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9642,
			573: 9642,
		},
	},
	{
		"minecraft:carrots[age=4]",
		nil,
		NewMapping{
			477: 5798,
			573: 5798,
			4:   2260,
			393: 5291,
			404: 5292,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6134,
			404: 6135,
			477: 6641,
			573: 6641,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=12,south=none,west=up]",
		nil,
		NewMapping{
			477: 2746,
			573: 2746,
			393: 2442,
			404: 2443,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=17,powered=false]",
		nil,
		NewMapping{
			393: 533,
			404: 533,
			477: 533,
			573: 533,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9344,
			573: 9344,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=11,south=none,west=up]",
		nil,
		NewMapping{
			404: 2002,
			477: 2305,
			573: 2305,
			393: 2001,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6061,
			404: 6062,
			477: 6568,
			573: 6568,
		},
	},
	{
		"minecraft:cobblestone_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5160,
			404: 5161,
			477: 5664,
			573: 5664,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=7,powered=true]",
		nil,
		NewMapping{
			573: 1012,
			477: 1012,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 9314,
			477: 9314,
		},
	},
	{
		"minecraft:oak_fence[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 3471,
			404: 3472,
			477: 3975,
			573: 3975,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10119,
			573: 10119,
		},
	},
	{
		"minecraft:andesite_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9989,
			573: 9989,
		},
	},
	{
		"minecraft:cactus[age=14]",
		nil,
		NewMapping{
			573: 3943,
			4:   1310,
			393: 3439,
			404: 3440,
			477: 3943,
		},
	},
	{
		"minecraft:light_blue_glazed_terracotta[facing=south]",
		nil,
		NewMapping{
			404: 8327,
			477: 8851,
			573: 8851,
			4:   3808,
			393: 8326,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=false,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			477: 4607,
			573: 4607,
			4:   1601,
			393: 4103,
			404: 4104,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			404: 5911,
			477: 6417,
			573: 6417,
			393: 5910,
		},
	},
	{
		"minecraft:acacia_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6344,
			404: 6345,
			477: 6851,
			573: 6851,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9260,
			573: 9260,
		},
	},
	{
		"minecraft:scaffolding[bottom=false,distance=1,waterlogged=false]",
		nil,
		NewMapping{
			477: 11118,
			573: 11118,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9379,
			573: 9379,
		},
	},
	{
		"minecraft:jungle_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			404: 5116,
			477: 5619,
			573: 5619,
			4:   2176,
			393: 5115,
		},
	},
	{
		"minecraft:orange_bed[facing=south,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 769,
			404: 769,
			477: 1069,
			573: 1069,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=15,south=up,west=side]",
		nil,
		NewMapping{
			573: 2912,
			393: 2608,
			404: 2609,
			477: 2912,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=10,south=side,west=up]",
		nil,
		NewMapping{
			393: 2997,
			404: 2998,
			477: 3301,
			573: 3301,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=8,south=up,west=side]",
		nil,
		NewMapping{
			393: 2545,
			404: 2546,
			477: 2849,
			573: 2849,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9242,
			573: 9242,
		},
	},
	{
		"minecraft:chain_command_block[conditional=true,facing=down]",
		nil,
		NewMapping{
			4:   3384,
			393: 8181,
			404: 8182,
			477: 8706,
			573: 8706,
		},
	},
	{
		"minecraft:purple_glazed_terracotta[facing=east]",
		nil,
		NewMapping{
			4:   3923,
			393: 8356,
			404: 8357,
			477: 8881,
			573: 8881,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=west,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3957,
			404: 3958,
			477: 4461,
			573: 4461,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 6096,
			477: 6602,
			573: 6602,
			393: 6095,
		},
	},
	{
		"minecraft:iron_door[facing=west,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			573: 3853,
			393: 3349,
			404: 3350,
			477: 3853,
		},
	},
	{
		"minecraft:player_head[rotation=0]",
		nil,
		NewMapping{
			404: 5512,
			477: 6014,
			573: 6014,
			393: 5511,
		},
	},
	{
		"minecraft:bone_block[axis=y]",
		nil,
		NewMapping{
			477: 8721,
			573: 8721,
			4:   3456,
			393: 8196,
			404: 8197,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			404: 6143,
			477: 6649,
			573: 6649,
			393: 6142,
		},
	},
	{
		"minecraft:birch_door[facing=east,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7801,
			404: 7802,
			477: 8326,
			573: 8326,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=east,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7414,
			404: 7415,
			477: 7939,
			573: 7939,
		},
	},
	{
		"minecraft:magenta_bed[facing=north,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 780,
			404: 780,
			477: 1080,
			573: 1080,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9261,
			573: 9261,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10416,
			573: 10416,
		},
	},
	{
		"minecraft:barrel[facing=up,open=false]",
		nil,
		NewMapping{
			477: 11144,
			573: 11144,
		},
	},
	{
		"minecraft:fire[age=5,east=false,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			573: 1624,
			393: 1320,
			404: 1321,
			477: 1624,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=14,south=side,west=none]",
		nil,
		NewMapping{
			404: 2460,
			477: 2763,
			573: 2763,
			393: 2459,
		},
	},
	{
		"minecraft:quartz_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 5740,
			404: 5741,
			477: 6247,
			573: 6247,
		},
	},
	{
		"minecraft:fire[age=3,east=true,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1245,
			404: 1246,
			477: 1549,
			573: 1549,
		},
	},
	{
		"minecraft:sign[rotation=3,waterlogged=true]",
		nil,
		NewMapping{
			393: 3081,
		},
	},
	{
		"minecraft:jungle_door[facing=north,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7811,
			404: 7812,
			477: 8336,
			573: 8336,
		},
	},
	{
		"minecraft:nether_brick_fence[east=true,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 4501,
			404: 4502,
			477: 5005,
			573: 5005,
		},
	},
	{
		"minecraft:granite_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9930,
			573: 9930,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10751,
			573: 10751,
		},
	},
	{
		"minecraft:white_tulip",
		nil,
		NewMapping{
			4:   614,
			393: 1118,
			404: 1118,
			477: 1418,
			573: 1418,
		},
	},
	{
		"minecraft:fire[age=2,east=false,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1221,
			404: 1222,
			477: 1525,
			573: 1525,
		},
	},
	{
		"minecraft:fire[age=5,east=true,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1310,
			404: 1311,
			477: 1614,
			573: 1614,
		},
	},
	{
		"minecraft:oak_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 1990,
			393: 1686,
			404: 1687,
			477: 1990,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=false,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 11034,
			573: 11034,
		},
	},
	{
		"minecraft:dispenser[facing=down,triggered=false]",
		nil,
		NewMapping{
			4:   368,
			393: 244,
			404: 244,
			477: 244,
			573: 244,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   1747,
			393: 4423,
			404: 4424,
			477: 4927,
			573: 4927,
		},
	},
	{
		"minecraft:spruce_button[face=floor,facing=east,powered=true]",
		nil,
		NewMapping{
			393: 5333,
			404: 5334,
			477: 5840,
			573: 5840,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=east,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9835,
			573: 9835,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9740,
			573: 9740,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 5706,
			393: 5202,
			404: 5203,
			477: 5706,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 6812,
			573: 6812,
			393: 6305,
			404: 6306,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9536,
			573: 9536,
		},
	},
	{
		"minecraft:jungle_wall_sign[facing=north,waterlogged=true]",
		nil,
		NewMapping{
			477: 3765,
			573: 3765,
		},
	},
	{
		"minecraft:brick_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   1731,
			393: 4343,
			404: 4344,
			477: 4847,
			573: 4847,
		},
	},
	{
		"minecraft:cyan_banner[rotation=13]",
		nil,
		NewMapping{
			404: 7012,
			477: 7518,
			573: 7518,
			393: 7011,
		},
	},
	{
		"minecraft:orange_bed[facing=east,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 776,
			404: 776,
			477: 1076,
			573: 1076,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=4,powered=true]",
		nil,
		NewMapping{
			393: 356,
			404: 356,
			477: 356,
			573: 356,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=11,south=up,west=none]",
		nil,
		NewMapping{
			477: 3309,
			573: 3309,
			393: 3005,
			404: 3006,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=east,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 6781,
			404: 6782,
			477: 7288,
			573: 7288,
		},
	},
	{
		"minecraft:sticky_piston[extended=false,facing=north]",
		nil,
		NewMapping{
			404: 1034,
			477: 1334,
			573: 1334,
			4:   466,
			393: 1034,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=7,south=side,west=up]",
		nil,
		NewMapping{
			393: 1818,
			404: 1819,
			477: 2122,
			573: 2122,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6588,
			404: 6589,
			477: 7095,
			573: 7095,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 6754,
			573: 6754,
			393: 6247,
			404: 6248,
		},
	},
	{
		"minecraft:jungle_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 5070,
			404: 5071,
			477: 5574,
			573: 5574,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=2,waterlogged=true]",
		nil,
		NewMapping{
			477: 3543,
			573: 3543,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10571,
			573: 10571,
		},
	},
	{
		"minecraft:infested_mossy_stone_bricks",
		nil,
		NewMapping{
			4:   1555,
			393: 3980,
			404: 3981,
			477: 4488,
			573: 4488,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=6,south=up,west=up]",
		nil,
		NewMapping{
			393: 2958,
			404: 2959,
			477: 3262,
			573: 3262,
		},
	},
	{
		"minecraft:cobblestone_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 5650,
			393: 5146,
			404: 5147,
			477: 5650,
		},
	},
	{
		"minecraft:dark_oak_door[facing=south,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7958,
			404: 7959,
			477: 8483,
			573: 8483,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=10,powered=false]",
		nil,
		NewMapping{
			573: 519,
			393: 519,
			404: 519,
			477: 519,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=7,south=up,west=side]",
		nil,
		NewMapping{
			404: 2681,
			477: 2984,
			573: 2984,
			393: 2680,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10808,
			573: 10808,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=18,powered=true]",
		nil,
		NewMapping{
			477: 884,
			573: 884,
		},
	},
	{
		"minecraft:comparator[facing=east,mode=compare,powered=false]",
		nil,
		NewMapping{
			4:   2403,
			393: 5648,
			404: 5649,
			477: 6155,
			573: 6155,
		},
	},
	{
		"minecraft:gray_shulker_box[facing=west]",
		nil,
		NewMapping{
			404: 8263,
			477: 8787,
			573: 8787,
			4:   3620,
			393: 8262,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=12,south=side,west=up]",
		nil,
		NewMapping{
			404: 1864,
			477: 2167,
			573: 2167,
			393: 1863,
		},
	},
	{
		"minecraft:fire[age=8,east=true,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			404: 1396,
			477: 1699,
			573: 1699,
			393: 1395,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=12,south=side,west=side]",
		nil,
		NewMapping{
			573: 2888,
			393: 2584,
			404: 2585,
			477: 2888,
		},
	},
	{
		"minecraft:diorite_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 10206,
			477: 10206,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10245,
			573: 10245,
		},
	},
	{
		"minecraft:oak_fence_gate[facing=north,in_wall=false,open=false,powered=false]",
		nil,
		NewMapping{
			4:   1714,
			393: 4307,
			404: 4308,
			477: 4811,
			573: 4811,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=13,powered=true]",
		nil,
		NewMapping{
			393: 524,
			404: 524,
			477: 524,
			573: 524,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9459,
			573: 9459,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10379,
			573: 10379,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=false,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 8061,
			404: 8062,
			477: 8586,
			573: 8586,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=true,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10754,
			573: 10754,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9528,
			573: 9528,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10381,
			573: 10381,
		},
	},
	{
		"minecraft:cobblestone_wall[east=true,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			404: 5156,
			477: 5659,
			573: 5659,
			393: 5155,
		},
	},
	{
		"minecraft:stone_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			477: 7800,
			573: 7800,
			393: 7293,
		},
	},
	{
		"minecraft:jungle_door[facing=north,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			477: 8334,
			573: 8334,
			393: 7809,
			404: 7810,
		},
	},
	{
		"minecraft:jungle_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5071,
			404: 5072,
			477: 5575,
			573: 5575,
		},
	},
	{
		"minecraft:bubble_coral[waterlogged=false]",
		nil,
		NewMapping{
			477: 8999,
			573: 8999,
			404: 8475,
		},
	},
	{
		"minecraft:acacia_button[face=wall,facing=west,powered=true]",
		nil,
		NewMapping{
			393: 5411,
			404: 5412,
			477: 5918,
			573: 5918,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=3,south=up,west=up]",
		nil,
		NewMapping{
			393: 1923,
			404: 1924,
			477: 2227,
			573: 2227,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9710,
			573: 9710,
		},
	},
	{
		"minecraft:grindstone[face=ceiling,facing=east]",
		nil,
		NewMapping{
			477: 11176,
			573: 11176,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9292,
			573: 9292,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=false,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			477: 4580,
			573: 4580,
			393: 4076,
			404: 4077,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=24,powered=true]",
		nil,
		NewMapping{
			393: 396,
			404: 396,
			477: 396,
			573: 396,
		},
	},
	{
		"minecraft:mossy_stone_brick_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			477: 10266,
			573: 10266,
		},
	},
	{
		"minecraft:yellow_shulker_box[facing=south]",
		nil,
		NewMapping{
			4:   3571,
			393: 8243,
			404: 8244,
			477: 8768,
			573: 8768,
		},
	},
	{
		"minecraft:fire[age=6,east=true,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 1330,
			404: 1331,
			477: 1634,
			573: 1634,
		},
	},
	{
		"minecraft:quartz_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 5756,
			404: 5757,
			477: 6263,
			573: 6263,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=11,south=side,west=none]",
		nil,
		NewMapping{
			393: 2432,
			404: 2433,
			477: 2736,
			573: 2736,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10529,
			573: 10529,
		},
	},
	{
		"minecraft:fire[age=3,east=false,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 1248,
			404: 1249,
			477: 1552,
			573: 1552,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 5175,
			477: 5678,
			573: 5678,
			393: 5174,
		},
	},
	{
		"minecraft:spruce_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4941,
			404: 4942,
			477: 5445,
			573: 5445,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9711,
			573: 9711,
		},
	},
	{
		"minecraft:andesite_stairs[facing=east,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9996,
			573: 9996,
		},
	},
	{
		"minecraft:birch_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			393: 7273,
			404: 7274,
			477: 7780,
			573: 7780,
		},
	},
	{
		"minecraft:fire[age=12,east=true,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1534,
			404: 1535,
			477: 1838,
			573: 1838,
		},
	},
	{
		"minecraft:jungle_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 5122,
			404: 5123,
			477: 5626,
			573: 5626,
		},
	},
	{
		"minecraft:stone_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9652,
			573: 9652,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=14,powered=false]",
		nil,
		NewMapping{
			477: 1027,
			573: 1027,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10171,
			573: 10171,
		},
	},
	{
		"minecraft:birch_sign[rotation=15,waterlogged=false]",
		nil,
		NewMapping{
			477: 3474,
			573: 3474,
		},
	},
	{
		"minecraft:brown_bed[facing=east,occupied=false,part=head]",
		nil,
		NewMapping{
			404: 954,
			477: 1254,
			573: 1254,
			393: 954,
		},
	},
	{
		"minecraft:oak_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 1706,
			404: 1707,
			477: 2010,
			573: 2010,
		},
	},
	{
		"minecraft:dark_oak_fence[east=false,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 8198,
			393: 7673,
			404: 7674,
			477: 8198,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=true,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10435,
			573: 10435,
		},
	},
	{
		"minecraft:acacia_wall_sign[facing=south,waterlogged=true]",
		nil,
		NewMapping{
			477: 3759,
			573: 3759,
		},
	},
	{
		"minecraft:birch_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 5533,
			573: 5533,
			393: 5029,
			404: 5030,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10887,
			573: 10887,
		},
	},
	{
		"minecraft:diorite_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10177,
			573: 10177,
		},
	},
	{
		"minecraft:detector_rail[powered=true,shape=ascending_east]",
		nil,
		NewMapping{
			573: 1318,
			4:   458,
			393: 1018,
			404: 1018,
			477: 1318,
		},
	},
	{
		"minecraft:red_sandstone_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			4:   2912,
			393: 7344,
			404: 7345,
			477: 7863,
			573: 7863,
		},
	},
	{
		"minecraft:acacia_door[facing=west,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			573: 8428,
			393: 7903,
			404: 7904,
			477: 8428,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=10,south=none,west=up]",
		nil,
		NewMapping{
			404: 2281,
			477: 2584,
			573: 2584,
			393: 2280,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=true,north=true,powered=false,south=false,west=true]",
		nil,
		NewMapping{
			393: 4857,
			404: 4858,
			477: 5361,
			573: 5361,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=2,powered=true]",
		nil,
		NewMapping{
			477: 752,
			573: 752,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10170,
			573: 10170,
		},
	},
	{
		"minecraft:bell[attachment=ceiling,facing=south,powered=false]",
		nil,
		NewMapping{
			573: 11209,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=14,south=up,west=up]",
		nil,
		NewMapping{
			393: 2022,
			404: 2023,
			477: 2326,
			573: 2326,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=east,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3897,
			404: 3898,
			477: 4401,
			573: 4401,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9576,
			477: 9576,
		},
	},
	{
		"minecraft:dead_horn_coral_wall_fan[facing=north,waterlogged=false]",
		nil,
		NewMapping{
			477: 9057,
			573: 9057,
			393: 8497,
			404: 8513,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6793,
			404: 6794,
			477: 7300,
			573: 7300,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=west,half=top,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			477: 4322,
			573: 4322,
			393: 3818,
			404: 3819,
		},
	},
	{
		"minecraft:gold_block",
		nil,
		NewMapping{
			4:   656,
			393: 1123,
			404: 1123,
			477: 1426,
			573: 1426,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 4671,
			404: 4672,
			477: 5175,
			573: 5175,
			4:   2054,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 5093,
			4:   1828,
			393: 4589,
			404: 4590,
			477: 5093,
		},
	},
	{
		"minecraft:acacia_wood[axis=x]",
		nil,
		NewMapping{
			573: 120,
			393: 120,
			404: 120,
			477: 120,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=0,south=up,west=up]",
		nil,
		NewMapping{
			393: 1896,
			404: 1897,
			477: 2200,
			573: 2200,
		},
	},
	{
		"minecraft:trapped_chest[facing=south,type=left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5588,
			404: 5589,
			477: 6095,
			573: 6095,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=15,powered=true]",
		nil,
		NewMapping{
			477: 828,
			573: 828,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=false,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10518,
			573: 10518,
		},
	},
	{
		"minecraft:birch_sign[rotation=15,waterlogged=true]",
		nil,
		NewMapping{
			477: 3473,
			573: 3473,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9332,
			573: 9332,
		},
	},
	{
		"minecraft:dark_oak_fence_gate[facing=west,in_wall=false,open=false,powered=true]",
		nil,
		NewMapping{
			4:   2985,
			393: 7507,
			404: 7508,
			477: 8032,
			573: 8032,
		},
	},
	{
		"minecraft:jungle_door[facing=south,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7823,
			404: 7824,
			477: 8348,
			573: 8348,
		},
	},
	{
		"minecraft:fire[age=6,east=false,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 1344,
			404: 1345,
			477: 1648,
			573: 1648,
		},
	},
	{
		"minecraft:fire[age=2,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1228,
			404: 1229,
			477: 1532,
			573: 1532,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9506,
			573: 9506,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=7,south=up,west=none]",
		nil,
		NewMapping{
			393: 2537,
			404: 2538,
			477: 2841,
			573: 2841,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 11031,
			573: 11031,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10079,
			573: 10079,
		},
	},
	{
		"minecraft:stone_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9664,
			573: 9664,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=7,waterlogged=false]",
		nil,
		NewMapping{
			573: 3554,
			477: 3554,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=1,powered=true]",
		nil,
		NewMapping{
			477: 900,
			573: 900,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=9,south=none,west=none]",
		nil,
		NewMapping{
			573: 2433,
			393: 2129,
			404: 2130,
			477: 2433,
		},
	},
	{
		"minecraft:spruce_door[facing=north,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7686,
			404: 7687,
			477: 8211,
			573: 8211,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 10431,
			477: 10431,
		},
	},
	{
		"minecraft:spruce_wall_sign[facing=west,waterlogged=true]",
		nil,
		NewMapping{
			477: 3745,
			573: 3745,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10349,
			573: 10349,
		},
	},
	{
		"minecraft:podzol[snowy=true]",
		nil,
		NewMapping{
			477: 12,
			573: 12,
			393: 12,
			404: 12,
		},
	},
	{
		"minecraft:birch_fence[east=true,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7559,
			404: 7560,
			477: 8084,
			573: 8084,
		},
	},
	{
		"minecraft:acacia_door[facing=east,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			404: 7929,
			477: 8453,
			573: 8453,
			393: 7928,
		},
	},
	{
		"minecraft:acacia_door[facing=north,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			4:   3139,
			393: 7884,
			404: 7885,
			477: 8409,
			573: 8409,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=south,in_wall=false,open=true,powered=false]",
		nil,
		NewMapping{
			573: 7959,
			4:   2964,
			393: 7434,
			404: 7435,
			477: 7959,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=1,powered=false]",
		nil,
		NewMapping{
			404: 301,
			477: 301,
			573: 301,
			393: 301,
		},
	},
	{
		"minecraft:white_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5835,
			404: 5836,
			477: 6342,
			573: 6342,
		},
	},
	{
		"minecraft:fire[age=3,east=true,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1231,
			404: 1232,
			477: 1535,
			573: 1535,
		},
	},
	{
		"minecraft:jungle_sign[rotation=11,waterlogged=false]",
		nil,
		NewMapping{
			477: 3530,
			573: 3530,
		},
	},
	{
		"minecraft:andesite_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9992,
			573: 9992,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=true,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10498,
			573: 10498,
		},
	},
	{
		"minecraft:dark_prismarine_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			393: 6818,
			404: 6819,
			477: 7325,
			573: 7325,
		},
	},
	{
		"minecraft:fire[age=10,east=true,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			404: 1458,
			477: 1761,
			573: 1761,
			393: 1457,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9286,
			573: 9286,
		},
	},
	{
		"minecraft:red_concrete_powder",
		nil,
		NewMapping{
			4:   4046,
			393: 8407,
			404: 8408,
			477: 8932,
			573: 8932,
		},
	},
	{
		"minecraft:spruce_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4957,
			404: 4958,
			477: 5461,
			573: 5461,
		},
	},
	{
		"minecraft:dead_bubble_coral_wall_fan[facing=south,waterlogged=true]",
		nil,
		NewMapping{
			477: 9042,
			573: 9042,
			393: 8482,
			404: 8498,
		},
	},
	{
		"minecraft:scaffolding[bottom=false,distance=0,waterlogged=true]",
		nil,
		NewMapping{
			477: 11115,
			573: 11115,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=10,south=none,west=side]",
		nil,
		NewMapping{
			393: 2569,
			404: 2570,
			477: 2873,
			573: 2873,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			573: 9173,
			477: 9173,
		},
	},
	{
		"minecraft:diorite_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 10214,
			573: 10214,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=17,powered=false]",
		nil,
		NewMapping{
			393: 633,
			404: 633,
			477: 633,
			573: 633,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=14,south=none,west=none]",
		nil,
		NewMapping{
			393: 2318,
			404: 2319,
			477: 2622,
			573: 2622,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5858,
			404: 5859,
			477: 6365,
			573: 6365,
		},
	},
	{
		"minecraft:fire[age=2,east=true,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			573: 1514,
			393: 1210,
			404: 1211,
			477: 1514,
		},
	},
	{
		"minecraft:dark_oak_door[facing=south,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7954,
			404: 7955,
			477: 8479,
			573: 8479,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 11016,
			573: 11016,
		},
	},
	{
		"minecraft:cobblestone_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			4:   707,
			393: 7314,
			404: 7315,
			477: 7833,
			573: 7833,
		},
	},
	{
		"minecraft:light_gray_bed[facing=south,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 881,
			404: 881,
			477: 1181,
			573: 1181,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=18,powered=true]",
		nil,
		NewMapping{
			393: 384,
			404: 384,
			477: 384,
			573: 384,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=8,south=side,west=side]",
		nil,
		NewMapping{
			393: 2548,
			404: 2549,
			477: 2852,
			573: 2852,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=1,south=none,west=side]",
		nil,
		NewMapping{
			393: 2776,
			404: 2777,
			477: 3080,
			573: 3080,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=0,powered=true]",
		nil,
		NewMapping{
			477: 848,
			573: 848,
		},
	},
	{
		"minecraft:vine[east=true,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			4:   1708,
			393: 4275,
			404: 4276,
			477: 4779,
			573: 4779,
		},
	},
	{
		"minecraft:purple_banner[rotation=7]",
		nil,
		NewMapping{
			393: 7021,
			404: 7022,
			477: 7528,
			573: 7528,
		},
	},
	{
		"minecraft:fire[age=3,east=true,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1238,
			404: 1239,
			477: 1542,
			573: 1542,
		},
	},
	{
		"minecraft:trapped_chest[facing=north,type=single,waterlogged=true]",
		nil,
		NewMapping{
			393: 5579,
			404: 5580,
			477: 6086,
			573: 6086,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=3,powered=false]",
		nil,
		NewMapping{
			393: 255,
			404: 255,
			477: 255,
			573: 255,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=2,south=up,west=up]",
		nil,
		NewMapping{
			393: 1770,
			404: 1771,
			477: 2074,
			573: 2074,
		},
	},
	{
		"minecraft:red_nether_brick_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			477: 10318,
			573: 10318,
		},
	},
	{
		"minecraft:dragon_egg",
		nil,
		NewMapping{
			477: 5139,
			573: 5139,
			4:   1952,
			393: 4635,
			404: 4636,
		},
	},
	{
		"minecraft:repeater[delay=4,facing=east,locked=false,powered=false]",
		nil,
		NewMapping{
			393: 3576,
			404: 3577,
			477: 4080,
			573: 4080,
			4:   1503,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4415,
			404: 4416,
			477: 4919,
			573: 4919,
		},
	},
	{
		"minecraft:quartz_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 5748,
			404: 5749,
			477: 6255,
			573: 6255,
		},
	},
	{
		"minecraft:fire[age=0,east=false,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			477: 1456,
			573: 1456,
			393: 1152,
			404: 1153,
		},
	},
	{
		"minecraft:andesite_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9985,
			573: 9985,
		},
	},
	{
		"minecraft:lantern[hanging=false]",
		nil,
		NewMapping{
			477: 11215,
			573: 11231,
		},
	},
	{
		"minecraft:bee_nest[facing=east,honey_level=4]",
		nil,
		NewMapping{
			573: 11309,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=6,south=up,west=side]",
		nil,
		NewMapping{
			477: 2975,
			573: 2975,
			393: 2671,
			404: 2672,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=20,powered=false]",
		nil,
		NewMapping{
			393: 439,
			404: 439,
			477: 439,
			573: 439,
		},
	},
	{
		"minecraft:chest[facing=south,type=left,waterlogged=true]",
		nil,
		NewMapping{
			393: 1736,
			404: 1737,
			477: 2040,
			573: 2040,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=8,south=up,west=up]",
		nil,
		NewMapping{
			393: 2832,
			404: 2833,
			477: 3136,
			573: 3136,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=3,south=side,west=none]",
		nil,
		NewMapping{
			477: 2376,
			573: 2376,
			393: 2072,
			404: 2073,
		},
	},
	{
		"minecraft:acacia_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			393: 7281,
			404: 7282,
			477: 7788,
			573: 7788,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=1,south=side,west=side]",
		nil,
		NewMapping{
			573: 2213,
			393: 1909,
			404: 1910,
			477: 2213,
		},
	},
	{
		"minecraft:smooth_stone_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			404: 7299,
			477: 7811,
			573: 7811,
		},
	},
	{
		"minecraft:spruce_wall_sign[facing=south,waterlogged=false]",
		nil,
		NewMapping{
			573: 3744,
			477: 3744,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10158,
			573: 10158,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=true,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 3991,
			404: 3992,
			477: 4495,
			573: 4495,
		},
	},
	{
		"minecraft:vine[east=false,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			404: 4290,
			477: 4793,
			573: 4793,
			393: 4289,
		},
	},
	{
		"minecraft:jungle_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 5604,
			393: 5100,
			404: 5101,
			477: 5604,
		},
	},
	{
		"minecraft:chest[facing=east,type=left,waterlogged=false]",
		nil,
		NewMapping{
			393: 1749,
			404: 1750,
			477: 2053,
			573: 2053,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=8,south=up,west=side]",
		nil,
		NewMapping{
			393: 2833,
			404: 2834,
			477: 3137,
			573: 3137,
		},
	},
	{
		"minecraft:nether_wart_block",
		nil,
		NewMapping{
			477: 8718,
			573: 8718,
			4:   3424,
			393: 8193,
			404: 8194,
		},
	},
	{
		"minecraft:vine[east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 4299,
			404: 4300,
			477: 4803,
			573: 4803,
			4:   1696,
		},
	},
	{
		"minecraft:spruce_button[face=floor,facing=west,powered=false]",
		nil,
		NewMapping{
			393: 5332,
			404: 5333,
			477: 5839,
			573: 5839,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9701,
			573: 9701,
		},
	},
	{
		"minecraft:redstone_wall_torch[facing=west,lit=false]",
		nil,
		NewMapping{
			4:   1202,
			393: 3388,
			404: 3389,
			477: 3892,
			573: 3892,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=3,south=side,west=up]",
		nil,
		NewMapping{
			393: 2502,
			404: 2503,
			477: 2806,
			573: 2806,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10984,
			573: 10984,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10912,
			477: 10912,
		},
	},
	{
		"minecraft:acacia_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6370,
			404: 6371,
			477: 6877,
			573: 6877,
		},
	},
	{
		"minecraft:jungle_door[facing=north,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7814,
			404: 7815,
			477: 8339,
			573: 8339,
		},
	},
	{
		"minecraft:jungle_stairs[facing=east,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5107,
			404: 5108,
			477: 5611,
			573: 5611,
		},
	},
	{
		"minecraft:fire[age=7,east=false,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1385,
			404: 1386,
			477: 1689,
			573: 1689,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=east,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3718,
			404: 3719,
			477: 4222,
			573: 4222,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=11,waterlogged=false]",
		nil,
		NewMapping{
			477: 3562,
			573: 3562,
		},
	},
	{
		"minecraft:diorite_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10180,
			573: 10180,
		},
	},
	{
		"minecraft:dark_oak_log[axis=z]",
		nil,
		NewMapping{
			4:   2601,
			393: 89,
			404: 89,
			477: 89,
			573: 89,
		},
	},
	{
		"minecraft:lava[level=3]",
		nil,
		NewMapping{
			404: 53,
			477: 53,
			573: 53,
			4:   179,
			393: 53,
		},
	},
	{
		"minecraft:fire[age=2,east=false,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			477: 1526,
			573: 1526,
			393: 1222,
			404: 1223,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6286,
			404: 6287,
			477: 6793,
			573: 6793,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10542,
			573: 10542,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=false,north=true,powered=true,south=true,west=true]",
		nil,
		NewMapping{
			393: 4803,
			404: 4804,
			477: 5307,
			573: 5307,
		},
	},
	{
		"minecraft:spruce_door[facing=south,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7695,
			404: 7696,
			477: 8220,
			573: 8220,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9700,
			573: 9700,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 11013,
			573: 11013,
		},
	},
	{
		"minecraft:light_blue_bed[facing=east,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 809,
			404: 809,
			477: 1109,
			573: 1109,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 4650,
			404: 4651,
			477: 5154,
			573: 5154,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 7162,
			573: 7162,
			393: 6655,
			404: 6656,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=4,south=none,west=none]",
		nil,
		NewMapping{
			404: 2085,
			477: 2388,
			573: 2388,
			393: 2084,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=true,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10878,
			573: 10878,
		},
	},
	{
		"minecraft:birch_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			573: 7781,
			4:   2002,
			393: 7274,
			404: 7275,
			477: 7781,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=2,powered=false]",
		nil,
		NewMapping{
			573: 603,
			393: 603,
			404: 603,
			477: 603,
		},
	},
	{
		"minecraft:dark_oak_leaves[distance=6,persistent=true]",
		nil,
		NewMapping{
			393: 224,
			404: 224,
			477: 224,
			573: 224,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=4,south=side,west=none]",
		nil,
		NewMapping{
			393: 2945,
			404: 2946,
			477: 3249,
			573: 3249,
		},
	},
	{
		"minecraft:daylight_detector[inverted=true,power=13]",
		nil,
		NewMapping{
			4:   2861,
			393: 5664,
			404: 5665,
			477: 6171,
			573: 6171,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=north,in_wall=false,open=true,powered=false]",
		nil,
		NewMapping{
			4:   2966,
			393: 7426,
			404: 7427,
			477: 7951,
			573: 7951,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=5,south=side,west=up]",
		nil,
		NewMapping{
			393: 2520,
			404: 2521,
			477: 2824,
			573: 2824,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9859,
			573: 9859,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=0,powered=false]",
		nil,
		NewMapping{
			477: 849,
			573: 849,
		},
	},
	{
		"minecraft:fire[age=1,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			404: 1169,
			477: 1472,
			573: 1472,
			393: 1168,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=false,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 11001,
			573: 11001,
		},
	},
	{
		"minecraft:barrel[facing=south,open=false]",
		nil,
		NewMapping{
			477: 11140,
			573: 11140,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=3,south=up,west=up]",
		nil,
		NewMapping{
			393: 2931,
			404: 2932,
			477: 3235,
			573: 3235,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=north,half=top,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3599,
			404: 3600,
			477: 4103,
			573: 4103,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=5,powered=true]",
		nil,
		NewMapping{
			393: 708,
			404: 708,
			477: 708,
			573: 708,
		},
	},
	{
		"minecraft:campfire[facing=west,lit=true,signal_fire=true,waterlogged=true]",
		nil,
		NewMapping{
			477: 11232,
			573: 11248,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10512,
			573: 10512,
		},
	},
	{
		"minecraft:iron_door[facing=north,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			573: 3822,
			4:   1139,
			393: 3318,
			404: 3319,
			477: 3822,
		},
	},
	{
		"minecraft:pink_stained_glass",
		nil,
		NewMapping{
			4:   1526,
			393: 3583,
			404: 3584,
			477: 4087,
			573: 4087,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6140,
			404: 6141,
			477: 6647,
			573: 6647,
		},
	},
	{
		"minecraft:acacia_fence[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 7620,
			404: 7621,
			477: 8145,
			573: 8145,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=9,south=up,west=none]",
		nil,
		NewMapping{
			393: 2555,
			404: 2556,
			477: 2859,
			573: 2859,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=true,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5989,
			404: 5990,
			477: 6496,
			573: 6496,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=23,powered=false]",
		nil,
		NewMapping{
			393: 545,
			404: 545,
			477: 545,
			573: 545,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10948,
			573: 10948,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9763,
			573: 9763,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10719,
			573: 10719,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10509,
			573: 10509,
		},
	},
	{
		"minecraft:acacia_button[face=floor,facing=west,powered=false]",
		nil,
		NewMapping{
			477: 5911,
			573: 5911,
			393: 5404,
			404: 5405,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=0,powered=true]",
		nil,
		NewMapping{
			404: 398,
			477: 398,
			573: 398,
			393: 398,
		},
	},
	{
		"minecraft:fire[age=5,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1324,
			404: 1325,
			477: 1628,
			573: 1628,
		},
	},
	{
		"minecraft:fire[age=1,east=true,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 1170,
			404: 1171,
			477: 1474,
			573: 1474,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10708,
			573: 10708,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9492,
			573: 9492,
		},
	},
	{
		"minecraft:purpur_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			393: 7347,
			404: 7348,
			477: 7872,
			573: 7872,
		},
	},
	{
		"minecraft:iron_bars[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 4683,
			393: 4179,
			404: 4180,
			477: 4683,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=0,south=up,west=side]",
		nil,
		NewMapping{
			393: 2617,
			404: 2618,
			477: 2921,
			573: 2921,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=17,powered=false]",
		nil,
		NewMapping{
			573: 883,
			477: 883,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9872,
			573: 9872,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=9,powered=true]",
		nil,
		NewMapping{
			477: 966,
			573: 966,
		},
	},
	{
		"minecraft:oxeye_daisy",
		nil,
		NewMapping{
			393: 1120,
			404: 1120,
			477: 1420,
			573: 1420,
			4:   616,
		},
	},
	{
		"minecraft:iron_door[facing=east,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			4:   1146,
			393: 3353,
			404: 3354,
			477: 3857,
			573: 3857,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=true,north=false,powered=true,south=true,west=true]",
		nil,
		NewMapping{
			404: 4796,
			477: 5299,
			573: 5299,
			393: 4795,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6313,
			404: 6314,
			477: 6820,
			573: 6820,
		},
	},
	{
		"minecraft:trapped_chest[facing=north,type=right,waterlogged=false]",
		nil,
		NewMapping{
			477: 6091,
			573: 6091,
			393: 5584,
			404: 5585,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10495,
			573: 10495,
		},
	},
	{
		"minecraft:diorite_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10178,
			573: 10178,
		},
	},
	{
		"minecraft:birch_wall_sign[facing=south,waterlogged=false]",
		nil,
		NewMapping{
			477: 3752,
			573: 3752,
		},
	},
	{
		"minecraft:command_block[conditional=false,facing=north]",
		nil,
		NewMapping{
			404: 5131,
			477: 5634,
			573: 5634,
			4:   2194,
			393: 5130,
		},
	},
	{
		"minecraft:wall_sign[facing=west,waterlogged=true]",
		nil,
		NewMapping{
			393: 3273,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=5,south=side,west=none]",
		nil,
		NewMapping{
			393: 2090,
			404: 2091,
			477: 2394,
			573: 2394,
		},
	},
	{
		"minecraft:chest[facing=south,type=left,waterlogged=false]",
		nil,
		NewMapping{
			393: 1737,
			404: 1738,
			477: 2041,
			573: 2041,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=true,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5244,
			404: 5245,
			477: 5748,
			573: 5748,
		},
	},
	{
		"minecraft:end_stone_bricks",
		nil,
		NewMapping{
			4:   3296,
			393: 8157,
			404: 8158,
			477: 8682,
			573: 8682,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=1,south=up,west=side]",
		nil,
		NewMapping{
			573: 2786,
			393: 2482,
			404: 2483,
			477: 2786,
		},
	},
	{
		"minecraft:tripwire_hook[attached=true,facing=west,powered=true]",
		nil,
		NewMapping{
			573: 5247,
			4:   2109,
			393: 4743,
			404: 4744,
			477: 5247,
		},
	},
	{
		"minecraft:stone_stairs[facing=west,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9663,
			573: 9663,
		},
	},
	{
		"minecraft:andesite_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10010,
			573: 10010,
		},
	},
	{
		"minecraft:andesite_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9969,
			573: 9969,
		},
	},
	{
		"minecraft:bone_block[axis=x]",
		nil,
		NewMapping{
			4:   3460,
			393: 8195,
			404: 8196,
			477: 8720,
			573: 8720,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=10,south=none,west=none]",
		nil,
		NewMapping{
			393: 1850,
			404: 1851,
			477: 2154,
			573: 2154,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 10993,
			477: 10993,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 9805,
			477: 9805,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=0,south=side,west=up]",
		nil,
		NewMapping{
			393: 2763,
			404: 2764,
			477: 3067,
			573: 3067,
		},
	},
	{
		"minecraft:oak_sign[rotation=0,waterlogged=true]",
		nil,
		NewMapping{
			404: 3076,
			477: 3379,
			573: 3379,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=11,powered=true]",
		nil,
		NewMapping{
			477: 820,
			573: 820,
		},
	},
	{
		"minecraft:blue_terracotta",
		nil,
		NewMapping{
			4:   2555,
			393: 5815,
			404: 5816,
			477: 6322,
			573: 6322,
		},
	},
	{
		"minecraft:tripwire_hook[attached=false,facing=south,powered=true]",
		nil,
		NewMapping{
			404: 4750,
			477: 5253,
			573: 5253,
			4:   2104,
			393: 4749,
		},
	},
	{
		"minecraft:acacia_button[face=floor,facing=north,powered=false]",
		nil,
		NewMapping{
			477: 5907,
			573: 5907,
			393: 5400,
			404: 5401,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=10,south=up,west=side]",
		nil,
		NewMapping{
			393: 2275,
			404: 2276,
			477: 2579,
			573: 2579,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 5689,
			573: 5689,
			393: 5185,
			404: 5186,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10975,
			573: 10975,
		},
	},
	{
		"minecraft:bell[attachment=floor,facing=east,powered=false]",
		nil,
		NewMapping{
			573: 11205,
		},
	},
	{
		"minecraft:trapped_chest[facing=south,type=single,waterlogged=false]",
		nil,
		NewMapping{
			573: 6093,
			4:   2339,
			393: 5586,
			404: 5587,
			477: 6093,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=true,north=false,powered=true,south=true,west=true]",
		nil,
		NewMapping{
			404: 4828,
			477: 5331,
			573: 5331,
			393: 4827,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=2,powered=false]",
		nil,
		NewMapping{
			393: 653,
			404: 653,
			477: 653,
			573: 653,
		},
	},
	{
		"minecraft:fire[age=12,east=true,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1532,
			404: 1533,
			477: 1836,
			573: 1836,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9834,
			573: 9834,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=north,half=top,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			4:   1544,
			393: 3600,
			404: 3601,
			477: 4104,
			573: 4104,
		},
	},
	{
		"minecraft:jungle_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 5090,
			404: 5091,
			477: 5594,
			573: 5594,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6737,
			404: 6738,
			477: 7244,
			573: 7244,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 6218,
			477: 6724,
			573: 6724,
			393: 6217,
		},
	},
	{
		"minecraft:andesite_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 10012,
			477: 10012,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=17,powered=false]",
		nil,
		NewMapping{
			477: 783,
			573: 783,
		},
	},
	{
		"minecraft:oak_fence_gate[facing=west,in_wall=false,open=false,powered=true]",
		nil,
		NewMapping{
			573: 4826,
			4:   1721,
			393: 4322,
			404: 4323,
			477: 4826,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=19,powered=true]",
		nil,
		NewMapping{
			393: 536,
			404: 536,
			477: 536,
			573: 536,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5952,
			404: 5953,
			477: 6459,
			573: 6459,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=9,south=up,west=side]",
		nil,
		NewMapping{
			477: 2714,
			573: 2714,
			393: 2410,
			404: 2411,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10981,
			573: 10981,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10859,
			573: 10859,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=east,in_wall=true,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7480,
			404: 7481,
			477: 8005,
			573: 8005,
		},
	},
	{
		"minecraft:acacia_stairs[facing=west,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			404: 6383,
			477: 6889,
			573: 6889,
			393: 6382,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=south,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 6741,
			404: 6742,
			477: 7248,
			573: 7248,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=true,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10694,
			573: 10694,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10603,
			573: 10603,
		},
	},
	{
		"minecraft:chorus_flower[age=1]",
		nil,
		NewMapping{
			4:   3201,
			393: 8068,
			404: 8069,
			477: 8593,
			573: 8593,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=12,south=side,west=none]",
		nil,
		NewMapping{
			393: 2153,
			404: 2154,
			477: 2457,
			573: 2457,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=9,south=none,west=none]",
		nil,
		NewMapping{
			404: 1986,
			477: 2289,
			573: 2289,
			393: 1985,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9304,
			573: 9304,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=11,south=up,west=up]",
		nil,
		NewMapping{
			393: 1995,
			404: 1996,
			477: 2299,
			573: 2299,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=true,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 4033,
			404: 4034,
			477: 4537,
			573: 4537,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9844,
			573: 9844,
		},
	},
	{
		"minecraft:red_tulip",
		nil,
		NewMapping{
			4:   612,
			393: 1116,
			404: 1116,
			477: 1416,
			573: 1416,
		},
	},
	{
		"minecraft:birch_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 5041,
			477: 5544,
			573: 5544,
			393: 5040,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=4,south=side,west=none]",
		nil,
		NewMapping{
			393: 2081,
			404: 2082,
			477: 2385,
			573: 2385,
		},
	},
	{
		"minecraft:spruce_door[facing=north,half=upper,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7684,
			404: 7685,
			477: 8209,
			573: 8209,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=false,north=false,powered=true,south=true,west=false]",
		nil,
		NewMapping{
			393: 4780,
			404: 4781,
			477: 5284,
			573: 5284,
		},
	},
	{
		"minecraft:jungle_fence[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7583,
			404: 7584,
			477: 8108,
			573: 8108,
		},
	},
	{
		"minecraft:iron_door[facing=west,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			573: 3849,
			393: 3345,
			404: 3346,
			477: 3849,
		},
	},
	{
		"minecraft:jungle_sign[rotation=7,waterlogged=true]",
		nil,
		NewMapping{
			477: 3521,
			573: 3521,
		},
	},
	{
		"minecraft:oak_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 1713,
			404: 1714,
			477: 2017,
			573: 2017,
		},
	},
	{
		"minecraft:dark_oak_planks",
		nil,
		NewMapping{
			573: 20,
			4:   85,
			393: 20,
			404: 20,
			477: 20,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4587,
			404: 4588,
			477: 5091,
			573: 5091,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=10,south=up,west=up]",
		nil,
		NewMapping{
			393: 1842,
			404: 1843,
			477: 2146,
			573: 2146,
		},
	},
	{
		"minecraft:pink_bed[facing=north,occupied=false,part=head]",
		nil,
		NewMapping{
			477: 1146,
			573: 1146,
			393: 846,
			404: 846,
		},
	},
	{
		"minecraft:brown_banner[rotation=15]",
		nil,
		NewMapping{
			573: 7568,
			393: 7061,
			404: 7062,
			477: 7568,
		},
	},
	{
		"minecraft:activator_rail[powered=false,shape=ascending_south]",
		nil,
		NewMapping{
			4:   2517,
			393: 5791,
			404: 5792,
			477: 6298,
			573: 6298,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6167,
			404: 6168,
			477: 6674,
			573: 6674,
		},
	},
	{
		"minecraft:magenta_bed[facing=east,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 795,
			404: 795,
			477: 1095,
			573: 1095,
		},
	},
	{
		"minecraft:fire[age=5,east=false,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1318,
			404: 1319,
			477: 1622,
			573: 1622,
		},
	},
	{
		"minecraft:wither_skeleton_skull[rotation=4]",
		nil,
		NewMapping{
			573: 5978,
			393: 5475,
			404: 5476,
			477: 5978,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=5,south=side,west=side]",
		nil,
		NewMapping{
			393: 2953,
			404: 2954,
			477: 3257,
			573: 3257,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=true,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 6432,
			573: 6432,
			393: 5925,
			404: 5926,
		},
	},
	{
		"minecraft:fire[age=13,east=true,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1551,
			404: 1552,
			477: 1855,
			573: 1855,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=10,south=side,west=side]",
		nil,
		NewMapping{
			393: 2422,
			404: 2423,
			477: 2726,
			573: 2726,
		},
	},
	{
		"minecraft:turtle_egg[eggs=3,hatch=0]",
		nil,
		NewMapping{
			573: 8968,
			393: 8443,
			404: 8444,
			477: 8968,
		},
	},
	{
		"minecraft:spruce_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			393: 7267,
			404: 7268,
			477: 7774,
			573: 7774,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=0,south=none,west=side]",
		nil,
		NewMapping{
			393: 1903,
			404: 1904,
			477: 2207,
			573: 2207,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9154,
			573: 9154,
		},
	},
	{
		"minecraft:smooth_sandstone_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			573: 10289,
			477: 10289,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=north,in_wall=false,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7395,
			404: 7396,
			477: 7920,
			573: 7920,
			4:   2954,
		},
	},
	{
		"minecraft:cyan_bed[facing=south,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 896,
			404: 896,
			477: 1196,
			573: 1196,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=west,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			477: 7041,
			573: 7041,
			393: 6534,
			404: 6535,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6100,
			404: 6101,
			477: 6607,
			573: 6607,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=true,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 10826,
			477: 10826,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4583,
			404: 4584,
			477: 5087,
			573: 5087,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 3199,
			404: 3200,
			477: 3663,
			573: 3663,
		},
	},
	{
		"minecraft:fire[age=15,east=false,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			573: 1943,
			393: 1639,
			404: 1640,
			477: 1943,
		},
	},
	{
		"minecraft:black_banner[rotation=6]",
		nil,
		NewMapping{
			393: 7100,
			404: 7101,
			477: 7607,
			573: 7607,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10631,
			573: 10631,
		},
	},
	{
		"minecraft:barrel[facing=down,open=false]",
		nil,
		NewMapping{
			477: 11146,
			573: 11146,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=false,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 11030,
			573: 11030,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=20,powered=false]",
		nil,
		NewMapping{
			393: 639,
			404: 639,
			477: 639,
			573: 639,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=false,north=false,powered=true,south=true,west=true]",
		nil,
		NewMapping{
			393: 4811,
			404: 4812,
			477: 5315,
			573: 5315,
		},
	},
	{
		"minecraft:pink_bed[facing=east,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 859,
			404: 859,
			477: 1159,
			573: 1159,
		},
	},
	{
		"minecraft:birch_button[face=wall,facing=north,powered=false]",
		nil,
		NewMapping{
			404: 5361,
			477: 5867,
			573: 5867,
			393: 5360,
		},
	},
	{
		"minecraft:red_banner[rotation=10]",
		nil,
		NewMapping{
			393: 7088,
			404: 7089,
			477: 7595,
			573: 7595,
		},
	},
	{
		"minecraft:light_weighted_pressure_plate[power=13]",
		nil,
		NewMapping{
			393: 5616,
			404: 5617,
			477: 6123,
			573: 6123,
			4:   2365,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=3,south=none,west=side]",
		nil,
		NewMapping{
			393: 1786,
			404: 1787,
			477: 2090,
			573: 2090,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=10,south=up,west=up]",
		nil,
		NewMapping{
			477: 2866,
			573: 2866,
			393: 2562,
			404: 2563,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=false,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10874,
			573: 10874,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=15,south=up,west=none]",
		nil,
		NewMapping{
			393: 2177,
			404: 2178,
			477: 2481,
			573: 2481,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=5,powered=false]",
		nil,
		NewMapping{
			393: 659,
			404: 659,
			477: 659,
			573: 659,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=false,north=true,powered=true,south=false,west=false]",
		nil,
		NewMapping{
			404: 4807,
			477: 5310,
			573: 5310,
			393: 4806,
		},
	},
	{
		"minecraft:acacia_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 6368,
			477: 6874,
			573: 6874,
			393: 6367,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6115,
			404: 6116,
			477: 6622,
			573: 6622,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5237,
			404: 5238,
			477: 5741,
			573: 5741,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=false,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			573: 4521,
			393: 4017,
			404: 4018,
			477: 4521,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=9,south=none,west=side]",
		nil,
		NewMapping{
			404: 2417,
			477: 2720,
			573: 2720,
			393: 2416,
		},
	},
	{
		"minecraft:birch_wall_sign[facing=east,waterlogged=false]",
		nil,
		NewMapping{
			477: 3756,
			573: 3756,
		},
	},
	{
		"minecraft:vine[east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			573: 4790,
			4:   1703,
			393: 4286,
			404: 4287,
			477: 4790,
		},
	},
	{
		"minecraft:white_carpet",
		nil,
		NewMapping{
			4:   2736,
			393: 6823,
			404: 6824,
			477: 7330,
			573: 7330,
		},
	},
	{
		"minecraft:spruce_fence[east=true,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 8046,
			393: 7521,
			404: 7522,
			477: 8046,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=8,south=up,west=none]",
		nil,
		NewMapping{
			393: 2258,
			404: 2259,
			477: 2562,
			573: 2562,
		},
	},
	{
		"minecraft:diorite_stairs[facing=north,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10175,
			573: 10175,
		},
	},
	{
		"minecraft:fire[age=5,east=true,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1301,
			404: 1302,
			477: 1605,
			573: 1605,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=10,south=none,west=up]",
		nil,
		NewMapping{
			393: 3000,
			404: 3001,
			477: 3304,
			573: 3304,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			573: 4620,
			393: 4116,
			404: 4117,
			477: 4620,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=east,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6783,
			404: 6784,
			477: 7290,
			573: 7290,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5955,
			404: 5956,
			477: 6462,
			573: 6462,
		},
	},
	{
		"minecraft:oak_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 1682,
			404: 1683,
			477: 1986,
			573: 1986,
		},
	},
	{
		"minecraft:cake[bites=0]",
		nil,
		NewMapping{
			404: 3507,
			477: 4010,
			573: 4010,
			4:   1472,
			393: 3506,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 7256,
			404: 7257,
			477: 7763,
			573: 7763,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			404: 6103,
			477: 6609,
			573: 6609,
			393: 6102,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=7,south=none,west=up]",
		nil,
		NewMapping{
			393: 2397,
			404: 2398,
			477: 2701,
			573: 2701,
		},
	},
	{
		"minecraft:horn_coral_wall_fan[facing=west,waterlogged=false]",
		nil,
		NewMapping{
			393: 8541,
			404: 8557,
			477: 9101,
			573: 9101,
		},
	},
	{
		"minecraft:brown_glazed_terracotta[facing=south]",
		nil,
		NewMapping{
			4:   3952,
			393: 8362,
			404: 8363,
			477: 8887,
			573: 8887,
		},
	},
	{
		"minecraft:observer[facing=west,powered=false]",
		nil,
		NewMapping{
			4:   3492,
			393: 8206,
			404: 8207,
			477: 8731,
			573: 8731,
		},
	},
	{
		"minecraft:chorus_plant[down=true,east=false,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			477: 8553,
			573: 8553,
			393: 8028,
			404: 8029,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10957,
			573: 10957,
		},
	},
	{
		"minecraft:oak_button[face=wall,facing=east,powered=false]",
		nil,
		NewMapping{
			4:   2289,
			393: 5318,
			404: 5319,
			477: 5825,
			573: 5825,
		},
	},
	{
		"minecraft:fire[age=6,east=false,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			477: 1650,
			573: 1650,
			393: 1346,
			404: 1347,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=9,south=up,west=up]",
		nil,
		NewMapping{
			393: 2265,
			404: 2266,
			477: 2569,
			573: 2569,
		},
	},
	{
		"minecraft:fire[age=13,east=false,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			573: 1880,
			393: 1576,
			404: 1577,
			477: 1880,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10677,
			573: 10677,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10997,
			573: 10997,
		},
	},
	{
		"minecraft:stone_bricks",
		nil,
		NewMapping{
			393: 3983,
			404: 3984,
			477: 4481,
			573: 4481,
			4:   1568,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4430,
			404: 4431,
			477: 4934,
			573: 4934,
		},
	},
	{
		"minecraft:fire[age=8,east=true,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1401,
			404: 1402,
			477: 1705,
			573: 1705,
		},
	},
	{
		"minecraft:fire[age=5,east=false,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1311,
			404: 1312,
			477: 1615,
			573: 1615,
		},
	},
	{
		"minecraft:oak_door[facing=west,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 3153,
			404: 3154,
			477: 3617,
			573: 3617,
		},
	},
	{
		"minecraft:spruce_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4913,
			404: 4914,
			477: 5417,
			573: 5417,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=10,powered=true]",
		nil,
		NewMapping{
			393: 268,
			404: 268,
			477: 268,
			573: 268,
		},
	},
	{
		"minecraft:oak_fence_gate[facing=south,in_wall=true,open=false,powered=true]",
		nil,
		NewMapping{
			393: 4310,
			404: 4311,
			477: 4814,
			573: 4814,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=24,powered=false]",
		nil,
		NewMapping{
			573: 997,
			477: 997,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10977,
			573: 10977,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=13,south=side,west=none]",
		nil,
		NewMapping{
			477: 2610,
			573: 2610,
			393: 2306,
			404: 2307,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=1,south=none,west=up]",
		nil,
		NewMapping{
			477: 3079,
			573: 3079,
			393: 2775,
			404: 2776,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=5,south=none,west=side]",
		nil,
		NewMapping{
			573: 3116,
			393: 2812,
			404: 2813,
			477: 3116,
		},
	},
	{
		"minecraft:stone_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9680,
			573: 9680,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10576,
			477: 10576,
		},
	},
	{
		"minecraft:campfire[facing=north,lit=false,signal_fire=false,waterlogged=true]",
		nil,
		NewMapping{
			477: 11222,
			573: 11238,
		},
	},
	{
		"minecraft:purpur_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 8614,
			393: 8089,
			404: 8090,
			477: 8614,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=true,north=false,powered=true,south=false,west=false]",
		nil,
		NewMapping{
			393: 4766,
			404: 4767,
			477: 5270,
			573: 5270,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 7236,
			404: 7237,
			477: 7743,
			573: 7743,
		},
	},
	{
		"minecraft:diorite_stairs[facing=west,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10222,
			573: 10222,
		},
	},
	{
		"minecraft:stone_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9677,
			573: 9677,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=4,south=up,west=none]",
		nil,
		NewMapping{
			404: 1791,
			477: 2094,
			573: 2094,
			393: 1790,
		},
	},
	{
		"minecraft:iron_door[facing=west,half=upper,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 3335,
			404: 3336,
			477: 3839,
			573: 3839,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9211,
			573: 9211,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 6692,
			573: 6692,
			393: 6185,
			404: 6186,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=east,half=top,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3964,
			404: 3965,
			477: 4468,
			573: 4468,
		},
	},
	{
		"minecraft:repeater[delay=2,facing=south,locked=true,powered=true]",
		nil,
		NewMapping{
			393: 3533,
			404: 3534,
			477: 4037,
			573: 4037,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=north,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3787,
			404: 3788,
			477: 4291,
			573: 4291,
		},
	},
	{
		"minecraft:brick_stairs[facing=south,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 4861,
			393: 4357,
			404: 4358,
			477: 4861,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=west,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			404: 7438,
			477: 7962,
			573: 7962,
			393: 7437,
		},
	},
	{
		"minecraft:fire[age=11,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			573: 1809,
			393: 1505,
			404: 1506,
			477: 1809,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 9599,
			477: 9599,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10116,
			573: 10116,
		},
	},
	{
		"minecraft:andesite_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9941,
			573: 9941,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=6,south=none,west=up]",
		nil,
		NewMapping{
			477: 3124,
			573: 3124,
			393: 2820,
			404: 2821,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 6744,
			393: 6237,
			404: 6238,
			477: 6744,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6713,
			404: 6714,
			477: 7220,
			573: 7220,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10380,
			573: 10380,
		},
	},
	{
		"minecraft:light_gray_shulker_box[facing=up]",
		nil,
		NewMapping{
			477: 8794,
			573: 8794,
			4:   3633,
			393: 8269,
			404: 8270,
		},
	},
	{
		"minecraft:quartz_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   2499,
			393: 5711,
			404: 5712,
			477: 6218,
			573: 6218,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=true,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10601,
			573: 10601,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9387,
			573: 9387,
		},
	},
	{
		"minecraft:light_gray_glazed_terracotta[facing=south]",
		nil,
		NewMapping{
			477: 8871,
			573: 8871,
			4:   3888,
			393: 8346,
			404: 8347,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6120,
			404: 6121,
			477: 6627,
			573: 6627,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=true,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 4097,
			404: 4098,
			477: 4601,
			573: 4601,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 10124,
			477: 10124,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9421,
			573: 9421,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9479,
			573: 9479,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=false,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10489,
			573: 10489,
		},
	},
	{
		"minecraft:jungle_stairs[facing=south,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 5074,
			404: 5075,
			477: 5578,
			573: 5578,
		},
	},
	{
		"minecraft:quartz_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 5716,
			404: 5717,
			477: 6223,
			573: 6223,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=false,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 4072,
			404: 4073,
			477: 4576,
			573: 4576,
		},
	},
	{
		"minecraft:purpur_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 8107,
			477: 8631,
			573: 8631,
			393: 8106,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=5,powered=true]",
		nil,
		NewMapping{
			393: 508,
			404: 508,
			477: 508,
			573: 508,
		},
	},
	{
		"minecraft:campfire[facing=north,lit=true,signal_fire=true,waterlogged=false]",
		nil,
		NewMapping{
			477: 11217,
			573: 11233,
		},
	},
	{
		"minecraft:birch_fence[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 7573,
			404: 7574,
			477: 8098,
			573: 8098,
		},
	},
	{
		"minecraft:creeper_head[rotation=3]",
		nil,
		NewMapping{
			393: 5534,
			404: 5535,
			477: 6037,
			573: 6037,
		},
	},
	{
		"minecraft:oak_fence[east=true,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 3967,
			393: 3463,
			404: 3464,
			477: 3967,
		},
	},
	{
		"minecraft:cauldron[level=2]",
		nil,
		NewMapping{
			573: 5127,
			4:   1890,
			393: 4623,
			404: 4624,
			477: 5127,
		},
	},
	{
		"minecraft:pink_banner[rotation=2]",
		nil,
		NewMapping{
			477: 7459,
			573: 7459,
			393: 6952,
			404: 6953,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6088,
			404: 6089,
			477: 6595,
			573: 6595,
		},
	},
	{
		"minecraft:jungle_leaves[distance=7,persistent=true]",
		nil,
		NewMapping{
			404: 198,
			477: 198,
			573: 198,
			393: 198,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 5160,
			393: 4656,
			404: 4657,
			477: 5160,
		},
	},
	{
		"minecraft:jungle_sign[rotation=2,waterlogged=false]",
		nil,
		NewMapping{
			477: 3512,
			573: 3512,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=true,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5184,
			404: 5185,
			477: 5688,
			573: 5688,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 7296,
			393: 6789,
			404: 6790,
			477: 7296,
		},
	},
	{
		"minecraft:dark_oak_leaves[distance=3,persistent=false]",
		nil,
		NewMapping{
			573: 219,
			393: 219,
			404: 219,
			477: 219,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=west,half=top,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3696,
			404: 3697,
			477: 4200,
			573: 4200,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6160,
			404: 6161,
			477: 6667,
			573: 6667,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10877,
			573: 10877,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10744,
			573: 10744,
		},
	},
	{
		"minecraft:green_shulker_box[facing=north]",
		nil,
		NewMapping{
			4:   3714,
			393: 8295,
			404: 8296,
			477: 8820,
			573: 8820,
		},
	},
	{
		"minecraft:dark_oak_fence_gate[facing=south,in_wall=false,open=false,powered=false]",
		nil,
		NewMapping{
			573: 8025,
			4:   2976,
			393: 7500,
			404: 7501,
			477: 8025,
		},
	},
	{
		"minecraft:oak_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 1657,
			404: 1658,
			477: 1961,
			573: 1961,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 4133,
			404: 4134,
			477: 4637,
			573: 4637,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=west,half=top,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3756,
			404: 3757,
			477: 4260,
			573: 4260,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9552,
			573: 9552,
		},
	},
	{
		"minecraft:grass_path",
		nil,
		NewMapping{
			4:   3328,
			393: 8162,
			404: 8163,
			477: 8687,
			573: 8687,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=false,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			404: 6035,
			477: 6541,
			573: 6541,
			393: 6034,
		},
	},
	{
		"minecraft:nether_brick_fence[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 4511,
			404: 4512,
			477: 5015,
			573: 5015,
		},
	},
	{
		"minecraft:oak_fence[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 3993,
			393: 3489,
			404: 3490,
			477: 3993,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10606,
			573: 10606,
		},
	},
	{
		"minecraft:oak_fence_gate[facing=north,in_wall=false,open=true,powered=false]",
		nil,
		NewMapping{
			393: 4305,
			404: 4306,
			477: 4809,
			573: 4809,
			4:   1718,
		},
	},
	{
		"minecraft:potatoes[age=0]",
		nil,
		NewMapping{
			4:   2272,
			393: 5295,
			404: 5296,
			477: 5802,
			573: 5802,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4702,
			404: 4703,
			477: 5206,
			573: 5206,
		},
	},
	{
		"minecraft:acacia_door[facing=west,half=upper,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			573: 8433,
			393: 7908,
			404: 7909,
			477: 8433,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=20,powered=true]",
		nil,
		NewMapping{
			393: 688,
			404: 688,
			477: 688,
			573: 688,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10898,
			573: 10898,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9250,
			477: 9250,
		},
	},
	{
		"minecraft:daylight_detector[inverted=true,power=11]",
		nil,
		NewMapping{
			404: 5663,
			477: 6169,
			573: 6169,
			4:   2859,
			393: 5662,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=0,powered=false]",
		nil,
		NewMapping{
			393: 299,
			404: 299,
			477: 299,
			573: 299,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=false,north=true,powered=false,south=true,west=true]",
		nil,
		NewMapping{
			393: 4775,
			404: 4776,
			477: 5279,
			573: 5279,
		},
	},
	{
		"minecraft:acacia_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6349,
			404: 6350,
			477: 6856,
			573: 6856,
		},
	},
	{
		"minecraft:jungle_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 5055,
			404: 5056,
			477: 5559,
			573: 5559,
			4:   2179,
		},
	},
	{
		"minecraft:jungle_leaves[distance=6,persistent=false]",
		nil,
		NewMapping{
			393: 197,
			404: 197,
			477: 197,
			573: 197,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 6827,
			393: 6320,
			404: 6321,
			477: 6827,
		},
	},
	{
		"minecraft:daylight_detector[inverted=true,power=9]",
		nil,
		NewMapping{
			573: 6167,
			4:   2857,
			393: 5660,
			404: 5661,
			477: 6167,
		},
	},
	{
		"minecraft:spruce_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   2149,
			393: 4925,
			404: 4926,
			477: 5429,
			573: 5429,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=north,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3668,
			404: 3669,
			477: 4172,
			573: 4172,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10844,
			573: 10844,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=west,in_wall=true,open=false,powered=true]",
		nil,
		NewMapping{
			477: 7964,
			573: 7964,
			393: 7439,
			404: 7440,
		},
	},
	{
		"minecraft:sign[rotation=8,waterlogged=false]",
		nil,
		NewMapping{
			4:   1016,
			393: 3092,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=west,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3826,
			404: 3827,
			477: 4330,
			573: 4330,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4464,
			404: 4465,
			477: 4968,
			573: 4968,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=11,south=up,west=up]",
		nil,
		NewMapping{
			404: 2572,
			477: 2875,
			573: 2875,
			393: 2571,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6590,
			404: 6591,
			477: 7097,
			573: 7097,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=5,south=up,west=up]",
		nil,
		NewMapping{
			393: 2085,
			404: 2086,
			477: 2389,
			573: 2389,
		},
	},
	{
		"minecraft:spruce_button[face=ceiling,facing=south,powered=true]",
		nil,
		NewMapping{
			404: 5346,
			477: 5852,
			573: 5852,
			393: 5345,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=24,powered=false]",
		nil,
		NewMapping{
			573: 947,
			477: 947,
		},
	},
	{
		"minecraft:pink_shulker_box[facing=south]",
		nil,
		NewMapping{
			393: 8255,
			404: 8256,
			477: 8780,
			573: 8780,
			4:   3603,
		},
	},
	{
		"minecraft:green_glazed_terracotta[facing=south]",
		nil,
		NewMapping{
			4:   3968,
			393: 8366,
			404: 8367,
			477: 8891,
			573: 8891,
		},
	},
	{
		"minecraft:jungle_leaves[distance=1,persistent=true]",
		nil,
		NewMapping{
			404: 186,
			477: 186,
			573: 186,
			4:   295,
			393: 186,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=north,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3595,
			404: 3596,
			477: 4099,
			573: 4099,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=21,powered=false]",
		nil,
		NewMapping{
			393: 541,
			404: 541,
			477: 541,
			573: 541,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=13,south=up,west=side]",
		nil,
		NewMapping{
			393: 1870,
			404: 1871,
			477: 2174,
			573: 2174,
		},
	},
	{
		"minecraft:oak_door[facing=north,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			404: 3114,
			477: 3577,
			573: 3577,
			393: 3113,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 6946,
			393: 6439,
			404: 6440,
			477: 6946,
		},
	},
	{
		"minecraft:sign[rotation=13,waterlogged=true]",
		nil,
		NewMapping{
			393: 3101,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9527,
			573: 9527,
		},
	},
	{
		"minecraft:lava[level=9]",
		nil,
		NewMapping{
			4:   169,
			393: 59,
			404: 59,
			477: 59,
			573: 59,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=south,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			4:   2677,
			393: 6521,
			404: 6522,
			477: 7028,
			573: 7028,
		},
	},
	{
		"minecraft:fire[age=9,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			4:   825,
			393: 1454,
			404: 1455,
			477: 1758,
			573: 1758,
		},
	},
	{
		"minecraft:moving_piston[facing=north,type=normal]",
		nil,
		NewMapping{
			4:   578,
			393: 1099,
			404: 1099,
			477: 1399,
			573: 1399,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=west,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6690,
			404: 6691,
			477: 7197,
			573: 7197,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=south,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9558,
			573: 9558,
		},
	},
	{
		"minecraft:granite_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9898,
			573: 9898,
		},
	},
	{
		"minecraft:fire[age=14,east=false,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			404: 1611,
			477: 1914,
			573: 1914,
			393: 1610,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=4,south=none,west=up]",
		nil,
		NewMapping{
			404: 2659,
			477: 2962,
			573: 2962,
			393: 2658,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=13,south=side,west=side]",
		nil,
		NewMapping{
			477: 3041,
			573: 3041,
			393: 2737,
			404: 2738,
		},
	},
	{
		"minecraft:acacia_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6335,
			404: 6336,
			477: 6842,
			573: 6842,
		},
	},
	{
		"minecraft:oak_wall_sign[facing=south,waterlogged=false]",
		nil,
		NewMapping{
			477: 3736,
			573: 3736,
			404: 3273,
		},
	},
	{
		"minecraft:acacia_door[facing=south,half=upper,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7888,
			404: 7889,
			477: 8413,
			573: 8413,
		},
	},
	{
		"minecraft:quartz_stairs[facing=east,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 5769,
			404: 5770,
			477: 6276,
			573: 6276,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=11,powered=true]",
		nil,
		NewMapping{
			393: 520,
			404: 520,
			477: 520,
			573: 520,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 10074,
			573: 10074,
		},
	},
	{
		"minecraft:pink_banner[rotation=9]",
		nil,
		NewMapping{
			393: 6959,
			404: 6960,
			477: 7466,
			573: 7466,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=south,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 6515,
			404: 6516,
			477: 7022,
			573: 7022,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10525,
			573: 10525,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=9,south=side,west=side]",
		nil,
		NewMapping{
			477: 3149,
			573: 3149,
			393: 2845,
			404: 2846,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=3,south=up,west=side]",
		nil,
		NewMapping{
			573: 2228,
			393: 1924,
			404: 1925,
			477: 2228,
		},
	},
	{
		"minecraft:oak_fence_gate[facing=east,in_wall=false,open=true,powered=true]",
		nil,
		NewMapping{
			404: 4329,
			477: 4832,
			573: 4832,
			4:   1727,
			393: 4328,
		},
	},
	{
		"minecraft:white_wall_banner[facing=west]",
		nil,
		NewMapping{
			573: 7619,
			4:   2836,
			393: 7112,
			404: 7113,
			477: 7619,
		},
	},
	{
		"minecraft:spruce_stairs[facing=south,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 4914,
			404: 4915,
			477: 5418,
			573: 5418,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=15,south=none,west=up]",
		nil,
		NewMapping{
			393: 2757,
			404: 2758,
			477: 3061,
			573: 3061,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=7,south=none,west=none]",
		nil,
		NewMapping{
			573: 2847,
			393: 2543,
			404: 2544,
			477: 2847,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10099,
			573: 10099,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10406,
			573: 10406,
		},
	},
	{
		"minecraft:acacia_door[facing=east,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			4:   3146,
			393: 7919,
			404: 7920,
			477: 8444,
			573: 8444,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=8,south=side,west=side]",
		nil,
		NewMapping{
			573: 2708,
			393: 2404,
			404: 2405,
			477: 2708,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6156,
			404: 6157,
			477: 6663,
			573: 6663,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=5,south=none,west=side]",
		nil,
		NewMapping{
			573: 2396,
			393: 2092,
			404: 2093,
			477: 2396,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=11,south=side,west=up]",
		nil,
		NewMapping{
			393: 3006,
			404: 3007,
			477: 3310,
			573: 3310,
		},
	},
	{
		"minecraft:nether_brick_fence[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 5025,
			573: 5025,
			393: 4521,
			404: 4522,
		},
	},
	{
		"minecraft:brick_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4344,
			404: 4345,
			477: 4848,
			573: 4848,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=east,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 4710,
			404: 4711,
			477: 5214,
			573: 5214,
		},
	},
	{
		"minecraft:spruce_leaves[distance=3,persistent=true]",
		nil,
		NewMapping{
			573: 162,
			393: 162,
			404: 162,
			477: 162,
		},
	},
	{
		"minecraft:red_banner[rotation=6]",
		nil,
		NewMapping{
			393: 7084,
			404: 7085,
			477: 7591,
			573: 7591,
		},
	},
	{
		"minecraft:lime_banner[rotation=15]",
		nil,
		NewMapping{
			404: 6950,
			477: 7456,
			573: 7456,
			393: 6949,
		},
	},
	{
		"minecraft:acacia_wall_sign[facing=west,waterlogged=true]",
		nil,
		NewMapping{
			477: 3761,
			573: 3761,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=east,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			573: 4273,
			393: 3769,
			404: 3770,
			477: 4273,
		},
	},
	{
		"minecraft:oak_sign[rotation=7,waterlogged=true]",
		nil,
		NewMapping{
			404: 3090,
			477: 3393,
			573: 3393,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 10781,
			477: 10781,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=7,south=none,west=side]",
		nil,
		NewMapping{
			573: 3278,
			393: 2974,
			404: 2975,
			477: 3278,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 6567,
			477: 7073,
			573: 7073,
			393: 6566,
		},
	},
	{
		"minecraft:fire[age=8,east=false,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			573: 1716,
			393: 1412,
			404: 1413,
			477: 1716,
		},
	},
	{
		"minecraft:light_blue_banner[rotation=11]",
		nil,
		NewMapping{
			477: 7420,
			573: 7420,
			393: 6913,
			404: 6914,
		},
	},
	{
		"minecraft:fire[age=7,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			404: 1378,
			477: 1681,
			573: 1681,
			393: 1377,
		},
	},
	{
		"minecraft:trapped_chest[facing=east,type=right,waterlogged=false]",
		nil,
		NewMapping{
			393: 5602,
			404: 5603,
			477: 6109,
			573: 6109,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5869,
			404: 5870,
			477: 6376,
			573: 6376,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6182,
			404: 6183,
			477: 6689,
			573: 6689,
		},
	},
	{
		"minecraft:chain_command_block[conditional=true,facing=east]",
		nil,
		NewMapping{
			4:   3389,
			393: 8177,
			404: 8178,
			477: 8702,
			573: 8702,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 7210,
			404: 7211,
			477: 7717,
			573: 7717,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=4,south=up,west=none]",
		nil,
		NewMapping{
			393: 2798,
			404: 2799,
			477: 3102,
			573: 3102,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=6,powered=true]",
		nil,
		NewMapping{
			393: 360,
			404: 360,
			477: 360,
			573: 360,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=true,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			404: 3991,
			477: 4494,
			573: 4494,
			393: 3990,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=false,north=false,powered=false,south=true,west=false]",
		nil,
		NewMapping{
			393: 4880,
			404: 4881,
			477: 5384,
			573: 5384,
		},
	},
	{
		"minecraft:fire[age=9,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1435,
			404: 1436,
			477: 1739,
			573: 1739,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9749,
			573: 9749,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10402,
			573: 10402,
		},
	},
	{
		"minecraft:white_banner[rotation=5]",
		nil,
		NewMapping{
			573: 7366,
			4:   2821,
			393: 6859,
			404: 6860,
			477: 7366,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=true,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 4086,
			404: 4087,
			477: 4590,
			573: 4590,
		},
	},
	{
		"minecraft:trapped_chest[facing=west,type=right,waterlogged=false]",
		nil,
		NewMapping{
			573: 6103,
			393: 5596,
			404: 5597,
			477: 6103,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10765,
			573: 10765,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10907,
			573: 10907,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=11,south=none,west=none]",
		nil,
		NewMapping{
			393: 2867,
			404: 2868,
			477: 3171,
			573: 3171,
		},
	},
	{
		"minecraft:blast_furnace[facing=east,lit=true]",
		nil,
		NewMapping{
			477: 11161,
			573: 11161,
		},
	},
	{
		"minecraft:spruce_button[face=ceiling,facing=east,powered=true]",
		nil,
		NewMapping{
			477: 5856,
			573: 5856,
			393: 5349,
			404: 5350,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 5701,
			393: 5197,
			404: 5198,
			477: 5701,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=south,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 6522,
			404: 6523,
			477: 7029,
			573: 7029,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=0,south=side,west=side]",
		nil,
		NewMapping{
			393: 2908,
			404: 2909,
			477: 3212,
			573: 3212,
		},
	},
	{
		"minecraft:fire[age=2,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			477: 1515,
			573: 1515,
			393: 1211,
			404: 1212,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=6,south=none,west=up]",
		nil,
		NewMapping{
			393: 2244,
			404: 2245,
			477: 2548,
			573: 2548,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=3,south=up,west=up]",
		nil,
		NewMapping{
			393: 2211,
			404: 2212,
			477: 2515,
			573: 2515,
		},
	},
	{
		"minecraft:yellow_banner[rotation=15]",
		nil,
		NewMapping{
			393: 6933,
			404: 6934,
			477: 7440,
			573: 7440,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10340,
			573: 10340,
		},
	},
	{
		"minecraft:detector_rail[powered=false,shape=north_south]",
		nil,
		NewMapping{
			404: 1022,
			477: 1322,
			573: 1322,
			4:   448,
			393: 1022,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4427,
			404: 4428,
			477: 4931,
			573: 4931,
		},
	},
	{
		"minecraft:dead_horn_coral_wall_fan[facing=west,waterlogged=true]",
		nil,
		NewMapping{
			393: 8500,
			404: 8516,
			477: 9060,
			573: 9060,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9577,
			573: 9577,
		},
	},
	{
		"minecraft:pink_bed[facing=east,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 858,
			404: 858,
			477: 1158,
			573: 1158,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=south,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 4548,
			404: 4549,
			477: 5052,
			573: 5052,
		},
	},
	{
		"minecraft:grindstone[face=wall,facing=west]",
		nil,
		NewMapping{
			477: 11171,
			573: 11171,
		},
	},
	{
		"minecraft:hopper[enabled=true,facing=south]",
		nil,
		NewMapping{
			404: 5688,
			477: 6194,
			573: 6194,
			4:   2467,
			393: 5687,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4431,
			404: 4432,
			477: 4935,
			573: 4935,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=22,powered=true]",
		nil,
		NewMapping{
			393: 492,
			404: 492,
			477: 492,
			573: 492,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=9,powered=false]",
		nil,
		NewMapping{
			393: 267,
			404: 267,
			477: 267,
			573: 267,
		},
	},
	{
		"minecraft:birch_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4969,
			404: 4970,
			477: 5473,
			573: 5473,
		},
	},
	{
		"minecraft:dead_bubble_coral_wall_fan[facing=north,waterlogged=false]",
		nil,
		NewMapping{
			393: 8481,
			404: 8497,
			477: 9041,
			573: 9041,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=false,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 4067,
			404: 4068,
			477: 4571,
			573: 4571,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=5,south=up,west=none]",
		nil,
		NewMapping{
			393: 2087,
			404: 2088,
			477: 2391,
			573: 2391,
		},
	},
	{
		"minecraft:acacia_button[face=wall,facing=north,powered=false]",
		nil,
		NewMapping{
			393: 5408,
			404: 5409,
			477: 5915,
			573: 5915,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10702,
			573: 10702,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9507,
			573: 9507,
		},
	},
	{
		"minecraft:andesite_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9951,
			573: 9951,
		},
	},
	{
		"minecraft:brick_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   1728,
			393: 4403,
			404: 4404,
			477: 4907,
			573: 4907,
		},
	},
	{
		"minecraft:fire[age=9,east=true,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1427,
			404: 1428,
			477: 1731,
			573: 1731,
		},
	},
	{
		"minecraft:oak_door[facing=west,half=upper,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			404: 3140,
			477: 3603,
			573: 3603,
			393: 3139,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=south,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			404: 3740,
			477: 4243,
			573: 4243,
			393: 3739,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=1,powered=false]",
		nil,
		NewMapping{
			477: 751,
			573: 751,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9382,
			573: 9382,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 7138,
			573: 7138,
			393: 6631,
			404: 6632,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6039,
			404: 6040,
			477: 6546,
			573: 6546,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6188,
			404: 6189,
			477: 6695,
			573: 6695,
		},
	},
	{
		"minecraft:green_banner[rotation=7]",
		nil,
		NewMapping{
			393: 7069,
			404: 7070,
			477: 7576,
			573: 7576,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=12,south=up,west=up]",
		nil,
		NewMapping{
			393: 3012,
			404: 3013,
			477: 3316,
			573: 3316,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=south,in_wall=true,open=false,powered=true]",
		nil,
		NewMapping{
			573: 7924,
			393: 7399,
			404: 7400,
			477: 7924,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10352,
			573: 10352,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=false,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10777,
			573: 10777,
		},
	},
	{
		"minecraft:acacia_door[facing=east,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			4:   3147,
			393: 7923,
			404: 7924,
			477: 8448,
			573: 8448,
		},
	},
	{
		"minecraft:birch_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			393: 7270,
			404: 7271,
			477: 7777,
			573: 7777,
			4:   2026,
		},
	},
	{
		"minecraft:chorus_plant[down=true,east=false,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			477: 8552,
			573: 8552,
			393: 8027,
			404: 8028,
		},
	},
	{
		"minecraft:acacia_fence[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 7642,
			404: 7643,
			477: 8167,
			573: 8167,
		},
	},
	{
		"minecraft:brain_coral",
		nil,
		NewMapping{
			393: 8460,
		},
	},
	{
		"minecraft:campfire[facing=south,lit=true,signal_fire=true,waterlogged=true]",
		nil,
		NewMapping{
			477: 11224,
			573: 11240,
		},
	},
	{
		"minecraft:acacia_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6411,
			404: 6412,
			477: 6918,
			573: 6918,
		},
	},
	{
		"minecraft:red_banner[rotation=14]",
		nil,
		NewMapping{
			477: 7599,
			573: 7599,
			393: 7092,
			404: 7093,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=15,south=side,west=up]",
		nil,
		NewMapping{
			404: 2323,
			477: 2626,
			573: 2626,
			393: 2322,
		},
	},
	{
		"minecraft:smooth_stone_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			404: 7294,
			477: 7806,
			573: 7806,
		},
	},
	{
		"minecraft:iron_door[facing=south,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			404: 3324,
			477: 3827,
			573: 3827,
			393: 3323,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=9,powered=false]",
		nil,
		NewMapping{
			393: 517,
			404: 517,
			477: 517,
			573: 517,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 11058,
			573: 11058,
		},
	},
	{
		"minecraft:damaged_anvil[facing=east]",
		nil,
		NewMapping{
			4:   2331,
			393: 5578,
			404: 5579,
			477: 6085,
			573: 6085,
		},
	},
	{
		"minecraft:gray_wall_banner[facing=east]",
		nil,
		NewMapping{
			393: 7141,
			404: 7142,
			477: 7648,
			573: 7648,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=true,north=false,powered=false,south=true,west=false]",
		nil,
		NewMapping{
			393: 4864,
			404: 4865,
			477: 5368,
			573: 5368,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=2,south=up,west=up]",
		nil,
		NewMapping{
			404: 2347,
			477: 2650,
			573: 2650,
			393: 2346,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=6,south=side,west=up]",
		nil,
		NewMapping{
			573: 3265,
			393: 2961,
			404: 2962,
			477: 3265,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10680,
			573: 10680,
		},
	},
	{
		"minecraft:bell[attachment=floor,facing=south,powered=false]",
		nil,
		NewMapping{
			573: 11201,
		},
	},
	{
		"minecraft:iron_door[facing=north,half=upper,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			393: 3310,
			404: 3311,
			477: 3814,
			573: 3814,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9584,
			573: 9584,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=south,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			573: 4313,
			393: 3809,
			404: 3810,
			477: 4313,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10443,
			573: 10443,
		},
	},
	{
		"minecraft:magenta_bed[facing=south,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 786,
			404: 786,
			477: 1086,
			573: 1086,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 3237,
			404: 3238,
			477: 3701,
			573: 3701,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=east,half=top,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3711,
			404: 3712,
			477: 4215,
			573: 4215,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=13,south=side,west=side]",
		nil,
		NewMapping{
			393: 2305,
			404: 2306,
			477: 2609,
			573: 2609,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 6564,
			393: 6057,
			404: 6058,
			477: 6564,
		},
	},
	{
		"minecraft:sandstone",
		nil,
		NewMapping{
			573: 245,
			4:   384,
			393: 245,
			404: 245,
			477: 245,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 4127,
			404: 4128,
			477: 4631,
			573: 4631,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9541,
			573: 9541,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=16,powered=false]",
		nil,
		NewMapping{
			477: 981,
			573: 981,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10786,
			573: 10786,
		},
	},
	{
		"minecraft:dispenser[facing=east,triggered=true]",
		nil,
		NewMapping{
			4:   381,
			393: 235,
			404: 235,
			477: 235,
			573: 235,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=west,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			573: 4271,
			393: 3767,
			404: 3768,
			477: 4271,
		},
	},
	{
		"minecraft:jungle_stairs[facing=east,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 5104,
			404: 5105,
			477: 5608,
			573: 5608,
		},
	},
	{
		"minecraft:spruce_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4939,
			404: 4940,
			477: 5443,
			573: 5443,
		},
	},
	{
		"minecraft:light_gray_banner[rotation=14]",
		nil,
		NewMapping{
			573: 7503,
			393: 6996,
			404: 6997,
			477: 7503,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=4,powered=false]",
		nil,
		NewMapping{
			393: 407,
			404: 407,
			477: 407,
			573: 407,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			404: 6119,
			477: 6625,
			573: 6625,
			393: 6118,
		},
	},
	{
		"minecraft:turtle_egg[eggs=2,hatch=0]",
		nil,
		NewMapping{
			393: 8440,
			404: 8441,
			477: 8965,
			573: 8965,
		},
	},
	{
		"minecraft:granite_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			573: 10305,
			477: 10305,
		},
	},
	{
		"minecraft:birch_sign[rotation=7,waterlogged=true]",
		nil,
		NewMapping{
			477: 3457,
			573: 3457,
		},
	},
	{
		"minecraft:fire[age=1,east=false,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1187,
			404: 1188,
			477: 1491,
			573: 1491,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=11,powered=false]",
		nil,
		NewMapping{
			477: 771,
			573: 771,
		},
	},
	{
		"minecraft:orange_terracotta",
		nil,
		NewMapping{
			4:   2545,
			393: 5805,
			404: 5806,
			477: 6312,
			573: 6312,
		},
	},
	{
		"minecraft:acacia_fence[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 7624,
			404: 7625,
			477: 8149,
			573: 8149,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6264,
			404: 6265,
			477: 6771,
			573: 6771,
		},
	},
	{
		"minecraft:oak_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 1688,
			404: 1689,
			477: 1992,
			573: 1992,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=true,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			404: 4084,
			477: 4587,
			573: 4587,
			393: 4083,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9759,
			573: 9759,
		},
	},
	{
		"minecraft:repeating_command_block[conditional=true,facing=west]",
		nil,
		NewMapping{
			4:   3372,
			393: 8167,
			404: 8168,
			477: 8692,
			573: 8692,
		},
	},
	{
		"minecraft:brown_shulker_box[facing=south]",
		nil,
		NewMapping{
			477: 8816,
			573: 8816,
			4:   3699,
			393: 8291,
			404: 8292,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=north,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3597,
			404: 3598,
			477: 4101,
			573: 4101,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=false,north=false,powered=false,south=true,west=false]",
		nil,
		NewMapping{
			477: 5288,
			573: 5288,
			393: 4784,
			404: 4785,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=11,powered=true]",
		nil,
		NewMapping{
			393: 620,
			404: 620,
			477: 620,
			573: 620,
		},
	},
	{
		"minecraft:spruce_leaves[distance=4,persistent=false]",
		nil,
		NewMapping{
			393: 165,
			404: 165,
			477: 165,
			573: 165,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6289,
			404: 6290,
			477: 6796,
			573: 6796,
		},
	},
	{
		"minecraft:horn_coral[waterlogged=true]",
		nil,
		NewMapping{
			404: 8478,
			477: 9002,
			573: 9002,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10436,
			573: 10436,
		},
	},
	{
		"minecraft:sign[rotation=5,waterlogged=false]",
		nil,
		NewMapping{
			4:   1013,
			393: 3086,
		},
	},
	{
		"minecraft:player_wall_head[facing=north]",
		nil,
		NewMapping{
			393: 5507,
			404: 5508,
			477: 6030,
			573: 6030,
		},
	},
	{
		"minecraft:oak_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 1959,
			573: 1959,
			393: 1655,
			404: 1656,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=south,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3943,
			404: 3944,
			477: 4447,
			573: 4447,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=11,powered=false]",
		nil,
		NewMapping{
			393: 721,
			404: 721,
			477: 721,
			573: 721,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6654,
			404: 6655,
			477: 7161,
			573: 7161,
		},
	},
	{
		"minecraft:fire[age=0,east=true,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1149,
			404: 1150,
			477: 1453,
			573: 1453,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 6520,
			393: 6013,
			404: 6014,
			477: 6520,
		},
	},
	{
		"minecraft:andesite_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9984,
			573: 9984,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9297,
			573: 9297,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=20,powered=false]",
		nil,
		NewMapping{
			404: 539,
			477: 539,
			573: 539,
			393: 539,
		},
	},
	{
		"minecraft:bell[attachment=floor,facing=south]",
		nil,
		NewMapping{
			477: 11199,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9307,
			573: 9307,
		},
	},
	{
		"minecraft:dead_tube_coral_wall_fan[facing=south,waterlogged=true]",
		nil,
		NewMapping{
			573: 9026,
			393: 8466,
			404: 8482,
			477: 9026,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=13,south=side,west=none]",
		nil,
		NewMapping{
			393: 2450,
			404: 2451,
			477: 2754,
			573: 2754,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=false,north=true,powered=true,south=true,west=false]",
		nil,
		NewMapping{
			393: 4868,
			404: 4869,
			477: 5372,
			573: 5372,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 7706,
			393: 7199,
			404: 7200,
			477: 7706,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=5,south=side,west=none]",
		nil,
		NewMapping{
			393: 1946,
			404: 1947,
			477: 2250,
			573: 2250,
		},
	},
	{
		"minecraft:black_concrete_powder",
		nil,
		NewMapping{
			4:   4047,
			393: 8408,
			404: 8409,
			477: 8933,
			573: 8933,
		},
	},
	{
		"minecraft:birch_button[face=ceiling,facing=east,powered=false]",
		nil,
		NewMapping{
			393: 5374,
			404: 5375,
			477: 5881,
			573: 5881,
		},
	},
	{
		"minecraft:cut_red_sandstone_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			573: 7866,
			477: 7866,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 4015,
			404: 4016,
			477: 4519,
			573: 4519,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 5691,
			573: 5691,
			393: 5187,
			404: 5188,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=13,powered=false]",
		nil,
		NewMapping{
			573: 575,
			393: 575,
			404: 575,
			477: 575,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9768,
			573: 9768,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6184,
			404: 6185,
			477: 6691,
			573: 6691,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=15,south=side,west=none]",
		nil,
		NewMapping{
			393: 2612,
			404: 2613,
			477: 2916,
			573: 2916,
		},
	},
	{
		"minecraft:fire[age=9,east=true,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1425,
			404: 1426,
			477: 1729,
			573: 1729,
		},
	},
	{
		"minecraft:spruce_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4921,
			404: 4922,
			477: 5425,
			573: 5425,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 3204,
			404: 3205,
			477: 3668,
			573: 3668,
		},
	},
	{
		"minecraft:red_wall_banner[facing=south]",
		nil,
		NewMapping{
			393: 7167,
			404: 7168,
			477: 7674,
			573: 7674,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=21,powered=false]",
		nil,
		NewMapping{
			573: 791,
			477: 791,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10534,
			573: 10534,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=15,powered=true]",
		nil,
		NewMapping{
			477: 328,
			573: 328,
			393: 328,
			404: 328,
		},
	},
	{
		"minecraft:birch_door[facing=south,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7771,
			404: 7772,
			477: 8296,
			573: 8296,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=13,south=side,west=none]",
		nil,
		NewMapping{
			393: 2738,
			404: 2739,
			477: 3042,
			573: 3042,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=true,north=false,powered=true,south=false,west=true]",
		nil,
		NewMapping{
			393: 4861,
			404: 4862,
			477: 5365,
			573: 5365,
		},
	},
	{
		"minecraft:dark_oak_door[facing=south,half=upper,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7956,
			404: 7957,
			477: 8481,
			573: 8481,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=12,south=none,west=up]",
		nil,
		NewMapping{
			393: 2154,
			404: 2155,
			477: 2458,
			573: 2458,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5899,
			404: 5900,
			477: 6406,
			573: 6406,
		},
	},
	{
		"minecraft:quartz_pillar[axis=z]",
		nil,
		NewMapping{
			573: 6206,
			4:   2484,
			393: 5699,
			404: 5700,
			477: 6206,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   2628,
			393: 6473,
			404: 6474,
			477: 6980,
			573: 6980,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=11,south=none,west=side]",
		nil,
		NewMapping{
			393: 2002,
			404: 2003,
			477: 2306,
			573: 2306,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=3,south=side,west=side]",
		nil,
		NewMapping{
			393: 1927,
			404: 1928,
			477: 2231,
			573: 2231,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=23,powered=true]",
		nil,
		NewMapping{
			393: 694,
			404: 694,
			477: 694,
			573: 694,
		},
	},
	{
		"minecraft:attached_pumpkin_stem[facing=south]",
		nil,
		NewMapping{
			393: 4245,
			404: 4246,
			477: 4749,
			573: 4749,
		},
	},
	{
		"minecraft:purpur_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 8649,
			393: 8124,
			404: 8125,
			477: 8649,
		},
	},
	{
		"minecraft:purpur_stairs[facing=east,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 8139,
			404: 8140,
			477: 8664,
			573: 8664,
		},
	},
	{
		"minecraft:brown_bed[facing=north,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 941,
			404: 941,
			477: 1241,
			573: 1241,
		},
	},
	{
		"minecraft:diorite_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 11089,
			573: 11089,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=1,south=side,west=up]",
		nil,
		NewMapping{
			393: 2052,
			404: 2053,
			477: 2356,
			573: 2356,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=20,powered=true]",
		nil,
		NewMapping{
			477: 788,
			573: 788,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 10662,
			477: 10662,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=1,powered=false]",
		nil,
		NewMapping{
			477: 801,
			573: 801,
		},
	},
	{
		"minecraft:light_blue_stained_glass",
		nil,
		NewMapping{
			393: 3580,
			404: 3581,
			477: 4084,
			573: 4084,
			4:   1523,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=2,south=side,west=side]",
		nil,
		NewMapping{
			393: 2206,
			404: 2207,
			477: 2510,
			573: 2510,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=true,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 4126,
			404: 4127,
			477: 4630,
			573: 4630,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=8,powered=true]",
		nil,
		NewMapping{
			573: 614,
			393: 614,
			404: 614,
			477: 614,
		},
	},
	{
		"minecraft:fire[age=2,east=true,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			477: 1516,
			573: 1516,
			393: 1212,
			404: 1213,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			573: 9813,
			477: 9813,
		},
	},
	{
		"minecraft:andesite_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9978,
			573: 9978,
		},
	},
	{
		"minecraft:blue_shulker_box[facing=east]",
		nil,
		NewMapping{
			404: 8285,
			477: 8809,
			573: 8809,
			4:   3685,
			393: 8284,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=west,in_wall=false,open=true,powered=false]",
		nil,
		NewMapping{
			4:   2949,
			393: 7410,
			404: 7411,
			477: 7935,
			573: 7935,
		},
	},
	{
		"minecraft:spruce_fence[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 7533,
			477: 8057,
			573: 8057,
			393: 7532,
		},
	},
	{
		"minecraft:brown_bed[facing=west,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 950,
			404: 950,
			477: 1250,
			573: 1250,
		},
	},
	{
		"minecraft:oak_door[facing=west,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			477: 3613,
			573: 3613,
			393: 3149,
			404: 3150,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 3208,
			404: 3209,
			477: 3672,
			573: 3672,
		},
	},
	{
		"minecraft:lime_wall_banner[facing=west]",
		nil,
		NewMapping{
			477: 7639,
			573: 7639,
			393: 7132,
			404: 7133,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=9,powered=true]",
		nil,
		NewMapping{
			477: 366,
			573: 366,
			393: 366,
			404: 366,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=13,south=none,west=none]",
		nil,
		NewMapping{
			477: 3333,
			573: 3333,
			4:   893,
			393: 3029,
			404: 3030,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=west,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3635,
			404: 3636,
			477: 4139,
			573: 4139,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=east,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3975,
			404: 3976,
			477: 4479,
			573: 4479,
		},
	},
	{
		"minecraft:gray_wall_banner[facing=south]",
		nil,
		NewMapping{
			393: 7139,
			404: 7140,
			477: 7646,
			573: 7646,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 6708,
			393: 6201,
			404: 6202,
			477: 6708,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10546,
			573: 10546,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9381,
			573: 9381,
		},
	},
	{
		"minecraft:spruce_door[facing=north,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			4:   3095,
			393: 7690,
			404: 7691,
			477: 8215,
			573: 8215,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=12,south=side,west=side]",
		nil,
		NewMapping{
			573: 2456,
			393: 2152,
			404: 2153,
			477: 2456,
		},
	},
	{
		"minecraft:fire[age=0,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			404: 1154,
			477: 1457,
			573: 1457,
			393: 1153,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 5066,
			573: 5066,
			393: 4562,
			404: 4563,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=south,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 7205,
			404: 7206,
			477: 7712,
			573: 7712,
		},
	},
	{
		"minecraft:fire[age=10,east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1483,
			404: 1484,
			477: 1787,
			573: 1787,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=true,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5861,
			404: 5862,
			477: 6368,
			573: 6368,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10375,
			573: 10375,
		},
	},
	{
		"minecraft:bee_nest[facing=east,honey_level=5]",
		nil,
		NewMapping{
			573: 11310,
		},
	},
	{
		"minecraft:quartz_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 5718,
			404: 5719,
			477: 6225,
			573: 6225,
		},
	},
	{
		"minecraft:pink_banner[rotation=3]",
		nil,
		NewMapping{
			393: 6953,
			404: 6954,
			477: 7460,
			573: 7460,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5255,
			404: 5256,
			477: 5759,
			573: 5759,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=12,south=up,west=up]",
		nil,
		NewMapping{
			393: 2868,
			404: 2869,
			477: 3172,
			573: 3172,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9463,
			573: 9463,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10391,
			573: 10391,
		},
	},
	{
		"minecraft:blue_orchid",
		nil,
		NewMapping{
			4:   609,
			393: 1113,
			404: 1113,
			477: 1413,
			573: 1413,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=1,south=side,west=none]",
		nil,
		NewMapping{
			393: 2486,
			404: 2487,
			477: 2790,
			573: 2790,
		},
	},
	{
		"minecraft:jungle_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 5089,
			404: 5090,
			477: 5593,
			573: 5593,
		},
	},
	{
		"minecraft:stone_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			393: 7295,
			477: 7802,
			573: 7802,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=north,half=top,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3916,
			404: 3917,
			477: 4420,
			573: 4420,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10959,
			573: 10959,
		},
	},
	{
		"minecraft:jungle_sapling[stage=0]",
		nil,
		NewMapping{
			573: 27,
			4:   99,
			393: 27,
			404: 27,
			477: 27,
		},
	},
	{
		"minecraft:repeater[delay=3,facing=east,locked=false,powered=false]",
		nil,
		NewMapping{
			573: 4064,
			4:   1499,
			393: 3560,
			404: 3561,
			477: 4064,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=7,powered=false]",
		nil,
		NewMapping{
			404: 363,
			477: 363,
			573: 363,
			393: 363,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=1,powered=false]",
		nil,
		NewMapping{
			477: 851,
			573: 851,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=6,powered=true]",
		nil,
		NewMapping{
			477: 960,
			573: 960,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6219,
			404: 6220,
			477: 6726,
			573: 6726,
		},
	},
	{
		"minecraft:fire[age=14,east=true,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1583,
			404: 1584,
			477: 1887,
			573: 1887,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10766,
			573: 10766,
		},
	},
	{
		"minecraft:bamboo[age=0,leaves=small,stage=0]",
		nil,
		NewMapping{
			477: 9118,
			573: 9118,
		},
	},
	{
		"minecraft:activator_rail[powered=false,shape=north_south]",
		nil,
		NewMapping{
			4:   2512,
			393: 5786,
			404: 5787,
			477: 6293,
			573: 6293,
		},
	},
	{
		"minecraft:turtle_egg[eggs=3,hatch=1]",
		nil,
		NewMapping{
			393: 8444,
			404: 8445,
			477: 8969,
			573: 8969,
		},
	},
	{
		"minecraft:oak_sign[rotation=4,waterlogged=true]",
		nil,
		NewMapping{
			477: 3387,
			573: 3387,
			404: 3084,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=false,north=false,powered=true,south=false,west=true]",
		nil,
		NewMapping{
			573: 5381,
			393: 4877,
			404: 4878,
			477: 5381,
		},
	},
	{
		"minecraft:gray_banner[rotation=0]",
		nil,
		NewMapping{
			393: 6966,
			404: 6967,
			477: 7473,
			573: 7473,
		},
	},
	{
		"minecraft:acacia_button[face=ceiling,facing=north,powered=false]",
		nil,
		NewMapping{
			477: 5923,
			573: 5923,
			393: 5416,
			404: 5417,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10769,
			573: 10769,
		},
	},
	{
		"minecraft:carrots[age=2]",
		nil,
		NewMapping{
			404: 5290,
			477: 5796,
			573: 5796,
			4:   2258,
			393: 5289,
		},
	},
	{
		"minecraft:dark_oak_fence_gate[facing=east,in_wall=true,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7511,
			404: 7512,
			477: 8036,
			573: 8036,
		},
	},
	{
		"minecraft:purple_bed[facing=south,occupied=false,part=head]",
		nil,
		NewMapping{
			404: 914,
			477: 1214,
			573: 1214,
			393: 914,
		},
	},
	{
		"minecraft:jungle_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			404: 5115,
			477: 5618,
			573: 5618,
			393: 5114,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=0,south=none,west=side]",
		nil,
		NewMapping{
			573: 2351,
			393: 2047,
			404: 2048,
			477: 2351,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9311,
			573: 9311,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10249,
			573: 10249,
		},
	},
	{
		"minecraft:red_glazed_terracotta[facing=north]",
		nil,
		NewMapping{
			4:   3986,
			393: 8369,
			404: 8370,
			477: 8894,
			573: 8894,
		},
	},
	{
		"minecraft:purpur_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 8095,
			404: 8096,
			477: 8620,
			573: 8620,
		},
	},
	{
		"minecraft:brick_stairs[facing=east,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 4898,
			573: 4898,
			393: 4394,
			404: 4395,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 7232,
			393: 6725,
			404: 6726,
			477: 7232,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9312,
			573: 9312,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=10,south=none,west=none]",
		nil,
		NewMapping{
			393: 2426,
			404: 2427,
			477: 2730,
			573: 2730,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=west,half=top,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3824,
			404: 3825,
			477: 4328,
			573: 4328,
		},
	},
	{
		"minecraft:oak_sign[rotation=8,waterlogged=false]",
		nil,
		NewMapping{
			573: 3396,
			404: 3093,
			477: 3396,
		},
	},
	{
		"minecraft:jungle_door[facing=east,half=upper,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			477: 8378,
			573: 8378,
			393: 7853,
			404: 7854,
		},
	},
	{
		"minecraft:dark_oak_door[facing=south,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			404: 7954,
			477: 8478,
			573: 8478,
			393: 7953,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 7201,
			477: 7707,
			573: 7707,
			393: 7200,
		},
	},
	{
		"minecraft:nether_brick_fence[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 5016,
			573: 5016,
			393: 4512,
			404: 4513,
		},
	},
	{
		"minecraft:dark_oak_leaves[distance=7,persistent=true]",
		nil,
		NewMapping{
			393: 226,
			404: 226,
			477: 226,
			573: 226,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10963,
			573: 10963,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=17,powered=false]",
		nil,
		NewMapping{
			477: 833,
			573: 833,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 7757,
			393: 7250,
			404: 7251,
			477: 7757,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=5,south=none,west=side]",
		nil,
		NewMapping{
			393: 1948,
			404: 1949,
			477: 2252,
			573: 2252,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=18,powered=false]",
		nil,
		NewMapping{
			393: 535,
			404: 535,
			477: 535,
			573: 535,
		},
	},
	{
		"minecraft:purpur_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 8081,
			404: 8082,
			477: 8606,
			573: 8606,
		},
	},
	{
		"minecraft:diorite_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10216,
			573: 10216,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10929,
			573: 10929,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9380,
			573: 9380,
		},
	},
	{
		"minecraft:birch_door[facing=east,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			4:   3115,
			393: 7795,
			404: 7796,
			477: 8320,
			573: 8320,
		},
	},
	{
		"minecraft:fire[age=15,east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			477: 1947,
			573: 1947,
			393: 1643,
			404: 1644,
		},
	},
	{
		"minecraft:oak_door[facing=west,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			573: 3615,
			393: 3151,
			404: 3152,
			477: 3615,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=north,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3859,
			404: 3860,
			477: 4363,
			573: 4363,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6593,
			404: 6594,
			477: 7100,
			573: 7100,
		},
	},
	{
		"minecraft:granite_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9877,
			573: 9877,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=true,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 4030,
			404: 4031,
			477: 4534,
			573: 4534,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=4,south=none,west=side]",
		nil,
		NewMapping{
			393: 2371,
			404: 2372,
			477: 2675,
			573: 2675,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 6476,
			393: 5969,
			404: 5970,
			477: 6476,
		},
	},
	{
		"minecraft:oak_door[facing=south,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 3133,
			404: 3134,
			477: 3597,
			573: 3597,
		},
	},
	{
		"minecraft:fire[age=1,east=false,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			573: 1494,
			393: 1190,
			404: 1191,
			477: 1494,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=10,south=side,west=up]",
		nil,
		NewMapping{
			404: 2278,
			477: 2581,
			573: 2581,
			393: 2277,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 10543,
			477: 10543,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9254,
			573: 9254,
		},
	},
	{
		"minecraft:jungle_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   2178,
			393: 5075,
			404: 5076,
			477: 5579,
			573: 5579,
		},
	},
	{
		"minecraft:frosted_ice[age=0]",
		nil,
		NewMapping{
			4:   3392,
			393: 8188,
			404: 8189,
			477: 8713,
			573: 8713,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 7724,
			573: 7724,
			393: 7217,
			404: 7218,
		},
	},
	{
		"minecraft:dark_oak_fence[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 7665,
			404: 7666,
			477: 8190,
			573: 8190,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=west,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3896,
			404: 3897,
			477: 4400,
			573: 4400,
		},
	},
	{
		"minecraft:granite_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9888,
			573: 9888,
		},
	},
	{
		"minecraft:stone_button[face=wall,facing=south,powered=true]",
		nil,
		NewMapping{
			4:   1243,
			393: 3401,
			404: 3402,
			477: 3905,
			573: 3905,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4476,
			404: 4477,
			477: 4980,
			573: 4980,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5205,
			404: 5206,
			477: 5709,
			573: 5709,
		},
	},
	{
		"minecraft:cave_air",
		nil,
		NewMapping{
			393: 8575,
			404: 8592,
			477: 9130,
			573: 9130,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=7,south=side,west=up]",
		nil,
		NewMapping{
			393: 2970,
			404: 2971,
			477: 3274,
			573: 3274,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9827,
			573: 9827,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=12,south=up,west=none]",
		nil,
		NewMapping{
			477: 2166,
			573: 2166,
			393: 1862,
			404: 1863,
		},
	},
	{
		"minecraft:stone_brick_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			393: 7323,
			404: 7324,
			477: 7842,
			573: 7842,
		},
	},
	{
		"minecraft:prismarine_brick_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			573: 7314,
			393: 6807,
			404: 6808,
			477: 7314,
		},
	},
	{
		"minecraft:spruce_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4929,
			404: 4930,
			477: 5433,
			573: 5433,
		},
	},
	{
		"minecraft:andesite_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 9945,
			477: 9945,
		},
	},
	{
		"minecraft:fire[age=3,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1249,
			404: 1250,
			477: 1553,
			573: 1553,
		},
	},
	{
		"minecraft:acacia_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6396,
			404: 6397,
			477: 6903,
			573: 6903,
		},
	},
	{
		"minecraft:yellow_bed[facing=west,occupied=true,part=foot]",
		nil,
		NewMapping{
			404: 821,
			477: 1121,
			573: 1121,
			393: 821,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 6727,
			393: 6220,
			404: 6221,
			477: 6727,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=2,south=side,west=none]",
		nil,
		NewMapping{
			393: 2639,
			404: 2640,
			477: 2943,
			573: 2943,
		},
	},
	{
		"minecraft:granite_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9880,
			573: 9880,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9737,
			573: 9737,
		},
	},
	{
		"minecraft:jungle_door[facing=south,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7835,
			404: 7836,
			477: 8360,
			573: 8360,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6700,
			404: 6701,
			477: 7207,
			573: 7207,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=2,south=none,west=side]",
		nil,
		NewMapping{
			573: 2945,
			393: 2641,
			404: 2642,
			477: 2945,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=15,south=up,west=up]",
		nil,
		NewMapping{
			573: 3343,
			393: 3039,
			404: 3040,
			477: 3343,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 10715,
			477: 10715,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9419,
			573: 9419,
		},
	},
	{
		"minecraft:oak_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 1707,
			404: 1708,
			477: 2011,
			573: 2011,
		},
	},
	{
		"minecraft:light_blue_bed[facing=south,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 801,
			404: 801,
			477: 1101,
			573: 1101,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=5,south=up,west=side]",
		nil,
		NewMapping{
			477: 2102,
			573: 2102,
			393: 1798,
			404: 1799,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10785,
			573: 10785,
		},
	},
	{
		"minecraft:fire[age=0,east=true,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1142,
			404: 1143,
			477: 1446,
			573: 1446,
		},
	},
	{
		"minecraft:chorus_plant[down=true,east=true,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			573: 8536,
			393: 8011,
			404: 8012,
			477: 8536,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=2,powered=true]",
		nil,
		NewMapping{
			404: 702,
			477: 702,
			573: 702,
			393: 702,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=13,powered=true]",
		nil,
		NewMapping{
			477: 824,
			573: 824,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4419,
			404: 4420,
			477: 4923,
			573: 4923,
		},
	},
	{
		"minecraft:diorite_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 11091,
			573: 11091,
		},
	},
	{
		"minecraft:scaffolding[bottom=false,distance=2,waterlogged=true]",
		nil,
		NewMapping{
			477: 11119,
			573: 11119,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10772,
			477: 10772,
		},
	},
	{
		"minecraft:command_block[conditional=true,facing=west]",
		nil,
		NewMapping{
			404: 5128,
			477: 5631,
			573: 5631,
			4:   2204,
			393: 5127,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=true,north=true,powered=true,south=true,west=true]",
		nil,
		NewMapping{
			393: 4819,
			404: 4820,
			477: 5323,
			573: 5323,
		},
	},
	{
		"minecraft:dark_oak_button[face=wall,facing=south,powered=false]",
		nil,
		NewMapping{
			393: 5434,
			404: 5435,
			477: 5941,
			573: 5941,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=false,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 4014,
			404: 4015,
			477: 4518,
			573: 4518,
		},
	},
	{
		"minecraft:oak_sign[rotation=7,waterlogged=false]",
		nil,
		NewMapping{
			404: 3091,
			477: 3394,
			573: 3394,
		},
	},
	{
		"minecraft:gray_banner[rotation=4]",
		nil,
		NewMapping{
			573: 7477,
			393: 6970,
			404: 6971,
			477: 7477,
		},
	},
	{
		"minecraft:sign[rotation=5,waterlogged=true]",
		nil,
		NewMapping{
			393: 3085,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=15,south=none,west=up]",
		nil,
		NewMapping{
			393: 2901,
			404: 2902,
			477: 3205,
			573: 3205,
		},
	},
	{
		"minecraft:quartz_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			404: 5732,
			477: 6238,
			573: 6238,
			4:   2498,
			393: 5731,
		},
	},
	{
		"minecraft:flower_pot",
		nil,
		NewMapping{
			4:   2248,
			393: 5265,
			404: 5266,
			477: 5769,
			573: 5769,
		},
	},
	{
		"minecraft:brewing_stand[has_bottle_0=false,has_bottle_1=false,has_bottle_2=false]",
		nil,
		NewMapping{
			477: 5124,
			573: 5124,
			4:   1872,
			393: 4620,
			404: 4621,
		},
	},
	{
		"minecraft:light_weighted_pressure_plate[power=14]",
		nil,
		NewMapping{
			4:   2366,
			393: 5617,
			404: 5618,
			477: 6124,
			573: 6124,
		},
	},
	{
		"minecraft:end_rod[facing=east]",
		nil,
		NewMapping{
			404: 7999,
			477: 8523,
			573: 8523,
			4:   3173,
			393: 7998,
		},
	},
	{
		"minecraft:acacia_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6348,
			404: 6349,
			477: 6855,
			573: 6855,
		},
	},
	{
		"minecraft:pink_bed[facing=west,occupied=true,part=foot]",
		nil,
		NewMapping{
			573: 1153,
			393: 853,
			404: 853,
			477: 1153,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10971,
			573: 10971,
		},
	},
	{
		"minecraft:hopper[enabled=false,facing=north]",
		nil,
		NewMapping{
			477: 6198,
			573: 6198,
			4:   2474,
			393: 5691,
			404: 5692,
		},
	},
	{
		"minecraft:jungle_button[face=ceiling,facing=west,powered=false]",
		nil,
		NewMapping{
			393: 5396,
			404: 5397,
			477: 5903,
			573: 5903,
		},
	},
	{
		"minecraft:dark_oak_door[facing=south,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			573: 8480,
			393: 7955,
			404: 7956,
			477: 8480,
		},
	},
	{
		"minecraft:birch_door[facing=east,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7803,
			404: 7804,
			477: 8328,
			573: 8328,
		},
	},
	{
		"minecraft:spruce_fence[east=true,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 8043,
			573: 8043,
			393: 7518,
			404: 7519,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9299,
			573: 9299,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=21,powered=false]",
		nil,
		NewMapping{
			393: 591,
			404: 591,
			477: 591,
			573: 591,
		},
	},
	{
		"minecraft:fire[age=6,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			404: 1346,
			477: 1649,
			573: 1649,
			393: 1345,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6038,
			404: 6039,
			477: 6545,
			573: 6545,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=10,south=none,west=up]",
		nil,
		NewMapping{
			404: 2857,
			477: 3160,
			573: 3160,
			393: 2856,
		},
	},
	{
		"minecraft:orange_wool",
		nil,
		NewMapping{
			4:   561,
			393: 1084,
			404: 1084,
			477: 1384,
			573: 1384,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=4,south=side,west=up]",
		nil,
		NewMapping{
			404: 2656,
			477: 2959,
			573: 2959,
			393: 2655,
		},
	},
	{
		"minecraft:dragon_head[rotation=7]",
		nil,
		NewMapping{
			393: 5558,
			404: 5559,
			477: 6061,
			573: 6061,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6687,
			404: 6688,
			477: 7194,
			573: 7194,
		},
	},
	{
		"minecraft:jungle_sign[rotation=6,waterlogged=true]",
		nil,
		NewMapping{
			573: 3519,
			477: 3519,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   1746,
			393: 4443,
			404: 4444,
			477: 4947,
			573: 4947,
		},
	},
	{
		"minecraft:potatoes[age=2]",
		nil,
		NewMapping{
			4:   2274,
			393: 5297,
			404: 5298,
			477: 5804,
			573: 5804,
		},
	},
	{
		"minecraft:cactus[age=2]",
		nil,
		NewMapping{
			4:   1298,
			393: 3427,
			404: 3428,
			477: 3931,
			573: 3931,
		},
	},
	{
		"minecraft:fire[age=1,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1196,
			404: 1197,
			477: 1500,
			573: 1500,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=6,south=side,west=side]",
		nil,
		NewMapping{
			573: 2402,
			393: 2098,
			404: 2099,
			477: 2402,
		},
	},
	{
		"minecraft:furnace[facing=south,lit=true]",
		nil,
		NewMapping{
			477: 3373,
			573: 3373,
			4:   995,
			393: 3069,
			404: 3070,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=9,south=none,west=none]",
		nil,
		NewMapping{
			393: 2705,
			404: 2706,
			477: 3009,
			573: 3009,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=14,powered=false]",
		nil,
		NewMapping{
			393: 427,
			404: 427,
			477: 427,
			573: 427,
		},
	},
	{
		"minecraft:brick_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4374,
			404: 4375,
			477: 4878,
			573: 4878,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=east,in_wall=false,open=true,powered=false]",
		nil,
		NewMapping{
			4:   2999,
			393: 7482,
			404: 7483,
			477: 8007,
			573: 8007,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=north,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			573: 4233,
			393: 3729,
			404: 3730,
			477: 4233,
		},
	},
	{
		"minecraft:fire[age=10,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			573: 1760,
			393: 1456,
			404: 1457,
			477: 1760,
		},
	},
	{
		"minecraft:birch_leaves[distance=3,persistent=false]",
		nil,
		NewMapping{
			393: 177,
			404: 177,
			477: 177,
			573: 177,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=7,powered=true]",
		nil,
		NewMapping{
			393: 412,
			404: 412,
			477: 412,
			573: 412,
		},
	},
	{
		"minecraft:fire[age=8,east=false,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1411,
			404: 1412,
			477: 1715,
			573: 1715,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=13,south=side,west=up]",
		nil,
		NewMapping{
			404: 2737,
			477: 3040,
			573: 3040,
			393: 2736,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 6572,
			573: 6572,
			393: 6065,
			404: 6066,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=west,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 6540,
			404: 6541,
			477: 7047,
			573: 7047,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=west,half=top,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3820,
			404: 3821,
			477: 4324,
			573: 4324,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=13,powered=false]",
		nil,
		NewMapping{
			393: 675,
			404: 675,
			477: 675,
			573: 675,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=east,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3908,
			404: 3909,
			477: 4412,
			573: 4412,
		},
	},
	{
		"minecraft:composter[level=5]",
		nil,
		NewMapping{
			477: 11267,
			573: 11283,
		},
	},
	{
		"minecraft:brown_shulker_box[facing=north]",
		nil,
		NewMapping{
			477: 8814,
			573: 8814,
			4:   3698,
			393: 8289,
			404: 8290,
		},
	},
	{
		"minecraft:lava[level=7]",
		nil,
		NewMapping{
			573: 57,
			4:   167,
			393: 57,
			404: 57,
			477: 57,
		},
	},
	{
		"minecraft:dark_oak_fence[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 8184,
			393: 7659,
			404: 7660,
			477: 8184,
		},
	},
	{
		"minecraft:fire[age=15,east=true,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1621,
			404: 1622,
			477: 1925,
			573: 1925,
		},
	},
	{
		"minecraft:oak_door[facing=north,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			393: 3112,
			404: 3113,
			477: 3576,
			573: 3576,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6756,
			404: 6757,
			477: 7263,
			573: 7263,
		},
	},
	{
		"minecraft:jungle_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 5108,
			404: 5109,
			477: 5612,
			573: 5612,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6136,
			404: 6137,
			477: 6643,
			573: 6643,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9450,
			573: 9450,
		},
	},
	{
		"minecraft:quartz_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 6246,
			393: 5739,
			404: 5740,
			477: 6246,
		},
	},
	{
		"minecraft:quartz_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 5759,
			404: 5760,
			477: 6266,
			573: 6266,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=south,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			404: 3743,
			477: 4246,
			573: 4246,
			393: 3742,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=true,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6148,
			404: 6149,
			477: 6655,
			573: 6655,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9330,
			573: 9330,
		},
	},
	{
		"minecraft:acacia_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6390,
			404: 6391,
			477: 6897,
			573: 6897,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=south,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9243,
			573: 9243,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5871,
			404: 5872,
			477: 6378,
			573: 6378,
		},
	},
	{
		"minecraft:diorite_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10200,
			573: 10200,
		},
	},
	{
		"minecraft:jungle_sign[rotation=4,waterlogged=false]",
		nil,
		NewMapping{
			573: 3516,
			477: 3516,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=2,south=side,west=up]",
		nil,
		NewMapping{
			573: 2221,
			393: 1917,
			404: 1918,
			477: 2221,
		},
	},
	{
		"minecraft:wither_skeleton_skull[rotation=5]",
		nil,
		NewMapping{
			393: 5476,
			404: 5477,
			477: 5979,
			573: 5979,
		},
	},
	{
		"minecraft:kelp[age=4]",
		nil,
		NewMapping{
			393: 8413,
			404: 8414,
			477: 8938,
			573: 8938,
		},
	},
	{
		"minecraft:fire[age=15,east=false,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			404: 1646,
			477: 1949,
			573: 1949,
			393: 1645,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=11,south=none,west=side]",
		nil,
		NewMapping{
			393: 2866,
			404: 2867,
			477: 3170,
			573: 3170,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=17,powered=true]",
		nil,
		NewMapping{
			477: 1032,
			573: 1032,
		},
	},
	{
		"minecraft:jungle_sign[rotation=13,waterlogged=true]",
		nil,
		NewMapping{
			477: 3533,
			573: 3533,
		},
	},
	{
		"minecraft:bee_nest[facing=north,honey_level=5]",
		nil,
		NewMapping{
			573: 11292,
		},
	},
	{
		"minecraft:skeleton_wall_skull[facing=south]",
		nil,
		NewMapping{
			4:   2315,
			393: 5448,
			404: 5449,
			477: 5971,
			573: 5971,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=false,north=false,powered=false,south=false,west=false]",
		nil,
		NewMapping{
			4:   2118,
			393: 4818,
			404: 4819,
			477: 5322,
			573: 5322,
		},
	},
	{
		"minecraft:fire[age=0,east=false,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1157,
			404: 1158,
			477: 1461,
			573: 1461,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=5,south=side,west=none]",
		nil,
		NewMapping{
			393: 2954,
			404: 2955,
			477: 3258,
			573: 3258,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=13,south=up,west=side]",
		nil,
		NewMapping{
			393: 3022,
			404: 3023,
			477: 3326,
			573: 3326,
		},
	},
	{
		"minecraft:melon_stem[age=0]",
		nil,
		NewMapping{
			477: 4764,
			573: 4764,
			4:   1680,
			393: 4260,
			404: 4261,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=14,south=up,west=up]",
		nil,
		NewMapping{
			393: 3030,
			404: 3031,
			477: 3334,
			573: 3334,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=10,south=side,west=side]",
		nil,
		NewMapping{
			393: 2710,
			404: 2711,
			477: 3014,
			573: 3014,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=6,powered=false]",
		nil,
		NewMapping{
			477: 761,
			573: 761,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=9,south=none,west=up]",
		nil,
		NewMapping{
			573: 2719,
			393: 2415,
			404: 2416,
			477: 2719,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=7,south=up,west=side]",
		nil,
		NewMapping{
			393: 2248,
			404: 2249,
			477: 2552,
			573: 2552,
		},
	},
	{
		"minecraft:sign[rotation=9,waterlogged=true]",
		nil,
		NewMapping{
			393: 3093,
		},
	},
	{
		"minecraft:purpur_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 8136,
			404: 8137,
			477: 8661,
			573: 8661,
		},
	},
	{
		"minecraft:cut_sandstone_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			477: 7820,
			573: 7820,
		},
	},
	{
		"minecraft:oak_leaves[distance=1,persistent=false]",
		nil,
		NewMapping{
			477: 145,
			573: 145,
			4:   288,
			393: 145,
			404: 145,
		},
	},
	{
		"minecraft:tripwire_hook[attached=true,facing=north,powered=false]",
		nil,
		NewMapping{
			477: 5244,
			573: 5244,
			4:   2102,
			393: 4740,
			404: 4741,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=6,south=none,west=none]",
		nil,
		NewMapping{
			393: 1814,
			404: 1815,
			477: 2118,
			573: 2118,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=south,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			404: 3881,
			477: 4384,
			573: 4384,
			393: 3880,
		},
	},
	{
		"minecraft:fire[age=13,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			404: 1564,
			477: 1867,
			573: 1867,
			393: 1563,
		},
	},
	{
		"minecraft:blue_shulker_box[facing=down]",
		nil,
		NewMapping{
			4:   3680,
			393: 8288,
			404: 8289,
			477: 8813,
			573: 8813,
		},
	},
	{
		"minecraft:wall_sign[facing=south,waterlogged=true]",
		nil,
		NewMapping{
			393: 3271,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=9,powered=false]",
		nil,
		NewMapping{
			393: 467,
			404: 467,
			477: 467,
			573: 467,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=22,powered=true]",
		nil,
		NewMapping{
			393: 292,
			404: 292,
			477: 292,
			573: 292,
		},
	},
	{
		"minecraft:oak_door[facing=north,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 3111,
			404: 3112,
			477: 3575,
			573: 3575,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=13,powered=true]",
		nil,
		NewMapping{
			477: 924,
			573: 924,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 10629,
			477: 10629,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=false,north=false,powered=true,south=false,west=false]",
		nil,
		NewMapping{
			573: 5382,
			4:   2115,
			393: 4878,
			404: 4879,
			477: 5382,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6076,
			404: 6077,
			477: 6583,
			573: 6583,
		},
	},
	{
		"minecraft:pink_bed[facing=north,occupied=true,part=foot]",
		nil,
		NewMapping{
			477: 1145,
			573: 1145,
			393: 845,
			404: 845,
		},
	},
	{
		"minecraft:gray_bed[facing=north,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 860,
			404: 860,
			477: 1160,
			573: 1160,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10653,
			573: 10653,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10246,
			573: 10246,
		},
	},
	{
		"minecraft:scaffolding[bottom=true,distance=4,waterlogged=false]",
		nil,
		NewMapping{
			477: 11108,
			573: 11108,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10065,
			573: 10065,
		},
	},
	{
		"minecraft:lime_bed[facing=north,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 830,
			404: 830,
			477: 1130,
			573: 1130,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=18,powered=false]",
		nil,
		NewMapping{
			393: 485,
			404: 485,
			477: 485,
			573: 485,
		},
	},
	{
		"minecraft:nether_brick_fence[east=false,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 4515,
			404: 4516,
			477: 5019,
			573: 5019,
		},
	},
	{
		"minecraft:light_gray_banner[rotation=6]",
		nil,
		NewMapping{
			393: 6988,
			404: 6989,
			477: 7495,
			573: 7495,
		},
	},
	{
		"minecraft:granite_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9905,
			573: 9905,
		},
	},
	{
		"minecraft:smooth_red_sandstone_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			477: 10262,
			573: 10262,
		},
	},
	{
		"minecraft:chain_command_block[conditional=false,facing=down]",
		nil,
		NewMapping{
			573: 8712,
			4:   3376,
			393: 8187,
			404: 8188,
			477: 8712,
		},
	},
	{
		"minecraft:black_concrete",
		nil,
		NewMapping{
			4:   4031,
			393: 8392,
			404: 8393,
			477: 8917,
			573: 8917,
		},
	},
	{
		"minecraft:dark_oak_fence_gate[facing=north,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7486,
			404: 7487,
			477: 8011,
			573: 8011,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9225,
			573: 9225,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=true,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 11011,
			573: 11011,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6478,
			404: 6479,
			477: 6985,
			573: 6985,
		},
	},
	{
		"minecraft:stone_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9674,
			573: 9674,
		},
	},
	{
		"minecraft:damaged_anvil[facing=south]",
		nil,
		NewMapping{
			4:   2328,
			393: 5576,
			404: 5577,
			477: 6083,
			573: 6083,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 3215,
			404: 3216,
			477: 3679,
			573: 3679,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=14,powered=false]",
		nil,
		NewMapping{
			393: 377,
			404: 377,
			477: 377,
			573: 377,
		},
	},
	{
		"minecraft:birch_door[facing=west,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7779,
			404: 7780,
			477: 8304,
			573: 8304,
		},
	},
	{
		"minecraft:jungle_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 5054,
			404: 5055,
			477: 5558,
			573: 5558,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=6,south=up,west=up]",
		nil,
		NewMapping{
			393: 2526,
			404: 2527,
			477: 2830,
			573: 2830,
		},
	},
	{
		"minecraft:oak_door[facing=north,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 3116,
			404: 3117,
			477: 3580,
			573: 3580,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=12,powered=false]",
		nil,
		NewMapping{
			477: 823,
			573: 823,
		},
	},
	{
		"minecraft:green_concrete_powder",
		nil,
		NewMapping{
			4:   4045,
			393: 8406,
			404: 8407,
			477: 8931,
			573: 8931,
		},
	},
	{
		"minecraft:quartz_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 5700,
			404: 5701,
			477: 6207,
			573: 6207,
		},
	},
	{
		"minecraft:tube_coral_wall_fan[facing=north,waterlogged=false]",
		nil,
		NewMapping{
			477: 9065,
			573: 9065,
			393: 8505,
			404: 8521,
		},
	},
	{
		"minecraft:beehive[facing=east,honey_level=4]",
		nil,
		NewMapping{
			573: 11333,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=true,north=true,powered=false,south=false,west=true]",
		nil,
		NewMapping{
			393: 4793,
			404: 4794,
			477: 5297,
			573: 5297,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10960,
			477: 10960,
		},
	},
	{
		"minecraft:light_blue_carpet",
		nil,
		NewMapping{
			4:   2739,
			393: 6826,
			404: 6827,
			477: 7333,
			573: 7333,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=1,south=up,west=up]",
		nil,
		NewMapping{
			393: 2337,
			404: 2338,
			477: 2641,
			573: 2641,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=24,powered=false]",
		nil,
		NewMapping{
			393: 347,
			404: 347,
			477: 347,
			573: 347,
		},
	},
	{
		"minecraft:light_blue_banner[rotation=15]",
		nil,
		NewMapping{
			573: 7424,
			393: 6917,
			404: 6918,
			477: 7424,
		},
	},
	{
		"minecraft:acacia_fence[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 7634,
			404: 7635,
			477: 8159,
			573: 8159,
		},
	},
	{
		"minecraft:cactus[age=15]",
		nil,
		NewMapping{
			393: 3440,
			404: 3441,
			477: 3944,
			573: 3944,
			4:   1311,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=14,south=side,west=none]",
		nil,
		NewMapping{
			404: 2604,
			477: 2907,
			573: 2907,
			393: 2603,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=10,south=side,west=none]",
		nil,
		NewMapping{
			393: 2279,
			404: 2280,
			477: 2583,
			573: 2583,
		},
	},
	{
		"minecraft:stripped_spruce_wood[axis=z]",
		nil,
		NewMapping{
			393: 131,
			404: 131,
			477: 131,
			573: 131,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 3233,
			477: 3696,
			573: 3696,
			393: 3232,
		},
	},
	{
		"minecraft:birch_log[axis=z]",
		nil,
		NewMapping{
			393: 80,
			404: 80,
			477: 80,
			573: 80,
			4:   282,
		},
	},
	{
		"minecraft:blue_bed[facing=north,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 924,
			404: 924,
			477: 1224,
			573: 1224,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9812,
			573: 9812,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=22,powered=false]",
		nil,
		NewMapping{
			477: 943,
			573: 943,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=8,south=none,west=none]",
		nil,
		NewMapping{
			477: 2712,
			573: 2712,
			393: 2408,
			404: 2409,
		},
	},
	{
		"minecraft:fire[age=5,east=true,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1305,
			404: 1306,
			477: 1609,
			573: 1609,
		},
	},
	{
		"minecraft:quartz_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 6229,
			393: 5722,
			404: 5723,
			477: 6229,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=5,south=up,west=none]",
		nil,
		NewMapping{
			477: 2823,
			573: 2823,
			393: 2519,
			404: 2520,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5962,
			404: 5963,
			477: 6469,
			573: 6469,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 10990,
			477: 10990,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=false,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10740,
			573: 10740,
		},
	},
	{
		"minecraft:rail[shape=north_south]",
		nil,
		NewMapping{
			404: 3180,
			477: 3643,
			573: 3643,
			4:   1056,
			393: 3179,
		},
	},
	{
		"minecraft:repeater[delay=3,facing=north,locked=false,powered=false]",
		nil,
		NewMapping{
			393: 3548,
			404: 3549,
			477: 4052,
			573: 4052,
			4:   1498,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=true,north=true,powered=false,south=false,west=true]",
		nil,
		NewMapping{
			477: 5329,
			573: 5329,
			393: 4825,
			404: 4826,
		},
	},
	{
		"minecraft:pink_banner[rotation=15]",
		nil,
		NewMapping{
			393: 6965,
			404: 6966,
			477: 7472,
			573: 7472,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9760,
			573: 9760,
		},
	},
	{
		"minecraft:white_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5833,
			404: 5834,
			477: 6340,
			573: 6340,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=west,half=top,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3754,
			404: 3755,
			477: 4258,
			573: 4258,
		},
	},
	{
		"minecraft:oak_wood[axis=y]",
		nil,
		NewMapping{
			573: 109,
			4:   284,
			393: 109,
			404: 109,
			477: 109,
		},
	},
	{
		"minecraft:black_banner[rotation=10]",
		nil,
		NewMapping{
			393: 7104,
			404: 7105,
			477: 7611,
			573: 7611,
		},
	},
	{
		"minecraft:lectern[facing=north,has_book=false,powered=true]",
		nil,
		NewMapping{
			573: 11179,
			477: 11179,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=13,powered=false]",
		nil,
		NewMapping{
			477: 275,
			573: 275,
			393: 275,
			404: 275,
		},
	},
	{
		"minecraft:birch_door[facing=west,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7787,
			404: 7788,
			477: 8312,
			573: 8312,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=10,south=none,west=none]",
		nil,
		NewMapping{
			393: 2282,
			404: 2283,
			477: 2586,
			573: 2586,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=false,north=false,powered=false,south=true,west=true]",
		nil,
		NewMapping{
			393: 4815,
			404: 4816,
			477: 5319,
			573: 5319,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=true,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 11010,
			573: 11010,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=south,in_wall=false,open=false,powered=false]",
		nil,
		NewMapping{
			4:   2928,
			393: 7372,
			404: 7373,
			477: 7897,
			573: 7897,
		},
	},
	{
		"minecraft:diorite_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10197,
			573: 10197,
		},
	},
	{
		"minecraft:stone_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9626,
			573: 9626,
		},
	},
	{
		"minecraft:bell[attachment=double_wall,facing=west,powered=true]",
		nil,
		NewMapping{
			573: 11226,
		},
	},
	{
		"minecraft:trapped_chest[facing=north,type=single,waterlogged=false]",
		nil,
		NewMapping{
			404: 5581,
			477: 6087,
			573: 6087,
			4:   2338,
			393: 5580,
		},
	},
	{
		"minecraft:dead_tube_coral_wall_fan[facing=east,waterlogged=false]",
		nil,
		NewMapping{
			393: 8471,
			404: 8487,
			477: 9031,
			573: 9031,
		},
	},
	{
		"minecraft:oak_fence[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 3474,
			404: 3475,
			477: 3978,
			573: 3978,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6304,
			404: 6305,
			477: 6811,
			573: 6811,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=13,south=side,west=side]",
		nil,
		NewMapping{
			393: 2017,
			404: 2018,
			477: 2321,
			573: 2321,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=west,in_wall=false,open=false,powered=true]",
		nil,
		NewMapping{
			404: 7412,
			477: 7936,
			573: 7936,
			4:   2953,
			393: 7411,
		},
	},
	{
		"minecraft:yellow_shulker_box[facing=west]",
		nil,
		NewMapping{
			4:   3572,
			393: 8244,
			404: 8245,
			477: 8769,
			573: 8769,
		},
	},
	{
		"minecraft:acacia_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			404: 6403,
			477: 6909,
			573: 6909,
			393: 6402,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=13,south=none,west=side]",
		nil,
		NewMapping{
			393: 3028,
			404: 3029,
			477: 3332,
			573: 3332,
		},
	},
	{
		"minecraft:fire[age=12,east=false,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1545,
			404: 1546,
			477: 1849,
			573: 1849,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5881,
			404: 5882,
			477: 6388,
			573: 6388,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9280,
			573: 9280,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10519,
			573: 10519,
		},
	},
	{
		"minecraft:melon_stem[age=5]",
		nil,
		NewMapping{
			393: 4265,
			404: 4266,
			477: 4769,
			573: 4769,
			4:   1685,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=west,in_wall=false,open=false,powered=false]",
		nil,
		NewMapping{
			4:   2993,
			393: 7476,
			404: 7477,
			477: 8001,
			573: 8001,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=9,south=up,west=side]",
		nil,
		NewMapping{
			393: 1978,
			404: 1979,
			477: 2282,
			573: 2282,
		},
	},
	{
		"minecraft:oak_fence[east=false,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 3482,
			404: 3483,
			477: 3986,
			573: 3986,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=11,south=none,west=none]",
		nil,
		NewMapping{
			404: 1860,
			477: 2163,
			573: 2163,
			393: 1859,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=12,south=side,west=side]",
		nil,
		NewMapping{
			393: 2440,
			404: 2441,
			477: 2744,
			573: 2744,
		},
	},
	{
		"minecraft:acacia_button[face=floor,facing=south,powered=false]",
		nil,
		NewMapping{
			393: 5402,
			404: 5403,
			477: 5909,
			573: 5909,
		},
	},
	{
		"minecraft:spruce_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4903,
			404: 4904,
			477: 5407,
			573: 5407,
		},
	},
	{
		"minecraft:spruce_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			573: 5398,
			393: 4894,
			404: 4895,
			477: 5398,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=north,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7357,
			404: 7358,
			477: 7882,
			573: 7882,
		},
	},
	{
		"minecraft:oak_button[face=ceiling,facing=south,powered=false]",
		nil,
		NewMapping{
			573: 5829,
			393: 5322,
			404: 5323,
			477: 5829,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6471,
			404: 6472,
			477: 6978,
			573: 6978,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=west,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 6539,
			404: 6540,
			477: 7046,
			573: 7046,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9733,
			573: 9733,
		},
	},
	{
		"minecraft:repeater[delay=3,facing=west,locked=false,powered=false]",
		nil,
		NewMapping{
			4:   1497,
			393: 3556,
			404: 3557,
			477: 4060,
			573: 4060,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=22,powered=false]",
		nil,
		NewMapping{
			404: 743,
			477: 743,
			573: 743,
			393: 743,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6257,
			404: 6258,
			477: 6764,
			573: 6764,
		},
	},
	{
		"minecraft:acacia_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 6905,
			573: 6905,
			393: 6398,
			404: 6399,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4567,
			404: 4568,
			477: 5071,
			573: 5071,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10387,
			573: 10387,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5251,
			404: 5252,
			477: 5755,
			573: 5755,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9281,
			573: 9281,
		},
	},
	{
		"minecraft:pink_glazed_terracotta[facing=north]",
		nil,
		NewMapping{
			4:   3858,
			393: 8337,
			404: 8338,
			477: 8862,
			573: 8862,
		},
	},
	{
		"minecraft:white_terracotta",
		nil,
		NewMapping{
			4:   2544,
			393: 5804,
			404: 5805,
			477: 6311,
			573: 6311,
		},
	},
	{
		"minecraft:jungle_leaves[distance=4,persistent=false]",
		nil,
		NewMapping{
			393: 193,
			404: 193,
			477: 193,
			573: 193,
		},
	},
	{
		"minecraft:piston_head[facing=south,short=true,type=sticky]",
		nil,
		NewMapping{
			573: 1368,
			393: 1068,
			404: 1068,
			477: 1368,
		},
	},
	{
		"minecraft:dark_oak_button[face=wall,facing=east,powered=true]",
		nil,
		NewMapping{
			393: 5437,
			404: 5438,
			477: 5944,
			573: 5944,
		},
	},
	{
		"minecraft:lime_bed[facing=south,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 833,
			404: 833,
			477: 1133,
			573: 1133,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=west,in_wall=false,open=false,powered=false]",
		nil,
		NewMapping{
			573: 7937,
			4:   2945,
			393: 7412,
			404: 7413,
			477: 7937,
		},
	},
	{
		"minecraft:brain_coral[waterlogged=false]",
		nil,
		NewMapping{
			404: 8473,
			477: 8997,
			573: 8997,
		},
	},
	{
		"minecraft:spruce_sign[rotation=0,waterlogged=true]",
		nil,
		NewMapping{
			477: 3411,
			573: 3411,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   2880,
			393: 7248,
			404: 7249,
			477: 7755,
			573: 7755,
		},
	},
	{
		"minecraft:stripped_jungle_log[axis=y]",
		nil,
		NewMapping{
			393: 97,
			404: 97,
			477: 97,
			573: 97,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=false,north=true,powered=true,south=true,west=true]",
		nil,
		NewMapping{
			393: 4771,
			404: 4772,
			477: 5275,
			573: 5275,
		},
	},
	{
		"minecraft:stone_button[face=ceiling,facing=east,powered=false]",
		nil,
		NewMapping{
			404: 3415,
			477: 3918,
			573: 3918,
			393: 3414,
		},
	},
	{
		"minecraft:nether_brick_fence[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 4510,
			404: 4511,
			477: 5014,
			573: 5014,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9320,
			573: 9320,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=south,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 6524,
			404: 6525,
			477: 7031,
			573: 7031,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=11,south=side,west=none]",
		nil,
		NewMapping{
			393: 2576,
			404: 2577,
			477: 2880,
			573: 2880,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=13,south=up,west=none]",
		nil,
		NewMapping{
			393: 2447,
			404: 2448,
			477: 2751,
			573: 2751,
		},
	},
	{
		"minecraft:gray_bed[facing=north,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 863,
			404: 863,
			477: 1163,
			573: 1163,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=21,powered=false]",
		nil,
		NewMapping{
			477: 441,
			573: 441,
			393: 441,
			404: 441,
		},
	},
	{
		"minecraft:stone_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9624,
			573: 9624,
		},
	},
	{
		"minecraft:activator_rail[powered=true,shape=east_west]",
		nil,
		NewMapping{
			4:   2521,
			393: 5781,
			404: 5782,
			477: 6288,
			573: 6288,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 6662,
			573: 6662,
			393: 6155,
			404: 6156,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=15,south=side,west=side]",
		nil,
		NewMapping{
			393: 2035,
			404: 2036,
			477: 2339,
			573: 2339,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 6412,
			573: 6412,
			393: 5905,
			404: 5906,
		},
	},
	{
		"minecraft:jungle_fence[east=true,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 7584,
			404: 7585,
			477: 8109,
			573: 8109,
		},
	},
	{
		"minecraft:light_blue_wall_banner[facing=west]",
		nil,
		NewMapping{
			573: 7631,
			393: 7124,
			404: 7125,
			477: 7631,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10338,
			573: 10338,
		},
	},
	{
		"minecraft:cut_red_sandstone_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			477: 7871,
			573: 7871,
		},
	},
	{
		"minecraft:light_weighted_pressure_plate[power=5]",
		nil,
		NewMapping{
			4:   2357,
			393: 5608,
			404: 5609,
			477: 6115,
			573: 6115,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=west,in_wall=false,open=true,powered=true]",
		nil,
		NewMapping{
			573: 7934,
			4:   2957,
			393: 7409,
			404: 7410,
			477: 7934,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 3193,
			404: 3194,
			477: 3657,
			573: 3657,
		},
	},
	{
		"minecraft:spruce_door[facing=east,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			404: 7740,
			477: 8264,
			573: 8264,
			393: 7739,
		},
	},
	{
		"minecraft:zombie_head[rotation=8]",
		nil,
		NewMapping{
			393: 5499,
			404: 5500,
			477: 6002,
			573: 6002,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10467,
			573: 10467,
		},
	},
	{
		"minecraft:brown_banner[rotation=6]",
		nil,
		NewMapping{
			393: 7052,
			404: 7053,
			477: 7559,
			573: 7559,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=3,powered=true]",
		nil,
		NewMapping{
			404: 354,
			477: 354,
			573: 354,
			393: 354,
		},
	},
	{
		"minecraft:fire[age=10,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			404: 1468,
			477: 1771,
			573: 1771,
			393: 1467,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=11,powered=true]",
		nil,
		NewMapping{
			477: 870,
			573: 870,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=9,waterlogged=false]",
		nil,
		NewMapping{
			477: 3558,
			573: 3558,
		},
	},
	{
		"minecraft:skeleton_wall_skull[facing=east]",
		nil,
		NewMapping{
			4:   2309,
			393: 5450,
			404: 5451,
			477: 5973,
			573: 5973,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=north,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			404: 3605,
			477: 4108,
			573: 4108,
			4:   1540,
			393: 3604,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=6,south=up,west=side]",
		nil,
		NewMapping{
			404: 2240,
			477: 2543,
			573: 2543,
			393: 2239,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=8,south=up,west=up]",
		nil,
		NewMapping{
			393: 2400,
			404: 2401,
			477: 2704,
			573: 2704,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=21,powered=true]",
		nil,
		NewMapping{
			477: 840,
			573: 840,
		},
	},
	{
		"minecraft:andesite_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 9949,
			477: 9949,
		},
	},
	{
		"minecraft:diorite_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 11084,
			573: 11084,
		},
	},
	{
		"minecraft:fire[age=8,east=true,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			573: 1701,
			393: 1397,
			404: 1398,
			477: 1701,
		},
	},
	{
		"minecraft:lime_banner[rotation=3]",
		nil,
		NewMapping{
			393: 6937,
			404: 6938,
			477: 7444,
			573: 7444,
		},
	},
	{
		"minecraft:blue_banner[rotation=7]",
		nil,
		NewMapping{
			404: 7038,
			477: 7544,
			573: 7544,
			393: 7037,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=23,powered=false]",
		nil,
		NewMapping{
			477: 645,
			573: 645,
			393: 645,
			404: 645,
		},
	},
	{
		"minecraft:dark_oak_wall_sign[facing=south,waterlogged=true]",
		nil,
		NewMapping{
			477: 3775,
			573: 3775,
		},
	},
	{
		"minecraft:lever[face=wall,facing=south,powered=true]",
		nil,
		NewMapping{
			477: 3791,
			573: 3791,
			4:   1115,
			393: 3287,
			404: 3288,
		},
	},
	{
		"minecraft:diorite_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10202,
			573: 10202,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9245,
			573: 9245,
		},
	},
	{
		"minecraft:orange_shulker_box[facing=south]",
		nil,
		NewMapping{
			4:   3523,
			393: 8225,
			404: 8226,
			477: 8750,
			573: 8750,
		},
	},
	{
		"minecraft:fire[age=8,east=true,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1399,
			404: 1400,
			477: 1703,
			573: 1703,
		},
	},
	{
		"minecraft:quartz_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 5774,
			404: 5775,
			477: 6281,
			573: 6281,
		},
	},
	{
		"minecraft:jungle_button[face=wall,facing=east,powered=true]",
		nil,
		NewMapping{
			404: 5390,
			477: 5896,
			573: 5896,
			393: 5389,
		},
	},
	{
		"minecraft:end_stone_brick_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			477: 10287,
			573: 10287,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10336,
			573: 10336,
		},
	},
	{
		"minecraft:red_wool",
		nil,
		NewMapping{
			477: 1397,
			573: 1397,
			4:   574,
			393: 1097,
			404: 1097,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=0,south=none,west=none]",
		nil,
		NewMapping{
			573: 2064,
			393: 1760,
			404: 1761,
			477: 2064,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=3,south=up,west=up]",
		nil,
		NewMapping{
			393: 2499,
			404: 2500,
			477: 2803,
			573: 2803,
		},
	},
	{
		"minecraft:oak_leaves[distance=7,persistent=true]",
		nil,
		NewMapping{
			573: 156,
			393: 156,
			404: 156,
			477: 156,
		},
	},
	{
		"minecraft:spruce_door[facing=west,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7719,
			404: 7720,
			477: 8244,
			573: 8244,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9253,
			573: 9253,
		},
	},
	{
		"minecraft:birch_sign[rotation=1,waterlogged=true]",
		nil,
		NewMapping{
			477: 3445,
			573: 3445,
		},
	},
	{
		"minecraft:player_head[rotation=2]",
		nil,
		NewMapping{
			393: 5513,
			404: 5514,
			477: 6016,
			573: 6016,
		},
	},
	{
		"minecraft:cake[bites=3]",
		nil,
		NewMapping{
			393: 3509,
			404: 3510,
			477: 4013,
			573: 4013,
			4:   1475,
		},
	},
	{
		"minecraft:activator_rail[powered=false,shape=ascending_east]",
		nil,
		NewMapping{
			573: 6295,
			4:   2514,
			393: 5788,
			404: 5789,
			477: 6295,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4724,
			404: 4725,
			477: 5228,
			573: 5228,
		},
	},
	{
		"minecraft:acacia_fence[east=true,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 7625,
			404: 7626,
			477: 8150,
			573: 8150,
		},
	},
	{
		"minecraft:fire[age=3,east=true,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 1234,
			404: 1235,
			477: 1538,
			573: 1538,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=4,powered=false]",
		nil,
		NewMapping{
			393: 507,
			404: 507,
			477: 507,
			573: 507,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=10,south=up,west=none]",
		nil,
		NewMapping{
			393: 2996,
			404: 2997,
			477: 3300,
			573: 3300,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10639,
			573: 10639,
		},
	},
	{
		"minecraft:polished_diorite",
		nil,
		NewMapping{
			404: 5,
			477: 5,
			573: 5,
			4:   20,
			393: 5,
		},
	},
	{
		"minecraft:pink_wall_banner[facing=south]",
		nil,
		NewMapping{
			393: 7135,
			404: 7136,
			477: 7642,
			573: 7642,
		},
	},
	{
		"minecraft:purpur_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 8077,
			404: 8078,
			477: 8602,
			573: 8602,
		},
	},
	{
		"minecraft:spruce_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4906,
			404: 4907,
			477: 5410,
			573: 5410,
		},
	},
	{
		"minecraft:quartz_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 5765,
			404: 5766,
			477: 6272,
			573: 6272,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9853,
			573: 9853,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9247,
			573: 9247,
		},
	},
	{
		"minecraft:stone_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9614,
			573: 9614,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 7221,
			393: 6714,
			404: 6715,
			477: 7221,
		},
	},
	{
		"minecraft:purpur_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 8117,
			404: 8118,
			477: 8642,
			573: 8642,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=10,south=none,west=side]",
		nil,
		NewMapping{
			477: 2729,
			573: 2729,
			393: 2425,
			404: 2426,
		},
	},
	{
		"minecraft:gravel",
		nil,
		NewMapping{
			573: 68,
			4:   208,
			393: 68,
			404: 68,
			477: 68,
		},
	},
	{
		"minecraft:quartz_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   2502,
			393: 5721,
			404: 5722,
			477: 6228,
			573: 6228,
		},
	},
	{
		"minecraft:lime_terracotta",
		nil,
		NewMapping{
			477: 6316,
			573: 6316,
			4:   2549,
			393: 5809,
			404: 5810,
		},
	},
	{
		"minecraft:iron_bars[east=false,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			4:   1616,
			393: 4210,
			404: 4211,
			477: 4714,
			573: 4714,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=7,south=up,west=none]",
		nil,
		NewMapping{
			393: 1961,
			404: 1962,
			477: 2265,
			573: 2265,
		},
	},
	{
		"minecraft:dragon_head[rotation=1]",
		nil,
		NewMapping{
			393: 5552,
			404: 5553,
			477: 6055,
			573: 6055,
		},
	},
	{
		"minecraft:oak_door[facing=north,half=lower,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 3115,
			404: 3116,
			477: 3579,
			573: 3579,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=8,south=side,west=none]",
		nil,
		NewMapping{
			393: 2405,
			404: 2406,
			477: 2709,
			573: 2709,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=west,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9503,
			573: 9503,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9336,
			573: 9336,
		},
	},
	{
		"minecraft:purple_bed[facing=north,occupied=false,part=foot]",
		nil,
		NewMapping{
			404: 911,
			477: 1211,
			573: 1211,
			393: 911,
		},
	},
	{
		"minecraft:fire[age=0,east=true,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			404: 1139,
			477: 1442,
			573: 1442,
			393: 1138,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=12,south=side,west=side]",
		nil,
		NewMapping{
			404: 2729,
			477: 3032,
			573: 3032,
			393: 2728,
		},
	},
	{
		"minecraft:horn_coral_wall_fan[facing=south,waterlogged=true]",
		nil,
		NewMapping{
			393: 8538,
			404: 8554,
			477: 9098,
			573: 9098,
		},
	},
	{
		"minecraft:zombie_head[rotation=15]",
		nil,
		NewMapping{
			393: 5506,
			404: 5507,
			477: 6009,
			573: 6009,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=9,south=side,west=none]",
		nil,
		NewMapping{
			393: 2846,
			404: 2847,
			477: 3150,
			573: 3150,
		},
	},
	{
		"minecraft:jungle_door[facing=south,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7832,
			404: 7833,
			477: 8357,
			573: 8357,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=east,in_wall=false,open=false,powered=false]",
		nil,
		NewMapping{
			4:   2931,
			393: 7388,
			404: 7389,
			477: 7913,
			573: 7913,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=9,powered=true]",
		nil,
		NewMapping{
			393: 616,
			404: 616,
			477: 616,
			573: 616,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10384,
			573: 10384,
		},
	},
	{
		"minecraft:fire[age=8,east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1419,
			404: 1420,
			477: 1723,
			573: 1723,
		},
	},
	{
		"minecraft:iron_door[facing=east,half=upper,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			4:   1145,
			393: 3358,
			404: 3359,
			477: 3862,
			573: 3862,
		},
	},
	{
		"minecraft:fire[age=9,east=false,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1439,
			404: 1440,
			477: 1743,
			573: 1743,
		},
	},
	{
		"minecraft:fire[age=15,east=true,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1619,
			404: 1620,
			477: 1923,
			573: 1923,
		},
	},
	{
		"minecraft:brick_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4361,
			404: 4362,
			477: 4865,
			573: 4865,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=3,south=up,west=side]",
		nil,
		NewMapping{
			573: 2804,
			393: 2500,
			404: 2501,
			477: 2804,
		},
	},
	{
		"minecraft:fire[age=14,east=true,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1597,
			404: 1598,
			477: 1901,
			573: 1901,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=13,south=side,west=up]",
		nil,
		NewMapping{
			573: 2320,
			393: 2016,
			404: 2017,
			477: 2320,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 3708,
			573: 3708,
			393: 3244,
			404: 3245,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=east,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			573: 7970,
			393: 7445,
			404: 7446,
			477: 7970,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=south,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6586,
			404: 6587,
			477: 7093,
			573: 7093,
		},
	},
	{
		"minecraft:cyan_bed[facing=west,occupied=false,part=head]",
		nil,
		NewMapping{
			404: 902,
			477: 1202,
			573: 1202,
			393: 902,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 10369,
			477: 10369,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=south,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10041,
			573: 10041,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=7,powered=false]",
		nil,
		NewMapping{
			477: 813,
			573: 813,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=8,south=side,west=none]",
		nil,
		NewMapping{
			477: 2133,
			573: 2133,
			393: 1829,
			404: 1830,
		},
	},
	{
		"minecraft:black_wall_banner[facing=east]",
		nil,
		NewMapping{
			393: 7173,
			404: 7174,
			477: 7680,
			573: 7680,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=true,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5892,
			404: 5893,
			477: 6399,
			573: 6399,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=10,south=side,west=up]",
		nil,
		NewMapping{
			393: 2421,
			404: 2422,
			477: 2725,
			573: 2725,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=true,north=false,powered=false,south=true,west=true]",
		nil,
		NewMapping{
			393: 4767,
			404: 4768,
			477: 5271,
			573: 5271,
		},
	},
	{
		"minecraft:sweet_berry_bush[age=3]",
		nil,
		NewMapping{
			477: 11251,
			573: 11267,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10331,
			573: 10331,
		},
	},
	{
		"minecraft:bee_nest[facing=north,honey_level=4]",
		nil,
		NewMapping{
			573: 11291,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=11,powered=false]",
		nil,
		NewMapping{
			393: 671,
			404: 671,
			477: 671,
			573: 671,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			573: 10023,
			477: 10023,
		},
	},
	{
		"minecraft:lantern[hanging=true]",
		nil,
		NewMapping{
			477: 11214,
			573: 11230,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9428,
			573: 9428,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10867,
			573: 10867,
		},
	},
	{
		"minecraft:mossy_stone_brick_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			477: 10268,
			573: 10268,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=north,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9375,
			573: 9375,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9160,
			573: 9160,
		},
	},
	{
		"minecraft:stone_button[face=wall,facing=west,powered=true]",
		nil,
		NewMapping{
			4:   1242,
			393: 3403,
			404: 3404,
			477: 3907,
			573: 3907,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 6210,
			477: 6716,
			573: 6716,
			393: 6209,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=0,south=side,west=up]",
		nil,
		NewMapping{
			393: 2907,
			404: 2908,
			477: 3211,
			573: 3211,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4428,
			404: 4429,
			477: 4932,
			573: 4932,
		},
	},
	{
		"minecraft:jungle_door[facing=east,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7867,
			404: 7868,
			477: 8392,
			573: 8392,
		},
	},
	{
		"minecraft:fire[age=4,east=false,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 1288,
			404: 1289,
			477: 1592,
			573: 1592,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=13,south=none,west=none]",
		nil,
		NewMapping{
			573: 2469,
			393: 2165,
			404: 2166,
			477: 2469,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=3,powered=false]",
		nil,
		NewMapping{
			477: 755,
			573: 755,
		},
	},
	{
		"minecraft:pink_wall_banner[facing=west]",
		nil,
		NewMapping{
			477: 7643,
			573: 7643,
			393: 7136,
			404: 7137,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=16,powered=true]",
		nil,
		NewMapping{
			393: 430,
			404: 430,
			477: 430,
			573: 430,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=true,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 10942,
			477: 10942,
		},
	},
	{
		"minecraft:granite",
		nil,
		NewMapping{
			4:   17,
			393: 2,
			404: 2,
			477: 2,
			573: 2,
		},
	},
	{
		"minecraft:redstone_lamp[lit=true]",
		nil,
		NewMapping{
			477: 5140,
			573: 5140,
			4:   1984,
			393: 4636,
			404: 4637,
		},
	},
	{
		"minecraft:birch_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 5472,
			393: 4968,
			404: 4969,
			477: 5472,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6798,
			404: 6799,
			477: 7305,
			573: 7305,
		},
	},
	{
		"minecraft:purpur_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			404: 8102,
			477: 8626,
			573: 8626,
			393: 8101,
		},
	},
	{
		"minecraft:granite_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9891,
			573: 9891,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10040,
			573: 10040,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=east,half=top,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			477: 4402,
			573: 4402,
			393: 3898,
			404: 3899,
		},
	},
	{
		"minecraft:green_banner[rotation=12]",
		nil,
		NewMapping{
			393: 7074,
			404: 7075,
			477: 7581,
			573: 7581,
		},
	},
	{
		"minecraft:yellow_banner[rotation=13]",
		nil,
		NewMapping{
			393: 6931,
			404: 6932,
			477: 7438,
			573: 7438,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=12,south=side,west=none]",
		nil,
		NewMapping{
			393: 2297,
			404: 2298,
			477: 2601,
			573: 2601,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=south,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9161,
			573: 9161,
		},
	},
	{
		"minecraft:gray_banner[rotation=12]",
		nil,
		NewMapping{
			393: 6978,
			404: 6979,
			477: 7485,
			573: 7485,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9143,
			573: 9143,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9771,
			573: 9771,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=3,south=up,west=side]",
		nil,
		NewMapping{
			393: 2068,
			404: 2069,
			477: 2372,
			573: 2372,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5891,
			404: 5892,
			477: 6398,
			573: 6398,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10605,
			573: 10605,
		},
	},
	{
		"minecraft:stone_stairs[facing=east,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9673,
			573: 9673,
		},
	},
	{
		"minecraft:smooth_red_sandstone_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			477: 10259,
			573: 10259,
		},
	},
	{
		"minecraft:birch_sign[rotation=11,waterlogged=true]",
		nil,
		NewMapping{
			477: 3465,
			573: 3465,
		},
	},
	{
		"minecraft:spruce_fence[east=false,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7547,
			404: 7548,
			477: 8072,
			573: 8072,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 6427,
			393: 5920,
			404: 5921,
			477: 6427,
		},
	},
	{
		"minecraft:iron_bars[east=true,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 4189,
			477: 4692,
			573: 4692,
			393: 4188,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=false,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 4007,
			404: 4008,
			477: 4511,
			573: 4511,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10388,
			573: 10388,
		},
	},
	{
		"minecraft:bamboo[age=1,leaves=large,stage=0]",
		nil,
		NewMapping{
			477: 9126,
			573: 9126,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6197,
			404: 6198,
			477: 6704,
			573: 6704,
		},
	},
	{
		"minecraft:jungle_sign[rotation=0,waterlogged=false]",
		nil,
		NewMapping{
			477: 3508,
			573: 3508,
		},
	},
	{
		"minecraft:lectern[facing=east,has_book=true,powered=true]",
		nil,
		NewMapping{
			477: 11189,
			573: 11189,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6119,
			404: 6120,
			477: 6626,
			573: 6626,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=true,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 3994,
			404: 3995,
			477: 4498,
			573: 4498,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=south,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			404: 6513,
			477: 7019,
			573: 7019,
			393: 6512,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 6571,
			477: 7077,
			573: 7077,
			393: 6570,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6060,
			404: 6061,
			477: 6567,
			573: 6567,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=false,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10710,
			573: 10710,
		},
	},
	{
		"minecraft:white_shulker_box[facing=up]",
		nil,
		NewMapping{
			4:   3505,
			393: 8221,
			404: 8222,
			477: 8746,
			573: 8746,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=15,south=none,west=side]",
		nil,
		NewMapping{
			393: 2038,
			404: 2039,
			477: 2342,
			573: 2342,
		},
	},
	{
		"minecraft:yellow_bed[facing=east,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 824,
			404: 824,
			477: 1124,
			573: 1124,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=west,half=top,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 6527,
			404: 6528,
			477: 7034,
			573: 7034,
		},
	},
	{
		"minecraft:diorite_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10218,
			573: 10218,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=12,south=side,west=none]",
		nil,
		NewMapping{
			393: 2729,
			404: 2730,
			477: 3033,
			573: 3033,
		},
	},
	{
		"minecraft:oak_wall_sign[facing=north,waterlogged=true]",
		nil,
		NewMapping{
			404: 3270,
			477: 3733,
			573: 3733,
		},
	},
	{
		"minecraft:mossy_cobblestone_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			477: 10279,
			573: 10279,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=8,south=up,west=up]",
		nil,
		NewMapping{
			393: 2544,
			404: 2545,
			477: 2848,
			573: 2848,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=17,powered=false]",
		nil,
		NewMapping{
			477: 683,
			573: 683,
			393: 683,
			404: 683,
		},
	},
	{
		"minecraft:spruce_stairs[facing=south,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4909,
			404: 4910,
			477: 5413,
			573: 5413,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=18,powered=true]",
		nil,
		NewMapping{
			573: 734,
			393: 734,
			404: 734,
			477: 734,
		},
	},
	{
		"minecraft:brick_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4396,
			404: 4397,
			477: 4900,
			573: 4900,
		},
	},
	{
		"minecraft:spruce_sign[rotation=3,waterlogged=false]",
		nil,
		NewMapping{
			477: 3418,
			573: 3418,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10101,
			573: 10101,
		},
	},
	{
		"minecraft:observer[facing=east,powered=false]",
		nil,
		NewMapping{
			393: 8202,
			404: 8203,
			477: 8727,
			573: 8727,
			4:   3493,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=false,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 6742,
			573: 6742,
			4:   2572,
			393: 6235,
			404: 6236,
		},
	},
	{
		"minecraft:chipped_anvil[facing=north]",
		nil,
		NewMapping{
			404: 5572,
			477: 6078,
			573: 6078,
			4:   2326,
			393: 5571,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6617,
			404: 6618,
			477: 7124,
			573: 7124,
		},
	},
	{
		"minecraft:brick_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 4903,
			573: 4903,
			393: 4399,
			404: 4400,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   2884,
			393: 7238,
			404: 7239,
			477: 7745,
			573: 7745,
		},
	},
	{
		"minecraft:lever[face=floor,facing=south,powered=false]",
		nil,
		NewMapping{
			477: 3784,
			573: 3784,
			393: 3280,
			404: 3281,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=11,powered=false]",
		nil,
		NewMapping{
			404: 321,
			477: 321,
			573: 321,
			393: 321,
		},
	},
	{
		"minecraft:dark_oak_button[face=floor,facing=south,powered=true]",
		nil,
		NewMapping{
			573: 5932,
			393: 5425,
			404: 5426,
			477: 5932,
		},
	},
	{
		"minecraft:stone_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9648,
			573: 9648,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6583,
			404: 6584,
			477: 7090,
			573: 7090,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4581,
			404: 4582,
			477: 5085,
			573: 5085,
		},
	},
	{
		"minecraft:birch_fence[east=false,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7571,
			404: 7572,
			477: 8096,
			573: 8096,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9414,
			573: 9414,
		},
	},
	{
		"minecraft:nether_wart[age=0]",
		nil,
		NewMapping{
			4:   1840,
			393: 4608,
			404: 4609,
			477: 5112,
			573: 5112,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=east,in_wall=false,open=true,powered=true]",
		nil,
		NewMapping{
			477: 7910,
			573: 7910,
			4:   2943,
			393: 7385,
			404: 7386,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9796,
			477: 9796,
		},
	},
	{
		"minecraft:infested_stone_bricks",
		nil,
		NewMapping{
			477: 4487,
			573: 4487,
			4:   1554,
			393: 3979,
			404: 3980,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4593,
			404: 4594,
			477: 5097,
			573: 5097,
		},
	},
	{
		"minecraft:dead_brain_coral_wall_fan[facing=west,waterlogged=false]",
		nil,
		NewMapping{
			393: 8477,
			404: 8493,
			477: 9037,
			573: 9037,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=west,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3691,
			404: 3692,
			477: 4195,
			573: 4195,
		},
	},
	{
		"minecraft:fire[age=4,east=false,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1286,
			404: 1287,
			477: 1590,
			573: 1590,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9329,
			573: 9329,
		},
	},
	{
		"minecraft:bee_nest[facing=north,honey_level=0]",
		nil,
		NewMapping{
			573: 11287,
		},
	},
	{
		"minecraft:jungle_door[facing=east,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7855,
			404: 7856,
			477: 8380,
			573: 8380,
			4:   3130,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=false,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5914,
			404: 5915,
			477: 6421,
			573: 6421,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4557,
			404: 4558,
			477: 5061,
			573: 5061,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=7,south=up,west=up]",
		nil,
		NewMapping{
			393: 2967,
			404: 2968,
			477: 3271,
			573: 3271,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10763,
			573: 10763,
		},
	},
	{
		"minecraft:pink_wool",
		nil,
		NewMapping{
			393: 1089,
			404: 1089,
			477: 1389,
			573: 1389,
			4:   566,
		},
	},
	{
		"minecraft:brown_stained_glass",
		nil,
		NewMapping{
			4:   1532,
			393: 3589,
			404: 3590,
			477: 4093,
			573: 4093,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=false,north=true,powered=false,south=true,west=false]",
		nil,
		NewMapping{
			573: 5376,
			393: 4872,
			404: 4873,
			477: 5376,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=5,south=side,west=up]",
		nil,
		NewMapping{
			477: 2536,
			573: 2536,
			393: 2232,
			404: 2233,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=15,south=side,west=none]",
		nil,
		NewMapping{
			393: 2036,
			404: 2037,
			477: 2340,
			573: 2340,
		},
	},
	{
		"minecraft:purpur_stairs[facing=south,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 8102,
			404: 8103,
			477: 8627,
			573: 8627,
		},
	},
	{
		"minecraft:acacia_sign[rotation=9,waterlogged=false]",
		nil,
		NewMapping{
			477: 3494,
			573: 3494,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=west,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3755,
			404: 3756,
			477: 4259,
			573: 4259,
		},
	},
	{
		"minecraft:sea_pickle[pickles=1,waterlogged=false]",
		nil,
		NewMapping{
			404: 8581,
			477: 9105,
			573: 9105,
			393: 8565,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=13,south=up,west=none]",
		nil,
		NewMapping{
			393: 2591,
			404: 2592,
			477: 2895,
			573: 2895,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=10,south=up,west=none]",
		nil,
		NewMapping{
			393: 2852,
			404: 2853,
			477: 3156,
			573: 3156,
		},
	},
	{
		"minecraft:light_gray_banner[rotation=13]",
		nil,
		NewMapping{
			393: 6995,
			404: 6996,
			477: 7502,
			573: 7502,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10524,
			573: 10524,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10840,
			573: 10840,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10952,
			573: 10952,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10018,
			573: 10018,
		},
	},
	{
		"minecraft:smoker[facing=east,lit=true]",
		nil,
		NewMapping{
			477: 11153,
			573: 11153,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9866,
			573: 9866,
		},
	},
	{
		"minecraft:furnace[facing=north,lit=true]",
		nil,
		NewMapping{
			477: 3371,
			573: 3371,
			4:   994,
			393: 3067,
			404: 3068,
		},
	},
	{
		"minecraft:heavy_weighted_pressure_plate[power=6]",
		nil,
		NewMapping{
			404: 5626,
			477: 6132,
			573: 6132,
			4:   2374,
			393: 5625,
		},
	},
	{
		"minecraft:white_shulker_box[facing=down]",
		nil,
		NewMapping{
			573: 8747,
			4:   3504,
			393: 8222,
			404: 8223,
			477: 8747,
		},
	},
	{
		"minecraft:jungle_fence[east=true,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 7586,
			404: 7587,
			477: 8111,
			573: 8111,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=14,south=none,west=up]",
		nil,
		NewMapping{
			393: 1884,
			404: 1885,
			477: 2188,
			573: 2188,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5949,
			404: 5950,
			477: 6456,
			573: 6456,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 6792,
			404: 6793,
			477: 7299,
			573: 7299,
		},
	},
	{
		"minecraft:fire[age=2,east=true,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1207,
			404: 1208,
			477: 1511,
			573: 1511,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=south,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9801,
			573: 9801,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=north,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3593,
			404: 3594,
			477: 4097,
			573: 4097,
		},
	},
	{
		"minecraft:dark_oak_door[facing=north,half=upper,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			477: 8459,
			573: 8459,
			393: 7934,
			404: 7935,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=false,north=true,powered=false,south=false,west=true]",
		nil,
		NewMapping{
			393: 4809,
			404: 4810,
			477: 5313,
			573: 5313,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=14,powered=true]",
		nil,
		NewMapping{
			573: 326,
			393: 326,
			404: 326,
			477: 326,
		},
	},
	{
		"minecraft:fire[age=9,east=false,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 1448,
			404: 1449,
			477: 1752,
			573: 1752,
		},
	},
	{
		"minecraft:jungle_wall_sign[facing=south,waterlogged=true]",
		nil,
		NewMapping{
			477: 3767,
			573: 3767,
		},
	},
	{
		"minecraft:spruce_stairs[facing=east,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 5448,
			573: 5448,
			393: 4944,
			404: 4945,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=south,half=top,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			477: 4376,
			573: 4376,
			393: 3872,
			404: 3873,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=12,south=none,west=up]",
		nil,
		NewMapping{
			393: 2874,
			404: 2875,
			477: 3178,
			573: 3178,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=east,half=top,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 6543,
			404: 6544,
			477: 7050,
			573: 7050,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9306,
			477: 9306,
		},
	},
	{
		"minecraft:cobblestone_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 5139,
			477: 5642,
			573: 5642,
			393: 5138,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=7,south=up,west=side]",
		nil,
		NewMapping{
			477: 3128,
			573: 3128,
			393: 2824,
			404: 2825,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9717,
			573: 9717,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9389,
			573: 9389,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=6,south=up,west=up]",
		nil,
		NewMapping{
			393: 2814,
			404: 2815,
			477: 3118,
			573: 3118,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=8,south=side,west=side]",
		nil,
		NewMapping{
			393: 2692,
			404: 2693,
			477: 2996,
			573: 2996,
		},
	},
	{
		"minecraft:fire[age=11,east=true,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			477: 1793,
			573: 1793,
			393: 1489,
			404: 1490,
		},
	},
	{
		"minecraft:birch_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4973,
			404: 4974,
			477: 5477,
			573: 5477,
		},
	},
	{
		"minecraft:birch_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 5527,
			393: 5023,
			404: 5024,
			477: 5527,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=12,powered=false]",
		nil,
		NewMapping{
			477: 773,
			573: 773,
		},
	},
	{
		"minecraft:granite_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9906,
			573: 9906,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10734,
			573: 10734,
		},
	},
	{
		"minecraft:green_banner[rotation=10]",
		nil,
		NewMapping{
			573: 7579,
			393: 7072,
			404: 7073,
			477: 7579,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=false,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 4106,
			404: 4107,
			477: 4610,
			573: 4610,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5852,
			404: 5853,
			477: 6359,
			573: 6359,
		},
	},
	{
		"minecraft:lily_of_the_valley",
		nil,
		NewMapping{
			477: 1423,
			573: 1423,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=10,powered=false]",
		nil,
		NewMapping{
			477: 919,
			573: 919,
		},
	},
	{
		"minecraft:fire[age=8,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			573: 1726,
			4:   824,
			393: 1422,
			404: 1423,
			477: 1726,
		},
	},
	{
		"minecraft:oak_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 1674,
			404: 1675,
			477: 1978,
			573: 1978,
		},
	},
	{
		"minecraft:birch_sign[rotation=9,waterlogged=true]",
		nil,
		NewMapping{
			477: 3461,
			573: 3461,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 10084,
			573: 10084,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=false,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10996,
			573: 10996,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=12,south=side,west=side]",
		nil,
		NewMapping{
			393: 2296,
			404: 2297,
			477: 2600,
			573: 2600,
		},
	},
	{
		"minecraft:andesite_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9939,
			573: 9939,
		},
	},
	{
		"minecraft:granite_stairs[facing=west,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9902,
			573: 9902,
		},
	},
	{
		"minecraft:water[level=7]",
		nil,
		NewMapping{
			4:   151,
			393: 41,
			404: 41,
			477: 41,
			573: 41,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4729,
			404: 4730,
			477: 5233,
			573: 5233,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4708,
			404: 4709,
			477: 5212,
			573: 5212,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=0,south=side,west=none]",
		nil,
		NewMapping{
			393: 2909,
			404: 2910,
			477: 3213,
			573: 3213,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=10,powered=false]",
		nil,
		NewMapping{
			477: 669,
			573: 669,
			393: 669,
			404: 669,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10835,
			573: 10835,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10365,
			573: 10365,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=3,powered=false]",
		nil,
		NewMapping{
			477: 1005,
			573: 1005,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 7222,
			573: 7222,
			393: 6715,
			404: 6716,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=south,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			573: 4444,
			393: 3940,
			404: 3941,
			477: 4444,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=24,powered=false]",
		nil,
		NewMapping{
			477: 397,
			573: 397,
			393: 397,
			404: 397,
		},
	},
	{
		"minecraft:chest[facing=south,type=single,waterlogged=true]",
		nil,
		NewMapping{
			393: 1734,
			404: 1735,
			477: 2038,
			573: 2038,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=east,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			477: 7938,
			573: 7938,
			393: 7413,
			404: 7414,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=south,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			573: 4255,
			393: 3751,
			404: 3752,
			477: 4255,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5179,
			404: 5180,
			477: 5683,
			573: 5683,
		},
	},
	{
		"minecraft:light_blue_banner[rotation=12]",
		nil,
		NewMapping{
			393: 6914,
			404: 6915,
			477: 7421,
			573: 7421,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=6,south=none,west=none]",
		nil,
		NewMapping{
			393: 1958,
			404: 1959,
			477: 2262,
			573: 2262,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 6576,
			393: 6069,
			404: 6070,
			477: 6576,
		},
	},
	{
		"minecraft:polished_andesite_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			477: 10319,
			573: 10319,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=south,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3803,
			404: 3804,
			477: 4307,
			573: 4307,
		},
	},
	{
		"minecraft:stone_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9651,
			573: 9651,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=false,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 6300,
			477: 6806,
			573: 6806,
			4:   2574,
			393: 6299,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 4412,
			404: 4413,
			477: 4916,
			573: 4916,
		},
	},
	{
		"minecraft:light_gray_bed[facing=north,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 879,
			404: 879,
			477: 1179,
			573: 1179,
		},
	},
	{
		"minecraft:quartz_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			404: 5759,
			477: 6265,
			573: 6265,
			393: 5758,
		},
	},
	{
		"minecraft:fire[age=4,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1292,
			404: 1293,
			477: 1596,
			573: 1596,
		},
	},
	{
		"minecraft:dark_oak_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			4:   2029,
			393: 7288,
			404: 7289,
			477: 7795,
			573: 7795,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=east,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3912,
			404: 3913,
			477: 4416,
			573: 4416,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10237,
			573: 10237,
		},
	},
	{
		"minecraft:dark_oak_fence_gate[facing=west,in_wall=false,open=false,powered=false]",
		nil,
		NewMapping{
			4:   2977,
			393: 7508,
			404: 7509,
			477: 8033,
			573: 8033,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=north,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			477: 4230,
			573: 4230,
			393: 3726,
			404: 3727,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=9,south=none,west=up]",
		nil,
		NewMapping{
			573: 2287,
			393: 1983,
			404: 1984,
			477: 2287,
		},
	},
	{
		"minecraft:nether_brick_fence[east=true,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 4497,
			404: 4498,
			477: 5001,
			573: 5001,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=1,south=up,west=up]",
		nil,
		NewMapping{
			393: 2625,
			404: 2626,
			477: 2929,
			573: 2929,
		},
	},
	{
		"minecraft:heavy_weighted_pressure_plate[power=12]",
		nil,
		NewMapping{
			4:   2380,
			393: 5631,
			404: 5632,
			477: 6138,
			573: 6138,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=true,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10505,
			573: 10505,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=east,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9195,
			573: 9195,
		},
	},
	{
		"minecraft:orange_banner[rotation=12]",
		nil,
		NewMapping{
			477: 7389,
			573: 7389,
			393: 6882,
			404: 6883,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=west,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3764,
			404: 3765,
			477: 4268,
			573: 4268,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=5,south=up,west=none]",
		nil,
		NewMapping{
			573: 2967,
			393: 2663,
			404: 2664,
			477: 2967,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10412,
			573: 10412,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10445,
			573: 10445,
		},
	},
	{
		"minecraft:light_gray_banner[rotation=10]",
		nil,
		NewMapping{
			393: 6992,
			404: 6993,
			477: 7499,
			573: 7499,
		},
	},
	{
		"minecraft:fire[age=7,east=true,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1373,
			404: 1374,
			477: 1677,
			573: 1677,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9235,
			573: 9235,
		},
	},
	{
		"minecraft:daylight_detector[inverted=false,power=12]",
		nil,
		NewMapping{
			4:   2428,
			393: 5679,
			404: 5680,
			477: 6186,
			573: 6186,
		},
	},
	{
		"minecraft:cocoa[age=2,facing=west]",
		nil,
		NewMapping{
			4:   2041,
			393: 4648,
			404: 4649,
			477: 5152,
			573: 5152,
		},
	},
	{
		"minecraft:cake[bites=5]",
		nil,
		NewMapping{
			477: 4015,
			573: 4015,
			4:   1477,
			393: 3511,
			404: 3512,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9204,
			573: 9204,
		},
	},
	{
		"minecraft:tripwire_hook[attached=false,facing=east,powered=true]",
		nil,
		NewMapping{
			4:   2107,
			393: 4753,
			404: 4754,
			477: 5257,
			573: 5257,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=north,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			404: 3724,
			477: 4227,
			573: 4227,
			393: 3723,
		},
	},
	{
		"minecraft:iron_bars[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 4208,
			404: 4209,
			477: 4712,
			573: 4712,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6127,
			404: 6128,
			477: 6634,
			573: 6634,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=12,powered=false]",
		nil,
		NewMapping{
			573: 423,
			393: 423,
			404: 423,
			477: 423,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=true,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10876,
			573: 10876,
		},
	},
	{
		"minecraft:andesite_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9965,
			573: 9965,
		},
	},
	{
		"minecraft:water[level=11]",
		nil,
		NewMapping{
			4:   139,
			393: 45,
			404: 45,
			477: 45,
			573: 45,
		},
	},
	{
		"minecraft:black_shulker_box[facing=up]",
		nil,
		NewMapping{
			4:   3745,
			393: 8311,
			404: 8312,
			477: 8836,
			573: 8836,
		},
	},
	{
		"minecraft:acacia_button[face=floor,facing=west,powered=true]",
		nil,
		NewMapping{
			573: 5910,
			393: 5403,
			404: 5404,
			477: 5910,
		},
	},
	{
		"minecraft:kelp[age=1]",
		nil,
		NewMapping{
			393: 8410,
			404: 8411,
			477: 8935,
			573: 8935,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=3,south=up,west=none]",
		nil,
		NewMapping{
			404: 2646,
			477: 2949,
			573: 2949,
			393: 2645,
		},
	},
	{
		"minecraft:oak_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 1703,
			404: 1704,
			477: 2007,
			573: 2007,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=true,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6213,
			404: 6214,
			477: 6720,
			573: 6720,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9727,
			573: 9727,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=east,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3847,
			404: 3848,
			477: 4351,
			573: 4351,
		},
	},
	{
		"minecraft:purpur_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 8112,
			404: 8113,
			477: 8637,
			573: 8637,
		},
	},
	{
		"minecraft:lime_banner[rotation=2]",
		nil,
		NewMapping{
			393: 6936,
			404: 6937,
			477: 7443,
			573: 7443,
		},
	},
	{
		"minecraft:cyan_bed[facing=west,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 903,
			404: 903,
			477: 1203,
			573: 1203,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 6433,
			393: 5926,
			404: 5927,
			477: 6433,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=4,powered=false]",
		nil,
		NewMapping{
			393: 607,
			404: 607,
			477: 607,
			573: 607,
		},
	},
	{
		"minecraft:creeper_head[rotation=4]",
		nil,
		NewMapping{
			573: 6038,
			393: 5535,
			404: 5536,
			477: 6038,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=16,powered=true]",
		nil,
		NewMapping{
			573: 530,
			393: 530,
			404: 530,
			477: 530,
		},
	},
	{
		"minecraft:quartz_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 5776,
			404: 5777,
			477: 6283,
			573: 6283,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6026,
			404: 6027,
			477: 6533,
			573: 6533,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=10,south=up,west=up]",
		nil,
		NewMapping{
			393: 2130,
			404: 2131,
			477: 2434,
			573: 2434,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=10,south=up,west=side]",
		nil,
		NewMapping{
			477: 2867,
			573: 2867,
			393: 2563,
			404: 2564,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6151,
			404: 6152,
			477: 6658,
			573: 6658,
		},
	},
	{
		"minecraft:magenta_bed[facing=north,occupied=false,part=foot]",
		nil,
		NewMapping{
			404: 783,
			477: 1083,
			573: 1083,
			393: 783,
		},
	},
	{
		"minecraft:chest[facing=east,type=right,waterlogged=false]",
		nil,
		NewMapping{
			393: 1751,
			404: 1752,
			477: 2055,
			573: 2055,
		},
	},
	{
		"minecraft:lectern[facing=south,has_book=false,powered=true]",
		nil,
		NewMapping{
			477: 11183,
			573: 11183,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 10024,
			573: 10024,
		},
	},
	{
		"minecraft:jack_o_lantern[facing=east]",
		nil,
		NewMapping{
			4:   1459,
			393: 3505,
			404: 3506,
			477: 4009,
			573: 4009,
		},
	},
	{
		"minecraft:light_blue_shulker_box[facing=south]",
		nil,
		NewMapping{
			477: 8762,
			573: 8762,
			4:   3555,
			393: 8237,
			404: 8238,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=east,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3646,
			404: 3647,
			477: 4150,
			573: 4150,
		},
	},
	{
		"minecraft:quartz_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			393: 7337,
			404: 7338,
			477: 7856,
			573: 7856,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 7196,
			404: 7197,
			477: 7703,
			573: 7703,
		},
	},
	{
		"minecraft:campfire[facing=north,lit=true,signal_fire=false,waterlogged=true]",
		nil,
		NewMapping{
			477: 11218,
			573: 11234,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10866,
			573: 10866,
		},
	},
	{
		"minecraft:jungle_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   2180,
			393: 5105,
			404: 5106,
			477: 5609,
			573: 5609,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=false,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 6044,
			477: 6550,
			573: 6550,
			4:   2566,
			393: 6043,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=south,half=top,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 6511,
			404: 6512,
			477: 7018,
			573: 7018,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=10,south=none,west=none]",
		nil,
		NewMapping{
			393: 2714,
			404: 2715,
			477: 3018,
			573: 3018,
		},
	},
	{
		"minecraft:dark_oak_wall_sign[facing=north,waterlogged=true]",
		nil,
		NewMapping{
			477: 3773,
			573: 3773,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=15,powered=true]",
		nil,
		NewMapping{
			477: 1028,
			573: 1028,
		},
	},
	{
		"minecraft:acacia_door[facing=south,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7891,
			404: 7892,
			477: 8416,
			573: 8416,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=east,half=top,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3900,
			404: 3901,
			477: 4404,
			573: 4404,
		},
	},
	{
		"minecraft:acacia_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6355,
			404: 6356,
			477: 6862,
			573: 6862,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=false,north=true,powered=true,south=true,west=true]",
		nil,
		NewMapping{
			393: 4867,
			404: 4868,
			477: 5371,
			573: 5371,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6706,
			404: 6707,
			477: 7213,
			573: 7213,
		},
	},
	{
		"minecraft:birch_leaves[distance=5,persistent=true]",
		nil,
		NewMapping{
			393: 180,
			404: 180,
			477: 180,
			573: 180,
		},
	},
	{
		"minecraft:spruce_stairs[facing=west,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4933,
			404: 4934,
			477: 5437,
			573: 5437,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=east,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3901,
			404: 3902,
			477: 4405,
			573: 4405,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 10507,
			477: 10507,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=13,powered=true]",
		nil,
		NewMapping{
			477: 874,
			573: 874,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6014,
			404: 6015,
			477: 6521,
			573: 6521,
		},
	},
	{
		"minecraft:birch_stairs[facing=west,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 5013,
			404: 5014,
			477: 5517,
			573: 5517,
		},
	},
	{
		"minecraft:tube_coral[waterlogged=true]",
		nil,
		NewMapping{
			404: 8470,
			477: 8994,
			573: 8994,
		},
	},
	{
		"minecraft:gray_stained_glass",
		nil,
		NewMapping{
			4:   1527,
			393: 3584,
			404: 3585,
			477: 4088,
			573: 4088,
		},
	},
	{
		"minecraft:red_sandstone_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			4:   2920,
			393: 7342,
			404: 7343,
			477: 7861,
			573: 7861,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=5,south=up,west=none]",
		nil,
		NewMapping{
			393: 1943,
			404: 1944,
			477: 2247,
			573: 2247,
		},
	},
	{
		"minecraft:brick_stairs[facing=west,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 4382,
			404: 4383,
			477: 4886,
			573: 4886,
		},
	},
	{
		"minecraft:zombie_head[rotation=9]",
		nil,
		NewMapping{
			393: 5500,
			404: 5501,
			477: 6003,
			573: 6003,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=true,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10438,
			573: 10438,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10589,
			573: 10589,
		},
	},
	{
		"minecraft:sandstone_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			393: 7301,
			404: 7302,
			477: 7814,
			573: 7814,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=4,south=up,west=side]",
		nil,
		NewMapping{
			393: 2941,
			404: 2942,
			477: 3245,
			573: 3245,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=2,south=none,west=none]",
		nil,
		NewMapping{
			477: 2514,
			573: 2514,
			393: 2210,
			404: 2211,
		},
	},
	{
		"minecraft:purpur_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 8650,
			393: 8125,
			404: 8126,
			477: 8650,
		},
	},
	{
		"minecraft:diorite_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 10204,
			477: 10204,
		},
	},
	{
		"minecraft:black_glazed_terracotta[facing=south]",
		nil,
		NewMapping{
			4:   4000,
			393: 8374,
			404: 8375,
			477: 8899,
			573: 8899,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6193,
			404: 6194,
			477: 6700,
			573: 6700,
		},
	},
	{
		"minecraft:dragon_head[rotation=12]",
		nil,
		NewMapping{
			477: 6066,
			573: 6066,
			393: 5563,
			404: 5564,
		},
	},
	{
		"minecraft:wither_skeleton_skull[rotation=1]",
		nil,
		NewMapping{
			393: 5472,
			404: 5473,
			477: 5975,
			573: 5975,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=2,south=up,west=side]",
		nil,
		NewMapping{
			393: 2779,
			404: 2780,
			477: 3083,
			573: 3083,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 3197,
			404: 3198,
			477: 3661,
			573: 3661,
		},
	},
	{
		"minecraft:fire[age=9,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1452,
			404: 1453,
			477: 1756,
			573: 1756,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4653,
			404: 4654,
			477: 5157,
			573: 5157,
		},
	},
	{
		"minecraft:birch_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 5469,
			573: 5469,
			4:   2167,
			393: 4965,
			404: 4966,
		},
	},
	{
		"minecraft:birch_leaves[distance=1,persistent=false]",
		nil,
		NewMapping{
			477: 173,
			573: 173,
			4:   290,
			393: 173,
			404: 173,
		},
	},
	{
		"minecraft:oak_fence[east=true,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 3970,
			393: 3466,
			404: 3467,
			477: 3970,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=10,south=side,west=none]",
		nil,
		NewMapping{
			573: 3015,
			393: 2711,
			404: 2712,
			477: 3015,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=true,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 4129,
			404: 4130,
			477: 4633,
			573: 4633,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10951,
			573: 10951,
		},
	},
	{
		"minecraft:diorite_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 11088,
			573: 11088,
		},
	},
	{
		"minecraft:stone_brick_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			393: 7325,
			404: 7326,
			477: 7844,
			573: 7844,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=5,south=up,west=none]",
		nil,
		NewMapping{
			573: 3111,
			393: 2807,
			404: 2808,
			477: 3111,
		},
	},
	{
		"minecraft:iron_door[facing=west,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 3339,
			404: 3340,
			477: 3843,
			573: 3843,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=0,powered=false]",
		nil,
		NewMapping{
			393: 499,
			404: 499,
			477: 499,
			573: 499,
		},
	},
	{
		"minecraft:nether_portal[axis=z]",
		nil,
		NewMapping{
			4:   1442,
			393: 3497,
			404: 3498,
			477: 4001,
			573: 4001,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6080,
			404: 6081,
			477: 6587,
			573: 6587,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=2,south=none,west=side]",
		nil,
		NewMapping{
			404: 1778,
			477: 2081,
			573: 2081,
			393: 1777,
		},
	},
	{
		"minecraft:stone_button[face=wall,facing=east,powered=true]",
		nil,
		NewMapping{
			4:   1241,
			393: 3405,
			404: 3406,
			477: 3909,
			573: 3909,
		},
	},
	{
		"minecraft:furnace[facing=west,lit=true]",
		nil,
		NewMapping{
			393: 3071,
			404: 3072,
			477: 3375,
			573: 3375,
			4:   996,
		},
	},
	{
		"minecraft:red_shulker_box[facing=north]",
		nil,
		NewMapping{
			477: 8826,
			573: 8826,
			4:   3730,
			393: 8301,
			404: 8302,
		},
	},
	{
		"minecraft:gray_banner[rotation=5]",
		nil,
		NewMapping{
			393: 6971,
			404: 6972,
			477: 7478,
			573: 7478,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6664,
			404: 6665,
			477: 7171,
			573: 7171,
		},
	},
	{
		"minecraft:oak_fence[east=false,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 3486,
			404: 3487,
			477: 3990,
			573: 3990,
		},
	},
	{
		"minecraft:acacia_fence[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 7626,
			404: 7627,
			477: 8151,
			573: 8151,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=9,south=none,west=side]",
		nil,
		NewMapping{
			393: 2272,
			404: 2273,
			477: 2576,
			573: 2576,
		},
	},
	{
		"minecraft:brick_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 4889,
			573: 4889,
			393: 4385,
			404: 4386,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6087,
			404: 6088,
			477: 6594,
			573: 6594,
		},
	},
	{
		"minecraft:fire[age=0,east=false,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 1162,
			404: 1163,
			477: 1466,
			573: 1466,
		},
	},
	{
		"minecraft:glass_pane[east=false,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 4746,
			573: 4746,
			4:   1632,
			393: 4242,
			404: 4243,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=west,half=top,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			573: 4450,
			393: 3946,
			404: 3947,
			477: 4450,
		},
	},
	{
		"minecraft:yellow_wall_banner[facing=south]",
		nil,
		NewMapping{
			573: 7634,
			393: 7127,
			404: 7128,
			477: 7634,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4426,
			404: 4427,
			477: 4930,
			573: 4930,
		},
	},
	{
		"minecraft:repeater[delay=1,facing=north,locked=true,powered=false]",
		nil,
		NewMapping{
			404: 3515,
			477: 4018,
			573: 4018,
			393: 3514,
		},
	},
	{
		"minecraft:chest[facing=north,type=right,waterlogged=true]",
		nil,
		NewMapping{
			393: 1732,
			404: 1733,
			477: 2036,
			573: 2036,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=10,powered=true]",
		nil,
		NewMapping{
			393: 468,
			404: 468,
			477: 468,
			573: 468,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=3,powered=false]",
		nil,
		NewMapping{
			404: 555,
			477: 555,
			573: 555,
			393: 555,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 6674,
			477: 7180,
			573: 7180,
			393: 6673,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10136,
			573: 10136,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5214,
			404: 5215,
			477: 5718,
			573: 5718,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9334,
			573: 9334,
		},
	},
	{
		"minecraft:tripwire_hook[attached=true,facing=east,powered=true]",
		nil,
		NewMapping{
			4:   2111,
			393: 4745,
			404: 4746,
			477: 5249,
			573: 5249,
		},
	},
	{
		"minecraft:bookshelf",
		nil,
		NewMapping{
			573: 1431,
			4:   752,
			393: 1127,
			404: 1128,
			477: 1431,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=west,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3947,
			404: 3948,
			477: 4451,
			573: 4451,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6779,
			404: 6780,
			477: 7286,
			573: 7286,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6317,
			404: 6318,
			477: 6824,
			573: 6824,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=12,south=none,west=up]",
		nil,
		NewMapping{
			393: 2730,
			404: 2731,
			477: 3034,
			573: 3034,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10655,
			573: 10655,
		},
	},
	{
		"minecraft:green_banner[rotation=5]",
		nil,
		NewMapping{
			573: 7574,
			393: 7067,
			404: 7068,
			477: 7574,
		},
	},
	{
		"minecraft:spruce_button[face=wall,facing=east,powered=true]",
		nil,
		NewMapping{
			393: 5341,
			404: 5342,
			477: 5848,
			573: 5848,
		},
	},
	{
		"minecraft:red_banner[rotation=3]",
		nil,
		NewMapping{
			393: 7081,
			404: 7082,
			477: 7588,
			573: 7588,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 4052,
			404: 4053,
			477: 4556,
			573: 4556,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 6611,
			573: 6611,
			393: 6104,
			404: 6105,
		},
	},
	{
		"minecraft:acacia_wall_sign[facing=north,waterlogged=true]",
		nil,
		NewMapping{
			477: 3757,
			573: 3757,
		},
	},
	{
		"minecraft:beehive[facing=east,honey_level=2]",
		nil,
		NewMapping{
			573: 11331,
		},
	},
	{
		"minecraft:bee_nest[facing=south,honey_level=2]",
		nil,
		NewMapping{
			573: 11295,
		},
	},
	{
		"minecraft:repeater[delay=2,facing=north,locked=false,powered=true]",
		nil,
		NewMapping{
			393: 3531,
			404: 3532,
			477: 4035,
			573: 4035,
			4:   1510,
		},
	},
	{
		"minecraft:bubble_coral_wall_fan[facing=south,waterlogged=true]",
		nil,
		NewMapping{
			404: 8538,
			477: 9082,
			573: 9082,
			393: 8522,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=14,south=side,west=up]",
		nil,
		NewMapping{
			404: 2170,
			477: 2473,
			573: 2473,
			393: 2169,
		},
	},
	{
		"minecraft:oak_door[facing=east,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			477: 3624,
			573: 3624,
			393: 3160,
			404: 3161,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=2,south=side,west=none]",
		nil,
		NewMapping{
			393: 1919,
			404: 1920,
			477: 2223,
			573: 2223,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 11056,
			573: 11056,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 10699,
			477: 10699,
		},
	},
	{
		"minecraft:cyan_shulker_box[facing=up]",
		nil,
		NewMapping{
			4:   3649,
			393: 8275,
			404: 8276,
			477: 8800,
			573: 8800,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6796,
			404: 6797,
			477: 7303,
			573: 7303,
		},
	},
	{
		"minecraft:lime_banner[rotation=9]",
		nil,
		NewMapping{
			393: 6943,
			404: 6944,
			477: 7450,
			573: 7450,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9305,
			573: 9305,
		},
	},
	{
		"minecraft:bell[attachment=double_wall,facing=east]",
		nil,
		NewMapping{
			477: 11213,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=2,south=none,west=side]",
		nil,
		NewMapping{
			477: 2225,
			573: 2225,
			393: 1921,
			404: 1922,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=8,south=side,west=none]",
		nil,
		NewMapping{
			393: 2117,
			404: 2118,
			477: 2421,
			573: 2421,
		},
	},
	{
		"minecraft:fire_coral_block",
		nil,
		NewMapping{
			393: 8457,
			404: 8458,
			477: 8982,
			573: 8982,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=13,south=side,west=none]",
		nil,
		NewMapping{
			393: 2018,
			404: 2019,
			477: 2322,
			573: 2322,
		},
	},
	{
		"minecraft:spruce_fence[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 7517,
			404: 7518,
			477: 8042,
			573: 8042,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=west,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3700,
			404: 3701,
			477: 4204,
			573: 4204,
		},
	},
	{
		"minecraft:dark_oak_button[face=floor,facing=north,powered=false]",
		nil,
		NewMapping{
			477: 5931,
			573: 5931,
			393: 5424,
			404: 5425,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=4,powered=false]",
		nil,
		NewMapping{
			477: 707,
			573: 707,
			393: 707,
			404: 707,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=23,powered=false]",
		nil,
		NewMapping{
			477: 845,
			573: 845,
		},
	},
	{
		"minecraft:jungle_stairs[facing=west,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 5094,
			477: 5597,
			573: 5597,
			393: 5093,
		},
	},
	{
		"minecraft:sea_pickle[pickles=3,waterlogged=true]",
		nil,
		NewMapping{
			404: 8584,
			477: 9108,
			573: 9108,
			393: 8568,
		},
	},
	{
		"minecraft:wither_skeleton_wall_skull[facing=south]",
		nil,
		NewMapping{
			393: 5468,
			404: 5469,
			477: 5991,
			573: 5991,
		},
	},
	{
		"minecraft:white_banner[rotation=7]",
		nil,
		NewMapping{
			573: 7368,
			4:   2823,
			393: 6861,
			404: 6862,
			477: 7368,
		},
	},
	{
		"minecraft:jungle_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 5120,
			404: 5121,
			477: 5624,
			573: 5624,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4706,
			404: 4707,
			477: 5210,
			573: 5210,
		},
	},
	{
		"minecraft:glass_pane[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 4231,
			404: 4232,
			477: 4735,
			573: 4735,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=1,south=side,west=side]",
		nil,
		NewMapping{
			477: 2789,
			573: 2789,
			393: 2485,
			404: 2486,
		},
	},
	{
		"minecraft:acacia_sign[rotation=5,waterlogged=false]",
		nil,
		NewMapping{
			477: 3486,
			573: 3486,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=15,south=none,west=none]",
		nil,
		NewMapping{
			573: 2199,
			393: 1895,
			404: 1896,
			477: 2199,
		},
	},
	{
		"minecraft:jungle_button[face=floor,facing=east,powered=true]",
		nil,
		NewMapping{
			393: 5381,
			404: 5382,
			477: 5888,
			573: 5888,
		},
	},
	{
		"minecraft:composter[level=7]",
		nil,
		NewMapping{
			477: 11269,
			573: 11285,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=north,in_wall=false,open=true,powered=true]",
		nil,
		NewMapping{
			4:   3006,
			393: 7457,
			404: 7458,
			477: 7982,
			573: 7982,
		},
	},
	{
		"minecraft:spruce_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			404: 4886,
			477: 5389,
			573: 5389,
			4:   2151,
			393: 4885,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6679,
			404: 6680,
			477: 7186,
			573: 7186,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=10,south=up,west=up]",
		nil,
		NewMapping{
			393: 2418,
			404: 2419,
			477: 2722,
			573: 2722,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=1,south=none,west=side]",
		nil,
		NewMapping{
			404: 1913,
			477: 2216,
			573: 2216,
			393: 1912,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=4,south=up,west=up]",
		nil,
		NewMapping{
			393: 2364,
			404: 2365,
			477: 2668,
			573: 2668,
		},
	},
	{
		"minecraft:fire[age=9,east=false,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1446,
			404: 1447,
			477: 1750,
			573: 1750,
		},
	},
	{
		"minecraft:acacia_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6406,
			404: 6407,
			477: 6913,
			573: 6913,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=12,powered=false]",
		nil,
		NewMapping{
			573: 1023,
			477: 1023,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=north,in_wall=true,open=false,powered=true]",
		nil,
		NewMapping{
			404: 7424,
			477: 7948,
			573: 7948,
			393: 7423,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 7072,
			573: 7072,
			393: 6565,
			404: 6566,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=14,south=up,west=none]",
		nil,
		NewMapping{
			393: 2168,
			404: 2169,
			477: 2472,
			573: 2472,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=west,half=top,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3951,
			404: 3952,
			477: 4455,
			573: 4455,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=south,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			573: 4121,
			393: 3617,
			404: 3618,
			477: 4121,
		},
	},
	{
		"minecraft:infested_cracked_stone_bricks",
		nil,
		NewMapping{
			4:   1556,
			393: 3981,
			404: 3982,
			477: 4489,
			573: 4489,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 7185,
			393: 6678,
			404: 6679,
			477: 7185,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=east,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			573: 4346,
			393: 3842,
			404: 3843,
			477: 4346,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10587,
			573: 10587,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10992,
			573: 10992,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=4,south=up,west=side]",
		nil,
		NewMapping{
			393: 2077,
			404: 2078,
			477: 2381,
			573: 2381,
		},
	},
	{
		"minecraft:black_banner[rotation=12]",
		nil,
		NewMapping{
			393: 7106,
			404: 7107,
			477: 7613,
			573: 7613,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=14,south=up,west=side]",
		nil,
		NewMapping{
			573: 3047,
			393: 2743,
			404: 2744,
			477: 3047,
		},
	},
	{
		"minecraft:green_shulker_box[facing=south]",
		nil,
		NewMapping{
			404: 8298,
			477: 8822,
			573: 8822,
			4:   3715,
			393: 8297,
		},
	},
	{
		"minecraft:cyan_concrete_powder",
		nil,
		NewMapping{
			393: 8402,
			404: 8403,
			477: 8927,
			573: 8927,
			4:   4041,
		},
	},
	{
		"minecraft:fire[age=9,east=false,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			404: 1441,
			477: 1744,
			573: 1744,
			393: 1440,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=8,south=none,west=none]",
		nil,
		NewMapping{
			477: 2280,
			573: 2280,
			393: 1976,
			404: 1977,
		},
	},
	{
		"minecraft:zombie_head[rotation=3]",
		nil,
		NewMapping{
			573: 5997,
			393: 5494,
			404: 5495,
			477: 5997,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=10,south=side,west=side]",
		nil,
		NewMapping{
			573: 2582,
			393: 2278,
			404: 2279,
			477: 2582,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 11023,
			573: 11023,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4574,
			404: 4575,
			477: 5078,
			573: 5078,
		},
	},
	{
		"minecraft:birch_pressure_plate[powered=false]",
		nil,
		NewMapping{
			404: 3373,
			477: 3876,
			573: 3876,
			393: 3372,
		},
	},
	{
		"minecraft:spruce_leaves[distance=6,persistent=false]",
		nil,
		NewMapping{
			393: 169,
			404: 169,
			477: 169,
			573: 169,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=east,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 7130,
			393: 6623,
			404: 6624,
			477: 7130,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=3,powered=true]",
		nil,
		NewMapping{
			477: 954,
			573: 954,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=true,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 3992,
			404: 3993,
			477: 4496,
			573: 4496,
		},
	},
	{
		"minecraft:acacia_fence[east=false,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 8165,
			573: 8165,
			393: 7640,
			404: 7641,
		},
	},
	{
		"minecraft:dragon_head[rotation=5]",
		nil,
		NewMapping{
			393: 5556,
			404: 5557,
			477: 6059,
			573: 6059,
		},
	},
	{
		"minecraft:activator_rail[powered=true,shape=ascending_west]",
		nil,
		NewMapping{
			4:   2523,
			393: 5783,
			404: 5784,
			477: 6290,
			573: 6290,
		},
	},
	{
		"minecraft:fire[age=7,east=true,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			573: 1670,
			393: 1366,
			404: 1367,
			477: 1670,
		},
	},
	{
		"minecraft:iron_door[facing=east,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 3363,
			404: 3364,
			477: 3867,
			573: 3867,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=0,south=none,west=up]",
		nil,
		NewMapping{
			573: 2494,
			393: 2190,
			404: 2191,
			477: 2494,
		},
	},
	{
		"minecraft:jungle_stairs[facing=west,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 5598,
			573: 5598,
			393: 5094,
			404: 5095,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10414,
			573: 10414,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=east,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9436,
			573: 9436,
		},
	},
	{
		"minecraft:trapped_chest[facing=south,type=right,waterlogged=true]",
		nil,
		NewMapping{
			393: 5589,
			404: 5590,
			477: 6096,
			573: 6096,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=5,south=side,west=up]",
		nil,
		NewMapping{
			477: 2104,
			573: 2104,
			393: 1800,
			404: 1801,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=8,south=side,west=none]",
		nil,
		NewMapping{
			393: 2261,
			404: 2262,
			477: 2565,
			573: 2565,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10659,
			573: 10659,
		},
	},
	{
		"minecraft:lime_banner[rotation=6]",
		nil,
		NewMapping{
			393: 6940,
			404: 6941,
			477: 7447,
			573: 7447,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10611,
			573: 10611,
		},
	},
	{
		"minecraft:hopper[enabled=true,facing=east]",
		nil,
		NewMapping{
			573: 6196,
			4:   2469,
			393: 5689,
			404: 5690,
			477: 6196,
		},
	},
	{
		"minecraft:spruce_door[facing=east,half=upper,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			4:   3096,
			393: 7728,
			404: 7729,
			477: 8253,
			573: 8253,
		},
	},
	{
		"minecraft:yellow_banner[rotation=1]",
		nil,
		NewMapping{
			393: 6919,
			404: 6920,
			477: 7426,
			573: 7426,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=false,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6138,
			404: 6139,
			477: 6645,
			573: 6645,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9369,
			573: 9369,
		},
	},
	{
		"minecraft:stone_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9679,
			573: 9679,
		},
	},
	{
		"minecraft:sand",
		nil,
		NewMapping{
			573: 66,
			4:   192,
			393: 66,
			404: 66,
			477: 66,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6187,
			404: 6188,
			477: 6694,
			573: 6694,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6293,
			404: 6294,
			477: 6800,
			573: 6800,
		},
	},
	{
		"minecraft:birch_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 5020,
			404: 5021,
			477: 5524,
			573: 5524,
		},
	},
	{
		"minecraft:campfire[facing=north,lit=true,signal_fire=true,waterlogged=true]",
		nil,
		NewMapping{
			477: 11216,
			573: 11232,
		},
	},
	{
		"minecraft:repeater[delay=4,facing=east,locked=false,powered=true]",
		nil,
		NewMapping{
			4:   1519,
			393: 3575,
			404: 3576,
			477: 4079,
			573: 4079,
		},
	},
	{
		"minecraft:jungle_door[facing=west,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			4:   3122,
			393: 7852,
			404: 7853,
			477: 8377,
			573: 8377,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=5,south=side,west=up]",
		nil,
		NewMapping{
			393: 2088,
			404: 2089,
			477: 2392,
			573: 2392,
		},
	},
	{
		"minecraft:chest[facing=north,type=left,waterlogged=false]",
		nil,
		NewMapping{
			393: 1731,
			404: 1732,
			477: 2035,
			573: 2035,
		},
	},
	{
		"minecraft:grindstone[face=floor,facing=east]",
		nil,
		NewMapping{
			477: 11168,
			573: 11168,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=false,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5199,
			404: 5200,
			477: 5703,
			573: 5703,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=true,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 11006,
			573: 11006,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=south,in_wall=false,open=false,powered=true]",
		nil,
		NewMapping{
			4:   2968,
			393: 7435,
			404: 7436,
			477: 7960,
			573: 7960,
		},
	},
	{
		"minecraft:green_wool",
		nil,
		NewMapping{
			4:   573,
			393: 1096,
			404: 1096,
			477: 1396,
			573: 1396,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=0,south=none,west=side]",
		nil,
		NewMapping{
			404: 2336,
			477: 2639,
			573: 2639,
			393: 2335,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6302,
			404: 6303,
			477: 6809,
			573: 6809,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=19,powered=true]",
		nil,
		NewMapping{
			477: 436,
			573: 436,
			393: 436,
			404: 436,
		},
	},
	{
		"minecraft:redstone_wall_torch[facing=south,lit=true]",
		nil,
		NewMapping{
			4:   1219,
			393: 3385,
			404: 3386,
			477: 3889,
			573: 3889,
		},
	},
	{
		"minecraft:polished_granite",
		nil,
		NewMapping{
			4:   18,
			393: 3,
			404: 3,
			477: 3,
			573: 3,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=7,powered=false]",
		nil,
		NewMapping{
			393: 313,
			404: 313,
			477: 313,
			573: 313,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 7204,
			404: 7205,
			477: 7711,
			573: 7711,
		},
	},
	{
		"minecraft:fire[age=5,east=true,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 1304,
			404: 1305,
			477: 1608,
			573: 1608,
		},
	},
	{
		"minecraft:orange_bed[facing=south,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 770,
			404: 770,
			477: 1070,
			573: 1070,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5928,
			404: 5929,
			477: 6435,
			573: 6435,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 9824,
			477: 9824,
		},
	},
	{
		"minecraft:cut_red_sandstone_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			477: 7870,
			573: 7870,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=9,south=side,west=up]",
		nil,
		NewMapping{
			393: 2124,
			404: 2125,
			477: 2428,
			573: 2428,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6718,
			404: 6719,
			477: 7225,
			573: 7225,
		},
	},
	{
		"minecraft:nether_brick_fence[east=false,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 4519,
			404: 4520,
			477: 5023,
			573: 5023,
		},
	},
	{
		"minecraft:fire_coral_wall_fan[facing=north,waterlogged=false]",
		nil,
		NewMapping{
			404: 8545,
			477: 9089,
			573: 9089,
			393: 8529,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=north,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 7179,
			404: 7180,
			477: 7686,
			573: 7686,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10704,
			573: 10704,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9603,
			573: 9603,
		},
	},
	{
		"minecraft:purpur_pillar[axis=z]",
		nil,
		NewMapping{
			393: 8076,
			404: 8077,
			477: 8601,
			573: 8601,
			4:   3240,
		},
	},
	{
		"minecraft:tnt",
		nil,
		NewMapping{
			4:   737,
			393: 1126,
		},
	},
	{
		"minecraft:repeater[delay=4,facing=west,locked=false,powered=false]",
		nil,
		NewMapping{
			4:   1501,
			393: 3572,
			404: 3573,
			477: 4076,
			573: 4076,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=0,south=up,west=none]",
		nil,
		NewMapping{
			404: 2763,
			477: 3066,
			573: 3066,
			393: 2762,
		},
	},
	{
		"minecraft:chorus_plant[down=true,east=false,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			573: 8558,
			393: 8033,
			404: 8034,
			477: 8558,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=17,powered=true]",
		nil,
		NewMapping{
			477: 982,
			573: 982,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 11012,
			573: 11012,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=west,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9343,
			573: 9343,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=5,powered=false]",
		nil,
		NewMapping{
			393: 409,
			404: 409,
			477: 409,
			573: 409,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=east,half=top,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3776,
			404: 3777,
			477: 4280,
			573: 4280,
		},
	},
	{
		"minecraft:jungle_door[facing=north,half=upper,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7808,
			404: 7809,
			477: 8333,
			573: 8333,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 6530,
			393: 6023,
			404: 6024,
			477: 6530,
		},
	},
	{
		"minecraft:dark_oak_wall_sign[facing=south,waterlogged=false]",
		nil,
		NewMapping{
			477: 3776,
			573: 3776,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9496,
			573: 9496,
		},
	},
	{
		"minecraft:andesite_stairs[facing=east,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9993,
			573: 9993,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10797,
			573: 10797,
		},
	},
	{
		"minecraft:brick_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   1734,
			393: 4353,
			404: 4354,
			477: 4857,
			573: 4857,
		},
	},
	{
		"minecraft:brown_banner[rotation=4]",
		nil,
		NewMapping{
			393: 7050,
			404: 7051,
			477: 7557,
			573: 7557,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=0,south=up,west=up]",
		nil,
		NewMapping{
			477: 2920,
			573: 2920,
			393: 2616,
			404: 2617,
		},
	},
	{
		"minecraft:jungle_stairs[facing=south,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			404: 5065,
			477: 5568,
			573: 5568,
			393: 5064,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=8,powered=true]",
		nil,
		NewMapping{
			477: 814,
			573: 814,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=east,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3780,
			404: 3781,
			477: 4284,
			573: 4284,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=north,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			573: 7005,
			393: 6498,
			404: 6499,
			477: 7005,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=true,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10633,
			573: 10633,
		},
	},
	{
		"minecraft:white_banner[rotation=15]",
		nil,
		NewMapping{
			404: 6870,
			477: 7376,
			573: 7376,
			4:   2831,
			393: 6869,
		},
	},
	{
		"minecraft:jungle_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			477: 7787,
			573: 7787,
			4:   2003,
			393: 7280,
			404: 7281,
		},
	},
	{
		"minecraft:acacia_stairs[facing=east,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 6902,
			573: 6902,
			393: 6395,
			404: 6396,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=west,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			477: 4453,
			573: 4453,
			393: 3949,
			404: 3950,
		},
	},
	{
		"minecraft:birch_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 5038,
			477: 5541,
			573: 5541,
			393: 5037,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=true,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10985,
			573: 10985,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=6,powered=false]",
		nil,
		NewMapping{
			393: 511,
			404: 511,
			477: 511,
			573: 511,
		},
	},
	{
		"minecraft:oak_door[facing=north,half=upper,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			573: 3578,
			393: 3114,
			404: 3115,
			477: 3578,
		},
	},
	{
		"minecraft:carrots[age=6]",
		nil,
		NewMapping{
			404: 5294,
			477: 5800,
			573: 5800,
			4:   2262,
			393: 5293,
		},
	},
	{
		"minecraft:jungle_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5051,
			404: 5052,
			477: 5555,
			573: 5555,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			573: 4503,
			393: 3999,
			404: 4000,
			477: 4503,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=8,south=side,west=up]",
		nil,
		NewMapping{
			393: 2979,
			404: 2980,
			477: 3283,
			573: 3283,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=11,south=side,west=none]",
		nil,
		NewMapping{
			477: 2160,
			573: 2160,
			393: 1856,
			404: 1857,
		},
	},
	{
		"minecraft:quartz_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 6208,
			573: 6208,
			4:   2503,
			393: 5701,
			404: 5702,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6071,
			404: 6072,
			477: 6578,
			573: 6578,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5876,
			404: 5877,
			477: 6383,
			573: 6383,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=22,powered=false]",
		nil,
		NewMapping{
			573: 493,
			393: 493,
			404: 493,
			477: 493,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 5100,
			573: 5100,
			393: 4596,
			404: 4597,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10955,
			573: 10955,
		},
	},
	{
		"minecraft:acacia_sign[rotation=7,waterlogged=true]",
		nil,
		NewMapping{
			573: 3489,
			477: 3489,
		},
	},
	{
		"minecraft:spruce_sign[rotation=4,waterlogged=true]",
		nil,
		NewMapping{
			477: 3419,
			573: 3419,
		},
	},
	{
		"minecraft:moving_piston[facing=south,type=sticky]",
		nil,
		NewMapping{
			4:   587,
			393: 1104,
			404: 1104,
			477: 1404,
			573: 1404,
		},
	},
	{
		"minecraft:furnace[facing=east,lit=true]",
		nil,
		NewMapping{
			393: 3073,
			404: 3074,
			477: 3377,
			573: 3377,
			4:   997,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 6522,
			393: 6015,
			404: 6016,
			477: 6522,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=west,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			404: 6532,
			477: 7038,
			573: 7038,
			393: 6531,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9860,
			573: 9860,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=south,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			477: 4122,
			573: 4122,
			393: 3618,
			404: 3619,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9255,
			573: 9255,
		},
	},
	{
		"minecraft:purpur_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 8653,
			573: 8653,
			4:   3249,
			393: 8128,
			404: 8129,
		},
	},
	{
		"minecraft:brain_coral_wall_fan[facing=west,waterlogged=true]",
		nil,
		NewMapping{
			393: 8516,
			404: 8532,
			477: 9076,
			573: 9076,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=north,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3670,
			404: 3671,
			477: 4174,
			573: 4174,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=false,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5231,
			404: 5232,
			477: 5735,
			573: 5735,
		},
	},
	{
		"minecraft:vine[east=false,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			573: 4792,
			393: 4288,
			404: 4289,
			477: 4792,
		},
	},
	{
		"minecraft:moving_piston[facing=west,type=normal]",
		nil,
		NewMapping{
			4:   580,
			393: 1105,
			404: 1105,
			477: 1405,
			573: 1405,
		},
	},
	{
		"minecraft:iron_bars[east=true,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 4187,
			404: 4188,
			477: 4691,
			573: 4691,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6262,
			404: 6263,
			477: 6769,
			573: 6769,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9869,
			573: 9869,
		},
	},
	{
		"minecraft:purple_glazed_terracotta[facing=west]",
		nil,
		NewMapping{
			404: 8356,
			477: 8880,
			573: 8880,
			4:   3921,
			393: 8355,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=12,south=none,west=none]",
		nil,
		NewMapping{
			393: 2300,
			404: 2301,
			477: 2604,
			573: 2604,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=2,south=side,west=side]",
		nil,
		NewMapping{
			393: 2926,
			404: 2927,
			477: 3230,
			573: 3230,
		},
	},
	{
		"minecraft:dark_oak_door[facing=east,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7990,
			404: 7991,
			477: 8515,
			573: 8515,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 9857,
			477: 9857,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=5,powered=true]",
		nil,
		NewMapping{
			393: 308,
			404: 308,
			477: 308,
			573: 308,
		},
	},
	{
		"minecraft:oak_sign[rotation=6,waterlogged=true]",
		nil,
		NewMapping{
			404: 3088,
			477: 3391,
			573: 3391,
		},
	},
	{
		"minecraft:jungle_door[facing=east,half=upper,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			4:   3128,
			393: 7856,
			404: 7857,
			477: 8381,
			573: 8381,
		},
	},
	{
		"minecraft:prismarine_brick_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			393: 6811,
			404: 6812,
			477: 7318,
			573: 7318,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 5938,
			477: 6444,
			573: 6444,
			393: 5937,
		},
	},
	{
		"minecraft:fire[age=3,east=false,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			573: 1560,
			393: 1256,
			404: 1257,
			477: 1560,
		},
	},
	{
		"minecraft:stripped_jungle_log[axis=x]",
		nil,
		NewMapping{
			404: 96,
			477: 96,
			573: 96,
			393: 96,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=south,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3931,
			404: 3932,
			477: 4435,
			573: 4435,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=west,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 7277,
			393: 6770,
			404: 6771,
			477: 7277,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=18,powered=true]",
		nil,
		NewMapping{
			573: 634,
			393: 634,
			404: 634,
			477: 634,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10241,
			573: 10241,
		},
	},
	{
		"minecraft:stone",
		nil,
		NewMapping{
			4:   16,
			393: 1,
			404: 1,
			477: 1,
			573: 1,
		},
	},
	{
		"minecraft:lever[face=ceiling,facing=east,powered=false]",
		nil,
		NewMapping{
			393: 3300,
			404: 3301,
			477: 3804,
			573: 3804,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=13,south=side,west=none]",
		nil,
		NewMapping{
			573: 2898,
			393: 2594,
			404: 2595,
			477: 2898,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6283,
			404: 6284,
			477: 6790,
			573: 6790,
		},
	},
	{
		"minecraft:purpur_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 8151,
			404: 8152,
			477: 8676,
			573: 8676,
		},
	},
	{
		"minecraft:dark_oak_door[facing=west,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7975,
			404: 7976,
			477: 8500,
			573: 8500,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10968,
			477: 10968,
		},
	},
	{
		"minecraft:moving_piston[facing=west,type=sticky]",
		nil,
		NewMapping{
			573: 1406,
			4:   588,
			393: 1106,
			404: 1106,
			477: 1406,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=north,in_wall=false,open=false,powered=true]",
		nil,
		NewMapping{
			4:   2970,
			393: 7427,
			404: 7428,
			477: 7952,
			573: 7952,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=south,in_wall=false,open=false,powered=true]",
		nil,
		NewMapping{
			4:   2952,
			393: 7403,
			404: 7404,
			477: 7928,
			573: 7928,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=4,south=side,west=up]",
		nil,
		NewMapping{
			393: 2367,
			404: 2368,
			477: 2671,
			573: 2671,
		},
	},
	{
		"minecraft:jungle_door[facing=west,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			573: 8364,
			393: 7839,
			404: 7840,
			477: 8364,
		},
	},
	{
		"minecraft:stone_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9659,
			573: 9659,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=5,powered=true]",
		nil,
		NewMapping{
			477: 908,
			573: 908,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 9268,
			477: 9268,
		},
	},
	{
		"minecraft:oak_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 1651,
			404: 1652,
			477: 1955,
			573: 1955,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=6,south=none,west=up]",
		nil,
		NewMapping{
			404: 2677,
			477: 2980,
			573: 2980,
			393: 2676,
		},
	},
	{
		"minecraft:dark_oak_button[face=ceiling,facing=west,powered=false]",
		nil,
		NewMapping{
			393: 5444,
			404: 5445,
			477: 5951,
			573: 5951,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=21,powered=true]",
		nil,
		NewMapping{
			393: 340,
			404: 340,
			477: 340,
			573: 340,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10607,
			573: 10607,
		},
	},
	{
		"minecraft:cyan_shulker_box[facing=down]",
		nil,
		NewMapping{
			477: 8801,
			573: 8801,
			4:   3648,
			393: 8276,
			404: 8277,
		},
	},
	{
		"minecraft:oak_leaves[distance=7,persistent=false]",
		nil,
		NewMapping{
			573: 157,
			393: 157,
			404: 157,
			477: 157,
		},
	},
	{
		"minecraft:jungle_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 5086,
			404: 5087,
			477: 5590,
			573: 5590,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10933,
			573: 10933,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10914,
			573: 10914,
		},
	},
	{
		"minecraft:gray_shulker_box[facing=up]",
		nil,
		NewMapping{
			4:   3617,
			393: 8263,
			404: 8264,
			477: 8788,
			573: 8788,
		},
	},
	{
		"minecraft:dark_prismarine_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			573: 7320,
			393: 6813,
			404: 6814,
			477: 7320,
		},
	},
	{
		"minecraft:lime_bed[facing=north,occupied=true,part=head]",
		nil,
		NewMapping{
			404: 828,
			477: 1128,
			573: 1128,
			393: 828,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=4,south=none,west=up]",
		nil,
		NewMapping{
			573: 2242,
			393: 1938,
			404: 1939,
			477: 2242,
		},
	},
	{
		"minecraft:acacia_fence[east=false,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			404: 7640,
			477: 8164,
			573: 8164,
			393: 7639,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=15,south=side,west=none]",
		nil,
		NewMapping{
			573: 2772,
			393: 2468,
			404: 2469,
			477: 2772,
		},
	},
	{
		"minecraft:acacia_button[face=ceiling,facing=west,powered=true]",
		nil,
		NewMapping{
			404: 5420,
			477: 5926,
			573: 5926,
			393: 5419,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=1,powered=true]",
		nil,
		NewMapping{
			573: 600,
			393: 600,
			404: 600,
			477: 600,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=west,in_wall=false,open=true,powered=false]",
		nil,
		NewMapping{
			573: 7999,
			4:   2997,
			393: 7474,
			404: 7475,
			477: 7999,
		},
	},
	{
		"minecraft:light_blue_glazed_terracotta[facing=east]",
		nil,
		NewMapping{
			4:   3811,
			393: 8328,
			404: 8329,
			477: 8853,
			573: 8853,
		},
	},
	{
		"minecraft:vine[east=false,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			404: 4291,
			477: 4794,
			573: 4794,
			4:   1702,
			393: 4290,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=1,south=up,west=none]",
		nil,
		NewMapping{
			393: 2771,
			404: 2772,
			477: 3075,
			573: 3075,
		},
	},
	{
		"minecraft:cobblestone_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5150,
			404: 5151,
			477: 5654,
			573: 5654,
		},
	},
	{
		"minecraft:stripped_birch_wood[axis=z]",
		nil,
		NewMapping{
			404: 134,
			477: 134,
			573: 134,
			393: 134,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 3728,
			573: 3728,
			393: 3264,
			404: 3265,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=east,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9513,
			573: 9513,
		},
	},
	{
		"minecraft:fire[age=12,east=false,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			573: 1840,
			393: 1536,
			404: 1537,
			477: 1840,
		},
	},
	{
		"minecraft:jack_o_lantern[facing=south]",
		nil,
		NewMapping{
			4:   1456,
			393: 3503,
			404: 3504,
			477: 4007,
			573: 4007,
		},
	},
	{
		"minecraft:oak_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 1953,
			4:   855,
			393: 1649,
			404: 1650,
			477: 1953,
		},
	},
	{
		"minecraft:jungle_door[facing=east,half=lower,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7861,
			404: 7862,
			477: 8386,
			573: 8386,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9525,
			573: 9525,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10071,
			573: 10071,
		},
	},
	{
		"minecraft:birch_sign[rotation=8,waterlogged=false]",
		nil,
		NewMapping{
			477: 3460,
			573: 3460,
		},
	},
	{
		"minecraft:brown_bed[facing=west,occupied=true,part=foot]",
		nil,
		NewMapping{
			404: 949,
			477: 1249,
			573: 1249,
			393: 949,
		},
	},
	{
		"minecraft:vine[east=false,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 4292,
			404: 4293,
			477: 4796,
			573: 4796,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=east,half=top,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3903,
			404: 3904,
			477: 4407,
			573: 4407,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=2,south=up,west=up]",
		nil,
		NewMapping{
			404: 2779,
			477: 3082,
			573: 3082,
			393: 2778,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10588,
			573: 10588,
		},
	},
	{
		"minecraft:stone_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9649,
			573: 9649,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10450,
			573: 10450,
		},
	},
	{
		"minecraft:lever[face=wall,facing=west,powered=false]",
		nil,
		NewMapping{
			477: 3794,
			573: 3794,
			4:   1106,
			393: 3290,
			404: 3291,
		},
	},
	{
		"minecraft:command_block[conditional=true,facing=down]",
		nil,
		NewMapping{
			4:   2200,
			393: 5129,
			404: 5130,
			477: 5633,
			573: 5633,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 6621,
			393: 6114,
			404: 6115,
			477: 6621,
		},
	},
	{
		"minecraft:spruce_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4891,
			404: 4892,
			477: 5395,
			573: 5395,
		},
	},
	{
		"minecraft:tube_coral_wall_fan[facing=south,waterlogged=true]",
		nil,
		NewMapping{
			477: 9066,
			573: 9066,
			393: 8506,
			404: 8522,
		},
	},
	{
		"minecraft:red_sandstone",
		nil,
		NewMapping{
			393: 7174,
			404: 7175,
			477: 7681,
			573: 7681,
			4:   2864,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=14,powered=true]",
		nil,
		NewMapping{
			477: 626,
			573: 626,
			393: 626,
			404: 626,
		},
	},
	{
		"minecraft:fire[age=4,east=false,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			404: 1280,
			477: 1583,
			573: 1583,
			393: 1279,
		},
	},
	{
		"minecraft:oak_button[face=floor,facing=east,powered=true]",
		nil,
		NewMapping{
			393: 5309,
			404: 5310,
			477: 5816,
			573: 5816,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=false,north=false,powered=true,south=false,west=true]",
		nil,
		NewMapping{
			393: 4781,
			404: 4782,
			477: 5285,
			573: 5285,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5904,
			404: 5905,
			477: 6411,
			573: 6411,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10660,
			573: 10660,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=8,waterlogged=false]",
		nil,
		NewMapping{
			573: 3556,
			477: 3556,
		},
	},
	{
		"minecraft:water[level=12]",
		nil,
		NewMapping{
			4:   156,
			393: 46,
			404: 46,
			477: 46,
			573: 46,
		},
	},
	{
		"minecraft:oak_door[facing=south,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			477: 3600,
			573: 3600,
			4:   1029,
			393: 3136,
			404: 3137,
		},
	},
	{
		"minecraft:spruce_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 4884,
			404: 4885,
			477: 5388,
			573: 5388,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=6,south=side,west=side]",
		nil,
		NewMapping{
			393: 2962,
			404: 2963,
			477: 3266,
			573: 3266,
		},
	},
	{
		"minecraft:light_gray_bed[facing=north,occupied=false,part=head]",
		nil,
		NewMapping{
			573: 1178,
			393: 878,
			404: 878,
			477: 1178,
		},
	},
	{
		"minecraft:spruce_sign[rotation=11,waterlogged=false]",
		nil,
		NewMapping{
			477: 3434,
			573: 3434,
		},
	},
	{
		"minecraft:birch_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 5004,
			404: 5005,
			477: 5508,
			573: 5508,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6269,
			404: 6270,
			477: 6776,
			573: 6776,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9559,
			573: 9559,
		},
	},
	{
		"minecraft:water[level=13]",
		nil,
		NewMapping{
			4:   141,
			393: 47,
			404: 47,
			477: 47,
			573: 47,
		},
	},
	{
		"minecraft:birch_door[facing=west,half=upper,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			573: 8299,
			393: 7774,
			404: 7775,
			477: 8299,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 6826,
			573: 6826,
			393: 6319,
			404: 6320,
		},
	},
	{
		"minecraft:jungle_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5057,
			404: 5058,
			477: 5561,
			573: 5561,
		},
	},
	{
		"minecraft:fire[age=6,east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			477: 1659,
			573: 1659,
			393: 1355,
			404: 1356,
		},
	},
	{
		"minecraft:pink_banner[rotation=11]",
		nil,
		NewMapping{
			477: 7468,
			573: 7468,
			393: 6961,
			404: 6962,
		},
	},
	{
		"minecraft:tube_coral_wall_fan[facing=east,waterlogged=true]",
		nil,
		NewMapping{
			393: 8510,
			404: 8526,
			477: 9070,
			573: 9070,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=12,south=up,west=side]",
		nil,
		NewMapping{
			477: 2885,
			573: 2885,
			393: 2581,
			404: 2582,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 9218,
			477: 9218,
		},
	},
	{
		"minecraft:andesite_stairs[facing=east,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9995,
			573: 9995,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9411,
			573: 9411,
		},
	},
	{
		"minecraft:spruce_fence[east=false,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			4:   3008,
			393: 7548,
			404: 7549,
			477: 8073,
			573: 8073,
		},
	},
	{
		"minecraft:kelp[age=7]",
		nil,
		NewMapping{
			393: 8416,
			404: 8417,
			477: 8941,
			573: 8941,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=2,south=up,west=none]",
		nil,
		NewMapping{
			393: 2924,
			404: 2925,
			477: 3228,
			573: 3228,
		},
	},
	{
		"minecraft:dark_oak_door[facing=south,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7960,
			404: 7961,
			477: 8485,
			573: 8485,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 6458,
			393: 5951,
			404: 5952,
			477: 6458,
		},
	},
	{
		"minecraft:spruce_sign[rotation=1,waterlogged=false]",
		nil,
		NewMapping{
			573: 3414,
			477: 3414,
		},
	},
	{
		"minecraft:oak_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   850,
			393: 1679,
			404: 1680,
			477: 1983,
			573: 1983,
		},
	},
	{
		"minecraft:fire_coral[waterlogged=true]",
		nil,
		NewMapping{
			404: 8476,
			477: 9000,
			573: 9000,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10669,
			573: 10669,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=11,south=side,west=none]",
		nil,
		NewMapping{
			477: 3024,
			573: 3024,
			393: 2720,
			404: 2721,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=3,south=none,west=up]",
		nil,
		NewMapping{
			573: 3241,
			393: 2937,
			404: 2938,
			477: 3241,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10843,
			573: 10843,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9410,
			573: 9410,
		},
	},
	{
		"minecraft:repeater[delay=1,facing=west,locked=true,powered=true]",
		nil,
		NewMapping{
			393: 3521,
			404: 3522,
			477: 4025,
			573: 4025,
		},
	},
	{
		"minecraft:fire[age=4,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1275,
			404: 1276,
			477: 1579,
			573: 1579,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=10,south=none,west=side]",
		nil,
		NewMapping{
			573: 2297,
			393: 1993,
			404: 1994,
			477: 2297,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 7215,
			477: 7721,
			573: 7721,
			393: 7214,
		},
	},
	{
		"minecraft:kelp[age=25]",
		nil,
		NewMapping{
			393: 8434,
			404: 8435,
			477: 8959,
			573: 8959,
		},
	},
	{
		"minecraft:prismarine_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			404: 6802,
			477: 7308,
			573: 7308,
			393: 6801,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=4,south=side,west=up]",
		nil,
		NewMapping{
			393: 2223,
			404: 2224,
			477: 2527,
			573: 2527,
		},
	},
	{
		"minecraft:dead_brain_coral_wall_fan[facing=north,waterlogged=false]",
		nil,
		NewMapping{
			393: 8473,
			404: 8489,
			477: 9033,
			573: 9033,
		},
	},
	{
		"minecraft:dead_brain_coral_wall_fan[facing=west,waterlogged=true]",
		nil,
		NewMapping{
			393: 8476,
			404: 8492,
			477: 9036,
			573: 9036,
		},
	},
	{
		"minecraft:stone_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9672,
			573: 9672,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=8,south=side,west=up]",
		nil,
		NewMapping{
			477: 2707,
			573: 2707,
			393: 2403,
			404: 2404,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10672,
			573: 10672,
		},
	},
	{
		"minecraft:bee_nest[facing=west,honey_level=0]",
		nil,
		NewMapping{
			573: 11299,
		},
	},
	{
		"minecraft:wither_skeleton_skull[rotation=0]",
		nil,
		NewMapping{
			404: 5472,
			477: 5974,
			573: 5974,
			393: 5471,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=east,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3654,
			404: 3655,
			477: 4158,
			573: 4158,
		},
	},
	{
		"minecraft:green_bed[facing=north,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 956,
			404: 956,
			477: 1256,
			573: 1256,
		},
	},
	{
		"minecraft:acacia_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6358,
			404: 6359,
			477: 6865,
			573: 6865,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=14,south=side,west=side]",
		nil,
		NewMapping{
			393: 2746,
			404: 2747,
			477: 3050,
			573: 3050,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=false,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 5736,
			573: 5736,
			393: 5232,
			404: 5233,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=true,north=false,powered=false,south=false,west=false]",
		nil,
		NewMapping{
			477: 5306,
			573: 5306,
			393: 4802,
			404: 4803,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=15,powered=false]",
		nil,
		NewMapping{
			404: 479,
			477: 479,
			573: 479,
			393: 479,
		},
	},
	{
		"minecraft:acacia_sign[rotation=10,waterlogged=false]",
		nil,
		NewMapping{
			477: 3496,
			573: 3496,
		},
	},
	{
		"minecraft:orange_bed[facing=north,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 764,
			404: 764,
			477: 1064,
			573: 1064,
		},
	},
	{
		"minecraft:oak_sign[rotation=13,waterlogged=false]",
		nil,
		NewMapping{
			477: 3406,
			573: 3406,
			404: 3103,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10880,
			477: 10880,
		},
	},
	{
		"minecraft:spruce_door[facing=east,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			4:   3088,
			393: 7740,
			404: 7741,
			477: 8265,
			573: 8265,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=10,south=none,west=none]",
		nil,
		NewMapping{
			393: 2570,
			404: 2571,
			477: 2874,
			573: 2874,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 7215,
			404: 7216,
			477: 7722,
			573: 7722,
		},
	},
	{
		"minecraft:blue_bed[facing=south,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 929,
			404: 929,
			477: 1229,
			573: 1229,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 5707,
			393: 5203,
			404: 5204,
			477: 5707,
		},
	},
	{
		"minecraft:granite_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			477: 10304,
			573: 10304,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9446,
			573: 9446,
		},
	},
	{
		"minecraft:fire[age=7,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1371,
			404: 1372,
			477: 1675,
			573: 1675,
		},
	},
	{
		"minecraft:magenta_bed[facing=south,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 787,
			404: 787,
			477: 1087,
			573: 1087,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=20,powered=false]",
		nil,
		NewMapping{
			393: 589,
			404: 589,
			477: 589,
			573: 589,
		},
	},
	{
		"minecraft:farmland[moisture=3]",
		nil,
		NewMapping{
			4:   963,
			393: 3062,
			404: 3063,
			477: 3366,
			573: 3366,
		},
	},
	{
		"minecraft:yellow_carpet",
		nil,
		NewMapping{
			4:   2740,
			393: 6827,
			404: 6828,
			477: 7334,
			573: 7334,
		},
	},
	{
		"minecraft:spruce_button[face=ceiling,facing=north,powered=true]",
		nil,
		NewMapping{
			573: 5850,
			393: 5343,
			404: 5344,
			477: 5850,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=6,south=none,west=side]",
		nil,
		NewMapping{
			393: 2677,
			404: 2678,
			477: 2981,
			573: 2981,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6280,
			404: 6281,
			477: 6787,
			573: 6787,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 11042,
			573: 11042,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 9865,
			477: 9865,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 4688,
			477: 5191,
			573: 5191,
			393: 4687,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6223,
			404: 6224,
			477: 6730,
			573: 6730,
		},
	},
	{
		"minecraft:birch_fence[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 8103,
			393: 7578,
			404: 7579,
			477: 8103,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=false,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10425,
			573: 10425,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=24,powered=true]",
		nil,
		NewMapping{
			477: 946,
			573: 946,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=false,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10554,
			573: 10554,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=6,south=none,west=side]",
		nil,
		NewMapping{
			393: 2821,
			404: 2822,
			477: 3125,
			573: 3125,
		},
	},
	{
		"minecraft:dead_horn_coral_wall_fan[facing=north,waterlogged=true]",
		nil,
		NewMapping{
			393: 8496,
			404: 8512,
			477: 9056,
			573: 9056,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=8,powered=true]",
		nil,
		NewMapping{
			393: 314,
			404: 314,
			477: 314,
			573: 314,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6608,
			404: 6609,
			477: 7115,
			573: 7115,
		},
	},
	{
		"minecraft:birch_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 5506,
			393: 5002,
			404: 5003,
			477: 5506,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 6407,
			393: 5900,
			404: 5901,
			477: 6407,
		},
	},
	{
		"minecraft:acacia_fence[east=false,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 8156,
			573: 8156,
			393: 7631,
			404: 7632,
		},
	},
	{
		"minecraft:purple_stained_glass",
		nil,
		NewMapping{
			404: 3588,
			477: 4091,
			573: 4091,
			4:   1530,
			393: 3587,
		},
	},
	{
		"minecraft:redstone_wall_torch[facing=west,lit=true]",
		nil,
		NewMapping{
			4:   1218,
			393: 3387,
			404: 3388,
			477: 3891,
			573: 3891,
		},
	},
	{
		"minecraft:white_bed[facing=north,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 748,
			404: 748,
			477: 1048,
			573: 1048,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=12,south=side,west=side]",
		nil,
		NewMapping{
			393: 1864,
			404: 1865,
			477: 2168,
			573: 2168,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=east,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3781,
			404: 3782,
			477: 4285,
			573: 4285,
		},
	},
	{
		"minecraft:fire[age=9,east=true,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1431,
			404: 1432,
			477: 1735,
			573: 1735,
		},
	},
	{
		"minecraft:zombie_head[rotation=6]",
		nil,
		NewMapping{
			573: 6000,
			393: 5497,
			404: 5498,
			477: 6000,
		},
	},
	{
		"minecraft:lime_banner[rotation=10]",
		nil,
		NewMapping{
			477: 7451,
			573: 7451,
			393: 6944,
			404: 6945,
		},
	},
	{
		"minecraft:composter[level=0]",
		nil,
		NewMapping{
			477: 11262,
			573: 11278,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9757,
			573: 9757,
		},
	},
	{
		"minecraft:carrots[age=5]",
		nil,
		NewMapping{
			393: 5292,
			404: 5293,
			477: 5799,
			573: 5799,
			4:   2261,
		},
	},
	{
		"minecraft:observer[facing=south,powered=false]",
		nil,
		NewMapping{
			573: 8729,
			4:   3491,
			393: 8204,
			404: 8205,
			477: 8729,
		},
	},
	{
		"minecraft:cyan_bed[facing=east,occupied=false,part=foot]",
		nil,
		NewMapping{
			404: 907,
			477: 1207,
			573: 1207,
			393: 907,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=true,north=true,powered=true,south=false,west=false]",
		nil,
		NewMapping{
			393: 4854,
			404: 4855,
			477: 5358,
			573: 5358,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10042,
			573: 10042,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=true,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 10473,
			477: 10473,
		},
	},
	{
		"minecraft:iron_door[facing=north,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			477: 3809,
			573: 3809,
			393: 3305,
			404: 3306,
		},
	},
	{
		"minecraft:skeleton_skull[rotation=14]",
		nil,
		NewMapping{
			393: 5465,
			404: 5466,
			477: 5968,
			573: 5968,
		},
	},
	{
		"minecraft:chorus_plant[down=true,east=true,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 8018,
			404: 8019,
			477: 8543,
			573: 8543,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=south,half=top,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3610,
			404: 3611,
			477: 4114,
			573: 4114,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=0,waterlogged=false]",
		nil,
		NewMapping{
			573: 3540,
			477: 3540,
		},
	},
	{
		"minecraft:birch_leaves[distance=4,persistent=true]",
		nil,
		NewMapping{
			393: 178,
			404: 178,
			477: 178,
			573: 178,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9449,
			573: 9449,
		},
	},
	{
		"minecraft:light_gray_concrete_powder",
		nil,
		NewMapping{
			393: 8401,
			404: 8402,
			477: 8926,
			573: 8926,
			4:   4040,
		},
	},
	{
		"minecraft:acacia_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			4:   2028,
			393: 7282,
			404: 7283,
			477: 7789,
			573: 7789,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=11,south=side,west=side]",
		nil,
		NewMapping{
			393: 2287,
			404: 2288,
			477: 2591,
			573: 2591,
		},
	},
	{
		"minecraft:pink_bed[facing=south,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 848,
			404: 848,
			477: 1148,
			573: 1148,
		},
	},
	{
		"minecraft:gray_bed[facing=south,occupied=true,part=foot]",
		nil,
		NewMapping{
			404: 865,
			477: 1165,
			573: 1165,
			393: 865,
		},
	},
	{
		"minecraft:stripped_acacia_wood[axis=z]",
		nil,
		NewMapping{
			393: 140,
			404: 140,
			477: 140,
			573: 140,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9244,
			573: 9244,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9592,
			573: 9592,
		},
	},
	{
		"minecraft:blue_shulker_box[facing=north]",
		nil,
		NewMapping{
			4:   3682,
			393: 8283,
			404: 8284,
			477: 8808,
			573: 8808,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 6937,
			393: 6430,
			404: 6431,
			477: 6937,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6040,
			404: 6041,
			477: 6547,
			573: 6547,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=15,south=side,west=up]",
		nil,
		NewMapping{
			393: 2610,
			404: 2611,
			477: 2914,
			573: 2914,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=14,powered=true]",
		nil,
		NewMapping{
			477: 526,
			573: 526,
			393: 526,
			404: 526,
		},
	},
	{
		"minecraft:acacia_sign[rotation=10,waterlogged=true]",
		nil,
		NewMapping{
			477: 3495,
			573: 3495,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9172,
			573: 9172,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=true,north=false,powered=true,south=false,west=true]",
		nil,
		NewMapping{
			393: 4765,
			404: 4766,
			477: 5269,
			573: 5269,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=false,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			477: 4639,
			573: 4639,
			393: 4135,
			404: 4136,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=11,south=up,west=up]",
		nil,
		NewMapping{
			404: 2140,
			477: 2443,
			573: 2443,
			393: 2139,
		},
	},
	{
		"minecraft:chest[facing=east,type=right,waterlogged=true]",
		nil,
		NewMapping{
			393: 1750,
			404: 1751,
			477: 2054,
			573: 2054,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=15,south=side,west=up]",
		nil,
		NewMapping{
			393: 1890,
			404: 1891,
			477: 2194,
			573: 2194,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 9548,
			477: 9548,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=5,powered=true]",
		nil,
		NewMapping{
			477: 808,
			573: 808,
		},
	},
	{
		"minecraft:birch_button[face=floor,facing=west,powered=false]",
		nil,
		NewMapping{
			393: 5356,
			404: 5357,
			477: 5863,
			573: 5863,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=3,south=none,west=side]",
		nil,
		NewMapping{
			404: 2363,
			477: 2666,
			573: 2666,
			393: 2362,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 3231,
			404: 3232,
			477: 3695,
			573: 3695,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5193,
			404: 5194,
			477: 5697,
			573: 5697,
		},
	},
	{
		"minecraft:fire[age=14,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1595,
			404: 1596,
			477: 1899,
			573: 1899,
		},
	},
	{
		"minecraft:bell[attachment=ceiling,facing=south,powered=true]",
		nil,
		NewMapping{
			573: 11208,
		},
	},
	{
		"minecraft:fire[age=12,east=false,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1541,
			404: 1542,
			477: 1845,
			573: 1845,
		},
	},
	{
		"minecraft:oak_fence[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 3484,
			404: 3485,
			477: 3988,
			573: 3988,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4688,
			404: 4689,
			477: 5192,
			573: 5192,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10148,
			573: 10148,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10869,
			573: 10869,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=9,south=none,west=none]",
		nil,
		NewMapping{
			393: 2561,
			404: 2562,
			477: 2865,
			573: 2865,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9291,
			573: 9291,
		},
	},
	{
		"minecraft:granite_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9909,
			573: 9909,
		},
	},
	{
		"minecraft:acacia_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			4:   2020,
			393: 7284,
			404: 7285,
			477: 7791,
			573: 7791,
		},
	},
	{
		"minecraft:blue_banner[rotation=1]",
		nil,
		NewMapping{
			393: 7031,
			404: 7032,
			477: 7538,
			573: 7538,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			404: 5941,
			477: 6447,
			573: 6447,
			393: 5940,
		},
	},
	{
		"minecraft:birch_door[facing=west,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7785,
			404: 7786,
			477: 8310,
			573: 8310,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=22,powered=true]",
		nil,
		NewMapping{
			393: 692,
			404: 692,
			477: 692,
			573: 692,
		},
	},
	{
		"minecraft:dark_oak_door[facing=east,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			4:   3163,
			393: 7987,
			404: 7988,
			477: 8512,
			573: 8512,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=15,south=up,west=up]",
		nil,
		NewMapping{
			393: 2463,
			404: 2464,
			477: 2767,
			573: 2767,
		},
	},
	{
		"minecraft:oak_fence[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 3481,
			404: 3482,
			477: 3985,
			573: 3985,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=12,powered=false]",
		nil,
		NewMapping{
			393: 473,
			404: 473,
			477: 473,
			573: 473,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=west,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3954,
			404: 3955,
			477: 4458,
			573: 4458,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=13,south=side,west=side]",
		nil,
		NewMapping{
			393: 1873,
			404: 1874,
			477: 2177,
			573: 2177,
		},
	},
	{
		"minecraft:fire[age=0,east=false,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1151,
			404: 1152,
			477: 1455,
			573: 1455,
		},
	},
	{
		"minecraft:fire[age=10,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			404: 1485,
			477: 1788,
			573: 1788,
			393: 1484,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=14,south=side,west=none]",
		nil,
		NewMapping{
			477: 2619,
			573: 2619,
			393: 2315,
			404: 2316,
		},
	},
	{
		"minecraft:birch_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 5495,
			393: 4991,
			404: 4992,
			477: 5495,
		},
	},
	{
		"minecraft:oak_fence[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 3977,
			393: 3473,
			404: 3474,
			477: 3977,
		},
	},
	{
		"minecraft:end_portal_frame[eye=true,facing=west]",
		nil,
		NewMapping{
			477: 5132,
			573: 5132,
			4:   1925,
			393: 4628,
			404: 4629,
		},
	},
	{
		"minecraft:chipped_anvil[facing=west]",
		nil,
		NewMapping{
			393: 5573,
			404: 5574,
			477: 6080,
			573: 6080,
			4:   2325,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=6,south=up,west=side]",
		nil,
		NewMapping{
			393: 2383,
			404: 2384,
			477: 2687,
			573: 2687,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=9,south=none,west=side]",
		nil,
		NewMapping{
			393: 2992,
			404: 2993,
			477: 3296,
			573: 3296,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10904,
			573: 10904,
		},
	},
	{
		"minecraft:lever[face=ceiling,facing=north,powered=false]",
		nil,
		NewMapping{
			573: 3798,
			4:   1111,
			393: 3294,
			404: 3295,
			477: 3798,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			4:   1613,
			393: 4114,
			404: 4115,
			477: 4618,
			573: 4618,
		},
	},
	{
		"minecraft:structure_block[mode=corner]",
		nil,
		NewMapping{
			4:   4082,
			393: 8580,
			404: 8597,
			477: 11254,
			573: 11270,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6794,
			404: 6795,
			477: 7301,
			573: 7301,
		},
	},
	{
		"minecraft:fire[age=10,east=false,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			404: 1479,
			477: 1782,
			573: 1782,
			393: 1478,
		},
	},
	{
		"minecraft:birch_sign[rotation=4,waterlogged=true]",
		nil,
		NewMapping{
			477: 3451,
			573: 3451,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=22,powered=false]",
		nil,
		NewMapping{
			477: 993,
			573: 993,
		},
	},
	{
		"minecraft:sticky_piston[extended=false,facing=up]",
		nil,
		NewMapping{
			393: 1038,
			404: 1038,
			477: 1338,
			573: 1338,
			4:   465,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=10,powered=true]",
		nil,
		NewMapping{
			393: 568,
			404: 568,
			477: 568,
			573: 568,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=3,powered=false]",
		nil,
		NewMapping{
			404: 455,
			477: 455,
			573: 455,
			393: 455,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=true,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6085,
			404: 6086,
			477: 6592,
			573: 6592,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=3,powered=true]",
		nil,
		NewMapping{
			393: 254,
			404: 254,
			477: 254,
			573: 254,
		},
	},
	{
		"minecraft:black_wool",
		nil,
		NewMapping{
			404: 1098,
			477: 1398,
			573: 1398,
			4:   575,
			393: 1098,
		},
	},
	{
		"minecraft:spruce_leaves[distance=6,persistent=true]",
		nil,
		NewMapping{
			404: 168,
			477: 168,
			573: 168,
			393: 168,
		},
	},
	{
		"minecraft:oak_door[facing=west,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 3143,
			404: 3144,
			477: 3607,
			573: 3607,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=true,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10697,
			573: 10697,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9831,
			573: 9831,
		},
	},
	{
		"minecraft:cake[bites=4]",
		nil,
		NewMapping{
			477: 4014,
			573: 4014,
			4:   1476,
			393: 3510,
			404: 3511,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=1,south=side,west=up]",
		nil,
		NewMapping{
			393: 2340,
			404: 2341,
			477: 2644,
			573: 2644,
		},
	},
	{
		"minecraft:yellow_banner[rotation=8]",
		nil,
		NewMapping{
			393: 6926,
			404: 6927,
			477: 7433,
			573: 7433,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 5096,
			393: 4592,
			404: 4593,
			477: 5096,
		},
	},
	{
		"minecraft:purple_banner[rotation=11]",
		nil,
		NewMapping{
			393: 7025,
			404: 7026,
			477: 7532,
			573: 7532,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=12,powered=true]",
		nil,
		NewMapping{
			393: 672,
			404: 672,
			477: 672,
			573: 672,
		},
	},
	{
		"minecraft:fire[age=11,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 1488,
			404: 1489,
			477: 1792,
			573: 1792,
		},
	},
	{
		"minecraft:granite_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9920,
			573: 9920,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=22,powered=true]",
		nil,
		NewMapping{
			393: 442,
			404: 442,
			477: 442,
			573: 442,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=false,north=false,powered=false,south=true,west=true]",
		nil,
		NewMapping{
			393: 4879,
			404: 4880,
			477: 5383,
			573: 5383,
		},
	},
	{
		"minecraft:fire[age=10,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			404: 1474,
			477: 1777,
			573: 1777,
			393: 1473,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=7,south=none,west=none]",
		nil,
		NewMapping{
			404: 2832,
			477: 3135,
			573: 3135,
			393: 2831,
		},
	},
	{
		"minecraft:acacia_door[facing=north,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			404: 7884,
			477: 8408,
			573: 8408,
			393: 7883,
		},
	},
	{
		"minecraft:andesite_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9976,
			573: 9976,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9772,
			573: 9772,
		},
	},
	{
		"minecraft:dark_oak_wood[axis=z]",
		nil,
		NewMapping{
			393: 125,
			404: 125,
			477: 125,
			573: 125,
		},
	},
	{
		"minecraft:spruce_sign[rotation=8,waterlogged=true]",
		nil,
		NewMapping{
			477: 3427,
			573: 3427,
		},
	},
	{
		"minecraft:ender_chest[facing=south,waterlogged=false]",
		nil,
		NewMapping{
			477: 5238,
			573: 5238,
			4:   2083,
			393: 4734,
			404: 4735,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=false,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			4:   2561,
			393: 5883,
			404: 5884,
			477: 6390,
			573: 6390,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 6596,
			393: 6089,
			404: 6090,
			477: 6596,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9539,
			573: 9539,
		},
	},
	{
		"minecraft:acacia_leaves[distance=2,persistent=false]",
		nil,
		NewMapping{
			477: 203,
			573: 203,
			4:   2584,
			393: 203,
			404: 203,
		},
	},
	{
		"minecraft:spruce_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4889,
			404: 4890,
			477: 5393,
			573: 5393,
		},
	},
	{
		"minecraft:birch_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 5030,
			404: 5031,
			477: 5534,
			573: 5534,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 6742,
			404: 6743,
			477: 7249,
			573: 7249,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9168,
			573: 9168,
		},
	},
	{
		"minecraft:mushroom_stem[down=false,east=false,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 4171,
			404: 4172,
			477: 4675,
			573: 4675,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=3,south=side,west=up]",
		nil,
		NewMapping{
			393: 1926,
			404: 1927,
			477: 2230,
			573: 2230,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=14,south=side,west=none]",
		nil,
		NewMapping{
			393: 2747,
			404: 2748,
			477: 3051,
			573: 3051,
		},
	},
	{
		"minecraft:purpur_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 8143,
			404: 8144,
			477: 8668,
			573: 8668,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9585,
			573: 9585,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=west,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 10143,
			573: 10143,
		},
	},
	{
		"minecraft:mossy_stone_brick_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			477: 10269,
			573: 10269,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=9,south=none,west=side]",
		nil,
		NewMapping{
			573: 2144,
			393: 1840,
			404: 1841,
			477: 2144,
		},
	},
	{
		"minecraft:oak_door[facing=west,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			404: 3145,
			477: 3608,
			573: 3608,
			393: 3144,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=false,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 5227,
			477: 5730,
			573: 5730,
			393: 5226,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=11,south=side,west=side]",
		nil,
		NewMapping{
			573: 3023,
			393: 2719,
			404: 2720,
			477: 3023,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 7134,
			393: 6627,
			404: 6628,
			477: 7134,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9289,
			573: 9289,
		},
	},
	{
		"minecraft:end_portal",
		nil,
		NewMapping{
			4:   1904,
			393: 4625,
			404: 4626,
			477: 5129,
			573: 5129,
		},
	},
	{
		"minecraft:fire[age=11,east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			573: 1819,
			393: 1515,
			404: 1516,
			477: 1819,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 3206,
			404: 3207,
			477: 3670,
			573: 3670,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=7,powered=false]",
		nil,
		NewMapping{
			404: 413,
			477: 413,
			573: 413,
			393: 413,
		},
	},
	{
		"minecraft:jungle_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 5063,
			404: 5064,
			477: 5567,
			573: 5567,
		},
	},
	{
		"minecraft:fire[age=13,east=false,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			477: 1874,
			573: 1874,
			393: 1570,
			404: 1571,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 4929,
			573: 4929,
			393: 4425,
			404: 4426,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6657,
			404: 6658,
			477: 7164,
			573: 7164,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10152,
			573: 10152,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10913,
			573: 10913,
		},
	},
	{
		"minecraft:orange_wall_banner[facing=north]",
		nil,
		NewMapping{
			573: 7621,
			393: 7114,
			404: 7115,
			477: 7621,
		},
	},
	{
		"minecraft:oak_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 1648,
			404: 1649,
			477: 1952,
			573: 1952,
		},
	},
	{
		"minecraft:oak_fence[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 3989,
			393: 3485,
			404: 3486,
			477: 3989,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=east,in_wall=false,open=false,powered=false]",
		nil,
		NewMapping{
			4:   2995,
			393: 7484,
			404: 7485,
			477: 8009,
			573: 8009,
		},
	},
	{
		"minecraft:detector_rail[powered=true,shape=ascending_west]",
		nil,
		NewMapping{
			4:   459,
			393: 1019,
			404: 1019,
			477: 1319,
			573: 1319,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=1,south=up,west=none]",
		nil,
		NewMapping{
			477: 2067,
			573: 2067,
			393: 1763,
			404: 1764,
		},
	},
	{
		"minecraft:fire[age=6,east=false,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1350,
			404: 1351,
			477: 1654,
			573: 1654,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 3719,
			573: 3719,
			393: 3255,
			404: 3256,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 6792,
			393: 6285,
			404: 6286,
			477: 6792,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 11020,
			477: 11020,
		},
	},
	{
		"minecraft:mushroom_stem[down=false,east=false,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 4163,
			404: 4164,
			477: 4667,
			573: 4667,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10788,
			573: 10788,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=21,powered=true]",
		nil,
		NewMapping{
			477: 940,
			573: 940,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=south,half=top,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			404: 6518,
			477: 7024,
			573: 7024,
			4:   2681,
			393: 6517,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=west,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			477: 4449,
			573: 4449,
			393: 3945,
			404: 3946,
		},
	},
	{
		"minecraft:purpur_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 8083,
			477: 8607,
			573: 8607,
			393: 8082,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=false,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10901,
			573: 10901,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10368,
			573: 10368,
		},
	},
	{
		"minecraft:dark_oak_door[facing=east,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7992,
			404: 7993,
			477: 8517,
			573: 8517,
		},
	},
	{
		"minecraft:birch_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 5504,
			573: 5504,
			393: 5000,
			404: 5001,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5897,
			404: 5898,
			477: 6404,
			573: 6404,
		},
	},
	{
		"minecraft:quartz_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5747,
			404: 5748,
			477: 6254,
			573: 6254,
		},
	},
	{
		"minecraft:purple_banner[rotation=2]",
		nil,
		NewMapping{
			573: 7523,
			393: 7016,
			404: 7017,
			477: 7523,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9726,
			573: 9726,
		},
	},
	{
		"minecraft:orange_banner[rotation=5]",
		nil,
		NewMapping{
			393: 6875,
			404: 6876,
			477: 7382,
			573: 7382,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=2,south=side,west=up]",
		nil,
		NewMapping{
			393: 1773,
			404: 1774,
			477: 2077,
			573: 2077,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=8,south=up,west=none]",
		nil,
		NewMapping{
			393: 2690,
			404: 2691,
			477: 2994,
			573: 2994,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10552,
			477: 10552,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9477,
			573: 9477,
		},
	},
	{
		"minecraft:comparator[facing=north,mode=subtract,powered=false]",
		nil,
		NewMapping{
			4:   2406,
			393: 5638,
			404: 5639,
			477: 6145,
			573: 6145,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=7,south=up,west=side]",
		nil,
		NewMapping{
			393: 1816,
			404: 1817,
			477: 2120,
			573: 2120,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9447,
			573: 9447,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4566,
			404: 4567,
			477: 5070,
			573: 5070,
		},
	},
	{
		"minecraft:dark_oak_button[face=ceiling,facing=east,powered=false]",
		nil,
		NewMapping{
			393: 5446,
			404: 5447,
			477: 5953,
			573: 5953,
		},
	},
	{
		"minecraft:birch_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 5010,
			404: 5011,
			477: 5514,
			573: 5514,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5985,
			404: 5986,
			477: 6492,
			573: 6492,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=west,half=top,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 6532,
			404: 6533,
			477: 7039,
			573: 7039,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=14,south=side,west=side]",
		nil,
		NewMapping{
			477: 2618,
			573: 2618,
			393: 2314,
			404: 2315,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=7,south=side,west=side]",
		nil,
		NewMapping{
			404: 2108,
			477: 2411,
			573: 2411,
			393: 2107,
		},
	},
	{
		"minecraft:vine[east=false,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			4:   1699,
			393: 4294,
			404: 4295,
			477: 4798,
			573: 4798,
		},
	},
	{
		"minecraft:peony[half=lower]",
		nil,
		NewMapping{
			4:   2805,
			393: 6849,
			404: 6850,
			477: 7356,
			573: 7356,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=8,south=up,west=none]",
		nil,
		NewMapping{
			404: 1827,
			477: 2130,
			573: 2130,
			393: 1826,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=10,south=side,west=side]",
		nil,
		NewMapping{
			393: 2566,
			404: 2567,
			477: 2870,
			573: 2870,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4460,
			404: 4461,
			477: 4964,
			573: 4964,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9302,
			573: 9302,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=false,north=false,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 10586,
			477: 10586,
		},
	},
	{
		"minecraft:spruce_sign[rotation=14,waterlogged=false]",
		nil,
		NewMapping{
			477: 3440,
			573: 3440,
		},
	},
	{
		"minecraft:rail[shape=ascending_west]",
		nil,
		NewMapping{
			4:   1059,
			393: 3182,
			404: 3183,
			477: 3646,
			573: 3646,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=4,south=side,west=up]",
		nil,
		NewMapping{
			404: 1936,
			477: 2239,
			573: 2239,
			393: 1935,
		},
	},
	{
		"minecraft:oak_stairs[facing=east,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 1708,
			404: 1709,
			477: 2012,
			573: 2012,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=15,powered=false]",
		nil,
		NewMapping{
			477: 779,
			573: 779,
		},
	},
	{
		"minecraft:torch",
		nil,
		NewMapping{
			573: 1434,
			4:   805,
			393: 1130,
			404: 1131,
			477: 1434,
		},
	},
	{
		"minecraft:spruce_fence[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 8058,
			393: 7533,
			404: 7534,
			477: 8058,
		},
	},
	{
		"minecraft:piston_head[facing=west,short=true,type=normal]",
		nil,
		NewMapping{
			393: 1071,
			404: 1071,
			477: 1371,
			573: 1371,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10909,
			573: 10909,
		},
	},
	{
		"minecraft:birch_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			477: 7779,
			573: 7779,
			4:   2018,
			393: 7272,
			404: 7273,
		},
	},
	{
		"minecraft:podzol[snowy=false]",
		nil,
		NewMapping{
			4:   50,
			393: 13,
			404: 13,
			477: 13,
			573: 13,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 3262,
			404: 3263,
			477: 3726,
			573: 3726,
		},
	},
	{
		"minecraft:fire[age=15,east=true,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			573: 1929,
			393: 1625,
			404: 1626,
			477: 1929,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 10813,
			477: 10813,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 11063,
			573: 11063,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 10064,
			573: 10064,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 7243,
			404: 7244,
			477: 7750,
			573: 7750,
		},
	},
	{
		"minecraft:purple_wall_banner[facing=north]",
		nil,
		NewMapping{
			573: 7657,
			393: 7150,
			404: 7151,
			477: 7657,
		},
	},
	{
		"minecraft:nether_brick_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			477: 7850,
			573: 7850,
			393: 7331,
			404: 7332,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=0,south=side,west=side]",
		nil,
		NewMapping{
			393: 1756,
			404: 1757,
			477: 2060,
			573: 2060,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=4,south=none,west=up]",
		nil,
		NewMapping{
			404: 2947,
			477: 3250,
			573: 3250,
			393: 2946,
		},
	},
	{
		"minecraft:dark_oak_fence[east=false,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7671,
			404: 7672,
			477: 8196,
			573: 8196,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 6629,
			573: 6629,
			393: 6122,
			404: 6123,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=0,south=up,west=side]",
		nil,
		NewMapping{
			393: 2185,
			404: 2186,
			477: 2489,
			573: 2489,
		},
	},
	{
		"minecraft:jungle_sapling[stage=1]",
		nil,
		NewMapping{
			477: 28,
			573: 28,
			4:   107,
			393: 28,
			404: 28,
		},
	},
	{
		"minecraft:orange_glazed_terracotta[facing=north]",
		nil,
		NewMapping{
			4:   3778,
			393: 8317,
			404: 8318,
			477: 8842,
			573: 8842,
		},
	},
	{
		"minecraft:piston_head[facing=down,short=false,type=sticky]",
		nil,
		NewMapping{
			477: 1382,
			573: 1382,
			4:   552,
			393: 1082,
			404: 1082,
		},
	},
	{
		"minecraft:lapis_block",
		nil,
		NewMapping{
			573: 232,
			4:   352,
			393: 232,
			404: 232,
			477: 232,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=6,south=up,west=none]",
		nil,
		NewMapping{
			393: 2960,
			404: 2961,
			477: 3264,
			573: 3264,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=north,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7389,
			404: 7390,
			477: 7914,
			573: 7914,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=9,powered=true]",
		nil,
		NewMapping{
			477: 866,
			573: 866,
		},
	},
	{
		"minecraft:birch_wall_sign[facing=north,waterlogged=true]",
		nil,
		NewMapping{
			477: 3749,
			573: 3749,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5225,
			404: 5226,
			477: 5729,
			573: 5729,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9184,
			573: 9184,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9580,
			573: 9580,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=23,powered=false]",
		nil,
		NewMapping{
			477: 895,
			573: 895,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=3,powered=false]",
		nil,
		NewMapping{
			477: 955,
			573: 955,
		},
	},
	{
		"minecraft:acacia_door[facing=west,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			4:   3142,
			393: 7914,
			404: 7915,
			477: 8439,
			573: 8439,
		},
	},
	{
		"minecraft:oak_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 2009,
			573: 2009,
			393: 1705,
			404: 1706,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9397,
			573: 9397,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10477,
			573: 10477,
		},
	},
	{
		"minecraft:stone_stairs[facing=south,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9638,
			573: 9638,
		},
	},
	{
		"minecraft:stone_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9686,
			573: 9686,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 11038,
			573: 11038,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=2,south=none,west=none]",
		nil,
		NewMapping{
			4:   882,
			393: 2930,
			404: 2931,
			477: 3234,
			573: 3234,
		},
	},
	{
		"minecraft:fire[age=5,east=true,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			404: 1296,
			477: 1599,
			573: 1599,
			393: 1295,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=3,powered=true]",
		nil,
		NewMapping{
			393: 454,
			404: 454,
			477: 454,
			573: 454,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4580,
			404: 4581,
			477: 5084,
			573: 5084,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=9,powered=false]",
		nil,
		NewMapping{
			393: 567,
			404: 567,
			477: 567,
			573: 567,
		},
	},
	{
		"minecraft:spruce_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 5409,
			573: 5409,
			4:   2150,
			393: 4905,
			404: 4906,
		},
	},
	{
		"minecraft:fire[age=6,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1356,
			404: 1357,
			477: 1660,
			573: 1660,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=8,south=none,west=side]",
		nil,
		NewMapping{
			393: 2983,
			404: 2984,
			477: 3287,
			573: 3287,
		},
	},
	{
		"minecraft:nether_brick_fence[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 4509,
			404: 4510,
			477: 5013,
			573: 5013,
		},
	},
	{
		"minecraft:light_blue_wall_banner[facing=north]",
		nil,
		NewMapping{
			393: 7122,
			404: 7123,
			477: 7629,
			573: 7629,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=east,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3974,
			404: 3975,
			477: 4478,
			573: 4478,
		},
	},
	{
		"minecraft:spruce_fence[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 7537,
			404: 7538,
			477: 8062,
			573: 8062,
		},
	},
	{
		"minecraft:magenta_banner[rotation=4]",
		nil,
		NewMapping{
			477: 7397,
			573: 7397,
			393: 6890,
			404: 6891,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 6480,
			393: 5973,
			404: 5974,
			477: 6480,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=true,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			573: 4558,
			393: 4054,
			404: 4055,
			477: 4558,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=south,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3687,
			404: 3688,
			477: 4191,
			573: 4191,
		},
	},
	{
		"minecraft:dark_oak_leaves[distance=7,persistent=false]",
		nil,
		NewMapping{
			393: 227,
			404: 227,
			477: 227,
			573: 227,
		},
	},
	{
		"minecraft:granite_stairs[facing=west,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9903,
			573: 9903,
		},
	},
	{
		"minecraft:scaffolding[bottom=true,distance=0,waterlogged=false]",
		nil,
		NewMapping{
			477: 11100,
			573: 11100,
		},
	},
	{
		"minecraft:red_shulker_box[facing=down]",
		nil,
		NewMapping{
			4:   3728,
			393: 8306,
			404: 8307,
			477: 8831,
			573: 8831,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=11,south=side,west=none]",
		nil,
		NewMapping{
			477: 2592,
			573: 2592,
			393: 2288,
			404: 2289,
		},
	},
	{
		"minecraft:light_blue_banner[rotation=7]",
		nil,
		NewMapping{
			404: 6910,
			477: 7416,
			573: 7416,
			393: 6909,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=14,south=none,west=side]",
		nil,
		NewMapping{
			477: 2333,
			573: 2333,
			393: 2029,
			404: 2030,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=11,powered=false]",
		nil,
		NewMapping{
			393: 621,
			404: 621,
			477: 621,
			573: 621,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6224,
			404: 6225,
			477: 6731,
			573: 6731,
		},
	},
	{
		"minecraft:stripped_oak_log[axis=x]",
		nil,
		NewMapping{
			393: 105,
			404: 105,
			477: 105,
			573: 105,
		},
	},
	{
		"minecraft:ladder[facing=east,waterlogged=true]",
		nil,
		NewMapping{
			573: 3641,
			393: 3177,
			404: 3178,
			477: 3641,
		},
	},
	{
		"minecraft:mushroom_stem[down=false,east=false,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 4174,
			404: 4175,
			477: 4678,
			573: 4678,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=2,south=none,west=side]",
		nil,
		NewMapping{
			393: 2497,
			404: 2498,
			477: 2801,
			573: 2801,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10918,
			573: 10918,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=true,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10939,
			573: 10939,
		},
	},
	{
		"minecraft:sign[rotation=11,waterlogged=false]",
		nil,
		NewMapping{
			4:   1019,
			393: 3098,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=true,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6308,
			404: 6309,
			477: 6815,
			573: 6815,
		},
	},
	{
		"minecraft:spruce_leaves[distance=4,persistent=true]",
		nil,
		NewMapping{
			573: 164,
			393: 164,
			404: 164,
			477: 164,
		},
	},
	{
		"minecraft:oak_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 1652,
			404: 1653,
			477: 1956,
			573: 1956,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=8,south=none,west=none]",
		nil,
		NewMapping{
			393: 2552,
			404: 2553,
			477: 2856,
			573: 2856,
		},
	},
	{
		"minecraft:smooth_quartz",
		nil,
		NewMapping{
			477: 7880,
			573: 7880,
			4:   703,
			393: 7355,
			404: 7356,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9562,
			573: 9562,
		},
	},
	{
		"minecraft:spruce_sign[rotation=13,waterlogged=false]",
		nil,
		NewMapping{
			477: 3438,
			573: 3438,
		},
	},
	{
		"minecraft:lava[level=5]",
		nil,
		NewMapping{
			573: 55,
			4:   165,
			393: 55,
			404: 55,
			477: 55,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4582,
			404: 4583,
			477: 5086,
			573: 5086,
		},
	},
	{
		"minecraft:attached_pumpkin_stem[facing=west]",
		nil,
		NewMapping{
			393: 4246,
			404: 4247,
			477: 4750,
			573: 4750,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6686,
			404: 6687,
			477: 7193,
			573: 7193,
		},
	},
	{
		"minecraft:iron_door[facing=south,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			477: 3837,
			573: 3837,
			393: 3333,
			404: 3334,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 9734,
			477: 9734,
		},
	},
	{
		"minecraft:acacia_sign[rotation=6,waterlogged=true]",
		nil,
		NewMapping{
			573: 3487,
			477: 3487,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			404: 7179,
			477: 7685,
			573: 7685,
			4:   2887,
			393: 7178,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=21,powered=true]",
		nil,
		NewMapping{
			404: 540,
			477: 540,
			573: 540,
			393: 540,
		},
	},
	{
		"minecraft:cyan_bed[facing=east,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 906,
			404: 906,
			477: 1206,
			573: 1206,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10667,
			573: 10667,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10930,
			573: 10930,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10515,
			573: 10515,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10583,
			573: 10583,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=false,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10561,
			573: 10561,
		},
	},
	{
		"minecraft:detector_rail[powered=false,shape=east_west]",
		nil,
		NewMapping{
			573: 1323,
			4:   449,
			393: 1023,
			404: 1023,
			477: 1323,
		},
	},
	{
		"minecraft:birch_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 5481,
			573: 5481,
			393: 4977,
			404: 4978,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=2,south=side,west=up]",
		nil,
		NewMapping{
			477: 2653,
			573: 2653,
			393: 2349,
			404: 2350,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=4,south=none,west=up]",
		nil,
		NewMapping{
			393: 2226,
			404: 2227,
			477: 2530,
			573: 2530,
		},
	},
	{
		"minecraft:fire[age=1,east=false,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1193,
			404: 1194,
			477: 1497,
			573: 1497,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=18,powered=false]",
		nil,
		NewMapping{
			393: 685,
			404: 685,
			477: 685,
			573: 685,
		},
	},
	{
		"minecraft:dragon_wall_head[facing=east]",
		nil,
		NewMapping{
			393: 5550,
			404: 5551,
			477: 6073,
			573: 6073,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10359,
			573: 10359,
		},
	},
	{
		"minecraft:polished_granite_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			477: 10257,
			573: 10257,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=west,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			573: 9743,
			477: 9743,
		},
	},
	{
		"minecraft:iron_door[facing=west,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			4:   1142,
			393: 3348,
			404: 3349,
			477: 3852,
			573: 3852,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=8,south=side,west=none]",
		nil,
		NewMapping{
			573: 2853,
			393: 2549,
			404: 2550,
			477: 2853,
		},
	},
	{
		"minecraft:light_gray_bed[facing=east,occupied=true,part=head]",
		nil,
		NewMapping{
			573: 1188,
			393: 888,
			404: 888,
			477: 1188,
		},
	},
	{
		"minecraft:campfire[facing=south,lit=false,signal_fire=true,waterlogged=false]",
		nil,
		NewMapping{
			477: 11229,
			573: 11245,
		},
	},
	{
		"minecraft:light_gray_banner[rotation=5]",
		nil,
		NewMapping{
			393: 6987,
			404: 6988,
			477: 7494,
			573: 7494,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10934,
			573: 10934,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 7282,
			393: 6775,
			404: 6776,
			477: 7282,
		},
	},
	{
		"minecraft:scaffolding[bottom=true,distance=5,waterlogged=true]",
		nil,
		NewMapping{
			477: 11109,
			573: 11109,
		},
	},
	{
		"minecraft:cyan_glazed_terracotta[facing=east]",
		nil,
		NewMapping{
			4:   3907,
			393: 8352,
			404: 8353,
			477: 8877,
			573: 8877,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=west,in_wall=false,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7380,
			404: 7381,
			477: 7905,
			573: 7905,
			4:   2929,
		},
	},
	{
		"minecraft:fire[age=0,east=true,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1150,
			404: 1151,
			477: 1454,
			573: 1454,
		},
	},
	{
		"minecraft:acacia_pressure_plate[powered=false]",
		nil,
		NewMapping{
			393: 3376,
			404: 3377,
			477: 3880,
			573: 3880,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6047,
			404: 6048,
			477: 6554,
			573: 6554,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10879,
			573: 10879,
		},
	},
	{
		"minecraft:purpur_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 8135,
			404: 8136,
			477: 8660,
			573: 8660,
		},
	},
	{
		"minecraft:dark_oak_fence[east=true,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 8173,
			573: 8173,
			393: 7648,
			404: 7649,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6031,
			404: 6032,
			477: 6538,
			573: 6538,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=5,powered=false]",
		nil,
		NewMapping{
			404: 709,
			477: 709,
			573: 709,
			393: 709,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=15,powered=true]",
		nil,
		NewMapping{
			573: 978,
			477: 978,
		},
	},
	{
		"minecraft:piston_head[facing=down,short=false,type=normal]",
		nil,
		NewMapping{
			4:   544,
			393: 1081,
			404: 1081,
			477: 1381,
			573: 1381,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 7203,
			573: 7203,
			393: 6696,
			404: 6697,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=4,south=side,west=side]",
		nil,
		NewMapping{
			393: 2944,
			404: 2945,
			477: 3248,
			573: 3248,
		},
	},
	{
		"minecraft:acacia_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 6914,
			573: 6914,
			393: 6407,
			404: 6408,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=south,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 6581,
			404: 6582,
			477: 7088,
			573: 7088,
		},
	},
	{
		"minecraft:birch_button[face=wall,facing=south,powered=false]",
		nil,
		NewMapping{
			393: 5362,
			404: 5363,
			477: 5869,
			573: 5869,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 11008,
			573: 11008,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9405,
			573: 9405,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 7255,
			477: 7761,
			573: 7761,
			393: 7254,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=12,south=none,west=none]",
		nil,
		NewMapping{
			404: 2733,
			477: 3036,
			573: 3036,
			393: 2732,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=true,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 4021,
			404: 4022,
			477: 4525,
			573: 4525,
		},
	},
	{
		"minecraft:glass_pane[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 4744,
			573: 4744,
			393: 4240,
			404: 4241,
		},
	},
	{
		"minecraft:jungle_fence[east=false,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 8129,
			573: 8129,
			393: 7604,
			404: 7605,
		},
	},
	{
		"minecraft:bee_nest[facing=west,honey_level=1]",
		nil,
		NewMapping{
			573: 11300,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=west,in_wall=true,open=false,powered=false]",
		nil,
		NewMapping{
			404: 7473,
			477: 7997,
			573: 7997,
			393: 7472,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 9441,
			477: 9441,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9787,
			573: 9787,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=1,south=side,west=none]",
		nil,
		NewMapping{
			393: 1766,
			404: 1767,
			477: 2070,
			573: 2070,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6315,
			404: 6316,
			477: 6822,
			573: 6822,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=24,powered=false]",
		nil,
		NewMapping{
			404: 597,
			477: 597,
			573: 597,
			393: 597,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4448,
			404: 4449,
			477: 4952,
			573: 4952,
		},
	},
	{
		"minecraft:orange_bed[facing=west,occupied=true,part=foot]",
		nil,
		NewMapping{
			573: 1073,
			393: 773,
			404: 773,
			477: 1073,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=15,south=up,west=side]",
		nil,
		NewMapping{
			404: 2177,
			477: 2480,
			573: 2480,
			393: 2176,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=south,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			404: 3622,
			477: 4125,
			573: 4125,
			393: 3621,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10723,
			573: 10723,
		},
	},
	{
		"minecraft:daylight_detector[inverted=true,power=4]",
		nil,
		NewMapping{
			573: 6162,
			4:   2852,
			393: 5655,
			404: 5656,
			477: 6162,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5208,
			404: 5209,
			477: 5712,
			573: 5712,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6594,
			404: 6595,
			477: 7101,
			573: 7101,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=true,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10442,
			573: 10442,
		},
	},
	{
		"minecraft:bell[attachment=ceiling,facing=west,powered=true]",
		nil,
		NewMapping{
			573: 11210,
		},
	},
	{
		"minecraft:spruce_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4943,
			404: 4944,
			477: 5447,
			573: 5447,
		},
	},
	{
		"minecraft:fire[age=9,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1441,
			404: 1442,
			477: 1745,
			573: 1745,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=false,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 5876,
			477: 6382,
			573: 6382,
			393: 5875,
		},
	},
	{
		"minecraft:dead_horn_coral_wall_fan[facing=south,waterlogged=false]",
		nil,
		NewMapping{
			393: 8499,
			404: 8515,
			477: 9059,
			573: 9059,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=east,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9433,
			573: 9433,
		},
	},
	{
		"minecraft:rail[shape=north_west]",
		nil,
		NewMapping{
			4:   1064,
			393: 3187,
			404: 3188,
			477: 3651,
			573: 3651,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6178,
			404: 6179,
			477: 6685,
			573: 6685,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=12,south=up,west=none]",
		nil,
		NewMapping{
			393: 3014,
			404: 3015,
			477: 3318,
			573: 3318,
		},
	},
	{
		"minecraft:diorite_wall[east=false,north=true,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 11070,
			573: 11070,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=north,half=top,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			477: 7004,
			573: 7004,
			4:   2684,
			393: 6497,
			404: 6498,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5242,
			404: 5243,
			477: 5746,
			573: 5746,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 3253,
			404: 3254,
			477: 3717,
			573: 3717,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=north,half=top,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			404: 3920,
			477: 4423,
			573: 4423,
			393: 3919,
		},
	},
	{
		"minecraft:light_blue_banner[rotation=9]",
		nil,
		NewMapping{
			393: 6911,
			404: 6912,
			477: 7418,
			573: 7418,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=true,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 5752,
			573: 5752,
			393: 5248,
			404: 5249,
		},
	},
	{
		"minecraft:oak_door[facing=west,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 3141,
			404: 3142,
			477: 3605,
			573: 3605,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=false,north=false,powered=true,south=false,west=true]",
		nil,
		NewMapping{
			393: 4813,
			404: 4814,
			477: 5317,
			573: 5317,
		},
	},
	{
		"minecraft:bell[attachment=ceiling,facing=east]",
		nil,
		NewMapping{
			477: 11205,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 10865,
			477: 10865,
		},
	},
	{
		"minecraft:red_bed[facing=east,occupied=false,part=head]",
		nil,
		NewMapping{
			4:   427,
			393: 986,
			404: 986,
			477: 1286,
			573: 1286,
		},
	},
	{
		"minecraft:brown_banner[rotation=1]",
		nil,
		NewMapping{
			573: 7554,
			393: 7047,
			404: 7048,
			477: 7554,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=22,powered=false]",
		nil,
		NewMapping{
			404: 293,
			477: 293,
			573: 293,
			393: 293,
		},
	},
	{
		"minecraft:purpur_stairs[facing=south,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 8105,
			404: 8106,
			477: 8630,
			573: 8630,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 10134,
			573: 10134,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10122,
			573: 10122,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10455,
			573: 10455,
		},
	},
	{
		"minecraft:pink_shulker_box[facing=down]",
		nil,
		NewMapping{
			4:   3600,
			393: 8258,
			404: 8259,
			477: 8783,
			573: 8783,
		},
	},
	{
		"minecraft:cocoa[age=0,facing=east]",
		nil,
		NewMapping{
			4:   2035,
			393: 4641,
			404: 4642,
			477: 5145,
			573: 5145,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=6,south=side,west=side]",
		nil,
		NewMapping{
			393: 1810,
			404: 1811,
			477: 2114,
			573: 2114,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=3,south=up,west=none]",
		nil,
		NewMapping{
			404: 2790,
			477: 3093,
			573: 3093,
			393: 2789,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6265,
			404: 6266,
			477: 6772,
			573: 6772,
		},
	},
	{
		"minecraft:comparator[facing=west,mode=subtract,powered=false]",
		nil,
		NewMapping{
			4:   2389,
			393: 5646,
			404: 5647,
			477: 6153,
			573: 6153,
		},
	},
	{
		"minecraft:spruce_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			393: 7263,
			404: 7264,
			477: 7770,
			573: 7770,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10716,
			477: 10716,
		},
	},
	{
		"minecraft:dark_oak_door[facing=west,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7967,
			404: 7968,
			477: 8492,
			573: 8492,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			573: 5042,
			393: 4538,
			404: 4539,
			477: 5042,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=17,powered=false]",
		nil,
		NewMapping{
			393: 433,
			404: 433,
			477: 433,
			573: 433,
		},
	},
	{
		"minecraft:spruce_door[facing=east,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			573: 8260,
			393: 7735,
			404: 7736,
			477: 8260,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=false,north=true,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10570,
			573: 10570,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=true,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 8041,
			404: 8042,
			477: 8566,
			573: 8566,
		},
	},
	{
		"minecraft:spruce_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			477: 7772,
			573: 7772,
			393: 7265,
			404: 7266,
		},
	},
	{
		"minecraft:iron_door[facing=south,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			477: 3829,
			573: 3829,
			393: 3325,
			404: 3326,
		},
	},
	{
		"minecraft:spruce_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4931,
			404: 4932,
			477: 5435,
			573: 5435,
		},
	},
	{
		"minecraft:wither_skeleton_wall_skull[facing=east]",
		nil,
		NewMapping{
			393: 5470,
			404: 5471,
			477: 5993,
			573: 5993,
		},
	},
	{
		"minecraft:yellow_banner[rotation=10]",
		nil,
		NewMapping{
			393: 6928,
			404: 6929,
			477: 7435,
			573: 7435,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=true,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 8045,
			404: 8046,
			477: 8570,
			573: 8570,
		},
	},
	{
		"minecraft:snow[layers=2]",
		nil,
		NewMapping{
			404: 3417,
			477: 3920,
			573: 3920,
			4:   1249,
			393: 3416,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=10,south=none,west=up]",
		nil,
		NewMapping{
			393: 2424,
			404: 2425,
			477: 2728,
			573: 2728,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=19,powered=false]",
		nil,
		NewMapping{
			477: 987,
			573: 987,
		},
	},
	{
		"minecraft:ladder[facing=south,waterlogged=true]",
		nil,
		NewMapping{
			393: 3173,
			404: 3174,
			477: 3637,
			573: 3637,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=east,half=top,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3775,
			404: 3776,
			477: 4279,
			573: 4279,
		},
	},
	{
		"minecraft:fire[age=13,east=false,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1575,
			404: 1576,
			477: 1879,
			573: 1879,
		},
	},
	{
		"minecraft:gray_wall_banner[facing=north]",
		nil,
		NewMapping{
			404: 7139,
			477: 7645,
			573: 7645,
			393: 7138,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 6998,
			393: 6491,
			404: 6492,
			477: 6998,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 5207,
			477: 5710,
			573: 5710,
			393: 5206,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=15,south=none,west=up]",
		nil,
		NewMapping{
			393: 2469,
			404: 2470,
			477: 2773,
			573: 2773,
		},
	},
	{
		"minecraft:acacia_sign[rotation=14,waterlogged=true]",
		nil,
		NewMapping{
			477: 3503,
			573: 3503,
		},
	},
	{
		"minecraft:daylight_detector[inverted=true,power=7]",
		nil,
		NewMapping{
			4:   2855,
			393: 5658,
			404: 5659,
			477: 6165,
			573: 6165,
		},
	},
	{
		"minecraft:orange_banner[rotation=0]",
		nil,
		NewMapping{
			393: 6870,
			404: 6871,
			477: 7377,
			573: 7377,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=west,half=top,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3823,
			404: 3824,
			477: 4327,
			573: 4327,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=false,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			404: 8058,
			477: 8582,
			573: 8582,
			393: 8057,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=14,south=none,west=side]",
		nil,
		NewMapping{
			393: 3037,
			404: 3038,
			477: 3341,
			573: 3341,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10944,
			477: 10944,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=north,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			477: 7979,
			573: 7979,
			393: 7454,
			404: 7455,
		},
	},
	{
		"minecraft:kelp[age=15]",
		nil,
		NewMapping{
			393: 8424,
			404: 8425,
			477: 8949,
			573: 8949,
		},
	},
	{
		"minecraft:brick_stairs[facing=east,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 4392,
			404: 4393,
			477: 4896,
			573: 4896,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=south,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3749,
			404: 3750,
			477: 4253,
			573: 4253,
		},
	},
	{
		"minecraft:smooth_stone_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			404: 7298,
			477: 7810,
			573: 7810,
		},
	},
	{
		"minecraft:jungle_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			393: 7276,
			404: 7277,
			477: 7783,
			573: 7783,
			4:   2027,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			404: 7248,
			477: 7754,
			573: 7754,
			393: 7247,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=14,powered=true]",
		nil,
		NewMapping{
			477: 826,
			573: 826,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=17,powered=true]",
		nil,
		NewMapping{
			393: 432,
			404: 432,
			477: 432,
			573: 432,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10893,
			573: 10893,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10527,
			573: 10527,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10471,
			573: 10471,
		},
	},
	{
		"minecraft:white_wall_banner[facing=north]",
		nil,
		NewMapping{
			4:   2834,
			393: 7110,
			404: 7111,
			477: 7617,
			573: 7617,
		},
	},
	{
		"minecraft:blue_glazed_terracotta[facing=east]",
		nil,
		NewMapping{
			4:   3939,
			393: 8360,
			404: 8361,
			477: 8885,
			573: 8885,
		},
	},
	{
		"minecraft:fire_coral_fan[waterlogged=true]",
		nil,
		NewMapping{
			393: 8560,
			404: 8576,
			477: 9020,
			573: 9020,
		},
	},
	{
		"minecraft:iron_door[facing=east,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 3362,
			404: 3363,
			477: 3866,
			573: 3866,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10703,
			573: 10703,
		},
	},
	{
		"minecraft:chain_command_block[conditional=true,facing=north]",
		nil,
		NewMapping{
			4:   3386,
			393: 8176,
			404: 8177,
			477: 8701,
			573: 8701,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=10,south=up,west=none]",
		nil,
		NewMapping{
			393: 2276,
			404: 2277,
			477: 2580,
			573: 2580,
		},
	},
	{
		"minecraft:oak_button[face=ceiling,facing=east,powered=false]",
		nil,
		NewMapping{
			573: 5833,
			393: 5326,
			404: 5327,
			477: 5833,
		},
	},
	{
		"minecraft:piston_head[facing=west,short=true,type=sticky]",
		nil,
		NewMapping{
			573: 1372,
			393: 1072,
			404: 1072,
			477: 1372,
		},
	},
	{
		"minecraft:light_gray_wall_banner[facing=south]",
		nil,
		NewMapping{
			393: 7143,
			404: 7144,
			477: 7650,
			573: 7650,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9368,
			573: 9368,
		},
	},
	{
		"minecraft:purple_bed[facing=north,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 909,
			404: 909,
			477: 1209,
			573: 1209,
		},
	},
	{
		"minecraft:light_blue_bed[facing=east,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 810,
			404: 810,
			477: 1110,
			573: 1110,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=east,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4719,
			404: 4720,
			477: 5223,
			573: 5223,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=3,powered=true]",
		nil,
		NewMapping{
			393: 704,
			404: 704,
			477: 704,
			573: 704,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=8,south=up,west=none]",
		nil,
		NewMapping{
			404: 2547,
			477: 2850,
			573: 2850,
			393: 2546,
		},
	},
	{
		"minecraft:fire[age=9,east=false,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1449,
			404: 1450,
			477: 1753,
			573: 1753,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=0,powered=true]",
		nil,
		NewMapping{
			573: 798,
			477: 798,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=12,powered=false]",
		nil,
		NewMapping{
			393: 673,
			404: 673,
			477: 673,
			573: 673,
		},
	},
	{
		"minecraft:fire[age=10,east=false,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			477: 1786,
			573: 1786,
			393: 1482,
			404: 1483,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=12,south=side,west=side]",
		nil,
		NewMapping{
			393: 2872,
			404: 2873,
			477: 3176,
			573: 3176,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 7710,
			573: 7710,
			393: 7203,
			404: 7204,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			404: 5917,
			477: 6423,
			573: 6423,
			393: 5916,
		},
	},
	{
		"minecraft:bell[attachment=single_wall,facing=south,powered=true]",
		nil,
		NewMapping{
			573: 11216,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=12,south=none,west=up]",
		nil,
		NewMapping{
			393: 1866,
			404: 1867,
			477: 2170,
			573: 2170,
		},
	},
	{
		"minecraft:spruce_button[face=wall,facing=south,powered=true]",
		nil,
		NewMapping{
			477: 5844,
			573: 5844,
			393: 5337,
			404: 5338,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10828,
			573: 10828,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 11057,
			573: 11057,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10468,
			573: 10468,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=south,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			404: 3880,
			477: 4383,
			573: 4383,
			393: 3879,
		},
	},
	{
		"minecraft:trapped_chest[facing=east,type=single,waterlogged=true]",
		nil,
		NewMapping{
			393: 5597,
			404: 5598,
			477: 6104,
			573: 6104,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=14,powered=true]",
		nil,
		NewMapping{
			393: 476,
			404: 476,
			477: 476,
			573: 476,
		},
	},
	{
		"minecraft:dead_bubble_coral[waterlogged=true]",
		nil,
		NewMapping{
			573: 8988,
			404: 8464,
			477: 8988,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=15,powered=true]",
		nil,
		NewMapping{
			477: 778,
			573: 778,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6126,
			404: 6127,
			477: 6633,
			573: 6633,
		},
	},
	{
		"minecraft:spruce_fence[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 8063,
			573: 8063,
			393: 7538,
			404: 7539,
		},
	},
	{
		"minecraft:jungle_door[facing=east,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7862,
			404: 7863,
			477: 8387,
			573: 8387,
		},
	},
	{
		"minecraft:quartz_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   2497,
			393: 5751,
			404: 5752,
			477: 6258,
			573: 6258,
		},
	},
	{
		"minecraft:cocoa[age=2,facing=north]",
		nil,
		NewMapping{
			477: 5150,
			573: 5150,
			4:   2042,
			393: 4646,
			404: 4647,
		},
	},
	{
		"minecraft:daylight_detector[inverted=true,power=10]",
		nil,
		NewMapping{
			4:   2858,
			393: 5661,
			404: 5662,
			477: 6168,
			573: 6168,
		},
	},
	{
		"minecraft:brewing_stand[has_bottle_0=true,has_bottle_1=false,has_bottle_2=false]",
		nil,
		NewMapping{
			4:   1873,
			393: 4616,
			404: 4617,
			477: 5120,
			573: 5120,
		},
	},
	{
		"minecraft:fire[age=5,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			573: 1617,
			393: 1313,
			404: 1314,
			477: 1617,
		},
	},
	{
		"minecraft:acacia_door[facing=north,half=upper,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7869,
			404: 7870,
			477: 8394,
			573: 8394,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=5,waterlogged=true]",
		nil,
		NewMapping{
			477: 3549,
			573: 3549,
		},
	},
	{
		"minecraft:diorite_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 10174,
			573: 10174,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 3196,
			404: 3197,
			477: 3660,
			573: 3660,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=10,south=side,west=up]",
		nil,
		NewMapping{
			477: 2149,
			573: 2149,
			393: 1845,
			404: 1846,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4601,
			404: 4602,
			477: 5105,
			573: 5105,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=true,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10494,
			573: 10494,
		},
	},
	{
		"minecraft:bell[attachment=double_wall,facing=south]",
		nil,
		NewMapping{
			477: 11211,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=12,south=side,west=side]",
		nil,
		NewMapping{
			573: 2312,
			393: 2008,
			404: 2009,
			477: 2312,
		},
	},
	{
		"minecraft:diorite_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			573: 10183,
			477: 10183,
		},
	},
	{
		"minecraft:loom[facing=north]",
		nil,
		NewMapping{
			477: 11131,
			573: 11131,
		},
	},
	{
		"minecraft:end_rod[facing=north]",
		nil,
		NewMapping{
			4:   3170,
			393: 7997,
			404: 7998,
			477: 8522,
			573: 8522,
		},
	},
	{
		"minecraft:acacia_planks",
		nil,
		NewMapping{
			4:   84,
			393: 19,
			404: 19,
			477: 19,
			573: 19,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6186,
			404: 6187,
			477: 6693,
			573: 6693,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=9,powered=false]",
		nil,
		NewMapping{
			404: 717,
			477: 717,
			573: 717,
			393: 717,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=7,south=up,west=none]",
		nil,
		NewMapping{
			404: 2250,
			477: 2553,
			573: 2553,
			393: 2249,
		},
	},
	{
		"minecraft:acacia_door[facing=north,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			477: 8396,
			573: 8396,
			393: 7871,
			404: 7872,
		},
	},
	{
		"minecraft:sea_pickle[pickles=4,waterlogged=true]",
		nil,
		NewMapping{
			573: 9110,
			393: 8570,
			404: 8586,
			477: 9110,
		},
	},
	{
		"minecraft:glass_pane[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 4228,
			404: 4229,
			477: 4732,
			573: 4732,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=13,south=none,west=none]",
		nil,
		NewMapping{
			393: 2309,
			404: 2310,
			477: 2613,
			573: 2613,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=5,south=up,west=up]",
		nil,
		NewMapping{
			393: 2949,
			404: 2950,
			477: 3253,
			573: 3253,
		},
	},
	{
		"minecraft:stone_button[face=ceiling,facing=west,powered=false]",
		nil,
		NewMapping{
			404: 3413,
			477: 3916,
			573: 3916,
			393: 3412,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=east,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3709,
			404: 3710,
			477: 4213,
			573: 4213,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=true,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 4091,
			404: 4092,
			477: 4595,
			573: 4595,
		},
	},
	{
		"minecraft:stone_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9632,
			573: 9632,
		},
	},
	{
		"minecraft:birch_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4987,
			404: 4988,
			477: 5491,
			573: 5491,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9590,
			477: 9590,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=false,north=false,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10582,
			573: 10582,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=false,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			4:   2563,
			393: 5947,
			404: 5948,
			477: 6454,
			573: 6454,
		},
	},
	{
		"minecraft:dark_oak_log[axis=y]",
		nil,
		NewMapping{
			4:   2593,
			393: 88,
			404: 88,
			477: 88,
			573: 88,
		},
	},
	{
		"minecraft:iron_door[facing=south,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			4:   1141,
			393: 3332,
			404: 3333,
			477: 3836,
			573: 3836,
		},
	},
	{
		"minecraft:brown_bed[facing=east,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 955,
			404: 955,
			477: 1255,
			573: 1255,
		},
	},
	{
		"minecraft:acacia_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6351,
			404: 6352,
			477: 6858,
			573: 6858,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10047,
			573: 10047,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10497,
			573: 10497,
		},
	},
	{
		"minecraft:jungle_fence[east=true,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 8114,
			393: 7589,
			404: 7590,
			477: 8114,
		},
	},
	{
		"minecraft:tube_coral_wall_fan[facing=east,waterlogged=false]",
		nil,
		NewMapping{
			393: 8511,
			404: 8527,
			477: 9071,
			573: 9071,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=false,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 4071,
			404: 4072,
			477: 4575,
			573: 4575,
		},
	},
	{
		"minecraft:birch_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 4987,
			477: 5490,
			573: 5490,
			393: 4986,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10470,
			573: 10470,
		},
	},
	{
		"minecraft:lever[face=wall,facing=east,powered=false]",
		nil,
		NewMapping{
			573: 3796,
			4:   1105,
			393: 3292,
			404: 3293,
			477: 3796,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=east,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7477,
			404: 7478,
			477: 8002,
			573: 8002,
		},
	},
	{
		"minecraft:zombie_head[rotation=0]",
		nil,
		NewMapping{
			404: 5492,
			477: 5994,
			573: 5994,
			393: 5491,
		},
	},
	{
		"minecraft:lever[face=floor,facing=east,powered=false]",
		nil,
		NewMapping{
			393: 3284,
			404: 3285,
			477: 3788,
			573: 3788,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=12,south=up,west=side]",
		nil,
		NewMapping{
			393: 1861,
			404: 1862,
			477: 2165,
			573: 2165,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=17,powered=false]",
		nil,
		NewMapping{
			573: 333,
			393: 333,
			404: 333,
			477: 333,
		},
	},
	{
		"minecraft:acacia_sign[rotation=14,waterlogged=false]",
		nil,
		NewMapping{
			477: 3504,
			573: 3504,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=false,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			4:   2225,
			393: 5260,
			404: 5261,
			477: 5764,
			573: 5764,
		},
	},
	{
		"minecraft:pink_terracotta",
		nil,
		NewMapping{
			393: 5810,
			404: 5811,
			477: 6317,
			573: 6317,
			4:   2550,
		},
	},
	{
		"minecraft:blue_banner[rotation=6]",
		nil,
		NewMapping{
			393: 7036,
			404: 7037,
			477: 7543,
			573: 7543,
		},
	},
	{
		"minecraft:jungle_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			393: 7279,
			404: 7280,
			477: 7786,
			573: 7786,
		},
	},
	{
		"minecraft:spruce_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4950,
			404: 4951,
			477: 5454,
			573: 5454,
		},
	},
	{
		"minecraft:jungle_door[facing=north,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7819,
			404: 7820,
			477: 8344,
			573: 8344,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10115,
			573: 10115,
		},
	},
	{
		"minecraft:granite_stairs[facing=south,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9878,
			573: 9878,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 6935,
			573: 6935,
			393: 6428,
			404: 6429,
		},
	},
	{
		"minecraft:nether_brick_fence[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 5002,
			393: 4498,
			404: 4499,
			477: 5002,
		},
	},
	{
		"minecraft:light_gray_bed[facing=south,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 883,
			404: 883,
			477: 1183,
			573: 1183,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=9,south=side,west=none]",
		nil,
		NewMapping{
			404: 2559,
			477: 2862,
			573: 2862,
			393: 2558,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6033,
			404: 6034,
			477: 6540,
			573: 6540,
		},
	},
	{
		"minecraft:carved_pumpkin[facing=north]",
		nil,
		NewMapping{
			4:   1378,
			393: 3498,
			404: 3499,
			477: 4002,
			573: 4002,
		},
	},
	{
		"minecraft:detector_rail[powered=false,shape=ascending_south]",
		nil,
		NewMapping{
			477: 1327,
			573: 1327,
			4:   453,
			393: 1027,
			404: 1027,
		},
	},
	{
		"minecraft:oak_button[face=wall,facing=north,powered=true]",
		nil,
		NewMapping{
			477: 5818,
			573: 5818,
			4:   2300,
			393: 5311,
			404: 5312,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 6106,
			477: 6612,
			573: 6612,
			393: 6105,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=true,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 3993,
			404: 3994,
			477: 4497,
			573: 4497,
		},
	},
	{
		"minecraft:moving_piston[facing=up,type=normal]",
		nil,
		NewMapping{
			4:   577,
			393: 1107,
			404: 1107,
			477: 1407,
			573: 1407,
		},
	},
	{
		"minecraft:wheat[age=6]",
		nil,
		NewMapping{
			4:   950,
			393: 3057,
			404: 3058,
			477: 3361,
			573: 3361,
		},
	},
	{
		"minecraft:light_gray_banner[rotation=8]",
		nil,
		NewMapping{
			573: 7497,
			393: 6990,
			404: 6991,
			477: 7497,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=7,south=side,west=side]",
		nil,
		NewMapping{
			393: 1963,
			404: 1964,
			477: 2267,
			573: 2267,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10836,
			573: 10836,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10688,
			573: 10688,
		},
	},
	{
		"minecraft:andesite_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10008,
			573: 10008,
		},
	},
	{
		"minecraft:barrel[facing=west,open=false]",
		nil,
		NewMapping{
			477: 11142,
			573: 11142,
		},
	},
	{
		"minecraft:smooth_sandstone",
		nil,
		NewMapping{
			4:   697,
			393: 7354,
			404: 7355,
			477: 7879,
			573: 7879,
		},
	},
	{
		"minecraft:purpur_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			404: 8099,
			477: 8623,
			573: 8623,
			4:   3254,
			393: 8098,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 7728,
			573: 7728,
			393: 7221,
			404: 7222,
		},
	},
	{
		"minecraft:purpur_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			393: 7351,
			404: 7352,
			477: 7876,
			573: 7876,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=2,south=none,west=none]",
		nil,
		NewMapping{
			393: 2786,
			404: 2787,
			477: 3090,
			573: 3090,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=9,south=up,west=none]",
		nil,
		NewMapping{
			393: 1979,
			404: 1980,
			477: 2283,
			573: 2283,
		},
	},
	{
		"minecraft:purple_bed[facing=east,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 921,
			404: 921,
			477: 1221,
			573: 1221,
		},
	},
	{
		"minecraft:stone_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9627,
			573: 9627,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=east,in_wall=false,open=true,powered=true]",
		nil,
		NewMapping{
			404: 7482,
			477: 8006,
			573: 8006,
			4:   3007,
			393: 7481,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=9,powered=false]",
		nil,
		NewMapping{
			393: 367,
			404: 367,
			477: 367,
			573: 367,
		},
	},
	{
		"minecraft:chest[facing=south,type=right,waterlogged=true]",
		nil,
		NewMapping{
			393: 1738,
			404: 1739,
			477: 2042,
			573: 2042,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=14,powered=false]",
		nil,
		NewMapping{
			393: 527,
			404: 527,
			477: 527,
			573: 527,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9575,
			573: 9575,
		},
	},
	{
		"minecraft:acacia_door[facing=west,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7906,
			404: 7907,
			477: 8431,
			573: 8431,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=23,powered=true]",
		nil,
		NewMapping{
			393: 494,
			404: 494,
			477: 494,
			573: 494,
		},
	},
	{
		"minecraft:fire[age=10,east=true,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1469,
			404: 1470,
			477: 1773,
			573: 1773,
		},
	},
	{
		"minecraft:yellow_glazed_terracotta[facing=south]",
		nil,
		NewMapping{
			4:   3824,
			393: 8330,
			404: 8331,
			477: 8855,
			573: 8855,
		},
	},
	{
		"minecraft:kelp[age=11]",
		nil,
		NewMapping{
			393: 8420,
			404: 8421,
			477: 8945,
			573: 8945,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=13,south=side,west=up]",
		nil,
		NewMapping{
			393: 2592,
			404: 2593,
			477: 2896,
			573: 2896,
		},
	},
	{
		"minecraft:iron_door[facing=south,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			404: 3329,
			477: 3832,
			573: 3832,
			393: 3328,
		},
	},
	{
		"minecraft:shulker_box[facing=east]",
		nil,
		NewMapping{
			477: 8737,
			573: 8737,
			393: 8212,
			404: 8213,
		},
	},
	{
		"minecraft:polished_andesite_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			573: 10322,
			477: 10322,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9551,
			573: 9551,
		},
	},
	{
		"minecraft:sandstone_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			393: 7302,
			404: 7303,
			477: 7815,
			573: 7815,
			4:   705,
		},
	},
	{
		"minecraft:fire[age=14,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			404: 1615,
			477: 1918,
			573: 1918,
			4:   830,
			393: 1614,
		},
	},
	{
		"minecraft:trapped_chest[facing=west,type=left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5594,
			404: 5595,
			477: 6101,
			573: 6101,
		},
	},
	{
		"minecraft:purpur_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 8646,
			573: 8646,
			393: 8121,
			404: 8122,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10550,
			573: 10550,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10337,
			573: 10337,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=18,powered=false]",
		nil,
		NewMapping{
			477: 785,
			573: 785,
		},
	},
	{
		"minecraft:sugar_cane[age=12]",
		nil,
		NewMapping{
			393: 3454,
			404: 3455,
			477: 3958,
			573: 3958,
			4:   1340,
		},
	},
	{
		"minecraft:orange_banner[rotation=6]",
		nil,
		NewMapping{
			393: 6876,
			404: 6877,
			477: 7383,
			573: 7383,
		},
	},
	{
		"minecraft:acacia_door[facing=south,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			404: 7895,
			477: 8419,
			573: 8419,
			393: 7894,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=west,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			477: 4462,
			573: 4462,
			393: 3958,
			404: 3959,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=7,powered=true]",
		nil,
		NewMapping{
			477: 912,
			573: 912,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10725,
			573: 10725,
		},
	},
	{
		"minecraft:orange_bed[facing=north,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 766,
			404: 766,
			477: 1066,
			573: 1066,
		},
	},
	{
		"minecraft:fire[age=0,east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1163,
			404: 1164,
			477: 1467,
			573: 1467,
		},
	},
	{
		"minecraft:jungle_button[face=ceiling,facing=west,powered=true]",
		nil,
		NewMapping{
			393: 5395,
			404: 5396,
			477: 5902,
			573: 5902,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=14,powered=false]",
		nil,
		NewMapping{
			477: 927,
			573: 927,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9347,
			573: 9347,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=north,half=top,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			573: 4292,
			393: 3788,
			404: 3789,
			477: 4292,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=true,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 4094,
			404: 4095,
			477: 4598,
			573: 4598,
		},
	},
	{
		"minecraft:oak_sign[rotation=2,waterlogged=true]",
		nil,
		NewMapping{
			404: 3080,
			477: 3383,
			573: 3383,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=11,south=none,west=side]",
		nil,
		NewMapping{
			393: 2434,
			404: 2435,
			477: 2738,
			573: 2738,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=0,powered=true]",
		nil,
		NewMapping{
			393: 698,
			404: 698,
			477: 698,
			573: 698,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=true,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 4062,
			404: 4063,
			477: 4566,
			573: 4566,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=false,north=true,powered=false,south=false,west=true]",
		nil,
		NewMapping{
			573: 5377,
			393: 4873,
			404: 4874,
			477: 5377,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=south,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7430,
			404: 7431,
			477: 7955,
			573: 7955,
		},
	},
	{
		"minecraft:fire[age=11,east=false,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1517,
			404: 1518,
			477: 1821,
			573: 1821,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=east,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9516,
			573: 9516,
		},
	},
	{
		"minecraft:dead_bush",
		nil,
		NewMapping{
			4:   496,
			393: 1043,
			404: 1043,
			477: 1343,
			573: 1343,
		},
	},
	{
		"minecraft:birch_button[face=floor,facing=north,powered=false]",
		nil,
		NewMapping{
			393: 5352,
			404: 5353,
			477: 5859,
			573: 5859,
		},
	},
	{
		"minecraft:iron_door[facing=south,half=upper,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 3319,
			404: 3320,
			477: 3823,
			573: 3823,
		},
	},
	{
		"minecraft:green_shulker_box[facing=down]",
		nil,
		NewMapping{
			4:   3712,
			393: 8300,
			404: 8301,
			477: 8825,
			573: 8825,
		},
	},
	{
		"minecraft:spruce_leaves[distance=2,persistent=false]",
		nil,
		NewMapping{
			404: 161,
			477: 161,
			573: 161,
			4:   297,
			393: 161,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 6668,
			393: 6161,
			404: 6162,
			477: 6668,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=24,powered=true]",
		nil,
		NewMapping{
			393: 296,
			404: 296,
			477: 296,
			573: 296,
		},
	},
	{
		"minecraft:dried_kelp_block",
		nil,
		NewMapping{
			393: 8436,
			404: 8437,
			477: 8961,
			573: 8961,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=4,south=side,west=up]",
		nil,
		NewMapping{
			477: 2815,
			573: 2815,
			393: 2511,
			404: 2512,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=13,south=up,west=up]",
		nil,
		NewMapping{
			477: 2173,
			573: 2173,
			393: 1869,
			404: 1870,
		},
	},
	{
		"minecraft:spruce_button[face=wall,facing=west,powered=false]",
		nil,
		NewMapping{
			393: 5340,
			404: 5341,
			477: 5847,
			573: 5847,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6477,
			404: 6478,
			477: 6984,
			573: 6984,
		},
	},
	{
		"minecraft:spruce_door[facing=south,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			4:   3093,
			393: 7706,
			404: 7707,
			477: 8231,
			573: 8231,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=north,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3853,
			404: 3854,
			477: 4357,
			573: 4357,
		},
	},
	{
		"minecraft:jungle_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			393: 7275,
			404: 7276,
			477: 7782,
			573: 7782,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=north,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 6499,
			404: 6500,
			477: 7006,
			573: 7006,
		},
	},
	{
		"minecraft:iron_door[facing=south,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			393: 3324,
			404: 3325,
			477: 3828,
			573: 3828,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=13,south=none,west=up]",
		nil,
		NewMapping{
			393: 2019,
			404: 2020,
			477: 2323,
			573: 2323,
		},
	},
	{
		"minecraft:brown_bed[facing=south,occupied=false,part=foot]",
		nil,
		NewMapping{
			477: 1247,
			573: 1247,
			393: 947,
			404: 947,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9868,
			573: 9868,
		},
	},
	{
		"minecraft:campfire[facing=east,lit=true,signal_fire=true,waterlogged=true]",
		nil,
		NewMapping{
			573: 11256,
			477: 11240,
		},
	},
	{
		"minecraft:andesite_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9973,
			573: 9973,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=10,powered=true]",
		nil,
		NewMapping{
			477: 868,
			573: 868,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=east,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9282,
			573: 9282,
		},
	},
	{
		"minecraft:dark_oak_fence[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 7652,
			404: 7653,
			477: 8177,
			573: 8177,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=11,south=up,west=side]",
		nil,
		NewMapping{
			393: 2716,
			404: 2717,
			477: 3020,
			573: 3020,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=north,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9455,
			573: 9455,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10910,
			573: 10910,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=12,powered=true]",
		nil,
		NewMapping{
			477: 1022,
			573: 1022,
		},
	},
	{
		"minecraft:oak_fence_gate[facing=north,in_wall=false,open=false,powered=true]",
		nil,
		NewMapping{
			393: 4306,
			404: 4307,
			477: 4810,
			573: 4810,
			4:   1722,
		},
	},
	{
		"minecraft:fire[age=8,east=false,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			573: 1717,
			393: 1413,
			404: 1414,
			477: 1717,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=7,south=up,west=side]",
		nil,
		NewMapping{
			393: 1960,
			404: 1961,
			477: 2264,
			573: 2264,
		},
	},
	{
		"minecraft:turtle_egg[eggs=4,hatch=0]",
		nil,
		NewMapping{
			477: 8971,
			573: 8971,
			393: 8446,
			404: 8447,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10536,
			477: 10536,
		},
	},
	{
		"minecraft:structure_block[mode=data]",
		nil,
		NewMapping{
			393: 8581,
			404: 8598,
			477: 11255,
			573: 11271,
			4:   4083,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5929,
			404: 5930,
			477: 6436,
			573: 6436,
		},
	},
	{
		"minecraft:structure_block[mode=load]",
		nil,
		NewMapping{
			4:   4081,
			393: 8579,
			404: 8596,
			477: 11253,
			573: 11269,
		},
	},
	{
		"minecraft:magenta_glazed_terracotta[facing=east]",
		nil,
		NewMapping{
			4:   3795,
			393: 8324,
			404: 8325,
			477: 8849,
			573: 8849,
		},
	},
	{
		"minecraft:quartz_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 5774,
			477: 6280,
			573: 6280,
			393: 5773,
		},
	},
	{
		"minecraft:stone_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9656,
			573: 9656,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=east,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 6555,
			404: 6556,
			477: 7062,
			573: 7062,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=1,south=none,west=side]",
		nil,
		NewMapping{
			393: 2488,
			404: 2489,
			477: 2792,
			573: 2792,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9761,
			573: 9761,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=east,in_wall=false,open=true,powered=false]",
		nil,
		NewMapping{
			4:   2935,
			393: 7386,
			404: 7387,
			477: 7911,
			573: 7911,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9582,
			573: 9582,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=south,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3611,
			404: 3612,
			477: 4115,
			573: 4115,
		},
	},
	{
		"minecraft:green_banner[rotation=0]",
		nil,
		NewMapping{
			404: 7063,
			477: 7569,
			573: 7569,
			393: 7062,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=west,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			404: 3884,
			477: 4387,
			573: 4387,
			393: 3883,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=8,south=up,west=none]",
		nil,
		NewMapping{
			404: 2835,
			477: 3138,
			573: 3138,
			393: 2834,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=0,powered=true]",
		nil,
		NewMapping{
			404: 298,
			477: 298,
			573: 298,
			393: 298,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=south,half=top,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			477: 4247,
			573: 4247,
			393: 3743,
			404: 3744,
		},
	},
	{
		"minecraft:jungle_leaves[distance=3,persistent=true]",
		nil,
		NewMapping{
			573: 190,
			393: 190,
			404: 190,
			477: 190,
		},
	},
	{
		"minecraft:moving_piston[facing=east,type=normal]",
		nil,
		NewMapping{
			477: 1401,
			573: 1401,
			4:   581,
			393: 1101,
			404: 1101,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4726,
			404: 4727,
			477: 5230,
			573: 5230,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=1,south=none,west=side]",
		nil,
		NewMapping{
			393: 2632,
			404: 2633,
			477: 2936,
			573: 2936,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6233,
			404: 6234,
			477: 6740,
			573: 6740,
		},
	},
	{
		"minecraft:orange_banner[rotation=7]",
		nil,
		NewMapping{
			393: 6877,
			404: 6878,
			477: 7384,
			573: 7384,
		},
	},
	{
		"minecraft:granite_stairs[facing=east,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9913,
			573: 9913,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=1,south=none,west=none]",
		nil,
		NewMapping{
			477: 3225,
			573: 3225,
			4:   881,
			393: 2921,
			404: 2922,
		},
	},
	{
		"minecraft:brown_bed[facing=south,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 944,
			404: 944,
			477: 1244,
			573: 1244,
		},
	},
	{
		"minecraft:birch_door[facing=east,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			404: 7801,
			477: 8325,
			573: 8325,
			393: 7800,
		},
	},
	{
		"minecraft:blue_bed[facing=west,occupied=false,part=foot]",
		nil,
		NewMapping{
			573: 1235,
			393: 935,
			404: 935,
			477: 1235,
		},
	},
	{
		"minecraft:brick_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			393: 7321,
			404: 7322,
			477: 7840,
			573: 7840,
		},
	},
	{
		"minecraft:tube_coral_wall_fan[facing=west,waterlogged=true]",
		nil,
		NewMapping{
			393: 8508,
			404: 8524,
			477: 9068,
			573: 9068,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=17,powered=true]",
		nil,
		NewMapping{
			393: 282,
			404: 282,
			477: 282,
			573: 282,
		},
	},
	{
		"minecraft:red_bed[facing=east,occupied=true,part=foot]",
		nil,
		NewMapping{
			477: 1285,
			573: 1285,
			393: 985,
			404: 985,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=6,south=up,west=none]",
		nil,
		NewMapping{
			393: 2240,
			404: 2241,
			477: 2544,
			573: 2544,
		},
	},
	{
		"minecraft:glass_pane[east=false,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 4230,
			404: 4231,
			477: 4734,
			573: 4734,
		},
	},
	{
		"minecraft:spruce_fence[east=true,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7527,
			404: 7528,
			477: 8052,
			573: 8052,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 7274,
			393: 6767,
			404: 6768,
			477: 7274,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=true,north=false,powered=false,south=false,west=true]",
		nil,
		NewMapping{
			393: 4833,
			404: 4834,
			477: 5337,
			573: 5337,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9413,
			573: 9413,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 10022,
			477: 10022,
		},
	},
	{
		"minecraft:stone_button[face=wall,facing=north,powered=false]",
		nil,
		NewMapping{
			4:   1236,
			393: 3400,
			404: 3401,
			477: 3904,
			573: 3904,
		},
	},
	{
		"minecraft:oak_door[facing=east,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			4:   1028,
			393: 3168,
			404: 3169,
			477: 3632,
			573: 3632,
		},
	},
	{
		"minecraft:dark_oak_door[facing=west,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			477: 8502,
			573: 8502,
			393: 7977,
			404: 7978,
		},
	},
	{
		"minecraft:lime_bed[facing=west,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 837,
			404: 837,
			477: 1137,
			573: 1137,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9542,
			573: 9542,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=10,powered=false]",
		nil,
		NewMapping{
			477: 1019,
			573: 1019,
		},
	},
	{
		"minecraft:granite_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9926,
			573: 9926,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10533,
			573: 10533,
		},
	},
	{
		"minecraft:spruce_leaves[distance=1,persistent=true]",
		nil,
		NewMapping{
			4:   293,
			393: 158,
			404: 158,
			477: 158,
			573: 158,
		},
	},
	{
		"minecraft:magenta_terracotta",
		nil,
		NewMapping{
			477: 6313,
			573: 6313,
			4:   2546,
			393: 5806,
			404: 5807,
		},
	},
	{
		"minecraft:cyan_banner[rotation=14]",
		nil,
		NewMapping{
			477: 7519,
			573: 7519,
			393: 7012,
			404: 7013,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=13,south=none,west=up]",
		nil,
		NewMapping{
			477: 3043,
			573: 3043,
			393: 2739,
			404: 2740,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=west,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9423,
			573: 9423,
		},
	},
	{
		"minecraft:beehive[facing=north,honey_level=4]",
		nil,
		NewMapping{
			573: 11315,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=6,south=side,west=up]",
		nil,
		NewMapping{
			404: 2242,
			477: 2545,
			573: 2545,
			393: 2241,
		},
	},
	{
		"minecraft:stone_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9685,
			573: 9685,
		},
	},
	{
		"minecraft:fire[age=5,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			4:   821,
			393: 1326,
			404: 1327,
			477: 1630,
			573: 1630,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=12,south=none,west=side]",
		nil,
		NewMapping{
			477: 3323,
			573: 3323,
			393: 3019,
			404: 3020,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6145,
			404: 6146,
			477: 6652,
			573: 6652,
		},
	},
	{
		"minecraft:chest[facing=west,type=single,waterlogged=true]",
		nil,
		NewMapping{
			393: 1740,
			404: 1741,
			477: 2044,
			573: 2044,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=true,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 6719,
			393: 6212,
			404: 6213,
			477: 6719,
		},
	},
	{
		"minecraft:iron_door[facing=south,half=upper,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			404: 3327,
			477: 3830,
			573: 3830,
			393: 3326,
		},
	},
	{
		"minecraft:lime_bed[facing=east,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 842,
			404: 842,
			477: 1142,
			573: 1142,
		},
	},
	{
		"minecraft:fire[age=4,east=false,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1287,
			404: 1288,
			477: 1591,
			573: 1591,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9464,
			573: 9464,
		},
	},
	{
		"minecraft:granite_stairs[facing=south,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9883,
			573: 9883,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10059,
			573: 10059,
		},
	},
	{
		"minecraft:jungle_sign[rotation=0,waterlogged=true]",
		nil,
		NewMapping{
			477: 3507,
			573: 3507,
		},
	},
	{
		"minecraft:smoker[facing=east,lit=false]",
		nil,
		NewMapping{
			573: 11154,
			477: 11154,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=7,south=side,west=side]",
		nil,
		NewMapping{
			393: 2827,
			404: 2828,
			477: 3131,
			573: 3131,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 6703,
			573: 6703,
			393: 6196,
			404: 6197,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=11,powered=true]",
		nil,
		NewMapping{
			573: 720,
			393: 720,
			404: 720,
			477: 720,
		},
	},
	{
		"minecraft:quartz_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5757,
			404: 5758,
			477: 6264,
			573: 6264,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6110,
			404: 6111,
			477: 6617,
			573: 6617,
		},
	},
	{
		"minecraft:stone_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9650,
			573: 9650,
		},
	},
	{
		"minecraft:bee_nest[facing=west,honey_level=3]",
		nil,
		NewMapping{
			573: 11302,
		},
	},
	{
		"minecraft:comparator[facing=south,mode=compare,powered=true]",
		nil,
		NewMapping{
			4:   2392,
			393: 5639,
			404: 5640,
			477: 6146,
			573: 6146,
		},
	},
	{
		"minecraft:spruce_sapling[stage=0]",
		nil,
		NewMapping{
			4:   97,
			393: 23,
			404: 23,
			477: 23,
			573: 23,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=west,half=top,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			4:   1550,
			393: 3628,
			404: 3629,
			477: 4132,
			573: 4132,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=south,half=top,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			477: 4119,
			573: 4119,
			393: 3615,
			404: 3616,
		},
	},
	{
		"minecraft:black_banner[rotation=8]",
		nil,
		NewMapping{
			477: 7609,
			573: 7609,
			393: 7102,
			404: 7103,
		},
	},
	{
		"minecraft:birch_door[facing=east,half=upper,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			4:   3112,
			393: 7792,
			404: 7793,
			477: 8317,
			573: 8317,
		},
	},
	{
		"minecraft:acacia_door[facing=south,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7889,
			404: 7890,
			477: 8414,
			573: 8414,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=16,powered=true]",
		nil,
		NewMapping{
			477: 730,
			573: 730,
			393: 730,
			404: 730,
		},
	},
	{
		"minecraft:chorus_plant[down=true,east=false,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 8022,
			404: 8023,
			477: 8547,
			573: 8547,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10855,
			573: 10855,
		},
	},
	{
		"minecraft:spruce_leaves[distance=7,persistent=true]",
		nil,
		NewMapping{
			393: 170,
			404: 170,
			477: 170,
			573: 170,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=3,south=up,west=none]",
		nil,
		NewMapping{
			393: 2069,
			404: 2070,
			477: 2373,
			573: 2373,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=10,south=none,west=side]",
		nil,
		NewMapping{
			393: 2857,
			404: 2858,
			477: 3161,
			573: 3161,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=23,powered=false]",
		nil,
		NewMapping{
			477: 795,
			573: 795,
		},
	},
	{
		"minecraft:spruce_sign[rotation=2,waterlogged=false]",
		nil,
		NewMapping{
			477: 3416,
			573: 3416,
		},
	},
	{
		"minecraft:light_gray_bed[facing=west,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 884,
			404: 884,
			477: 1184,
			573: 1184,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=12,south=none,west=up]",
		nil,
		NewMapping{
			404: 2587,
			477: 2890,
			573: 2890,
			393: 2586,
		},
	},
	{
		"minecraft:brick_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			404: 4384,
			477: 4887,
			573: 4887,
			4:   1729,
			393: 4383,
		},
	},
	{
		"minecraft:dark_oak_fence_gate[facing=east,in_wall=false,open=false,powered=true]",
		nil,
		NewMapping{
			404: 7516,
			477: 8040,
			573: 8040,
			4:   2987,
			393: 7515,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=east,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			573: 4409,
			393: 3905,
			404: 3906,
			477: 4409,
		},
	},
	{
		"minecraft:spruce_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4956,
			404: 4957,
			477: 5460,
			573: 5460,
		},
	},
	{
		"minecraft:birch_fence[east=false,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7567,
			404: 7568,
			477: 8092,
			573: 8092,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=0,south=side,west=up]",
		nil,
		NewMapping{
			393: 1899,
			404: 1900,
			477: 2203,
			573: 2203,
		},
	},
	{
		"minecraft:oak_leaves[distance=3,persistent=false]",
		nil,
		NewMapping{
			393: 149,
			404: 149,
			477: 149,
			573: 149,
		},
	},
	{
		"minecraft:fletching_table",
		nil,
		NewMapping{
			477: 11164,
			573: 11164,
		},
	},
	{
		"minecraft:birch_log[axis=x]",
		nil,
		NewMapping{
			393: 78,
			404: 78,
			477: 78,
			573: 78,
			4:   278,
		},
	},
	{
		"minecraft:nether_brick_fence[east=false,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			404: 4519,
			477: 5022,
			573: 5022,
			393: 4518,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=false,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 4134,
			404: 4135,
			477: 4638,
			573: 4638,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4663,
			404: 4664,
			477: 5167,
			573: 5167,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=6,powered=false]",
		nil,
		NewMapping{
			393: 411,
			404: 411,
			477: 411,
			573: 411,
		},
	},
	{
		"minecraft:lectern[facing=north,has_book=false,powered=false]",
		nil,
		NewMapping{
			477: 11180,
			573: 11180,
		},
	},
	{
		"minecraft:repeater[delay=1,facing=east,locked=false,powered=false]",
		nil,
		NewMapping{
			4:   1491,
			393: 3528,
			404: 3529,
			477: 4032,
			573: 4032,
		},
	},
	{
		"minecraft:lime_banner[rotation=1]",
		nil,
		NewMapping{
			393: 6935,
			404: 6936,
			477: 7442,
			573: 7442,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9284,
			573: 9284,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=19,powered=false]",
		nil,
		NewMapping{
			477: 787,
			573: 787,
		},
	},
	{
		"minecraft:jack_o_lantern[facing=north]",
		nil,
		NewMapping{
			4:   1458,
			393: 3502,
			404: 3503,
			477: 4006,
			573: 4006,
		},
	},
	{
		"minecraft:nether_brick_fence[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 5024,
			393: 4520,
			404: 4521,
			477: 5024,
		},
	},
	{
		"minecraft:brick_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4405,
			404: 4406,
			477: 4909,
			573: 4909,
		},
	},
	{
		"minecraft:lectern[facing=east,has_book=true,powered=false]",
		nil,
		NewMapping{
			477: 11190,
			573: 11190,
		},
	},
	{
		"minecraft:granite_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9895,
			573: 9895,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=false,north=true,powered=true,south=false,west=false]",
		nil,
		NewMapping{
			393: 4838,
			404: 4839,
			477: 5342,
			573: 5342,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 6785,
			573: 6785,
			393: 6278,
			404: 6279,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5190,
			404: 5191,
			477: 5694,
			573: 5694,
		},
	},
	{
		"minecraft:jigsaw[facing=south]",
		nil,
		NewMapping{
			477: 11258,
			573: 11274,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=north,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 6502,
			404: 6503,
			477: 7009,
			573: 7009,
		},
	},
	{
		"minecraft:acacia_sign[rotation=13,waterlogged=false]",
		nil,
		NewMapping{
			573: 3502,
			477: 3502,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10978,
			573: 10978,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6615,
			404: 6616,
			477: 7122,
			573: 7122,
		},
	},
	{
		"minecraft:dead_fire_coral_wall_fan[facing=south,waterlogged=false]",
		nil,
		NewMapping{
			393: 8491,
			404: 8507,
			477: 9051,
			573: 9051,
		},
	},
	{
		"minecraft:gray_banner[rotation=13]",
		nil,
		NewMapping{
			573: 7486,
			393: 6979,
			404: 6980,
			477: 7486,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=16,powered=true]",
		nil,
		NewMapping{
			393: 630,
			404: 630,
			477: 630,
			573: 630,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=7,south=none,west=up]",
		nil,
		NewMapping{
			393: 2685,
			404: 2686,
			477: 2989,
			573: 2989,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5911,
			404: 5912,
			477: 6418,
			573: 6418,
		},
	},
	{
		"minecraft:mushroom_stem[down=false,east=false,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			573: 4670,
			393: 4166,
			404: 4167,
			477: 4670,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			404: 5897,
			477: 6403,
			573: 6403,
			393: 5896,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=11,south=none,west=up]",
		nil,
		NewMapping{
			393: 2433,
			404: 2434,
			477: 2737,
			573: 2737,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=7,south=up,west=side]",
		nil,
		NewMapping{
			393: 2536,
			404: 2537,
			477: 2840,
			573: 2840,
		},
	},
	{
		"minecraft:fire[age=3,east=false,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 1258,
			404: 1259,
			477: 1562,
			573: 1562,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=22,powered=true]",
		nil,
		NewMapping{
			573: 792,
			477: 792,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=17,powered=false]",
		nil,
		NewMapping{
			477: 1033,
			573: 1033,
		},
	},
	{
		"minecraft:spruce_door[facing=west,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7713,
			404: 7714,
			477: 8238,
			573: 8238,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 5180,
			393: 4676,
			404: 4677,
			477: 5180,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5977,
			404: 5978,
			477: 6484,
			573: 6484,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4551,
			404: 4552,
			477: 5055,
			573: 5055,
		},
	},
	{
		"minecraft:oak_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 1694,
			404: 1695,
			477: 1998,
			573: 1998,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=20,powered=false]",
		nil,
		NewMapping{
			393: 739,
			404: 739,
			477: 739,
			573: 739,
		},
	},
	{
		"minecraft:purpur_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 8149,
			404: 8150,
			477: 8674,
			573: 8674,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=west,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3825,
			404: 3826,
			477: 4329,
			573: 4329,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10399,
			573: 10399,
		},
	},
	{
		"minecraft:birch_wall_sign[facing=west,waterlogged=false]",
		nil,
		NewMapping{
			477: 3754,
			573: 3754,
		},
	},
	{
		"minecraft:heavy_weighted_pressure_plate[power=5]",
		nil,
		NewMapping{
			4:   2373,
			393: 5624,
			404: 5625,
			477: 6131,
			573: 6131,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=false,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 5197,
			477: 5700,
			573: 5700,
			4:   2224,
			393: 5196,
		},
	},
	{
		"minecraft:brown_banner[rotation=10]",
		nil,
		NewMapping{
			393: 7056,
			404: 7057,
			477: 7563,
			573: 7563,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=10,powered=true]",
		nil,
		NewMapping{
			393: 618,
			404: 618,
			477: 618,
			573: 618,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=11,south=none,west=side]",
		nil,
		NewMapping{
			477: 2450,
			573: 2450,
			393: 2146,
			404: 2147,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=7,powered=false]",
		nil,
		NewMapping{
			404: 513,
			477: 513,
			573: 513,
			393: 513,
		},
	},
	{
		"minecraft:purple_glazed_terracotta[facing=north]",
		nil,
		NewMapping{
			573: 8878,
			4:   3922,
			393: 8353,
			404: 8354,
			477: 8878,
		},
	},
	{
		"minecraft:chain_command_block[conditional=true,facing=up]",
		nil,
		NewMapping{
			393: 8180,
			404: 8181,
			477: 8705,
			573: 8705,
			4:   3385,
		},
	},
	{
		"minecraft:detector_rail[powered=false,shape=ascending_east]",
		nil,
		NewMapping{
			573: 1324,
			4:   450,
			393: 1024,
			404: 1024,
			477: 1324,
		},
	},
	{
		"minecraft:kelp[age=17]",
		nil,
		NewMapping{
			573: 8951,
			393: 8426,
			404: 8427,
			477: 8951,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=east,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			573: 7063,
			393: 6556,
			404: 6557,
			477: 7063,
		},
	},
	{
		"minecraft:oak_sign[rotation=10,waterlogged=true]",
		nil,
		NewMapping{
			404: 3096,
			477: 3399,
			573: 3399,
		},
	},
	{
		"minecraft:dark_oak_fence_gate[facing=west,in_wall=false,open=true,powered=true]",
		nil,
		NewMapping{
			404: 7506,
			477: 8030,
			573: 8030,
			4:   2989,
			393: 7505,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 6936,
			393: 6429,
			404: 6430,
			477: 6936,
		},
	},
	{
		"minecraft:acacia_wood[axis=z]",
		nil,
		NewMapping{
			573: 122,
			393: 122,
			404: 122,
			477: 122,
		},
	},
	{
		"minecraft:turtle_egg[eggs=1,hatch=0]",
		nil,
		NewMapping{
			393: 8437,
			404: 8438,
			477: 8962,
			573: 8962,
		},
	},
	{
		"minecraft:wall_sign[facing=east,waterlogged=true]",
		nil,
		NewMapping{
			393: 3275,
		},
	},
	{
		"minecraft:gray_bed[facing=west,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 871,
			404: 871,
			477: 1171,
			573: 1171,
		},
	},
	{
		"minecraft:acacia_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6399,
			404: 6400,
			477: 6906,
			573: 6906,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 7708,
			393: 7201,
			404: 7202,
			477: 7708,
		},
	},
	{
		"minecraft:acacia_fence[east=true,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 8143,
			573: 8143,
			393: 7618,
			404: 7619,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 10979,
			477: 10979,
		},
	},
	{
		"minecraft:fire[age=4,east=true,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1271,
			404: 1272,
			477: 1575,
			573: 1575,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9611,
			573: 9611,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=true,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10378,
			573: 10378,
		},
	},
	{
		"minecraft:brown_shulker_box[facing=down]",
		nil,
		NewMapping{
			573: 8819,
			4:   3696,
			393: 8294,
			404: 8295,
			477: 8819,
		},
	},
	{
		"minecraft:oak_fence_gate[facing=south,in_wall=false,open=true,powered=true]",
		nil,
		NewMapping{
			573: 4816,
			4:   1724,
			393: 4312,
			404: 4313,
			477: 4816,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=west,in_wall=true,open=false,powered=false]",
		nil,
		NewMapping{
			477: 7933,
			573: 7933,
			393: 7408,
			404: 7409,
		},
	},
	{
		"minecraft:dark_oak_fence[east=false,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7675,
			404: 7676,
			477: 8200,
			573: 8200,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=south,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3613,
			404: 3614,
			477: 4117,
			573: 4117,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 6615,
			573: 6615,
			393: 6108,
			404: 6109,
		},
	},
	{
		"minecraft:light_gray_bed[facing=north,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 876,
			404: 876,
			477: 1176,
			573: 1176,
		},
	},
	{
		"minecraft:oak_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 1681,
			477: 1984,
			573: 1984,
			393: 1680,
		},
	},
	{
		"minecraft:cyan_bed[facing=south,occupied=false,part=foot]",
		nil,
		NewMapping{
			573: 1199,
			393: 899,
			404: 899,
			477: 1199,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=10,waterlogged=true]",
		nil,
		NewMapping{
			477: 3559,
			573: 3559,
		},
	},
	{
		"minecraft:birch_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   2162,
			393: 4995,
			404: 4996,
			477: 5499,
			573: 5499,
		},
	},
	{
		"minecraft:white_banner[rotation=2]",
		nil,
		NewMapping{
			477: 7363,
			573: 7363,
			4:   2818,
			393: 6856,
			404: 6857,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=north,in_wall=false,open=false,powered=false]",
		nil,
		NewMapping{
			477: 7953,
			573: 7953,
			4:   2962,
			393: 7428,
			404: 7429,
		},
	},
	{
		"minecraft:green_bed[facing=west,occupied=true,part=foot]",
		nil,
		NewMapping{
			477: 1265,
			573: 1265,
			393: 965,
			404: 965,
		},
	},
	{
		"minecraft:fire[age=2,east=false,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1219,
			404: 1220,
			477: 1523,
			573: 1523,
		},
	},
	{
		"minecraft:wither_skeleton_skull[rotation=13]",
		nil,
		NewMapping{
			393: 5484,
			404: 5485,
			477: 5987,
			573: 5987,
		},
	},
	{
		"minecraft:spruce_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4927,
			404: 4928,
			477: 5431,
			573: 5431,
		},
	},
	{
		"minecraft:fire[age=2,east=true,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			477: 1517,
			573: 1517,
			393: 1213,
			404: 1214,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4469,
			404: 4470,
			477: 4973,
			573: 4973,
		},
	},
	{
		"minecraft:chiseled_quartz_block",
		nil,
		NewMapping{
			573: 6203,
			4:   2481,
			393: 5696,
			404: 5697,
			477: 6203,
		},
	},
	{
		"minecraft:jungle_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 5081,
			477: 5584,
			573: 5584,
			393: 5080,
		},
	},
	{
		"minecraft:fire[age=3,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			573: 1536,
			393: 1232,
			404: 1233,
			477: 1536,
		},
	},
	{
		"minecraft:fire[age=10,east=true,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1463,
			404: 1464,
			477: 1767,
			573: 1767,
		},
	},
	{
		"minecraft:petrified_oak_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			393: 7307,
			404: 7308,
			477: 7826,
			573: 7826,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10462,
			573: 10462,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10923,
			573: 10923,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=south,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3685,
			404: 3686,
			477: 4189,
			573: 4189,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5945,
			404: 5946,
			477: 6452,
			573: 6452,
		},
	},
	{
		"minecraft:nether_brick_fence[east=false,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 4523,
			404: 4524,
			477: 5027,
			573: 5027,
		},
	},
	{
		"minecraft:acacia_door[facing=north,half=upper,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7872,
			404: 7873,
			477: 8397,
			573: 8397,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=west,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			573: 4203,
			393: 3699,
			404: 3700,
			477: 4203,
		},
	},
	{
		"minecraft:tripwire_hook[attached=false,facing=east,powered=false]",
		nil,
		NewMapping{
			393: 4754,
			404: 4755,
			477: 5258,
			573: 5258,
			4:   2099,
		},
	},
	{
		"minecraft:blue_bed[facing=north,occupied=false,part=foot]",
		nil,
		NewMapping{
			477: 1227,
			573: 1227,
			393: 927,
			404: 927,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=9,south=up,west=none]",
		nil,
		NewMapping{
			573: 2139,
			393: 1835,
			404: 1836,
			477: 2139,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=15,south=side,west=side]",
		nil,
		NewMapping{
			393: 3043,
			404: 3044,
			477: 3347,
			573: 3347,
		},
	},
	{
		"minecraft:purpur_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 8635,
			393: 8110,
			404: 8111,
			477: 8635,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=south,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			477: 4250,
			573: 4250,
			393: 3746,
			404: 3747,
		},
	},
	{
		"minecraft:dark_oak_door[facing=west,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			404: 7970,
			477: 8494,
			573: 8494,
			393: 7969,
		},
	},
	{
		"minecraft:dark_oak_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			393: 7289,
			404: 7290,
			477: 7796,
			573: 7796,
		},
	},
	{
		"minecraft:birch_fence[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 8088,
			393: 7563,
			404: 7564,
			477: 8088,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 4069,
			404: 4070,
			477: 4573,
			573: 4573,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=north,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7453,
			404: 7454,
			477: 7978,
			573: 7978,
		},
	},
	{
		"minecraft:jungle_wood[axis=x]",
		nil,
		NewMapping{
			393: 117,
			404: 117,
			477: 117,
			573: 117,
		},
	},
	{
		"minecraft:brewing_stand[has_bottle_0=true,has_bottle_1=true,has_bottle_2=false]",
		nil,
		NewMapping{
			4:   1875,
			393: 4614,
			404: 4615,
			477: 5118,
			573: 5118,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=14,south=up,west=up]",
		nil,
		NewMapping{
			393: 2454,
			404: 2455,
			477: 2758,
			573: 2758,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=false,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6042,
			404: 6043,
			477: 6549,
			573: 6549,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=0,south=up,west=up]",
		nil,
		NewMapping{
			573: 2056,
			393: 1752,
			404: 1753,
			477: 2056,
		},
	},
	{
		"minecraft:fire[age=10,east=true,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1455,
			404: 1456,
			477: 1759,
			573: 1759,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4677,
			404: 4678,
			477: 5181,
			573: 5181,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5997,
			404: 5998,
			477: 6504,
			573: 6504,
		},
	},
	{
		"minecraft:birch_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4972,
			404: 4973,
			477: 5476,
			573: 5476,
		},
	},
	{
		"minecraft:fire[age=0,east=true,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			573: 1452,
			393: 1148,
			404: 1149,
			477: 1452,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=true,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10954,
			573: 10954,
		},
	},
	{
		"minecraft:brick_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4411,
			404: 4412,
			477: 4915,
			573: 4915,
		},
	},
	{
		"minecraft:acacia_door[facing=north,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			477: 8403,
			573: 8403,
			393: 7878,
			404: 7879,
		},
	},
	{
		"minecraft:bell[attachment=double_wall,facing=north,powered=true]",
		nil,
		NewMapping{
			573: 11222,
		},
	},
	{
		"minecraft:creeper_wall_head[facing=north]",
		nil,
		NewMapping{
			393: 5527,
			404: 5528,
			477: 6050,
			573: 6050,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9271,
			573: 9271,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9315,
			573: 9315,
		},
	},
	{
		"minecraft:acacia_door[facing=west,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			404: 7917,
			477: 8441,
			573: 8441,
			4:   3138,
			393: 7916,
		},
	},
	{
		"minecraft:quartz_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 6240,
			573: 6240,
			393: 5733,
			404: 5734,
		},
	},
	{
		"minecraft:gray_bed[facing=east,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 873,
			404: 873,
			477: 1173,
			573: 1173,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=true,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			404: 4054,
			477: 4557,
			573: 4557,
			393: 4053,
		},
	},
	{
		"minecraft:spruce_door[facing=east,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			404: 7737,
			477: 8261,
			573: 8261,
			393: 7736,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=11,powered=false]",
		nil,
		NewMapping{
			477: 1021,
			573: 1021,
		},
	},
	{
		"minecraft:repeater[delay=1,facing=south,locked=false,powered=false]",
		nil,
		NewMapping{
			404: 3521,
			477: 4024,
			573: 4024,
			4:   1488,
			393: 3520,
		},
	},
	{
		"minecraft:fire[age=8,east=true,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			404: 1397,
			477: 1700,
			573: 1700,
			393: 1396,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=15,south=up,west=up]",
		nil,
		NewMapping{
			393: 2751,
			404: 2752,
			477: 3055,
			573: 3055,
		},
	},
	{
		"minecraft:iron_bars[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 4694,
			393: 4190,
			404: 4191,
			477: 4694,
		},
	},
	{
		"minecraft:fire[age=5,east=true,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1299,
			404: 1300,
			477: 1603,
			573: 1603,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10439,
			573: 10439,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 10573,
			477: 10573,
		},
	},
	{
		"minecraft:cut_red_sandstone_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			477: 7867,
			573: 7867,
		},
	},
	{
		"minecraft:command_block[conditional=false,facing=west]",
		nil,
		NewMapping{
			4:   2196,
			393: 5133,
			404: 5134,
			477: 5637,
			573: 5637,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=13,south=up,west=none]",
		nil,
		NewMapping{
			393: 1871,
			404: 1872,
			477: 2175,
			573: 2175,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=west,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3881,
			404: 3882,
			477: 4385,
			573: 4385,
		},
	},
	{
		"minecraft:birch_button[face=floor,facing=south,powered=false]",
		nil,
		NewMapping{
			393: 5354,
			404: 5355,
			477: 5861,
			573: 5861,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9594,
			573: 9594,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=4,powered=true]",
		nil,
		NewMapping{
			477: 756,
			573: 756,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=14,powered=true]",
		nil,
		NewMapping{
			477: 926,
			573: 926,
		},
	},
	{
		"minecraft:spruce_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4896,
			404: 4897,
			477: 5400,
			573: 5400,
		},
	},
	{
		"minecraft:player_head[rotation=10]",
		nil,
		NewMapping{
			573: 6024,
			393: 5521,
			404: 5522,
			477: 6024,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=north,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3862,
			404: 3863,
			477: 4366,
			573: 4366,
		},
	},
	{
		"minecraft:birch_door[facing=north,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7746,
			404: 7747,
			477: 8271,
			573: 8271,
		},
	},
	{
		"minecraft:acacia_fence[east=true,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 7616,
			404: 7617,
			477: 8141,
			573: 8141,
		},
	},
	{
		"minecraft:fire[age=15,east=true,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1630,
			404: 1631,
			477: 1934,
			573: 1934,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=14,south=none,west=up]",
		nil,
		NewMapping{
			477: 3340,
			573: 3340,
			393: 3036,
			404: 3037,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=3,south=up,west=side]",
		nil,
		NewMapping{
			393: 2644,
			404: 2645,
			477: 2948,
			573: 2948,
		},
	},
	{
		"minecraft:prismarine_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			393: 6802,
			404: 6803,
			477: 7309,
			573: 7309,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10440,
			573: 10440,
		},
	},
	{
		"minecraft:jungle_leaves[distance=1,persistent=false]",
		nil,
		NewMapping{
			4:   291,
			393: 187,
			404: 187,
			477: 187,
			573: 187,
		},
	},
	{
		"minecraft:light_weighted_pressure_plate[power=6]",
		nil,
		NewMapping{
			477: 6116,
			573: 6116,
			4:   2358,
			393: 5609,
			404: 5610,
		},
	},
	{
		"minecraft:white_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5845,
			404: 5846,
			477: 6352,
			573: 6352,
		},
	},
	{
		"minecraft:nether_brick_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			393: 7329,
			404: 7330,
			477: 7848,
			573: 7848,
		},
	},
	{
		"minecraft:bricks",
		nil,
		NewMapping{
			393: 1125,
			404: 1125,
			477: 1428,
			573: 1428,
			4:   720,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=east,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			477: 4153,
			573: 4153,
			393: 3649,
			404: 3650,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=10,south=up,west=none]",
		nil,
		NewMapping{
			404: 2133,
			477: 2436,
			573: 2436,
			393: 2132,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			404: 5909,
			477: 6415,
			573: 6415,
			393: 5908,
		},
	},
	{
		"minecraft:fire[age=5,east=true,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1297,
			404: 1298,
			477: 1601,
			573: 1601,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 6440,
			393: 5933,
			404: 5934,
			477: 6440,
		},
	},
	{
		"minecraft:fire[age=3,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1243,
			404: 1244,
			477: 1547,
			573: 1547,
		},
	},
	{
		"minecraft:cyan_bed[facing=north,occupied=true,part=head]",
		nil,
		NewMapping{
			477: 1192,
			573: 1192,
			393: 892,
			404: 892,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=3,south=up,west=none]",
		nil,
		NewMapping{
			393: 1781,
			404: 1782,
			477: 2085,
			573: 2085,
		},
	},
	{
		"minecraft:brick_stairs[facing=south,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 4362,
			404: 4363,
			477: 4866,
			573: 4866,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=true,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			573: 4562,
			393: 4058,
			404: 4059,
			477: 4562,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=5,south=up,west=side]",
		nil,
		NewMapping{
			477: 2966,
			573: 2966,
			393: 2662,
			404: 2663,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=19,powered=true]",
		nil,
		NewMapping{
			393: 386,
			404: 386,
			477: 386,
			573: 386,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9171,
			573: 9171,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=17,powered=true]",
		nil,
		NewMapping{
			477: 832,
			573: 832,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10511,
			573: 10511,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=north,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			573: 4430,
			393: 3926,
			404: 3927,
			477: 4430,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=east,in_wall=true,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7384,
			404: 7385,
			477: 7909,
			573: 7909,
		},
	},
	{
		"minecraft:blue_banner[rotation=14]",
		nil,
		NewMapping{
			477: 7551,
			573: 7551,
			393: 7044,
			404: 7045,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9849,
			573: 9849,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9170,
			573: 9170,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=south,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3609,
			404: 3610,
			477: 4113,
			573: 4113,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9377,
			573: 9377,
		},
	},
	{
		"minecraft:brown_banner[rotation=12]",
		nil,
		NewMapping{
			393: 7058,
			404: 7059,
			477: 7565,
			573: 7565,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=true,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 4066,
			404: 4067,
			477: 4570,
			573: 4570,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=4,south=none,west=side]",
		nil,
		NewMapping{
			393: 1939,
			404: 1940,
			477: 2243,
			573: 2243,
		},
	},
	{
		"minecraft:fire[age=0,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			477: 1451,
			573: 1451,
			393: 1147,
			404: 1148,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=23,powered=true]",
		nil,
		NewMapping{
			393: 644,
			404: 644,
			477: 644,
			573: 644,
		},
	},
	{
		"minecraft:jungle_door[facing=east,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7865,
			404: 7866,
			477: 8390,
			573: 8390,
		},
	},
	{
		"minecraft:jungle_fence[east=true,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 7593,
			404: 7594,
			477: 8118,
			573: 8118,
		},
	},
	{
		"minecraft:yellow_bed[facing=east,occupied=false,part=head]",
		nil,
		NewMapping{
			477: 1126,
			573: 1126,
			393: 826,
			404: 826,
		},
	},
	{
		"minecraft:diorite_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10219,
			573: 10219,
		},
	},
	{
		"minecraft:fire[age=14,east=false,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1603,
			404: 1604,
			477: 1907,
			573: 1907,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 7696,
			393: 7189,
			404: 7190,
			477: 7696,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 4680,
			477: 5183,
			573: 5183,
			393: 4679,
		},
	},
	{
		"minecraft:blue_glazed_terracotta[facing=north]",
		nil,
		NewMapping{
			573: 8882,
			4:   3938,
			393: 8357,
			404: 8358,
			477: 8882,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 6544,
			573: 6544,
			393: 6037,
			404: 6038,
		},
	},
	{
		"minecraft:jungle_fence[east=false,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 8132,
			573: 8132,
			393: 7607,
			404: 7608,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=north,half=top,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3660,
			404: 3661,
			477: 4164,
			573: 4164,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=8,south=side,west=up]",
		nil,
		NewMapping{
			393: 2691,
			404: 2692,
			477: 2995,
			573: 2995,
		},
	},
	{
		"minecraft:diorite_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 11085,
			573: 11085,
		},
	},
	{
		"minecraft:acacia_door[facing=south,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7887,
			404: 7888,
			477: 8412,
			573: 8412,
		},
	},
	{
		"minecraft:conduit[waterlogged=false]",
		nil,
		NewMapping{
			573: 9114,
			404: 8590,
			477: 9114,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 11047,
			573: 11047,
		},
	},
	{
		"minecraft:carrots[age=0]",
		nil,
		NewMapping{
			404: 5288,
			477: 5794,
			573: 5794,
			4:   2256,
			393: 5287,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=8,powered=false]",
		nil,
		NewMapping{
			393: 565,
			404: 565,
			477: 565,
			573: 565,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=6,powered=false]",
		nil,
		NewMapping{
			477: 461,
			573: 461,
			393: 461,
			404: 461,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=11,south=up,west=side]",
		nil,
		NewMapping{
			573: 2300,
			393: 1996,
			404: 1997,
			477: 2300,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=2,south=side,west=side]",
		nil,
		NewMapping{
			393: 2782,
			404: 2783,
			477: 3086,
			573: 3086,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9303,
			573: 9303,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=east,half=top,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			404: 3707,
			477: 4210,
			573: 4210,
			393: 3706,
		},
	},
	{
		"minecraft:andesite_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			477: 10311,
			573: 10311,
		},
	},
	{
		"minecraft:stripped_acacia_wood[axis=x]",
		nil,
		NewMapping{
			393: 138,
			404: 138,
			477: 138,
			573: 138,
		},
	},
	{
		"minecraft:fire[age=12,east=false,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1549,
			404: 1550,
			477: 1853,
			573: 1853,
		},
	},
	{
		"minecraft:sweet_berry_bush[age=2]",
		nil,
		NewMapping{
			477: 11250,
			573: 11266,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10973,
			573: 10973,
		},
	},
	{
		"minecraft:smooth_sandstone_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			477: 10290,
			573: 10290,
		},
	},
	{
		"minecraft:daylight_detector[inverted=true,power=15]",
		nil,
		NewMapping{
			4:   2863,
			393: 5666,
			404: 5667,
			477: 6173,
			573: 6173,
		},
	},
	{
		"minecraft:black_stained_glass",
		nil,
		NewMapping{
			4:   1535,
			393: 3592,
			404: 3593,
			477: 4096,
			573: 4096,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 6798,
			477: 7304,
			573: 7304,
			393: 6797,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=east,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6475,
			404: 6476,
			477: 6982,
			573: 6982,
		},
	},
	{
		"minecraft:brick_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4384,
			404: 4385,
			477: 4888,
			573: 4888,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=south,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3816,
			404: 3817,
			477: 4320,
			573: 4320,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=4,south=none,west=up]",
		nil,
		NewMapping{
			393: 2802,
			404: 2803,
			477: 3106,
			573: 3106,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 10132,
			477: 10132,
		},
	},
	{
		"minecraft:jungle_door[facing=east,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7864,
			404: 7865,
			477: 8389,
			573: 8389,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 6639,
			573: 6639,
			393: 6132,
			404: 6133,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=19,powered=true]",
		nil,
		NewMapping{
			477: 786,
			573: 786,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 11037,
			573: 11037,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=true,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 11003,
			573: 11003,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=false,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10618,
			573: 10618,
		},
	},
	{
		"minecraft:spruce_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			404: 7267,
			477: 7773,
			573: 7773,
			4:   2017,
			393: 7266,
		},
	},
	{
		"minecraft:orange_banner[rotation=3]",
		nil,
		NewMapping{
			393: 6873,
			404: 6874,
			477: 7380,
			573: 7380,
		},
	},
	{
		"minecraft:dark_oak_button[face=floor,facing=south,powered=false]",
		nil,
		NewMapping{
			393: 5426,
			404: 5427,
			477: 5933,
			573: 5933,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=9,south=side,west=up]",
		nil,
		NewMapping{
			404: 2557,
			477: 2860,
			573: 2860,
			393: 2556,
		},
	},
	{
		"minecraft:acacia_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 6372,
			404: 6373,
			477: 6879,
			573: 6879,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4686,
			404: 4687,
			477: 5190,
			573: 5190,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=west,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3817,
			404: 3818,
			477: 4321,
			573: 4321,
		},
	},
	{
		"minecraft:bamboo[age=1,leaves=small,stage=1]",
		nil,
		NewMapping{
			477: 9125,
			573: 9125,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=false,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 6582,
			4:   2567,
			393: 6075,
			404: 6076,
			477: 6582,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=11,powered=true]",
		nil,
		NewMapping{
			477: 370,
			573: 370,
			393: 370,
			404: 370,
		},
	},
	{
		"minecraft:fire[age=2,east=true,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1206,
			404: 1207,
			477: 1510,
			573: 1510,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6176,
			404: 6177,
			477: 6683,
			573: 6683,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=north,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3721,
			404: 3722,
			477: 4225,
			573: 4225,
		},
	},
	{
		"minecraft:sign[rotation=0,waterlogged=true]",
		nil,
		NewMapping{
			393: 3075,
		},
	},
	{
		"minecraft:dark_oak_fence_gate[facing=south,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			477: 8018,
			573: 8018,
			393: 7493,
			404: 7494,
		},
	},
	{
		"minecraft:oak_sign[rotation=11,waterlogged=false]",
		nil,
		NewMapping{
			477: 3402,
			573: 3402,
			404: 3099,
		},
	},
	{
		"minecraft:scaffolding[bottom=true,distance=1,waterlogged=false]",
		nil,
		NewMapping{
			477: 11102,
			573: 11102,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9867,
			573: 9867,
		},
	},
	{
		"minecraft:acacia_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 6403,
			404: 6404,
			477: 6910,
			573: 6910,
			4:   2608,
		},
	},
	{
		"minecraft:blue_bed[facing=east,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 936,
			404: 936,
			477: 1236,
			573: 1236,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=17,powered=true]",
		nil,
		NewMapping{
			393: 632,
			404: 632,
			477: 632,
			573: 632,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=14,south=up,west=none]",
		nil,
		NewMapping{
			573: 2760,
			393: 2456,
			404: 2457,
			477: 2760,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=22,powered=true]",
		nil,
		NewMapping{
			477: 1042,
			573: 1042,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=false,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10902,
			573: 10902,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10738,
			573: 10738,
		},
	},
	{
		"minecraft:lime_shulker_box[facing=down]",
		nil,
		NewMapping{
			4:   3584,
			393: 8252,
			404: 8253,
			477: 8777,
			573: 8777,
		},
	},
	{
		"minecraft:fire[age=10,east=false,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			404: 1477,
			477: 1780,
			573: 1780,
			393: 1476,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=east,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4591,
			404: 4592,
			477: 5095,
			573: 5095,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 4422,
			404: 4423,
			477: 4926,
			573: 4926,
		},
	},
	{
		"minecraft:oak_fence[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 3460,
			404: 3461,
			477: 3964,
			573: 3964,
		},
	},
	{
		"minecraft:oak_sign[rotation=8,waterlogged=true]",
		nil,
		NewMapping{
			404: 3092,
			477: 3395,
			573: 3395,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=false,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10778,
			573: 10778,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9570,
			573: 9570,
		},
	},
	{
		"minecraft:white_banner[rotation=9]",
		nil,
		NewMapping{
			477: 7370,
			573: 7370,
			4:   2825,
			393: 6863,
			404: 6864,
		},
	},
	{
		"minecraft:iron_door[facing=south,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			4:   1137,
			393: 3334,
			404: 3335,
			477: 3838,
			573: 3838,
		},
	},
	{
		"minecraft:netherrack",
		nil,
		NewMapping{
			404: 3494,
			477: 3997,
			573: 3997,
			4:   1392,
			393: 3493,
		},
	},
	{
		"minecraft:bubble_coral_wall_fan[facing=east,waterlogged=false]",
		nil,
		NewMapping{
			393: 8527,
			404: 8543,
			477: 9087,
			573: 9087,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=true,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 4025,
			404: 4026,
			477: 4529,
			573: 4529,
		},
	},
	{
		"minecraft:beehive[facing=north,honey_level=0]",
		nil,
		NewMapping{
			573: 11311,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10856,
			573: 10856,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=12,powered=true]",
		nil,
		NewMapping{
			477: 772,
			573: 772,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9367,
			573: 9367,
		},
	},
	{
		"minecraft:spruce_door[facing=west,half=upper,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			477: 8235,
			573: 8235,
			393: 7710,
			404: 7711,
		},
	},
	{
		"minecraft:fire[age=1,east=true,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1182,
			404: 1183,
			477: 1486,
			573: 1486,
		},
	},
	{
		"minecraft:pink_wall_banner[facing=east]",
		nil,
		NewMapping{
			393: 7137,
			404: 7138,
			477: 7644,
			573: 7644,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=0,south=up,west=side]",
		nil,
		NewMapping{
			477: 3209,
			573: 3209,
			393: 2905,
			404: 2906,
		},
	},
	{
		"minecraft:fire[age=12,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1531,
			404: 1532,
			477: 1835,
			573: 1835,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=east,in_wall=true,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7416,
			404: 7417,
			477: 7941,
			573: 7941,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=12,south=side,west=none]",
		nil,
		NewMapping{
			393: 2009,
			404: 2010,
			477: 2313,
			573: 2313,
		},
	},
	{
		"minecraft:redstone_wall_torch[facing=north,lit=false]",
		nil,
		NewMapping{
			4:   1204,
			393: 3384,
			404: 3385,
			477: 3888,
			573: 3888,
		},
	},
	{
		"minecraft:oak_leaves[distance=2,persistent=true]",
		nil,
		NewMapping{
			404: 146,
			477: 146,
			573: 146,
			4:   300,
			393: 146,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=5,south=up,west=side]",
		nil,
		NewMapping{
			404: 1943,
			477: 2246,
			573: 2246,
			393: 1942,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=east,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 6621,
			404: 6622,
			477: 7128,
			573: 7128,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=false,north=true,powered=true,south=true,west=true]",
		nil,
		NewMapping{
			393: 4835,
			404: 4836,
			477: 5339,
			573: 5339,
		},
	},
	{
		"minecraft:beehive[facing=west,honey_level=3]",
		nil,
		NewMapping{
			573: 11326,
		},
	},
	{
		"minecraft:cocoa[age=2,facing=south]",
		nil,
		NewMapping{
			573: 5151,
			4:   2040,
			393: 4647,
			404: 4648,
			477: 5151,
		},
	},
	{
		"minecraft:fire[age=15,east=false,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			404: 1637,
			477: 1940,
			573: 1940,
			393: 1636,
		},
	},
	{
		"minecraft:dark_oak_door[facing=south,half=upper,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			477: 8475,
			573: 8475,
			393: 7950,
			404: 7951,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=8,south=none,west=side]",
		nil,
		NewMapping{
			573: 2999,
			393: 2695,
			404: 2696,
			477: 2999,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=true,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10364,
			477: 10364,
		},
	},
	{
		"minecraft:purple_bed[facing=east,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 923,
			404: 923,
			477: 1223,
			573: 1223,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9159,
			573: 9159,
		},
	},
	{
		"minecraft:acacia_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 6333,
			404: 6334,
			477: 6840,
			573: 6840,
			4:   2615,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6616,
			404: 6617,
			477: 7123,
			573: 7123,
		},
	},
	{
		"minecraft:fire[age=11,east=true,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			404: 1501,
			477: 1804,
			573: 1804,
			393: 1500,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 7133,
			573: 7133,
			393: 6626,
			404: 6627,
		},
	},
	{
		"minecraft:gray_bed[facing=west,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 870,
			404: 870,
			477: 1170,
			573: 1170,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=15,powered=true]",
		nil,
		NewMapping{
			393: 428,
			404: 428,
			477: 428,
			573: 428,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9457,
			573: 9457,
		},
	},
	{
		"minecraft:piston[extended=true,facing=up]",
		nil,
		NewMapping{
			4:   537,
			393: 1051,
			404: 1051,
			477: 1351,
			573: 1351,
		},
	},
	{
		"minecraft:red_bed[facing=west,occupied=false,part=head]",
		nil,
		NewMapping{
			477: 1282,
			573: 1282,
			4:   425,
			393: 982,
			404: 982,
		},
	},
	{
		"minecraft:acacia_door[facing=north,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			573: 8398,
			393: 7873,
			404: 7874,
			477: 8398,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=south,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3614,
			404: 3615,
			477: 4118,
			573: 4118,
		},
	},
	{
		"minecraft:acacia_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6359,
			404: 6360,
			477: 6866,
			573: 6866,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=1,south=up,west=up]",
		nil,
		NewMapping{
			477: 2497,
			573: 2497,
			393: 2193,
			404: 2194,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=18,powered=false]",
		nil,
		NewMapping{
			404: 585,
			477: 585,
			573: 585,
			393: 585,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=south,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 4432,
			404: 4433,
			477: 4936,
			573: 4936,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=east,in_wall=false,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7417,
			404: 7418,
			477: 7942,
			573: 7942,
			4:   2959,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 6654,
			393: 6147,
			404: 6148,
			477: 6654,
		},
	},
	{
		"minecraft:brick_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4358,
			404: 4359,
			477: 4862,
			573: 4862,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=2,south=side,west=up]",
		nil,
		NewMapping{
			393: 2637,
			404: 2638,
			477: 2941,
			573: 2941,
		},
	},
	{
		"minecraft:oak_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 1701,
			404: 1702,
			477: 2005,
			573: 2005,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9278,
			573: 9278,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9191,
			573: 9191,
		},
	},
	{
		"minecraft:birch_wall_sign[facing=west,waterlogged=true]",
		nil,
		NewMapping{
			477: 3753,
			573: 3753,
		},
	},
	{
		"minecraft:beehive[facing=south,honey_level=3]",
		nil,
		NewMapping{
			573: 11320,
		},
	},
	{
		"minecraft:repeating_command_block[conditional=false,facing=south]",
		nil,
		NewMapping{
			4:   3363,
			393: 8172,
			404: 8173,
			477: 8697,
			573: 8697,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=3,south=none,west=none]",
		nil,
		NewMapping{
			393: 1787,
			404: 1788,
			477: 2091,
			573: 2091,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6058,
			404: 6059,
			477: 6565,
			573: 6565,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 11032,
			573: 11032,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10400,
			573: 10400,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=true,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 10755,
			477: 10755,
		},
	},
	{
		"minecraft:frosted_ice[age=1]",
		nil,
		NewMapping{
			393: 8189,
			404: 8190,
			477: 8714,
			573: 8714,
			4:   3393,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=12,south=up,west=side]",
		nil,
		NewMapping{
			573: 2597,
			393: 2293,
			404: 2294,
			477: 2597,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=false,north=true,powered=true,south=false,west=true]",
		nil,
		NewMapping{
			393: 4805,
			404: 4806,
			477: 5309,
			573: 5309,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=4,south=none,west=side]",
		nil,
		NewMapping{
			477: 2531,
			573: 2531,
			393: 2227,
			404: 2228,
		},
	},
	{
		"minecraft:composter[level=8]",
		nil,
		NewMapping{
			477: 11270,
			573: 11286,
		},
	},
	{
		"minecraft:andesite_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9959,
			573: 9959,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 10671,
			477: 10671,
		},
	},
	{
		"minecraft:bee_nest[facing=east,honey_level=2]",
		nil,
		NewMapping{
			573: 11307,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   2052,
			393: 4711,
			404: 4712,
			477: 5215,
			573: 5215,
		},
	},
	{
		"minecraft:green_bed[facing=north,occupied=false,part=foot]",
		nil,
		NewMapping{
			477: 1259,
			573: 1259,
			393: 959,
			404: 959,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 4704,
			477: 5207,
			573: 5207,
			393: 4703,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=south,half=top,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3868,
			404: 3869,
			477: 4372,
			573: 4372,
		},
	},
	{
		"minecraft:jungle_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 5112,
			404: 5113,
			477: 5616,
			573: 5616,
		},
	},
	{
		"minecraft:repeater[delay=3,facing=east,locked=true,powered=true]",
		nil,
		NewMapping{
			404: 3558,
			477: 4061,
			573: 4061,
			393: 3557,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=south,half=top,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3802,
			404: 3803,
			477: 4306,
			573: 4306,
		},
	},
	{
		"minecraft:acacia_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6405,
			404: 6406,
			477: 6912,
			573: 6912,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 10437,
			477: 10437,
		},
	},
	{
		"minecraft:diorite_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10226,
			573: 10226,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9372,
			573: 9372,
		},
	},
	{
		"minecraft:cyan_stained_glass",
		nil,
		NewMapping{
			404: 3587,
			477: 4090,
			573: 4090,
			4:   1529,
			393: 3586,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=4,south=up,west=none]",
		nil,
		NewMapping{
			477: 2382,
			573: 2382,
			393: 2078,
			404: 2079,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 7253,
			404: 7254,
			477: 7760,
			573: 7760,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 9588,
			477: 9588,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 11022,
			477: 11022,
		},
	},
	{
		"minecraft:spruce_door[facing=north,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			4:   3091,
			393: 7692,
			404: 7693,
			477: 8217,
			573: 8217,
		},
	},
	{
		"minecraft:ender_chest[facing=west,waterlogged=true]",
		nil,
		NewMapping{
			393: 4735,
			404: 4736,
			477: 5239,
			573: 5239,
		},
	},
	{
		"minecraft:brick_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 4369,
			477: 4872,
			573: 4872,
			393: 4368,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=1,south=none,west=none]",
		nil,
		NewMapping{
			393: 2345,
			404: 2346,
			477: 2649,
			573: 2649,
		},
	},
	{
		"minecraft:spruce_sign[rotation=10,waterlogged=false]",
		nil,
		NewMapping{
			573: 3432,
			477: 3432,
		},
	},
	{
		"minecraft:blue_bed[facing=west,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 934,
			404: 934,
			477: 1234,
			573: 1234,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=11,south=up,west=none]",
		nil,
		NewMapping{
			477: 2733,
			573: 2733,
			393: 2429,
			404: 2430,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=2,powered=false]",
		nil,
		NewMapping{
			477: 803,
			573: 803,
		},
	},
	{
		"minecraft:redstone_lamp[lit=false]",
		nil,
		NewMapping{
			4:   1968,
			393: 4637,
			404: 4638,
			477: 5141,
			573: 5141,
		},
	},
	{
		"minecraft:potted_dandelion",
		nil,
		NewMapping{
			393: 5273,
			404: 5274,
			477: 5777,
			573: 5777,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6190,
			404: 6191,
			477: 6697,
			573: 6697,
		},
	},
	{
		"minecraft:black_bed[facing=west,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 996,
			404: 996,
			477: 1296,
			573: 1296,
		},
	},
	{
		"minecraft:white_bed[facing=east,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 761,
			404: 761,
			477: 1061,
			573: 1061,
		},
	},
	{
		"minecraft:potted_wither_rose",
		nil,
		NewMapping{
			477: 5789,
			573: 5789,
		},
	},
	{
		"minecraft:granite_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9892,
			573: 9892,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=3,south=none,west=side]",
		nil,
		NewMapping{
			393: 2650,
			404: 2651,
			477: 2954,
			573: 2954,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 6713,
			573: 6713,
			393: 6206,
			404: 6207,
		},
	},
	{
		"minecraft:oak_fence[east=false,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			404: 3489,
			477: 3992,
			573: 3992,
			393: 3488,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 10998,
			477: 10998,
		},
	},
	{
		"minecraft:skeleton_wall_skull[facing=west]",
		nil,
		NewMapping{
			4:   2316,
			393: 5449,
			404: 5450,
			477: 5972,
			573: 5972,
		},
	},
	{
		"minecraft:lever[face=ceiling,facing=north,powered=true]",
		nil,
		NewMapping{
			4:   1119,
			393: 3293,
			404: 3294,
			477: 3797,
			573: 3797,
		},
	},
	{
		"minecraft:dark_oak_button[face=wall,facing=north,powered=true]",
		nil,
		NewMapping{
			393: 5431,
			404: 5432,
			477: 5938,
			573: 5938,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6097,
			404: 6098,
			477: 6604,
			573: 6604,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9189,
			573: 9189,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9437,
			573: 9437,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=13,powered=false]",
		nil,
		NewMapping{
			573: 1025,
			477: 1025,
		},
	},
	{
		"minecraft:acacia_sign[rotation=11,waterlogged=false]",
		nil,
		NewMapping{
			477: 3498,
			573: 3498,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=north,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3666,
			404: 3667,
			477: 4170,
			573: 4170,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=3,south=side,west=up]",
		nil,
		NewMapping{
			573: 2518,
			393: 2214,
			404: 2215,
			477: 2518,
		},
	},
	{
		"minecraft:quartz_stairs[facing=south,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 5728,
			404: 5729,
			477: 6235,
			573: 6235,
		},
	},
	{
		"minecraft:glass_pane[east=false,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 4234,
			404: 4235,
			477: 4738,
			573: 4738,
		},
	},
	{
		"minecraft:spruce_fence[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 7530,
			404: 7531,
			477: 8055,
			573: 8055,
		},
	},
	{
		"minecraft:stonecutter[facing=west]",
		nil,
		NewMapping{
			477: 11196,
			573: 11196,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=east,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 7240,
			404: 7241,
			477: 7747,
			573: 7747,
		},
	},
	{
		"minecraft:rail[shape=north_east]",
		nil,
		NewMapping{
			4:   1065,
			393: 3188,
			404: 3189,
			477: 3652,
			573: 3652,
		},
	},
	{
		"minecraft:nether_brick_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			477: 7853,
			573: 7853,
			4:   702,
			393: 7334,
			404: 7335,
		},
	},
	{
		"minecraft:dispenser[facing=up,triggered=false]",
		nil,
		NewMapping{
			573: 242,
			4:   369,
			393: 242,
			404: 242,
			477: 242,
		},
	},
	{
		"minecraft:glass",
		nil,
		NewMapping{
			404: 230,
			477: 230,
			573: 230,
			4:   320,
			393: 230,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6444,
			404: 6445,
			477: 6951,
			573: 6951,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=true,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 11004,
			573: 11004,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10628,
			573: 10628,
		},
	},
	{
		"minecraft:quartz_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5743,
			404: 5744,
			477: 6250,
			573: 6250,
		},
	},
	{
		"minecraft:oak_stairs[facing=north,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 1954,
			393: 1650,
			404: 1651,
			477: 1954,
		},
	},
	{
		"minecraft:fire[age=2,east=false,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 1218,
			404: 1219,
			477: 1522,
			573: 1522,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=west,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3697,
			404: 3698,
			477: 4201,
			573: 4201,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10339,
			573: 10339,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=24,powered=true]",
		nil,
		NewMapping{
			477: 896,
			573: 896,
		},
	},
	{
		"minecraft:beehive[facing=west,honey_level=0]",
		nil,
		NewMapping{
			573: 11323,
		},
	},
	{
		"minecraft:oak_button[face=floor,facing=north,powered=false]",
		nil,
		NewMapping{
			4:   2293,
			393: 5304,
			404: 5305,
			477: 5811,
			573: 5811,
		},
	},
	{
		"minecraft:fire[age=4,east=false,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1289,
			404: 1290,
			477: 1593,
			573: 1593,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=8,powered=false]",
		nil,
		NewMapping{
			393: 715,
			404: 715,
			477: 715,
			573: 715,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 3712,
			393: 3248,
			404: 3249,
			477: 3712,
		},
	},
	{
		"minecraft:skeleton_skull[rotation=3]",
		nil,
		NewMapping{
			393: 5454,
			404: 5455,
			477: 5957,
			573: 5957,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=1,south=up,west=up]",
		nil,
		NewMapping{
			573: 3217,
			393: 2913,
			404: 2914,
			477: 3217,
		},
	},
	{
		"minecraft:acacia_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 6896,
			393: 6389,
			404: 6390,
			477: 6896,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=13,south=none,west=none]",
		nil,
		NewMapping{
			393: 2741,
			404: 2742,
			477: 3045,
			573: 3045,
		},
	},
	{
		"minecraft:scaffolding[bottom=true,distance=7,waterlogged=false]",
		nil,
		NewMapping{
			477: 11114,
			573: 11114,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=west,in_wall=false,open=false,powered=true]",
		nil,
		NewMapping{
			4:   2969,
			393: 7443,
			404: 7444,
			477: 7968,
			573: 7968,
		},
	},
	{
		"minecraft:prismarine_bricks",
		nil,
		NewMapping{
			404: 6560,
			477: 7066,
			573: 7066,
			4:   2689,
			393: 6559,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=13,south=side,west=up]",
		nil,
		NewMapping{
			393: 2448,
			404: 2449,
			477: 2752,
			573: 2752,
		},
	},
	{
		"minecraft:potted_cactus",
		nil,
		NewMapping{
			393: 5286,
			404: 5287,
			477: 5793,
			573: 5793,
		},
	},
	{
		"minecraft:lime_bed[facing=east,occupied=false,part=foot]",
		nil,
		NewMapping{
			477: 1143,
			573: 1143,
			393: 843,
			404: 843,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=11,south=up,west=up]",
		nil,
		NewMapping{
			404: 2284,
			477: 2587,
			573: 2587,
			393: 2283,
		},
	},
	{
		"minecraft:acacia_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 6845,
			573: 6845,
			393: 6338,
			404: 6339,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9783,
			573: 9783,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10640,
			573: 10640,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 9861,
			477: 9861,
		},
	},
	{
		"minecraft:fire[age=3,east=true,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			573: 1546,
			393: 1242,
			404: 1243,
			477: 1546,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=north,half=top,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			573: 4231,
			393: 3727,
			404: 3728,
			477: 4231,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6667,
			404: 6668,
			477: 7174,
			573: 7174,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=true,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			404: 5184,
			477: 5687,
			573: 5687,
			393: 5183,
		},
	},
	{
		"minecraft:lilac[half=upper]",
		nil,
		NewMapping{
			573: 7351,
			393: 6844,
			404: 6845,
			477: 7351,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=5,powered=true]",
		nil,
		NewMapping{
			477: 408,
			573: 408,
			393: 408,
			404: 408,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 3193,
			477: 3656,
			573: 3656,
			393: 3192,
		},
	},
	{
		"minecraft:gray_bed[facing=west,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 869,
			404: 869,
			477: 1169,
			573: 1169,
		},
	},
	{
		"minecraft:jungle_sign[rotation=4,waterlogged=true]",
		nil,
		NewMapping{
			477: 3515,
			573: 3515,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9300,
			573: 9300,
		},
	},
	{
		"minecraft:repeater[delay=1,facing=west,locked=false,powered=true]",
		nil,
		NewMapping{
			477: 4027,
			573: 4027,
			4:   1505,
			393: 3523,
			404: 3524,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4417,
			404: 4418,
			477: 4921,
			573: 4921,
		},
	},
	{
		"minecraft:green_bed[facing=north,occupied=false,part=head]",
		nil,
		NewMapping{
			573: 1258,
			393: 958,
			404: 958,
			477: 1258,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10845,
			573: 10845,
		},
	},
	{
		"minecraft:stone_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9666,
			573: 9666,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10604,
			573: 10604,
		},
	},
	{
		"minecraft:daylight_detector[inverted=true,power=14]",
		nil,
		NewMapping{
			4:   2862,
			393: 5665,
			404: 5666,
			477: 6172,
			573: 6172,
		},
	},
	{
		"minecraft:emerald_block",
		nil,
		NewMapping{
			4:   2128,
			393: 4883,
			404: 4884,
			477: 5387,
			573: 5387,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=3,south=none,west=side]",
		nil,
		NewMapping{
			393: 1930,
			404: 1931,
			477: 2234,
			573: 2234,
		},
	},
	{
		"minecraft:oak_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 1999,
			573: 1999,
			393: 1695,
			404: 1696,
		},
	},
	{
		"minecraft:turtle_egg[eggs=1,hatch=2]",
		nil,
		NewMapping{
			477: 8964,
			573: 8964,
			393: 8439,
			404: 8440,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=false,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10389,
			573: 10389,
		},
	},
	{
		"minecraft:bell[attachment=ceiling,facing=east,powered=false]",
		nil,
		NewMapping{
			573: 11213,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5855,
			404: 5856,
			477: 6362,
			573: 6362,
		},
	},
	{
		"minecraft:fire[age=14,east=false,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			404: 1600,
			477: 1903,
			573: 1903,
			393: 1599,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=1,south=side,west=side]",
		nil,
		NewMapping{
			404: 2054,
			477: 2357,
			573: 2357,
			393: 2053,
		},
	},
	{
		"minecraft:dead_bubble_coral_fan[waterlogged=true]",
		nil,
		NewMapping{
			393: 8548,
			404: 8564,
			477: 9008,
			573: 9008,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 6831,
			393: 6324,
			404: 6325,
			477: 6831,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=13,powered=true]",
		nil,
		NewMapping{
			393: 374,
			404: 374,
			477: 374,
			573: 374,
		},
	},
	{
		"minecraft:sticky_piston[extended=true,facing=east]",
		nil,
		NewMapping{
			573: 1329,
			4:   477,
			393: 1029,
			404: 1029,
			477: 1329,
		},
	},
	{
		"minecraft:lava[level=4]",
		nil,
		NewMapping{
			393: 54,
			404: 54,
			477: 54,
			573: 54,
			4:   180,
		},
	},
	{
		"minecraft:nether_brick_fence[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 4507,
			404: 4508,
			477: 5011,
			573: 5011,
		},
	},
	{
		"minecraft:jungle_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5047,
			404: 5048,
			477: 5551,
			573: 5551,
		},
	},
	{
		"minecraft:acacia_fence[east=true,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 8142,
			393: 7617,
			404: 7618,
			477: 8142,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=true,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6053,
			404: 6054,
			477: 6560,
			573: 6560,
		},
	},
	{
		"minecraft:dead_bubble_coral_wall_fan[facing=north,waterlogged=true]",
		nil,
		NewMapping{
			393: 8480,
			404: 8496,
			477: 9040,
			573: 9040,
		},
	},
	{
		"minecraft:light_gray_bed[facing=east,occupied=false,part=head]",
		nil,
		NewMapping{
			404: 890,
			477: 1190,
			573: 1190,
			393: 890,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=north,half=top,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			477: 4228,
			573: 4228,
			393: 3724,
			404: 3725,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=east,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3907,
			404: 3908,
			477: 4411,
			573: 4411,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=15,south=side,west=none]",
		nil,
		NewMapping{
			393: 2900,
			404: 2901,
			477: 3204,
			573: 3204,
		},
	},
	{
		"minecraft:jungle_door[facing=west,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			477: 8374,
			573: 8374,
			393: 7849,
			404: 7850,
		},
	},
	{
		"minecraft:stone_button[face=ceiling,facing=north,powered=true]",
		nil,
		NewMapping{
			4:   1240,
			393: 3407,
			404: 3408,
			477: 3911,
			573: 3911,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 6562,
			393: 6055,
			404: 6056,
			477: 6562,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=south,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 4680,
			404: 4681,
			477: 5184,
			573: 5184,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 4954,
			393: 4450,
			404: 4451,
			477: 4954,
		},
	},
	{
		"minecraft:sign[rotation=9,waterlogged=false]",
		nil,
		NewMapping{
			4:   1017,
			393: 3094,
		},
	},
	{
		"minecraft:oak_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 1672,
			404: 1673,
			477: 1976,
			573: 1976,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6431,
			404: 6432,
			477: 6938,
			573: 6938,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=east,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3720,
			404: 3721,
			477: 4224,
			573: 4224,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10821,
			573: 10821,
		},
	},
	{
		"minecraft:light_gray_carpet",
		nil,
		NewMapping{
			4:   2744,
			393: 6831,
			404: 6832,
			477: 7338,
			573: 7338,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=south,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 4941,
			573: 4941,
			393: 4437,
			404: 4438,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=true,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 4026,
			404: 4027,
			477: 4530,
			573: 4530,
		},
	},
	{
		"minecraft:sign[rotation=1,waterlogged=false]",
		nil,
		NewMapping{
			4:   1009,
			393: 3078,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=3,powered=false]",
		nil,
		NewMapping{
			573: 655,
			393: 655,
			404: 655,
			477: 655,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=0,powered=false]",
		nil,
		NewMapping{
			477: 899,
			573: 899,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 10110,
			477: 10110,
		},
	},
	{
		"minecraft:beehive[facing=west,honey_level=5]",
		nil,
		NewMapping{
			573: 11328,
		},
	},
	{
		"minecraft:glass_pane[east=true,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 4214,
			404: 4215,
			477: 4718,
			573: 4718,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=6,powered=false]",
		nil,
		NewMapping{
			573: 561,
			393: 561,
			404: 561,
			477: 561,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=8,powered=false]",
		nil,
		NewMapping{
			573: 465,
			393: 465,
			404: 465,
			477: 465,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9467,
			573: 9467,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9773,
			573: 9773,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=21,powered=false]",
		nil,
		NewMapping{
			477: 991,
			573: 991,
		},
	},
	{
		"minecraft:mossy_cobblestone_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			477: 10280,
			573: 10280,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5927,
			404: 5928,
			477: 6434,
			573: 6434,
		},
	},
	{
		"minecraft:fire[age=12,east=false,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1535,
			404: 1536,
			477: 1839,
			573: 1839,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			404: 6095,
			477: 6601,
			573: 6601,
			393: 6094,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=4,south=up,west=up]",
		nil,
		NewMapping{
			404: 1789,
			477: 2092,
			573: 2092,
			393: 1788,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			404: 5983,
			477: 6489,
			573: 6489,
			393: 5982,
		},
	},
	{
		"minecraft:andesite_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9998,
			573: 9998,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=16,powered=false]",
		nil,
		NewMapping{
			477: 781,
			573: 781,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=0,south=none,west=up]",
		nil,
		NewMapping{
			393: 2478,
			404: 2479,
			477: 2782,
			573: 2782,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=false,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 6068,
			477: 6574,
			573: 6574,
			393: 6067,
		},
	},
	{
		"minecraft:oak_sign[rotation=11,waterlogged=true]",
		nil,
		NewMapping{
			477: 3401,
			573: 3401,
			404: 3098,
		},
	},
	{
		"minecraft:diorite_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 10185,
			477: 10185,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10052,
			573: 10052,
		},
	},
	{
		"minecraft:stone_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9655,
			573: 9655,
		},
	},
	{
		"minecraft:light_blue_wool",
		nil,
		NewMapping{
			393: 1086,
			404: 1086,
			477: 1386,
			573: 1386,
			4:   563,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=east,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3844,
			404: 3845,
			477: 4348,
			573: 4348,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 6723,
			573: 6723,
			393: 6216,
			404: 6217,
		},
	},
	{
		"minecraft:spruce_fence[east=false,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 8068,
			393: 7543,
			404: 7544,
			477: 8068,
		},
	},
	{
		"minecraft:lever[face=ceiling,facing=east,powered=true]",
		nil,
		NewMapping{
			393: 3299,
			404: 3300,
			477: 3803,
			573: 3803,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9586,
			573: 9586,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10385,
			573: 10385,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=13,south=up,west=up]",
		nil,
		NewMapping{
			393: 2445,
			404: 2446,
			477: 2749,
			573: 2749,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=1,south=side,west=up]",
		nil,
		NewMapping{
			477: 3220,
			573: 3220,
			393: 2916,
			404: 2917,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4672,
			404: 4673,
			477: 5176,
			573: 5176,
		},
	},
	{
		"minecraft:fire[age=10,east=false,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 1472,
			404: 1473,
			477: 1776,
			573: 1776,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 10447,
			477: 10447,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=0,south=side,west=up]",
		nil,
		NewMapping{
			477: 2347,
			573: 2347,
			393: 2043,
			404: 2044,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=6,south=up,west=none]",
		nil,
		NewMapping{
			393: 1808,
			404: 1809,
			477: 2112,
			573: 2112,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6191,
			404: 6192,
			477: 6698,
			573: 6698,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=east,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3655,
			404: 3656,
			477: 4159,
			573: 4159,
		},
	},
	{
		"minecraft:acacia_door[facing=north,half=upper,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7876,
			404: 7877,
			477: 8401,
			573: 8401,
		},
	},
	{
		"minecraft:fire[age=15,east=true,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1628,
			404: 1629,
			477: 1932,
			573: 1932,
		},
	},
	{
		"minecraft:stone_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 9635,
			477: 9635,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=south,in_wall=false,open=true,powered=true]",
		nil,
		NewMapping{
			4:   2956,
			393: 7401,
			404: 7402,
			477: 7926,
			573: 7926,
		},
	},
	{
		"minecraft:purpur_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 8659,
			573: 8659,
			393: 8134,
			404: 8135,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=south,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			404: 3683,
			477: 4186,
			573: 4186,
			393: 3682,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 6426,
			477: 6932,
			573: 6932,
			393: 6425,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=3,south=side,west=none]",
		nil,
		NewMapping{
			393: 2936,
			404: 2937,
			477: 3240,
			573: 3240,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=south,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9478,
			573: 9478,
		},
	},
	{
		"minecraft:oak_door[facing=west,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			573: 3616,
			4:   1030,
			393: 3152,
			404: 3153,
			477: 3616,
		},
	},
	{
		"minecraft:birch_door[facing=south,half=lower,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			477: 8290,
			573: 8290,
			393: 7765,
			404: 7766,
		},
	},
	{
		"minecraft:birch_leaves[distance=3,persistent=true]",
		nil,
		NewMapping{
			393: 176,
			404: 176,
			477: 176,
			573: 176,
		},
	},
	{
		"minecraft:bamboo[age=0,leaves=small,stage=1]",
		nil,
		NewMapping{
			477: 9119,
			573: 9119,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10976,
			573: 10976,
		},
	},
	{
		"minecraft:hay_block[axis=z]",
		nil,
		NewMapping{
			4:   2728,
			393: 6822,
			404: 6823,
			477: 7329,
			573: 7329,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6635,
			404: 6636,
			477: 7142,
			573: 7142,
		},
	},
	{
		"minecraft:stone_button[face=ceiling,facing=south,powered=false]",
		nil,
		NewMapping{
			477: 3914,
			573: 3914,
			393: 3410,
			404: 3411,
		},
	},
	{
		"minecraft:trapped_chest[facing=east,type=right,waterlogged=true]",
		nil,
		NewMapping{
			393: 5601,
			404: 5602,
			477: 6108,
			573: 6108,
		},
	},
	{
		"minecraft:oak_door[facing=west,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			404: 3149,
			477: 3612,
			573: 3612,
			393: 3148,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=23,powered=true]",
		nil,
		NewMapping{
			393: 444,
			404: 444,
			477: 444,
			573: 444,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=12,powered=true]",
		nil,
		NewMapping{
			573: 972,
			477: 972,
		},
	},
	{
		"minecraft:jungle_button[face=floor,facing=south,powered=false]",
		nil,
		NewMapping{
			393: 5378,
			404: 5379,
			477: 5885,
			573: 5885,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=3,south=side,west=side]",
		nil,
		NewMapping{
			477: 2087,
			573: 2087,
			393: 1783,
			404: 1784,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6575,
			404: 6576,
			477: 7082,
			573: 7082,
		},
	},
	{
		"minecraft:birch_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 5009,
			404: 5010,
			477: 5513,
			573: 5513,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=false,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			477: 4512,
			573: 4512,
			393: 4008,
			404: 4009,
		},
	},
	{
		"minecraft:oak_fence_gate[facing=east,in_wall=false,open=true,powered=false]",
		nil,
		NewMapping{
			4:   1719,
			393: 4329,
			404: 4330,
			477: 4833,
			573: 4833,
		},
	},
	{
		"minecraft:light_weighted_pressure_plate[power=11]",
		nil,
		NewMapping{
			4:   2363,
			393: 5614,
			404: 5615,
			477: 6121,
			573: 6121,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5944,
			404: 5945,
			477: 6451,
			573: 6451,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9739,
			573: 9739,
		},
	},
	{
		"minecraft:jungle_wall_sign[facing=east,waterlogged=false]",
		nil,
		NewMapping{
			477: 3772,
			573: 3772,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 4957,
			4:   1749,
			393: 4453,
			404: 4454,
			477: 4957,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=false,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5874,
			404: 5875,
			477: 6381,
			573: 6381,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=22,powered=false]",
		nil,
		NewMapping{
			477: 443,
			573: 443,
			393: 443,
			404: 443,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=7,south=none,west=none]",
		nil,
		NewMapping{
			393: 2255,
			404: 2256,
			477: 2559,
			573: 2559,
		},
	},
	{
		"minecraft:daylight_detector[inverted=true,power=6]",
		nil,
		NewMapping{
			393: 5657,
			404: 5658,
			477: 6164,
			573: 6164,
			4:   2854,
		},
	},
	{
		"minecraft:lever[face=floor,facing=west,powered=true]",
		nil,
		NewMapping{
			4:   1118,
			393: 3281,
			404: 3282,
			477: 3785,
			573: 3785,
		},
	},
	{
		"minecraft:fire[age=5,east=false,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 1312,
			404: 1313,
			477: 1616,
			573: 1616,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=8,powered=true]",
		nil,
		NewMapping{
			393: 514,
			404: 514,
			477: 514,
			573: 514,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 6391,
			573: 6391,
			393: 5884,
			404: 5885,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=11,south=side,west=up]",
		nil,
		NewMapping{
			477: 2734,
			573: 2734,
			393: 2430,
			404: 2431,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10760,
			573: 10760,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=22,powered=false]",
		nil,
		NewMapping{
			477: 843,
			573: 843,
		},
	},
	{
		"minecraft:spruce_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 5391,
			393: 4887,
			404: 4888,
			477: 5391,
		},
	},
	{
		"minecraft:jungle_button[face=ceiling,facing=south,powered=false]",
		nil,
		NewMapping{
			393: 5394,
			404: 5395,
			477: 5901,
			573: 5901,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=5,south=up,west=up]",
		nil,
		NewMapping{
			393: 2373,
			404: 2374,
			477: 2677,
			573: 2677,
		},
	},
	{
		"minecraft:lilac[half=lower]",
		nil,
		NewMapping{
			4:   2801,
			393: 6845,
			404: 6846,
			477: 7352,
			573: 7352,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=false,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10548,
			573: 10548,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 11036,
			477: 11036,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10488,
			573: 10488,
		},
	},
	{
		"minecraft:polished_granite_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			477: 10255,
			573: 10255,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=south,in_wall=false,open=true,powered=true]",
		nil,
		NewMapping{
			4:   2972,
			393: 7433,
			404: 7434,
			477: 7958,
			573: 7958,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=west,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3703,
			404: 3704,
			477: 4207,
			573: 4207,
		},
	},
	{
		"minecraft:fire[age=15,east=true,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 1624,
			404: 1625,
			477: 1928,
			573: 1928,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=9,south=none,west=none]",
		nil,
		NewMapping{
			393: 2849,
			404: 2850,
			477: 3153,
			573: 3153,
		},
	},
	{
		"minecraft:iron_door[facing=east,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 3365,
			404: 3366,
			477: 3869,
			573: 3869,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9524,
			573: 9524,
		},
	},
	{
		"minecraft:acacia_stairs[facing=south,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 6362,
			404: 6363,
			477: 6869,
			573: 6869,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6128,
			404: 6129,
			477: 6635,
			573: 6635,
		},
	},
	{
		"minecraft:stone_stairs[facing=east,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9675,
			573: 9675,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10903,
			573: 10903,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=south,in_wall=true,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7463,
			404: 7464,
			477: 7988,
			573: 7988,
		},
	},
	{
		"minecraft:creeper_head[rotation=11]",
		nil,
		NewMapping{
			573: 6045,
			393: 5542,
			404: 5543,
			477: 6045,
		},
	},
	{
		"minecraft:lime_wall_banner[facing=north]",
		nil,
		NewMapping{
			477: 7637,
			573: 7637,
			393: 7130,
			404: 7131,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 4020,
			404: 4021,
			477: 4524,
			573: 4524,
		},
	},
	{
		"minecraft:purple_bed[facing=east,occupied=false,part=head]",
		nil,
		NewMapping{
			573: 1222,
			393: 922,
			404: 922,
			477: 1222,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10753,
			573: 10753,
		},
	},
	{
		"minecraft:bell[attachment=single_wall,facing=east]",
		nil,
		NewMapping{
			477: 11209,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=11,south=none,west=none]",
		nil,
		NewMapping{
			393: 2579,
			404: 2580,
			477: 2883,
			573: 2883,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=15,south=side,west=side]",
		nil,
		NewMapping{
			477: 2195,
			573: 2195,
			393: 1891,
			404: 1892,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=6,south=none,west=up]",
		nil,
		NewMapping{
			477: 2836,
			573: 2836,
			393: 2532,
			404: 2533,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=18,powered=false]",
		nil,
		NewMapping{
			477: 1035,
			573: 1035,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10770,
			573: 10770,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=10,powered=true]",
		nil,
		NewMapping{
			393: 418,
			404: 418,
			477: 418,
			573: 418,
		},
	},
	{
		"minecraft:purpur_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 8675,
			573: 8675,
			393: 8150,
			404: 8151,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=24,powered=false]",
		nil,
		NewMapping{
			477: 547,
			573: 547,
			393: 547,
			404: 547,
		},
	},
	{
		"minecraft:comparator[facing=south,mode=compare,powered=false]",
		nil,
		NewMapping{
			573: 6147,
			4:   2384,
			393: 5640,
			404: 5641,
			477: 6147,
		},
	},
	{
		"minecraft:orange_shulker_box[facing=north]",
		nil,
		NewMapping{
			4:   3522,
			393: 8223,
			404: 8224,
			477: 8748,
			573: 8748,
		},
	},
	{
		"minecraft:jungle_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 5623,
			393: 5119,
			404: 5120,
			477: 5623,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=4,south=side,west=none]",
		nil,
		NewMapping{
			573: 3105,
			393: 2801,
			404: 2802,
			477: 3105,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5252,
			404: 5253,
			477: 5756,
			573: 5756,
		},
	},
	{
		"minecraft:diorite_wall[east=false,north=true,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 11081,
			573: 11081,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=8,powered=false]",
		nil,
		NewMapping{
			477: 765,
			573: 765,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10775,
			573: 10775,
		},
	},
	{
		"minecraft:acacia_sign[rotation=11,waterlogged=true]",
		nil,
		NewMapping{
			477: 3497,
			573: 3497,
		},
	},
	{
		"minecraft:stone_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9684,
			573: 9684,
		},
	},
	{
		"minecraft:sticky_piston[extended=false,facing=east]",
		nil,
		NewMapping{
			573: 1335,
			4:   469,
			393: 1035,
			404: 1035,
			477: 1335,
		},
	},
	{
		"minecraft:blue_concrete",
		nil,
		NewMapping{
			393: 8388,
			404: 8389,
			477: 8913,
			573: 8913,
			4:   4027,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4696,
			404: 4697,
			477: 5200,
			573: 5200,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4434,
			404: 4435,
			477: 4938,
			573: 4938,
		},
	},
	{
		"minecraft:skeleton_skull[rotation=12]",
		nil,
		NewMapping{
			573: 5966,
			393: 5463,
			404: 5464,
			477: 5966,
		},
	},
	{
		"minecraft:white_banner[rotation=3]",
		nil,
		NewMapping{
			4:   2819,
			393: 6857,
			404: 6858,
			477: 7364,
			573: 7364,
		},
	},
	{
		"minecraft:pink_carpet",
		nil,
		NewMapping{
			4:   2742,
			393: 6829,
			404: 6830,
			477: 7336,
			573: 7336,
		},
	},
	{
		"minecraft:cyan_terracotta",
		nil,
		NewMapping{
			4:   2553,
			393: 5813,
			404: 5814,
			477: 6320,
			573: 6320,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=west,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7373,
			404: 7374,
			477: 7898,
			573: 7898,
		},
	},
	{
		"minecraft:spruce_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4917,
			404: 4918,
			477: 5421,
			573: 5421,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10035,
			573: 10035,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=false,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10564,
			573: 10564,
		},
	},
	{
		"minecraft:glass_pane[east=false,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 4229,
			404: 4230,
			477: 4733,
			573: 4733,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=14,south=up,west=none]",
		nil,
		NewMapping{
			393: 2600,
			404: 2601,
			477: 2904,
			573: 2904,
		},
	},
	{
		"minecraft:kelp[age=5]",
		nil,
		NewMapping{
			573: 8939,
			393: 8414,
			404: 8415,
			477: 8939,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=5,south=none,west=none]",
		nil,
		NewMapping{
			393: 2381,
			404: 2382,
			477: 2685,
			573: 2685,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4602,
			404: 4603,
			477: 5106,
			573: 5106,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=21,powered=false]",
		nil,
		NewMapping{
			477: 941,
			573: 941,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=6,south=side,west=none]",
		nil,
		NewMapping{
			393: 1955,
			404: 1956,
			477: 2259,
			573: 2259,
		},
	},
	{
		"minecraft:oak_sign[rotation=12,waterlogged=true]",
		nil,
		NewMapping{
			404: 3100,
			477: 3403,
			573: 3403,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9840,
			477: 9840,
		},
	},
	{
		"minecraft:birch_door[facing=north,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			4:   3111,
			393: 7754,
			404: 7755,
			477: 8279,
			573: 8279,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 7195,
			404: 7196,
			477: 7702,
			573: 7702,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=21,powered=false]",
		nil,
		NewMapping{
			404: 741,
			477: 741,
			573: 741,
			393: 741,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 5051,
			573: 5051,
			393: 4547,
			404: 4548,
		},
	},
	{
		"minecraft:sign[rotation=14,waterlogged=true]",
		nil,
		NewMapping{
			393: 3103,
		},
	},
	{
		"minecraft:granite_stairs[facing=east,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9922,
			573: 9922,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=false,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10966,
			573: 10966,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=9,south=side,west=up]",
		nil,
		NewMapping{
			393: 2700,
			404: 2701,
			477: 3004,
			573: 3004,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 6657,
			573: 6657,
			393: 6150,
			404: 6151,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			404: 6173,
			477: 6679,
			573: 6679,
			393: 6172,
		},
	},
	{
		"minecraft:oak_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 1964,
			393: 1660,
			404: 1661,
			477: 1964,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=5,south=none,west=up]",
		nil,
		NewMapping{
			393: 2811,
			404: 2812,
			477: 3115,
			573: 3115,
		},
	},
	{
		"minecraft:acacia_leaves[distance=6,persistent=false]",
		nil,
		NewMapping{
			393: 211,
			404: 211,
			477: 211,
			573: 211,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10456,
			573: 10456,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10908,
			573: 10908,
		},
	},
	{
		"minecraft:cyan_shulker_box[facing=north]",
		nil,
		NewMapping{
			4:   3650,
			393: 8271,
			404: 8272,
			477: 8796,
			573: 8796,
		},
	},
	{
		"minecraft:oak_fence[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 3476,
			404: 3477,
			477: 3980,
			573: 3980,
		},
	},
	{
		"minecraft:blue_banner[rotation=3]",
		nil,
		NewMapping{
			477: 7540,
			573: 7540,
			393: 7033,
			404: 7034,
		},
	},
	{
		"minecraft:oak_button[face=ceiling,facing=west,powered=true]",
		nil,
		NewMapping{
			573: 5830,
			393: 5323,
			404: 5324,
			477: 5830,
		},
	},
	{
		"minecraft:lime_banner[rotation=7]",
		nil,
		NewMapping{
			393: 6941,
			404: 6942,
			477: 7448,
			573: 7448,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=true,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10812,
			573: 10812,
		},
	},
	{
		"minecraft:magenta_wool",
		nil,
		NewMapping{
			4:   562,
			393: 1085,
			404: 1085,
			477: 1385,
			573: 1385,
		},
	},
	{
		"minecraft:blue_bed[facing=east,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 938,
			404: 938,
			477: 1238,
			573: 1238,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10701,
			573: 10701,
		},
	},
	{
		"minecraft:powered_rail[powered=false,shape=ascending_east]",
		nil,
		NewMapping{
			4:   434,
			393: 1012,
			404: 1012,
			477: 1312,
			573: 1312,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=0,powered=false]",
		nil,
		NewMapping{
			393: 699,
			404: 699,
			477: 699,
			573: 699,
		},
	},
	{
		"minecraft:purpur_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 8143,
			477: 8667,
			573: 8667,
			393: 8142,
		},
	},
	{
		"minecraft:acacia_sign[rotation=12,waterlogged=true]",
		nil,
		NewMapping{
			573: 3499,
			477: 3499,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10046,
			573: 10046,
		},
	},
	{
		"minecraft:sign[rotation=10,waterlogged=true]",
		nil,
		NewMapping{
			393: 3095,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=true,north=true,powered=false,south=false,west=true]",
		nil,
		NewMapping{
			393: 4761,
			404: 4762,
			477: 5265,
			573: 5265,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=5,south=none,west=none]",
		nil,
		NewMapping{
			393: 2237,
			404: 2238,
			477: 2541,
			573: 2541,
		},
	},
	{
		"minecraft:dark_oak_door[facing=north,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7942,
			404: 7943,
			477: 8467,
			573: 8467,
		},
	},
	{
		"minecraft:andesite_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9968,
			573: 9968,
		},
	},
	{
		"minecraft:red_banner[rotation=1]",
		nil,
		NewMapping{
			404: 7080,
			477: 7586,
			573: 7586,
			393: 7079,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=false,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10649,
			573: 10649,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=20,powered=true]",
		nil,
		NewMapping{
			477: 638,
			573: 638,
			393: 638,
			404: 638,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=18,powered=true]",
		nil,
		NewMapping{
			393: 534,
			404: 534,
			477: 534,
			573: 534,
		},
	},
	{
		"minecraft:yellow_bed[facing=east,occupied=false,part=foot]",
		nil,
		NewMapping{
			477: 1127,
			573: 1127,
			393: 827,
			404: 827,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=south,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			404: 3620,
			477: 4123,
			573: 4123,
			393: 3619,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=3,powered=true]",
		nil,
		NewMapping{
			404: 554,
			477: 554,
			573: 554,
			393: 554,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5972,
			404: 5973,
			477: 6479,
			573: 6479,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10584,
			573: 10584,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 10092,
			477: 10092,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9240,
			573: 9240,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 6920,
			573: 6920,
			4:   2631,
			393: 6413,
			404: 6414,
		},
	},
	{
		"minecraft:spruce_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			477: 7771,
			573: 7771,
			4:   2025,
			393: 7264,
			404: 7265,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=2,south=side,west=up]",
		nil,
		NewMapping{
			393: 2061,
			404: 2062,
			477: 2365,
			573: 2365,
		},
	},
	{
		"minecraft:spruce_pressure_plate[powered=true]",
		nil,
		NewMapping{
			393: 3369,
			404: 3370,
			477: 3873,
			573: 3873,
		},
	},
	{
		"minecraft:jungle_door[facing=south,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7827,
			404: 7828,
			477: 8352,
			573: 8352,
		},
	},
	{
		"minecraft:lime_bed[facing=west,occupied=false,part=foot]",
		nil,
		NewMapping{
			477: 1139,
			573: 1139,
			393: 839,
			404: 839,
		},
	},
	{
		"minecraft:prismarine_brick_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			393: 6808,
			404: 6809,
			477: 7315,
			573: 7315,
		},
	},
	{
		"minecraft:jungle_fence[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 7594,
			404: 7595,
			477: 8119,
			573: 8119,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 6817,
			573: 6817,
			393: 6310,
			404: 6311,
		},
	},
	{
		"minecraft:structure_void",
		nil,
		NewMapping{
			4:   3472,
			393: 8198,
			404: 8199,
			477: 8723,
			573: 8723,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=11,powered=true]",
		nil,
		NewMapping{
			404: 420,
			477: 420,
			573: 420,
			393: 420,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6768,
			404: 6769,
			477: 7275,
			573: 7275,
		},
	},
	{
		"minecraft:acacia_fence[east=false,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7643,
			404: 7644,
		},
	},