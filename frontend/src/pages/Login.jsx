import { useState } from 'react'
import { useNavigate } from 'react-router-dom'

export default function LoginPage() {
  const [login, setLogin] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')
  const navigate = useNavigate()

  const handleSubmit = async (e) => {
    e.preventDefault()
    setError('')

    try {
      const response = await fetch(`/auth/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ login, password })
      })

      if (!response.ok) throw new Error('Неверные данные')

      const data = await response.json()
      localStorage.setItem('token', data.token)
      navigate('/admin')
    } catch (err) {
      setError(err.message)
    }
  }

  return (
    <div style={styles.container}>
      <h2 style={styles.title}>Вход</h2>
      <form onSubmit={handleSubmit} style={styles.form}>
        <label style={styles.label}>Логин:</label>
        <input
          type="text"
          value={login}
          onChange={(e) => setLogin(e.target.value)}
          style={styles.input}
          required
        />

        <label style={styles.label}>Пароль:</label>
        <input
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          style={styles.input}
          required
        />

        <button type="submit" style={styles.button}>Войти</button>
        {error && <p style={styles.error}>{error}</p>}
      </form>
    </div>
  )
}

const styles = {
  container: {
    maxWidth: '400px',
    margin: '80px auto',
    padding: '30px',
    backgroundColor: '#f9f9f9',
    borderRadius: '12px',
    boxShadow: '0 0 8px rgba(0,0,0,0.1)',
    fontFamily: 'sans-serif'
  },
  title: {
    textAlign: 'center',
    marginBottom: '20px'
  },
  form: {
    display: 'flex',
    flexDirection: 'column',
    gap: '15px'
  },
  label: {
    fontWeight: '600'
  },
  input: {
    padding: '12px',
    borderRadius: '10px',
    border: '1px solid #ccc',
    backgroundColor: '#2d2d2d',
    color: '#fff',
    fontSize: '16px'
  },
  button: {
    padding: '12px',
    borderRadius: '10px',
    border: '2px solid #5b7bfa',
    backgroundColor: 'transparent',
    color: '#5b7bfa',
    fontWeight: '600',
    cursor: 'pointer',
    fontSize: '16px',
    transition: '0.2s'
  },
  error: {
    color: 'red',
    fontSize: '14px',
    marginTop: '10px',
    textAlign: 'center'
  }
}
