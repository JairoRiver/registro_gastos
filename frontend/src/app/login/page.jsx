'use client'

import Link from "next/link"
import React, { useState } from "react"
import styles from '../singup/Singup.module.css'
import { variables } from "../../../variables"

export default function Login() {
    const [login, setLogin] = useState('')

    const handleSubmit = async (event) => {
        // Stop the form from submitting and refreshing the page.
        event.preventDefault()

        // Get data from the form.
        const data = {
            username: event.target.username.value,
            password: event.target.password.value
        }

        // Send the data to the server in JSON format.
        const JSONdata = JSON.stringify(data)

        // API endpoint where we send form data.
        const endpoint = variables.backendURL + '/v1/user/login'

        // Form the request for sending data to the server.
        const options = {
            // The method is POST because we are sending data.
            method: 'POST',
            // Tell the server we're sending JSON.
            headers: {
                'Content-Type': 'application/json',
            },
            // Body of the request is the JSON data we created above.
            body: JSONdata,
        }

        // Send the form data to our forms API on Vercel and get a response.
        const response = await fetch(endpoint, options)

        // Get the response data from server as JSON.
        if (response.status === 200) {
            const result = await response.json()
            setLogin({
                id: result.user_id,
                username: result.user,
                aToken: result.access_token,
                rtoken: result.refresh_token
            })

            localStorage.setItem("aToken", login.aToken)
            localStorage.setItem("rtoken", login.rtoken)
        }
    }

    return (
        <section className={styles.section}>
            <div>
                <h2>Login</h2>
            </div>
            <div className={styles.form_container}>
                <form onSubmit={handleSubmit} className={styles.formColumn}>
                    <label htmlFor="username">Username</label>
                    <input type="text" id="username" name="username" required />

                    <label htmlFor="password">Password</label>
                    <input type="password" id="password" name="password" required />

                    <button type="submit">Submit</button>
                </form>
            </div>
            <div>
                <Link href="/singup">No tienes cuenta? Registrate!</Link>
            </div>
            {login != '' && <h2>Bienvenido {login.username} ya puedes crear nuevos gastos o ver tus gastos</h2>}
        </section>
    )
}