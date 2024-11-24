from flask import Flask, request, render_template
import hashlib

app = Flask(__name__)

# Данные для авторизации
stored_username = "admin"
with open('password.txt', 'r') as f:
    stored_password_hash = hashlib.sha256(f.read().encode()).hexdigest()

@app.route("/", methods=["GET", "POST"])
def login():
    error = None
    if request.method == "POST":
        username = request.form.get("username")
        password = request.form.get("password")
        password_hash = hashlib.sha256(password.encode()).hexdigest()

        # Проверка имени пользователя и хеша пароля
        if username == stored_username and password_hash[0:6] == stored_password_hash[0:6]:
            # Чтение флага из файла и удаление лишних пробелов
            with open('flag.txt', 'r') as flag_file:
                flag = flag_file.read().strip()
            return flag
        else:
            error = "Неправильное имя пользователя или пароль"

    return render_template("index.html", error=error)

if __name__ == "__main__":
    app.run(debug=True)
