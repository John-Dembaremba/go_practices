import random
import time

def generate_random_num(rand_nums):
    return random.randint(0, rand_nums - 1)

def is_prime(num):
    if num < 2:
        return False
    for i in range(num - 1, 1, -1):
        if num % i == 0:
            return False
    return True

def is_divisible_by_two(num):
    return num % 2 == 0

def array_nums(rand_nums, prime_nums, even_nums, odd_nums):
    prime_list, even_list, odd_list = [], [], []
    
    while True:
        rand_num = generate_random_num(rand_nums)
        if is_prime(rand_num) and len(prime_list) < prime_nums:
            prime_list.append(rand_num)
        elif is_divisible_by_two(rand_num) and len(even_list) < even_nums:
            even_list.append(rand_num)
        elif not is_divisible_by_two(rand_num) and len(odd_list) < odd_nums:
            odd_list.append(rand_num)
        
        if len(prime_list) == prime_nums and len(even_list) == even_nums and len(odd_list) == odd_nums:
            break
    
    return prime_list, even_list, odd_list

def main():
    start_time = time.time()
    prime_nums, even_nums, odd_nums = array_nums(1000000000, 3, 4, 4)

    print(prime_nums)
    print(even_nums)
    print(odd_nums)
    print(time.time() - start_time)

if __name__ == "__main__":
    main()
