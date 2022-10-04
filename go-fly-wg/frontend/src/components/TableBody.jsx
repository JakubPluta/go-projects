import { useState, useEffect } from 'react'
import TableRow from './TableRow'


const TableBody = () => {
    const [flights, setFlights] = useState([])



    const getFlights = () => {
        fetch("http://localhost:8000/flights").then(res => res.json()).then(
            f =>setFlights(Object.values(f))
        ).catch(err => console.log(err))
        
    }

    useEffect(() => getFlights(), [])

        return (
        <tbody>
        {flights?.map((flight, i) => (
            <TableRow key={i} flight={flight} />
        ))}
        </tbody>
    )
}


export default TableBody