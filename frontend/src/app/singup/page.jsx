'use client'

import Link from "next/link"
import React, { useState } from "react"
import { variables } from "../../../variables"
import styles from './Singup.module.css'

export default function Login() {
    const [user, setuser] = useState('')

    const handleSubmit = async (event) => {
        // Stop the form from submitting and refreshing the page.
        event.preventDefault()

        // Get data from the form.
        const data = {
            username: event.target.username.value,
            email: event.target.email.value,
            password: event.target.password.value
        }

        // Send the data to the server in JSON format.
        const JSONdata = JSON.stringify(data)

        // API endpoint where we send form data.
        const endpoint = variables.backendURL + '/v1/user'

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
            setuser({
                id: result.ID,
                username: result.Username,
                email: result.Email,
                createdAt: result.CreatedAt
            })
        }
    }

    return (
        <section className={styles.section}>
            <div>
                <h2>Sign up</h2>
            </div>
            <div className={styles.form_container}>
                <form onSubmit={handleSubmit} className={styles.formColumn}>
                    <label htmlFor="username">Username</label>
                    <input type="text" id="username" name="username" required />

                    <label htmlFor="email">Email</label>
                    <input type="email" id="email" name="email" required />

                    <label htmlFor="password">Password</label>
                    <input type="password" id="password" name="password" required />

                    <button type="submit">Submit</button>
                </form>
            </div>
            <div>
                <Link href="/login">Ya tienes cuenta?</Link>
            </div>
            {user != '' && <h2>Nevo usuario {user.username} creado</h2>}
        </section>
    )
}


