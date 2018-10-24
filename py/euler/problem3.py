def is_prime(i):
    if i == 1 or i == 2:
        return True

    for j in range(2, i//2):
        if i % j == 0:
            return False
    return True


def largest_prime_factor(num):
    for i in range(num // 2, 2, -1):
        if num % i == 0 and is_prime(i):
            return i


def prime_factors(num):
    return [i for i in range(2, num//2) if num % i == 0 and is_prime(i)]


print(largest_prime_factor(600851475143))