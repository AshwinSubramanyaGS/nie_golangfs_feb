import { useEffect, useState } from "react";
import PageHeader from "../header/PageHeader"
import axios from 'axios'
function FlightList() {
  const [flights, setFlights] = useState([{
    "id": "001",
    "number": "AI 845",
    "airline_name": "Air India",
    "source": "Mumbai",
    "destination": "Abu dhabi",
    "capacity": 300,
    "price": 5000
  }]);
  
  const readAllFlights = async () => {
    try {
      const baseUrl = 'http://127.0.0.1:8080';
      const response = await axios.get(`${baseUrl}/flights`);
      const queriedFlights = response.data;
      setFlights(queriedFlights);
    } catch(error) {
      alert("Server Error")
    }
  }

  useEffect(() => { readAllFlights(); } ,[]);

  const onDelete = async (id) => {
    if(!confirm("Are you sure to delete?")) {
      return;
    }
    try {
      const baseUrl = 'http://127.0.0.1:8080';
      const response = await axios.delete(`${baseUrl}/flights/${id}`);
      alert(response.data.message);
      readAllFlights();
    } catch(error) {
      alert("Server Error")
    }
  }
  return (
    <>
      <PageHeader PageNumber={1} />
      <h3>List of Flights</h3>
      <div className="container">
        <table className="table table-primary table-striped">
          <thead className="table-dark">
            <tr>
              <th scope="col">Flight Number</th>
              <th scope="col">Airline Name</th>
              <th scope="col">Source</th>
              <th scope="col">Destination</th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            {flights.map( (flight, index) => <>
                <tr key={index}>
                  <th scope="row">{flight.number}</th>
                  <td>{flight.airline_name}</td>
                  <td>{flight.source}</td>
                  <td>{flight.destination}</td>
                  <td>
                    <a href={`/flights/edit/${flight.id}`} className="btn btn-warning">Edit Price</a>
                    <button className="btn btn-danger" onClick={()=>onDelete(flight.id)}>Delete</button>
                  </td>
                </tr>
              </>
              )
            }
            
          </tbody>
        </table>
      </div>
    </>
  )
}

export default FlightList