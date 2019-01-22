import Login from "./components/Login";
import ListFiles from "./components/files/ListFiles"
import FileCreate from "./components/files/Create"

export default [
    {
        path: "",
        component: ListFiles
    },
    {
        name: "files.create",
        path: "/files/create",
        component: FileCreate
    },
    {
        name: "login",
        path: "/login",
        component: Login
    }
]