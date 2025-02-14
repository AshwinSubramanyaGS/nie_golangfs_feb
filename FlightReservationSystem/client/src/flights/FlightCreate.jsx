import axios from "axios";
import PageHeader from "../header/PageHeader"
import { useState } from "react";
import { useNavigate } from "react-router-dom";

function FlightCreate() {
  const [flight, setFlight] = useState({
    "id": "",
    "number": "",
    "airline_name": "",
    "source": "",
    "destination": "",
    "capacity": 0,
    "price": 0
  });
  const navigate = useNavigate();
  const OnChangeBox = (event) => {
    const newFlight = {... flight};
    newFlight[event.target.id] = event.target.value;
    setFlight(newFlight);
  }
  const OnCreate = async () => {
    try {
      const baseUrl = 'http://127.0.0.1:8080';
      const response = await axios.post(`${baseUrl}/flights`, {...flight, 
                                                            capacity: parseInt(flight.capacity), 
                                                            price: parseFloat(flight.price)});
      const json = response.data;
      setFlight(json.flight);
      alert(json.message);
      navigate("/flights/list");
    } catch(error) {
      alert("Server Error")
    }
  }
  return (
    <>
      <PageHeader PageNumber={2} />
      <h3><a href="/flights/list" className="btn btn-light">Go Back</a>New Flight</h3>
      <div className="container">
        <div className="form-group mb-3">
          <label htmlFor="number" className="form-label">Flight Number:</label>
          <input type="text" className="form-control" id="number" placeholder="Please enter flight number" 
            value={flight.number}
            onChange={OnChangeBox}/>
        </div>
        <div className="form-group mb-3">
          <label htmlFor="airline_name" className="form-label">Airline Name:</label>
          <input type="text" className="form-control" id="airline_name" placeholder="Please enter airline name" 
            value={flight.airline_name}
            onChange={OnChangeBox}/>
        </div>
        <div className="form-group mb-3">
          <label htmlFor="source" className="form-label">Source:</label>
          <input type="text" className="form-control" id="source" placeholder="Please enter source" 
            value={flight.source}
            onChange={OnChangeBox}/>
        </div>
        <div className="form-group mb-3">
          <label htmlFor="destination" className="form-label">Destination:</label>
          <input type="text" className="form-control" id="destination" placeholder="Please enter destination" 
            value={flight.destination}
            onChange={OnChangeBox}/>
        </div>
        <div className="form-group mb-3">
          <label htmlFor="capacity" className="form-label">Capacity(Number of Seats):</label>
          <input type="text" className="form-control" id="capacity" placeholder="Please enter capacity" 
            value={flight.capacity}
            onChange={OnChangeBox}/>
        </div>
        <div className="form-group mb-3">
          <label htmlFor="price" className="form-label">Ticket Price:</label>
          <input type="text" className="form-control" id="price" placeholder="Please enter ticket price" 
            value={flight.price}
            onChange={OnChangeBox}/>
        </div>
        <button className="btn btn-success"
          onClick={OnCreate}>Create Flight</button>
      </div>
    </>
  )
}

export default FlightCreate