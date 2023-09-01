def SieveOfEratosthenes(n):
    primes = [True for i in range(n+1)]
    p = 2
    while (p * p <= n):
        if (primes[p] == True):
            for i in range(p * p, n+1, p):
                primes[i] = False
        p += 1
    primeNumbers = [p for p in range(2, n) if primes[p]]
    return primeNumbers

def FindLargest6DigitPrime(primes):
    sixDigitPrimes = [prime for prime in primes if prime >= 100000 and prime <= 999999]
    return max(sixDigitPrimes)

def FindMiddle4DigitPrime(primes):
    fourDigitPrimes = sorted([prime for prime in primes if prime >= 1000 and prime <= 9999])
    middle = len(fourDigitPrimes) // 2
    if len(fourDigitPrimes) % 2:
        return fourDigitPrimes[middle]
    else:
        return (fourDigitPrimes[middle - 1] + fourDigitPrimes[middle]) / 2

def main():
    primes = SieveOfEratosthenes(1000000)
    largest6DigitPrime = FindLargest6DigitPrime(primes)
    print("Largest 6-digit prime:", largest6DigitPrime)
    middle4DigitPrime = FindMiddle4DigitPrime(primes)
    print("Middle 4-digit prime:", middle4DigitPrime)

if __name__ == "__main__":
    main()
