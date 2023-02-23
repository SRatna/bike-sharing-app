import { List, message } from 'antd';
import { useEffect, useState } from 'react';
import BikeDetailsModal from './BikeDetailsModal';

const BikeList = () => {
  const [bikes, setBikes] = useState([]);
  const [currentBike, setCurrentBike] = useState({});
  const [isBikeDetailsModalOpen, setIsBikeDetailsModalOpen] = useState(false);
  const [messageApi, contextHolder] = message.useMessage();

  const showBikeDetailsModal = (bike) => {
    setCurrentBike(bike);
    setIsBikeDetailsModalOpen(true);
  };

  const closeBikeDetailsModal = () => {
    setIsBikeDetailsModalOpen(false);
  };

  const fetchBikes = async () => {
    const bikesResponse = await fetch('/api/bikes');
    if (bikesResponse.status !== 200) {
      const msg = await bikesResponse.text();
      messageApi.error({ content: msg });
      return;
    }
    const bikesData = await bikesResponse.json();
    setBikes(bikesData);
  }

  const updateBikesList = (newBike) => {
    const updatedBikes = bikes.map((bike) => {
      if (bike.id === newBike.id) {
        return {
          ...bike,
          ...newBike
        }
      }
      return bike;
    });
    setBikes(updatedBikes);
  }

  useEffect(() => {
    fetchBikes();
  }, [])

  const getRentStatus = (bike) => {
    if (!bike.rented) return 'Available for rent';
    const sessionId = sessionStorage.getItem('sessionId');
    if (bike.sessionId !== sessionId) return 'Rented by SOMEONE ELSE';
    return 'Rented by YOU';
  }

  return (
    <>
      {contextHolder}
      <List
        itemLayout="horizontal"
        dataSource={bikes}
        renderItem={(item) => (
          <List.Item 
            style={{ background: item.rented ? '#f3eeee' : '' }}
            actions={[<a key="details" onClick={() => showBikeDetailsModal(item)}>View Details</a>]}
          >
            <List.Item.Meta
              title={<a onClick={() => showBikeDetailsModal(item)}>{item.name}</a>}
              description={getRentStatus(item)}
            />
          </List.Item>
        )}
      />
      <BikeDetailsModal 
        bike={currentBike}
        isModalOpen={isBikeDetailsModalOpen}
        closeModal={closeBikeDetailsModal}
        updateBikesList={updateBikesList}
      />
    </>
  );
}

export default BikeList;