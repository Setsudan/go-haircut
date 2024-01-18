import './assets/styles/root.scss'
import { useEffect, useState } from 'react'
export default function App() {

  const [serverStatus, setServerStatus] = useState({ code: 0, message: 'Loading...' } as { code: number, message: string })
  useEffect(() => {
    fetch('http://localhost:8080/')
      .then((res) => {
        console.log(res)
        setServerStatus({
          code: res.status,
          message: res.statusText,
        })
      })
      .catch((err) => setServerStatus(err))
  }, [])

  return (
    <div className="App">
      <h1>Go Haircut</h1>
      <pre>{JSON.stringify(serverStatus, null, 2)}</pre>
    </div>
  )
}
