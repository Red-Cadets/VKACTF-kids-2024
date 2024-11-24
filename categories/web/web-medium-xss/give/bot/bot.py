from selenium import webdriver
import time
import random
import string


def verify_solution(task, solution):

    try:    
        test = ":".join(solution.split(":")[3:9]).lower()

        print(f"Task: {task}")
        print(f"Solution: {solution}")
        print(f"Test: {test}")

        if task.lower() == test:

            
            return True
        else:
            return False
    except FileNotFoundError:
        print("Hashcash is not installed on this system.")
        return False
    
    
def generate_task(resource, bits):
    version = 1
    
    date = time.strftime("%y%m%d%H%M%S")
    
    id = ''.join(random.choices(string.ascii_letters + string.digits, k=10)) 
    extra_info = ''.join(random.choices(string.ascii_letters, k=5))
    
    return f"{version}:{bits}:{date}:{resource}:{id}:{extra_info}"


def visit(url):
    with open("flag", "r") as f:
        flag = f.read()
        
    try:
        url = "http://go_app_container:60888" + url
        driver.get(url)
        driver.add_cookie({"name": "flag", "value": flag})
        driver.get(url)
        time.sleep(3)
    except:
        pass
    finally:
        driver.quit()



if __name__ == "__main__":

    resource = "user_allow"
    bits = 20

    task = generate_task(resource, bits)
    print(f"Task generated. Solve the hashcash task:\n{task}\n")

    user_solution = input("Enter the solution hashcash task: ").strip()
    
    if not verify_solution(task, user_solution):
        print("Invalid solution")
        exit(1)        

    print("Solving the hashcash task...")
    time.sleep(3)
    print("Congratulations! You have solved the hashcash task.")

    options = webdriver.ChromeOptions()
    options.add_argument('--headless')
    options.add_argument('--no-sandbox')
    options.add_argument('--disable-dev-shm-usage')
    options.add_argument('--remote-debugging-port=9222')
    
    driver = webdriver.Chrome(options=options)
    url = input("input url, EXAMPLE: /MzQ1My0wNS0yMzIzOjUzY2FyZGlvbG9naXN0 > ")

    if url.startswith("/"):
        visit(url)
    else:
        print("Incorrect URL, please start with \'/\'")