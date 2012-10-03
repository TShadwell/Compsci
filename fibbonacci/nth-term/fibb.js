/*
 * A very succint solution that breaks esily because of all the function calls.
 * returns undefined for terms of the sequence that are, theoretically, undefined.
 */
function nthFibb(num){
	return (
		num<2?
			//By the definition of the sequence
			[0, 1][num]
		:
			//Sum of the previous two numbers in the sequence
			nthFibb(num-1)
			+
			nthFibb(num-2)
	);
}
/*
 * Another solution, including the 'negafibbonacci' numbers
 */
 function nthFibb(num){
	return (
		num<2?
			num>=0?
				num
			:
				//F[n]=F[n+1]-F[n+2]
				nthFibb(num+2)
				-
				nthFibb(num+1)
		:
			//F[n]=F[n-1]+F[n-2]
			nthFibb(num-1)
			+
			nthFibb(num-2)
	);
}

