import { Breadcrumb, Layout, theme } from 'antd';
import { useEffect } from 'react';
const { Header, Content, Footer } = Layout;
import './App.css';
import BikeList from './components/BikeList';
import { v4 as uuidv4 } from 'uuid';

const App = () => {
  const {
    token: { colorBgContainer },
  } = theme.useToken();
 
  useEffect(() => {
    const sessionId = sessionStorage.getItem('sessionId');
    if (!sessionId) {
      sessionStorage.setItem('sessionId', uuidv4());
    }
  }, [])

  return (
    <Layout className="layout">
      <Header>
        <div className="logo">
          Bike Sharing App
        </div>
      </Header>
      <Content
        style={{
          padding: '0 50px',
        }}
      >
        <Breadcrumb
          style={{
            margin: '16px 0',
          }}
        >
          <Breadcrumb.Item>Home</Breadcrumb.Item>
          <Breadcrumb.Item>Bikes</Breadcrumb.Item>
        </Breadcrumb>
        <div
          className="site-layout-content"
          style={{
            background: colorBgContainer,
          }}
        >
          <BikeList />
        </div>
      </Content>
      <Footer
        style={{
          textAlign: 'center',
        }}
      >
        Bike Sharing App ©2023 Created by SRatna
      </Footer>
    </Layout>
  );
};
export default App;