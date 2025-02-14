import { useEffect, useState } from "react";
import { useNavigate, useParams } from 'react-router-dom'
import PageHeader from "../header/PageHeader"
import axios from "axios";

function FlightEdit() {
  const [flight, setFlight] = useState({
    "id": "001",
    "number": "AI 845",
    "airline_name": "Air India",
    "source": "Mumbai",
    "destination": "Abu dhabi",
    "capacity": 300,
    "price": 5000
  });
  const params = useParams();
  const navigate = useNavigate();
  const readFlightById = async () => {
    try {
      const baseUrl = 'http://127.0.0.1:8080';
      const response = await axios.get(`${baseUrl}/flights/${params.id}`);
      const queriedFlight = response.data;
      setFlight(queriedFlight);
    } catch(error) {
      alert("Server Error")
    }
  };
  const OnChangeBox = (event) => {
    const newFlight = {... flight};
    newFlight[event.target.id] = event.target.value;
    setFlight(newFlight);
  };
  const OnUpdate = async () => {
    try {
      const baseUrl = 'http://127.0.0.1:8080';
      const response = await axios.put(`${baseUrl}/flights/${params.id}`, {...flight, 
                                                            capacity: parseInt(flight.capacity), 
                                                            price: parseFloat(flight.price)});
      const json = response.data;
      setFlight(json.flight);
      alert(json.message);
      navigate("/flights/list");
    } catch(error) {
      alert("Server Error")
    }
  };

  useEffect(() => { readFlightById(); } ,[]);
  return (
    <>
      <PageHeader PageNumber={1} />
      <h3><a href="/flights/list" className="btn btn-light">Go Back</a>Edit Flight Ticket Price</h3>
      <div className="container">
        <div className="form-group mb-3">
          <label htmlFor="number" className="form-label">Flight Number:</label>
          <div className="form-control" id="number">{flight.number}</div>
        </div>
        <div className="form-group mb-3">
          <label htmlFor="airline_name" className="form-label">Airline Name:</label>
          <div type="text" className="form-control" id="airline_name">{flight.airline_name}</div>
        </div>
        <div className="form-group mb-3">
          <label htmlFor="source" className="form-label">Source:</label>
          <div className="form-control" id="source">{flight.source}</div>
        </div>
        <div className="form-group mb-3">
          <label htmlFor="destination" className="form-label">Destination:</label>
          <div className="form-control" id="destination">{flight.destination}</div>
        </div>
        <div className="form-group mb-3">
          <label htmlFor="capacity" className="form-label">Capacity(Number of Seats):</label>
          <div className="form-control" id="capacity">{flight.capacity}</div>
        </div>
        <div className="form-group mb-3">
          <label htmlFor="price" className="form-label">Ticket Price:</label>
          <input type="text" className="form-control" id="price" placeholder="Please enter ticket price"
             value={flight.price}
             onChange={OnChangeBox}/>
        </div>
        <button className="btn btn-warning"
          onClick={OnUpdate}>Update Price</button>
      </div>
    </>
  )
}

export default FlightEdit