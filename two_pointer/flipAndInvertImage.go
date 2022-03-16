// Given an n x n binary matrix image, flip the image horizontally, then invert it, and return the resulting image.

// To flip an image horizontally means that each row of the image is reversed.

// For example, flipping [1,1,0] horizontally results in [0,1,1].
// To invert an image means that each 0 is replaced by 1, and each 1 is replaced by 0.

// For example, inverting [0,1,1] results in [1,0,0].

package twopointer

func flipAndInvertImage(image [][]int) [][]int {
	for i := 0; i < len(image); i++ {
		left, right := 0, len(image[i])-1
		for left <= right {
			image[i][left], image[i][right] = image[i][right], image[i][left]
			left++
			right--
		}
		for j := 0; j < len(image[i]); j++ {
			if image[i][j] == 0 {
				image[i][j] = 1
			} else {
				image[i][j] = 0
			}
		}
	}
	return image
}
