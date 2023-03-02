import './globals.css'
import { Navigation } from './components/Navigation'

export default function RootLayout({ children }) {
  return (
    <html lang="es">
      <title>cost control</title>
      <body>
        <Navigation />
        {children}
      </body>
    </html>
  )
}
