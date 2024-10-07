# Exoplanet Microservice

## Overview

This microservice is designed to support space voyagers embarking on a journey to study different exoplanets outside of
our solar system. Exoplanets are categorized into two types: Gas Giants and Terrestrial planets. The microservice
provides functionality to manage exoplanet data, including adding, listing, getting by ID, updating, and deleting
exoplanets. Additionally, it offers a fuel estimation feature to calculate the fuel cost for a trip to a specific
exoplanet based on crew capacity.

## Features

1. **Add Exoplanet**: Users can add a new exoplanet by providing its name, description, distance from Earth, radius,
   mass (for Terrestrial planets only), and type of exoplanet (Gas Giant or Terrestrial).

2. **List Exoplanets**: Retrieve a list of all available exoplanets.

3. **Get Exoplanet by ID**: Retrieve information about a specific exoplanet by its unique ID.

4. **Update Exoplanet**: Update the details of an existing exoplanet.

5. **Delete Exoplanet**: Remove an exoplanet from the catalog.

6. **Fuel Estimation**: Retrieve an overall fuel cost estimation for a trip to any particular exoplanet for a given crew
   capacity. The fuel estimation formula is provided.

    ```
    f = d / (g^2) * c units
    ```
   Where:
    - `d`: Distance of exoplanet from Earth
    - `g`: Gravity of exoplanet
    - `c`: Crew capacity (integer)

   The logic to calculate gravity for each type is as follows:
    - Gas Giant: `g = (0.5 / r^2)`
    - Terrestrial: `g = (m / r^2)` (where `m` is mass and `r` is radius)

## Setup

1. **Clone Repository**: Clone this repository to your local machine.

2. **Build Docker Image**: Build the Docker image for the microservice.

3. **Run Docker Container**: Run the Docker container from the built image.

4. **Access Endpoints**: Access the microservice endpoints using HTTP requests.

## Endpoints

- **Add Exoplanet**: `PUT /add-exoplanet`
    - Payload:
      ```json
      {
          "name": "Exoplanet Name",
          "description": "Description of the exoplanet",
          "distanceFromEarth": 123,
          "radius": 67.89,
          "mass": 123.45, (optional, required only for Terrestrial planets)
          "typeOfExoplanet": "GasGiant" or "Terrestrial"
      }
      ```

- **List Exoplanets**: `GET /list-exoplanet`

- **Get Exoplanet by ID**: `GET /get-exoplanet?id={id}`

  - **Update Exoplanet**: `POST /update-exoplanet?id={id}`
      - Payload:
        ```json
        {
            "inputFields": [
                {
                    "key": "name",
                    "value": "Proxima Centauri"
                },
                {
                    "key": "radius",
                    "value": 1.2
                }
            ]
        }
    ```

- **Delete Exoplanet**: `DELETE /delete-exoplanet?id={id}`

- **Fuel Estimation**: `GET /get-fuel-estimation?id={id}&crewSize={crew_capacity}`

## Dependencies

- Golang
- MySQL
- Docker

## Contributors

- [Meghansh Bansal](https://github.com/MeghanshBansal)

