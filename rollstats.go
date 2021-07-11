let _= 0
let rndcnt = 0

#define print(x) apps[0].put("dbg", concat(apps[0].get("dbg"), x ))

#define printOneDigit(n, outp) print(setbyte(outp, 0, n+48))

#define fgGreen print(\"\\x1B[38;5;2m") 
#define fgWhite print(\"\\x1B[38;5;15m") 
#define bgWhite print("\\x1B[48;5;15m")
#define fgBlack print("\\x1B[38;5;0m")
#define bgBlack print("\\x1B[48;5;0m")
#define bgBlue print("\\x1B[48;5;4m")
#define fgYellow print(\"\\x1B[38;5;11m") 
#define fgRed print(\"\\x1B[38;5;1m") 

#define getRand(var1, min, max) \
    bigRand = btoi(substring(hash,rndcnt,7+rndcnt)) \
    rndcnt = rndcnt + 1                             \
    let var1 = min + bigRand % (max - min + 1)
    
function printNum(n, outp) {   
    let i = 0
	for 1 {
		let digit = n % 10
		outp = setbyte(outp, 2-i, digit+48)
		n = n / 10		
		if n == 0 {
		    print(outp)
			return 1
		}		
		i = i + 1
	}	
	return 1
}

function logic() { 
    apps[0].put("dbg"," ")
    let drop = 99
	let hash = sha256(txn.TxId)
	let bigRand = 0
	getRand(roll1, 1,6)
	getRand(roll2, 1,6)
	getRand(roll3, 1,6)
	getRand(roll4, 1,6)

	printOneDigit(roll1, "    ")
	printOneDigit(roll2, "    ")
	printOneDigit(roll3, "    ")
	printOneDigit(roll4, "  = ")
	let stat = 0
	if roll1 < drop { drop = roll1 }
	if roll2 < drop { drop = roll2 }
	if roll3 < drop { drop = roll3 }
	if roll4 < drop { drop = roll4 }
	stat = roll1 + roll2 + roll3 + roll4 - drop
    let bonus = 0
    fgBlack
    bgWhite
    _= printNum(stat, "    ")
    print(" ")
    fgWhite
    bgBlack    
    print(" ")
    if stat > 11 {
    	bonus = (stat - 10) / 2
    	fgYellow
    	print("[+")
    	printOneDigit(bonus, " ", 1)
    	print("] ")    	
    }
    if stat < 10 {
    	bonus = ((10 - stat + 1) / 2)
    	fgRed
    	print("[-")
    	printOneDigit(bonus, " ", 1)
    	print("] ")
    }
    print(" ")
	return 1
}
