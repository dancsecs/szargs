package main

func process(numbers []float64, operation string) {
	sayPrintf(vLv1, "Verbose set to %d\n", verbose)
	sayPrintf(vLv2, "Read in %d numbers\n", len(numbers))
	sayPrintf(vLv2, "Operation: %s\n", operation)

	sum := float64(0)

	for i, n := range numbers {
		sayPrintf(vLv3, "Number (%d): %f\n", i, n)
		sum += n
	}

	if operation == "average" {
		avg := float64(0)
		if len(numbers) > 0 {
			avg = sum / float64(len(numbers))
		}

		sayPrintf(vAll, "Avg: %f\n", avg)
	} else {
		sayPrintf(vAll, "Sum: %f\n", sum)
	}
}
