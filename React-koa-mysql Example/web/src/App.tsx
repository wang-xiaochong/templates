import React, { useState, useEffect } from 'react';
// import ReactDOMServer from 'react-dom/server'
// import logo from './logo.svg';
import './App.css';
import { getAllBanners } from './models/banner'
import { BannerData } from 'models/banner'

function App() {

  const [banners, setBanners] = useState<BannerData[]>([]);
  useEffect(() => {
    (async () => {
      let banners = await getAllBanners()
      setBanners(banners)
    })();
  }, [])
  console.log(banners)



  return (
    <div>
      {/* aaa */}
      <ul>
        {banners.map(banner =>
          <li key={banner.ID}>{banner.href}</li>
        )}
      </ul>
    </div>
    // <div className="App">
    //   <header className="App-header">
    //     <img src={logo} className="App-logo" alt="logo" />
    //     <p>
    //       Edit <code>src/App.tsx</code> and save to reload.
    //     </p>
    //     <a
    //       className="App-link"
    //       href="https://reactjs.org"
    //       target="_blank"
    //       rel="noopener noreferrer"
    //     >
    //       Learn React
    //     </a>
    //   </header>
    // </div>
  );
}

export default App;

// export function render() {
//   return ReactDOMServer.renderToString(<App />)
//   // return ReactDOMServer.renderToStaticMarkup(<App/>)
// }
