const checkboxes = document.querySelectorAll('input[type="checkbox"]')
checkboxes.forEach((checkbox) => {
    checkbox.addEventListener('change', (event) => {
        let targetId = event.target.id
        let id = targetId.match(/\d+/)
        if (event.target.checked) {
            let url = "http://localhost:8080/todo/complete?id="
            fetch(url + id, {
                method: "post"
            }).then(response => {
                if (response.ok) {
                    document.getElementById("completed" + id).checked = true
                }
            }).catch((error) => {
                console.error(error)
            })
        } else {
            let url = "http://localhost:8080/todo/uncompleted?id="
            fetch(url + id, {
                method: "post"
            }).then(response => {
                if (response.ok) {
                    document.getElementById("completed" + id).checked = false
                }
            }).catch((error) => {
                console.error(error)
            })
        }
    })
})

function deleteTask(key) {
    let url = "http://localhost:8080/todo/delete?id=" + encodeURIComponent(key)
    fetch(url, {
        method: "delete"
    }).then(response => {
        if (response.ok) {
            document.querySelector(`[data-key='${key}']`).remove();
            document.querySelector(`.tasks[data-key='${key}']`).remove();
        } else {
            console.error("Error deleting task")
        }
    });
}

function updateTask(id) {
    const newTask = {
        id: id,
        title: document.getElementById("title" + id).value,
        description: document.getElementById("description" + id).value,
        dueDate: "2006-01-02T15:04:05Z"
    };
    fetch("http://localhost:8080/todo/create", {
        method: "PUT",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(newTask)
    }).then(response => {
        response.json()
    }).catch((error) => {
        console.error("Error: ", error)
    });
}

function createTask() {
    const newTask = {
        id: 4,
        title: document.getElementById("title").value,
        description: document.getElementById("description").value,
        dueDate: "2006-01-02T15:04:05Z"
    };
    fetch("http://localhost:8080/todo/create", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(newTask)
    }).then(response => {
        response.json()
    }).catch((error) => {
        console.error("Error: ", error)
    });
}
