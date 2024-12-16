package main

import (
	"fmt"
	"time"
)

func timeTaken() func() {
    start := time.Now()

    return func() {
        elapsed := time.Since(start)
        fmt.Println(elapsed)
    }
}


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2, len1, len2 := 0, 0, len(nums1), len(nums2)
	mid := ((len1 + len2)/2)
	zipped := []int{}

	for l1 < len1 && l2 < len2 {
		if nums1[l1] <= nums2[l2] {
			zipped = append(zipped, nums1[l1])
			l1++
		} else {
			zipped = append(zipped, nums2[l2])
			l2++
		}
	}
	if l1 < len1 {
		zipped = append(zipped, nums1[l1:]...)
		l1++
	}
	if l2 < len2 {
		zipped = append(zipped, nums2[l2:]...)
		l2++
	}
	// mid := len(zipped)/2
	// if (len1 + len2)%2 == 0 {
	// 	return (float64(zipped[mid-1])+float64(zipped[mid]))/2}
	// return float64(zipped[mid])
	if (len1 + len2)%2 == 1 {
		return float64(zipped[mid])
	}
	return (float64(zipped[mid-1])+float64(zipped[mid]))/2
}

func main() {
	nums1 := []int{1, 2, 6, 11, 14, 15, 15, 18, 20, 24, 26, 27, 29, 29, 32, 33, 33, 35, 37, 38, 42, 46, 49, 49, 50, 51, 53, 55, 56, 57, 58, 59, 61, 62, 63, 64, 66, 66, 66, 67, 70, 70, 72, 73, 75, 75, 76, 80, 81, 81, 82, 82, 83, 84, 85, 88, 88, 88, 88, 89, 90, 91, 93, 95, 99, 103, 104, 104, 106, 107, 108, 110, 112, 113, 114, 114, 115, 115, 116, 118, 123, 124, 126, 131, 132, 135, 135, 137, 139, 139, 139, 140, 146, 150, 150, 159, 161, 167, 167, 169, 170, 172, 174, 175, 181, 182, 183, 183, 186, 187, 191, 194, 194, 194, 195, 196, 200, 201, 205, 206, 207, 208, 209, 213, 215, 216, 216, 216, 217, 220, 221, 223, 224, 226, 227, 228, 229, 233, 233, 233, 234, 235, 235, 235, 237, 237, 239, 240, 241, 244, 246, 247, 247, 248, 249, 253, 257, 259, 261, 261, 261, 266, 267, 271, 272, 275, 282, 285, 287, 287, 288, 290, 292, 292, 292, 292, 295, 297, 298, 298, 298, 300, 301, 302, 302, 302, 306, 308, 311, 312, 313, 314, 319, 319, 321, 323, 326, 331, 331, 333, 335, 336, 339, 340, 344, 344, 348, 352, 353, 354, 355, 356, 358, 358, 358, 359, 361, 367, 367, 368, 371, 371, 372, 373, 375, 375, 376, 376, 377, 384, 384, 385, 385, 385, 386, 386, 386, 388, 388, 390, 391, 391, 393, 396, 396, 405, 411, 412, 415, 418, 419, 419, 420, 421, 424, 426, 426, 428, 430, 432, 432, 433, 433, 433, 434, 434, 434, 436, 437, 440, 443, 443, 444, 446, 449, 451, 457, 458, 459, 461, 461, 464, 464, 469, 469, 471, 480, 485, 487, 489, 490, 491, 491, 492, 494, 495, 496, 499, 500, 500}
	nums2 := []int{0, 1, 7, 8, 12, 12, 14, 23, 23, 24, 26, 27, 33, 33, 38, 41, 43, 44, 47, 48, 50, 51, 51, 54, 55, 58, 58, 58, 61, 65, 65, 66, 69, 74, 74, 77, 82, 86, 87, 87, 92, 95, 96, 101, 102, 104, 105, 105, 107, 109, 109, 114, 116, 117, 119, 123, 129, 135, 136, 137, 140, 140, 141, 141, 144, 145, 147, 151, 151, 154, 154, 156, 158, 158, 159, 162, 163, 164, 166, 167, 168, 168, 172, 173, 174, 174, 175, 177, 180, 181, 181, 182, 182, 182, 183, 185, 187, 187, 188, 188, 188, 194, 195, 197, 197, 198, 198, 198, 200, 201, 202, 203, 204, 206, 217, 219, 219, 220, 221, 223, 224, 225, 225, 227, 228, 230, 230, 239, 239, 240, 240, 242, 243, 243, 245, 245, 245, 247, 247, 247, 248, 249, 249, 253, 253, 255, 257, 257, 258, 259, 259, 259, 259, 261, 264, 264, 266, 267, 269, 272, 275, 275, 275, 279, 279, 281, 284, 285, 285, 286, 286, 286, 288, 288, 289, 296, 296, 297, 302, 304, 305, 309, 312, 313, 314, 319, 319, 322, 323, 325, 326, 327, 327, 329, 329, 332, 332, 334, 338, 339, 340, 340, 341, 342, 343, 344, 348, 349, 351, 353, 356, 360, 361, 362, 362, 362, 364, 364, 367, 368, 370, 375, 376, 378, 379, 381, 381, 382, 382, 385, 387, 388, 388, 390, 392, 392, 392, 392, 393, 394, 395, 395, 399, 399, 403, 405, 408, 409, 409, 409, 410, 411, 415, 416, 416, 419, 421, 422, 423, 424, 427, 427, 428, 429, 437, 437, 438, 438, 446, 447, 449, 452, 453, 454, 454, 457, 461, 467, 467, 471, 474, 474, 476, 478, 478, 478, 480, 481, 481, 482, 484, 491, 491, 493, 494, 494, 496, 496, 499, 500}
	nums3 := []int{0, 1, 2, 5, 6, 8, 9, 14, 15, 20, 23, 24, 25, 25, 26, 27, 30, 32, 36, 36, 37, 38, 41, 41, 42, 43, 44, 44, 46, 47, 47, 49, 51, 54, 54, 54, 55, 55, 55, 57, 57, 57, 58, 65, 66, 67, 67, 69, 70, 71, 72, 74, 75, 77, 77, 78, 80, 81, 82, 82, 86, 87, 87, 89, 90, 90, 90, 90, 90, 91, 91, 93, 95, 96, 103, 104, 104, 106, 109, 110, 115, 115, 117, 120, 121, 123, 125, 125, 125, 125, 128, 130, 133, 137, 137, 139, 141, 142, 146, 148, 151, 152, 154, 155, 156, 157, 157, 158, 161, 163, 165, 168, 169, 169, 169, 170, 172, 173, 179, 181, 181, 184, 185, 185, 185, 188, 190, 194, 195, 199, 199, 202, 202, 204, 204, 207, 207, 208, 212, 212, 214, 214, 218, 218, 219, 219, 222, 222, 225, 227, 227, 228, 231, 232, 233, 233, 234, 234, 236, 236, 236, 237, 237, 237, 238, 238, 240, 240, 241, 242, 243, 247, 248, 250, 252, 253, 259, 259, 262, 264, 265, 265, 268, 269, 274, 274, 279, 280, 282, 286, 286, 293, 297, 297, 299, 302, 303, 308, 309, 310, 313, 316, 316, 316, 317, 320, 323, 324, 325, 326, 328, 328, 328, 331, 331, 332, 332, 334, 335, 337, 338, 340, 348, 351, 355, 360, 362, 367, 367, 368, 368, 369, 369, 375, 379, 380, 380, 382, 385, 391, 392, 394, 398, 400, 402, 410, 411, 413, 414, 415, 415, 417, 417, 417, 419, 419, 424, 425, 427, 428, 428, 431, 435, 437, 439, 445, 445, 445, 448, 448, 450, 450, 450, 452, 453, 454, 457, 460, 462, 464, 465, 469, 470, 472, 472, 473, 474, 474, 476, 478, 482, 484, 486, 496, 496, 498, 498, 499, 500, 500}
	nums4 := []int{1, 1, 3, 4, 6, 6, 10, 14, 15, 17, 18, 22, 22, 25, 27, 29, 34, 34, 35, 35, 35, 38, 39, 39, 39, 40, 44, 47, 48, 48, 53, 53, 54, 55, 60, 61, 61, 63, 66, 67, 67, 71, 72, 72, 72, 73, 73, 74, 74, 75, 75, 76, 77, 79, 82, 85, 86, 90, 92, 92, 93, 93, 94, 98, 99, 100, 102, 106, 107, 108, 108, 111, 112, 113, 114, 114, 114, 114, 116, 120, 124, 132, 134, 134, 135, 136, 136, 137, 137, 138, 139, 141, 141, 142, 144, 150, 150, 153, 154, 157, 159, 161, 162, 162, 162, 162, 163, 163, 164, 164, 167, 167, 170, 172, 173, 176, 176, 179, 179, 180, 180, 188, 190, 191, 192, 192, 194, 194, 194, 195, 196, 199, 204, 207, 208, 209, 210, 211, 214, 214, 214, 218, 219, 221, 223, 223, 223, 224, 224, 225, 226, 226, 227, 235, 235, 236, 236, 239, 239, 242, 246, 249, 250, 250, 250, 252, 256, 258, 261, 261, 262, 262, 264, 268, 269, 271, 272, 275, 275, 275, 276, 278, 280, 280, 284, 284, 292, 292, 295, 295, 297, 300, 300, 302, 302, 302, 304, 308, 309, 310, 313, 314, 317, 319, 320, 320, 322, 324, 324, 325, 325, 333, 335, 336, 336, 337, 339, 341, 342, 347, 352, 355, 358, 363, 365, 365, 367, 367, 369, 371, 372, 374, 375, 383, 385, 385, 387, 391, 392, 393, 395, 396, 396, 396, 398, 399, 399, 399, 400, 403, 404, 409, 411, 419, 420, 423, 424, 426, 428, 430, 435, 435, 435, 437, 437, 437, 439, 439, 440, 442, 443, 452, 453, 456, 456, 457, 458, 458, 459, 459, 462, 466, 468, 470, 472, 474, 477, 479, 486, 486, 490, 492, 492, 493, 494, 496, 497, 498, 498, 499}
	nums5 := []int{1, 2, 6, 11, 14, 15, 15, 18, 20, 24, 26, 27, 29, 29, 32, 33, 33, 35, 37, 38, 42, 46, 49, 49, 50, 51, 53, 55, 56, 57, 58, 59, 61, 62, 63, 64, 66, 66, 66, 67, 70, 70, 72, 73, 75, 75, 76, 80, 81, 81, 82, 82, 83, 84, 85, 88, 88, 88, 88, 89, 90, 91, 93, 95, 99, 103, 104, 104, 106, 107, 108, 110, 112, 113, 114, 114, 115, 115, 116, 118, 123, 124, 126, 131, 132, 135, 135, 137, 139, 139, 139, 140, 146, 150, 150, 159, 161, 167, 167, 169, 170, 172, 174, 175, 181, 182, 183, 183, 186, 187, 191, 194, 194, 194, 195, 196, 200, 201, 205, 206, 207, 208, 209, 213, 215, 216, 216, 216, 217, 220, 221, 223, 224, 226, 227, 228, 229, 233, 233, 233, 234, 235, 235, 235, 237, 237, 239, 240, 241, 244, 246, 247, 247, 248, 249, 253, 257, 259, 261, 261, 261, 266, 267, 271, 272, 275, 282, 285, 287, 287, 288, 290, 292, 292, 292, 292, 295, 297, 298, 298, 298, 300, 301, 302, 302, 302, 306, 308, 311, 312, 313, 314, 319, 319, 321, 323, 326, 331, 331, 333, 335, 336, 339, 340, 344, 344, 348, 352, 353, 354, 355, 356, 358, 358, 358, 359, 361, 367, 367, 368, 371, 371, 372, 373, 375, 375, 376, 376, 377, 384, 384, 385, 385, 385, 386, 386, 386, 388, 388, 390, 391, 391, 393, 396, 396, 405, 411, 412, 415, 418, 419, 419, 420, 421, 424, 426, 426, 428, 430, 432, 432, 433, 433, 433, 434, 434, 434, 436, 437, 440, 443, 443, 444, 446, 449, 451, 457, 458, 459, 461, 461, 464, 464, 469, 469, 471, 480, 485, 487, 489, 490, 491, 491, 492, 494, 495, 496, 499, 500, 500}
	nums6 := []int{1, 1, 3, 4, 6, 6, 10, 14, 15, 17, 18, 22, 22, 25, 27, 29, 34, 34, 35, 35, 35, 38, 39, 39, 39, 40, 44, 47, 48, 48, 53, 53, 54, 55, 60, 61, 61, 63, 66, 67, 67, 71, 72, 72, 72, 73, 73, 74, 74, 75, 75, 76, 77, 79, 82, 85, 86, 90, 92, 92, 93, 93, 94, 98, 99, 100, 102, 106, 107, 108, 108, 111, 112, 113, 114, 114, 114, 114, 116, 120, 124, 132, 134, 134, 135, 136, 136, 137, 137, 138, 139, 141, 141, 142, 144, 150, 150, 153, 154, 157, 159, 161, 162, 162, 162, 162, 163, 163, 164, 164, 167, 167, 170, 172, 173, 176, 176, 179, 179, 180, 180, 188, 190, 191, 192, 192, 194, 194, 194, 195, 196, 199, 204, 207, 208, 209, 210, 211, 214, 214, 214, 218, 219, 221, 223, 223, 223, 224, 224, 225, 226, 226, 227, 235, 235, 236, 236, 239, 239, 242, 246, 249, 250, 250, 250, 252, 256, 258, 261, 261, 262, 262, 264, 268, 269, 271, 272, 275, 275, 275, 276, 278, 280, 280, 284, 284, 292, 292, 295, 295, 297, 300, 300, 302, 302, 302, 304, 308, 309, 310, 313, 314, 317, 319, 320, 320, 322, 324, 324, 325, 325, 333, 335, 336, 336, 337, 339, 341, 342, 347, 352, 355, 358, 363, 365, 365, 367, 367, 369, 371, 372, 374, 375, 383, 385, 385, 387, 391, 392, 393, 395, 396, 396, 396, 398, 399, 399, 399, 400, 403, 404, 409, 411, 419, 420, 423, 424, 426, 428, 430, 435, 435, 435, 437, 437, 437, 439, 439, 440, 442, 443, 452, 453, 456, 456, 457, 458, 458, 459, 459, 462, 466, 468, 470, 472, 474, 477, 479, 486, 486, 490, 492, 492, 493, 494, 496, 497, 498, 498, 499}
	nums7 := []int{0, 1, 2, 5, 6, 8, 9, 14, 15, 20, 23, 24, 25, 25, 26, 27, 30, 32, 36, 36, 37, 38, 41, 41, 42, 43, 44, 44, 46, 47, 47, 49, 51, 54, 54, 54, 55, 55, 55, 57, 57, 57, 58, 65, 66, 67, 67, 69, 70, 71, 72, 74, 75, 77, 77, 78, 80, 81, 82, 82, 86, 87, 87, 89, 90, 90, 90, 90, 90, 91, 91, 93, 95, 96, 103, 104, 104, 106, 109, 110, 115, 115, 117, 120, 121, 123, 125, 125, 125, 125, 128, 130, 133, 137, 137, 139, 141, 142, 146, 148, 151, 152, 154, 155, 156, 157, 157, 158, 161, 163, 165, 168, 169, 169, 169, 170, 172, 173, 179, 181, 181, 184, 185, 185, 185, 188, 190, 194, 195, 199, 199, 202, 202, 204, 204, 207, 207, 208, 212, 212, 214, 214, 218, 218, 219, 219, 222, 222, 225, 227, 227, 228, 231, 232, 233, 233, 234, 234, 236, 236, 236, 237, 237, 237, 238, 238, 240, 240, 241, 242, 243, 247, 248, 250, 252, 253, 259, 259, 262, 264, 265, 265, 268, 269, 274, 274, 279, 280, 282, 286, 286, 293, 297, 297, 299, 302, 303, 308, 309, 310, 313, 316, 316, 316, 317, 320, 323, 324, 325, 326, 328, 328, 328, 331, 331, 332, 332, 334, 335, 337, 338, 340, 348, 351, 355, 360, 362, 367, 367, 368, 368, 369, 369, 375, 379, 380, 380, 382, 385, 391, 392, 394, 398, 400, 402, 410, 411, 413, 414, 415, 415, 417, 417, 417, 419, 419, 424, 425, 427, 428, 428, 431, 435, 437, 439, 445, 445, 445, 448, 448, 450, 450, 450, 452, 453, 454, 457, 460, 462, 464, 465, 469, 470, 472, 472, 473, 474, 474, 476, 478, 482, 484, 486, 496, 496, 498, 498, 499, 500, 500}
	nums8 := []int{0, 1, 7, 8, 12, 12, 14, 23, 23, 24, 26, 27, 33, 33, 38, 41, 43, 44, 47, 48, 50, 51, 51, 54, 55, 58, 58, 58, 61, 65, 65, 66, 69, 74, 74, 77, 82, 86, 87, 87, 92, 95, 96, 101, 102, 104, 105, 105, 107, 109, 109, 114, 116, 117, 119, 123, 129, 135, 136, 137, 140, 140, 141, 141, 144, 145, 147, 151, 151, 154, 154, 156, 158, 158, 159, 162, 163, 164, 166, 167, 168, 168, 172, 173, 174, 174, 175, 177, 180, 181, 181, 182, 182, 182, 183, 185, 187, 187, 188, 188, 188, 194, 195, 197, 197, 198, 198, 198, 200, 201, 202, 203, 204, 206, 217, 219, 219, 220, 221, 223, 224, 225, 225, 227, 228, 230, 230, 239, 239, 240, 240, 242, 243, 243, 245, 245, 245, 247, 247, 247, 248, 249, 249, 253, 253, 255, 257, 257, 258, 259, 259, 259, 259, 261, 264, 264, 266, 267, 269, 272, 275, 275, 275, 279, 279, 281, 284, 285, 285, 286, 286, 286, 288, 288, 289, 296, 296, 297, 302, 304, 305, 309, 312, 313, 314, 319, 319, 322, 323, 325, 326, 327, 327, 329, 329, 332, 332, 334, 338, 339, 340, 340, 341, 342, 343, 344, 348, 349, 351, 353, 356, 360, 361, 362, 362, 362, 364, 364, 367, 368, 370, 375, 376, 378, 379, 381, 381, 382, 382, 385, 387, 388, 388, 390, 392, 392, 392, 392, 393, 394, 395, 395, 399, 399, 403, 405, 408, 409, 409, 409, 410, 411, 415, 416, 416, 419, 421, 422, 423, 424, 427, 427, 428, 429, 437, 437, 438, 438, 446, 447, 449, 452, 453, 454, 454, 457, 461, 467, 467, 471, 474, 474, 476, 478, 478, 478, 480, 481, 481, 482, 484, 491, 491, 493, 494, 494, 496, 496, 499, 500}
	
	defer timeTaken()()
	res := findMedianSortedArrays(nums1, nums2)
	fmt.Println(res)
	res = findMedianSortedArrays(nums3, nums4)
	fmt.Println(res)
	res = findMedianSortedArrays(nums5, nums6)
	fmt.Println(res)
	res = findMedianSortedArrays(nums7, nums8)
	fmt.Println(res)
}