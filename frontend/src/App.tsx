import { useEffect, useState } from "react";

import { createPromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { PaintService } from "./gen/paint/v1/paint_connect";
import { Color, GetColorStreamResponse } from "./gen/paint/v1/paint_pb";
import { proto3 } from "@bufbuild/protobuf";


function getColorFromStr(color:string) {
  return proto3.getEnumType(Color).findName("COLOR_"+color.toUpperCase())?.no
}

const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});

// Here we make the client itself, combining the service
// definition with the transport.
const client = createPromiseClient(PaintService, transport);

export default function App() {
  
  const [gridSize, setGridSize] = useState(3);
  const [gridColors, setGridColors] = useState<string[]>(Array(9).fill(''));
  const colors = ['red', 'blue', 'green'];

  useEffect(() => {
    // declare the async data fetching function
    const initStreams = async () => {
      // get the data from the api
      const colorStream = client.getColorStream({})
      for await (const value of colorStream) {
        console.log(Color[value.color]);
        setGridColors(Array(gridSize*gridSize).fill(Color[value.color]));

      }
      // set state with the result
      // setData(json);
    }

    // call the function
    initStreams()
      // make sure to catch any error
      .catch(console.error);;


  }, []);

  // const updateGridColors = async (color: Color) => {
  //   try {
  //     const apiResponse = await fetchColorsFromApi(color);
  //     setGridColors(apiResponse);
  //   } catch (error) {
  //     console.error('Error fetching colors:', error);
  //   }
  // };



  const handleButtonClick = async (color: string) => {
    const response = await client.changeColor({color: getColorFromStr(color)})
    console.log({response});
    
    // if (response.succeed) {
      // const newColor = await client.getColor({})
      // console.log(Color[newColor.color]);
      
      // setGridColors(Array(gridSize*gridSize).fill(Color[newColor.color]));
      // updateGridColors(color);
    // }
  };

  const handleGridSizeChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const newSize = parseInt(event.target.value, 10);
    setGridSize(newSize);
    // You may want to make an API call here to get new colors based on the updated grid size
    // For simplicity, updating colors with an empty array
    setGridColors(Array(newSize * newSize).fill('GREEN'));
  };

  return (
    <div>
      <div>
        <label>Grid Size:</label>
        <input
          type="number"
          min="1"
          max="10"
          value={gridSize}
          onChange={handleGridSizeChange}
        />
      </div>
      <div>
        {colors.map((color) => (
          <button key={color} onClick={() => handleButtonClick(color)}>
            {color}
          </button>
        ))}
      </div>
      <div style={{ display: 'grid', gridTemplateColumns: `repeat(${gridSize}, 1fr)` }}>
        {gridColors.map((gridColor, index) => (
          <div
            key={index}
            style={{
              width: '50px',
              height: '50px',
              backgroundColor: gridColor,
              border: '1px solid black',
            }}
          />
        ))}
      </div>
    </div>
  );
}
