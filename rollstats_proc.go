let _= 0
let rndcnt = 0




    
function printNum(n, outp) {   
    let i = 0
    for 1 {
        let digit = n % 10
        outp = setbyte(outp, 2-i, digit+48)
        n = n / 10		
        if n == 0 {
            apps[0].put("dbg", concat(apps[0].get("dbg"), outp ))
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
    
    bigRand = btoi(substring(hash,rndcnt,7+rndcnt)) 
    rndcnt = rndcnt + 1                             
    let roll1 =  1 + bigRand % (6 -  1 + 1)
    
    bigRand = btoi(substring(hash,rndcnt,7+rndcnt)) 
    rndcnt = rndcnt + 1                             
    let roll2 =  1 + bigRand % (6 -  1 + 1)
    
    bigRand = btoi(substring(hash,rndcnt,7+rndcnt)) 
    rndcnt = rndcnt + 1                             
    let roll3 =  1 + bigRand % (6 -  1 + 1)
    
    bigRand = btoi(substring(hash,rndcnt,7+rndcnt)) 
    rndcnt = rndcnt + 1                             
    let roll4 =  1 + bigRand % (6 -  1 + 1)

    apps[0].put("dbg", concat(apps[0].get("dbg"), setbyte( "    ", 0, roll1+48) ))
    apps[0].put("dbg", concat(apps[0].get("dbg"), setbyte( "    ", 0, roll2+48) ))
    apps[0].put("dbg", concat(apps[0].get("dbg"), setbyte( "    ", 0, roll3+48) ))
    apps[0].put("dbg", concat(apps[0].get("dbg"), setbyte( "  = ", 0, roll4+48) ))
    let stat = 0
    if roll1 < drop { drop = roll1 }
    if roll2 < drop { drop = roll2 }
    if roll3 < drop { drop = roll3 }
    if roll4 < drop { drop = roll4 }
    stat = roll1 + roll2 + roll3 + roll4 - drop
    let bonus = 0
    apps[0].put("dbg", concat(apps[0].get("dbg"), "\x1B[38;5;0m" ))
    apps[0].put("dbg", concat(apps[0].get("dbg"), "\x1B[48;5;15m" ))
    _= printNum(stat, "    ")
    apps[0].put("dbg", concat(apps[0].get("dbg"), " " ))
    apps[0].put("dbg", concat(apps[0].get("dbg"), "\x1B[38;5;15m" )) 
    apps[0].put("dbg", concat(apps[0].get("dbg"), "\x1B[48;5;0m" ))    
    apps[0].put("dbg", concat(apps[0].get("dbg"), " " ))
    if stat > 11 {
        bonus = (stat - 10) / 2
        apps[0].put("dbg", concat(apps[0].get("dbg"), "\x1B[38;5;11m" )) 
        apps[0].put("dbg", concat(apps[0].get("dbg"), "[+" ))
        apps[0].put("dbg", concat(apps[0].get("dbg"), setbyte( " ", 0, bonus+48) ))
        apps[0].put("dbg", concat(apps[0].get("dbg"), "] " ))    	
    }
    if stat < 10 {
        bonus = ((10 - stat + 1) / 2)
        apps[0].put("dbg", concat(apps[0].get("dbg"), "\x1B[38;5;1m" )) 
        apps[0].put("dbg", concat(apps[0].get("dbg"), "[-" ))
        apps[0].put("dbg", concat(apps[0].get("dbg"), setbyte( " ", 0, bonus+48) ))
        apps[0].put("dbg", concat(apps[0].get("dbg"), "] " ))
    }
    apps[0].put("dbg", concat(apps[0].get("dbg"), " " ))
    return 1
}
