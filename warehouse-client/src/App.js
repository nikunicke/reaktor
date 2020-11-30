import React, { useState, useEffect } from 'react'
import jacketsService from './services/jackets'
import shirtsService from './services/shirts'
import accessoriesService from './services/accessories'
import Container from 'react-bootstrap/Container'
import Tabs from 'react-bootstrap/Tabs'
import Tab from 'react-bootstrap/Tab'
import BootstrapTable from 'react-bootstrap-table-next';
import paginationFactory from 'react-bootstrap-table2-paginator';


const App = () => {
  const [ jackets, setJackets ] = useState([])
  const [ shirts, setShirts ] = useState([])
  const [ accessories, setAccessories ] = useState([])

  useEffect(() => {
    jacketsService.getAll().then(
      jackets => setJackets(jackets)
    )
  }, [])
  useEffect(() => {
    shirtsService.getAll().then(
      shirts => setShirts(shirts)
    )
  }, [])
  useEffect(() => {
    accessoriesService.getAll().then(
      accessories => setAccessories(accessories)
    )
  }, [])

  const columns = [{
    dataField: "id",
    text: "ID"
  }, {
    dataField: "name",
    text: "Name"
  }, {
    dataField: "color",
    text: "Color"
  }, {
    dataField: "price",
    text: "Price"
  }, {
    dataField: "manufacturer",
    text: "Manufacturer"
  }, {
    dataField: "availability",
    text: "Availability",
    style: function callback(cell, row, rowIndex, colIndex) {
      if (cell === "INSTOCK") {
        return ({ backgroundColor: "#5cb85c" })
      } else if (cell === "LESSTHAN10") {
        return ({ backgroundColor: "#f0ad4e" })
      } else if (cell === "OUTOFSTOCK") {
        return ({ backgroundColor: "#d9534f"})
      }
    }
  }]



  return (
    <div className="App">
      <Container fluid style={{maxWidth: "1300px"}}>
        <center><h1>reaktor warehouse</h1></center>
        <Tabs defaultActiveKey="jackets" id="uncontrolled-tab">
          <Tab eventKey="jackets" title="Jackets">
            <BootstrapTable keyField="id" data={jackets} columns={columns} pagination={paginationFactory({})}/>
          </Tab>
          <Tab eventKey="shirts" title="Shirts">
           <BootstrapTable keyField="id" data={shirts} columns={columns} pagination={paginationFactory({})}/>
          </Tab>
          <Tab eventKey="accessories" title="Accessories">
           <BootstrapTable keyField="id" data={accessories} columns={columns} pagination={paginationFactory({})}/>
          </Tab>
        </Tabs>
      </Container>
    </div>
  )
}

export default App;
