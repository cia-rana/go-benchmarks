// The average filter.
sum = 0
for i := 0; i < bpp; i++ {
	cdat3[i] = cdat0[i] - pdat[i]/2
	sum += abs8(cdat3[i])
}
for i := bpp; i < n; i++ {
	cdat3[i] = cdat0[i] - uint8((int(cdat0[i-bpp])+int(pdat[i]))/2)
	sum += abs8(cdat3[i])
	if sum >= best {
		break
	}
}
if sum < best {
	best = sum
	filter = ftAverage
}

return filter
