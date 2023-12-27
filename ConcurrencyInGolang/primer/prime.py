import random
import time


def generate_and_check_primes(n):
    def is_prime(num):
        if num < 2:
            return False
        for i in range(2, int(num**0.5) + 1):
            if num % i == 0:
                return False
        return True

    start_time = time.time()

    for _ in range(n):
        random_num = random.randint(0, n)
        if is_prime(random_num):
            print(random_num)

    end_time = time.time()
    execution_time = end_time - start_time
    print(f"Execution time: {execution_time} seconds")


# Example usage:
generate_and_check_primes(500000000)
