import random
import string

length = 35
characters = string.ascii_lowercase + string.digits

random_string = ''.join(random.choices(characters, k=length))

print(random_string)
