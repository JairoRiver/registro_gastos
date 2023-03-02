import Link from "next/link"

export default function Login() {
    return (
        <section>
            <div>
                <h2>Login</h2>
            </div>
            <div>
                <p>Aquí vamos a añadir el formulario</p>
            </div>
            <div>
                <p>Aquí vamos a añadir la opción de recuperar contraseña</p>
                <Link href="/singup">Crear una cuenta</Link>
            </div>
        </section>
    )
}