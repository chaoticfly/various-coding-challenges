import re

regex = re.compile("^[a-z0-9]+$")

def check_palindrome(strinp):

    strinp = strinp.lower()
    l = len(strinp) - 1 
    s = 0
    while s<l:
        first = None
        last = None
        if regex.match(strinp[s]):
            first = strinp[s]
        else:
            s += 1
        
        if regex.match(strinp[l]):
            last = strinp[l]
        else:
            l -= 1

        if first and last:
            if first != last:
                return False

            s += 1
            l -= 1
    return True


if __name__ == "__main__":
    print("\"a used a car\" is a palindrome? %s" % check_palindrome("a used a car"))
    print("\"A man, a plan, a canal: Panama\" is a palindrome? %s" % check_palindrome("A man, a plan, a canal: Panama"))
    print("\"Malayalam\" is a palindrome? %s" % check_palindrome("Malayalam"))
    