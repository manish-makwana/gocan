import 'primereact/resources/themes/saga-blue/theme.css';
import 'primereact/resources/primereact.min.css';
import 'primeicons/primeicons.css';
import 'primeflex/primeflex.css';
import {Menu} from "./Menu";
import {Scenes} from "./screens/Scenes";
import {BrowserRouter, Route, Switch} from "react-router-dom";
import {Scene} from "./screens/Scene";
import {AppDetails} from "./screens/AppDetails";

function App() {
  return (
      <BrowserRouter basename="/">
        <div className="App layout-wrapper">
          <div className="layout-topbar">
            <Menu/>
            <Switch>
              <Route path="/scenes/:sceneId/apps/:appId">
                <AppDetails />
              </Route>
              <Route path="/scenes/:sceneId">
                <Scene />
              </Route>
              <Route path="/scenes">
                <Scenes/>
              </Route>
              <Route path="/">
                <Scenes/>
              </Route>
            </Switch>
          </div>
        </div>
      </BrowserRouter>
  );
}

export default App;
