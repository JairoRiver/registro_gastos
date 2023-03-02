import { Inter } from 'next/font/google'

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  return (
    <>
      <h1>Control de gastos APP</h1>

      <h3>Para que sirve control de Gastos</h3>

      <p>Con control de gastos podrás subir tus gastos diarios para tener mejor visibilidad y control de tus gastos</p>

      <h4>Proximos Pasos</h4>
      <ul>
        <li>
          <input type='checkbox' /> Dashboard persolizados con tus metricas de gastos
        </li>

        <li>
          <input type='checkbox' /> Creación de grupos para compartir la visibilidad de gastos
        </li>

        <li>
          <input type='checkbox' /> Subida de gastos desde una foto de un ticket
        </li>
      </ul>
    </>
  )
}
