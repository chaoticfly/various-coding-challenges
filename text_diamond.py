


def text_diamond(num,space=" ", control="*"):
    if num and num % 2 == 0:
        return None
    diamond = ''
    for i in range( num * 2, 0, -2):
        if i >= num:
            i = i - num - 1
            spaces = repeater(space,i/2)
            stars = repeater(control,num - i)
            diamond += spaces + stars + spaces + "\n"
        else:
            i = num - i + 1
            spaces = repeater(space,i/2)
            stars = repeater(control,num - i)
            diamond += spaces + stars + spaces + "\n"
    return diamond.strip("\n")

def repeater(character, number):
    temp = ''
    for i in range(number):
        temp = temp + character
    return temp


if __name__ == "__main__":
    k = input("Enter the number of Stars (odd number):")
    print(text_diamond(k))

