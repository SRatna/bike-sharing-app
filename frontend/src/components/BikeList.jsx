import { List } from 'antd';
import { useEffect, useState } from 'react';
import BikeDetailsModal from './BikeDetailsModal';

const BikeList = () => {
  const [bikes, setBikes] = useState([]);
  const [currentBike, setCurrentBike] = useState(null);
  const [isBikeDetailsModalOpen, setIsBikeDetailsModalOpen] = useState(false);

  const showBikeDetailsModal = (bike) => {
    setCurrentBike(bike);
    setIsBikeDetailsModalOpen(true);
  };

  const closeBikeDetailsModal = () => {
    setIsBikeDetailsModalOpen(false);
    setCurrentBike(null);
  };

  const fetchBikes = async () => {
    const bikesResponse = await fetch('/api/bikes');
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

  return (
    <>
      <List
        itemLayout="horizontal"
        dataSource={bikes}
        renderItem={(item) => (
          <List.Item 
            style={{ background: item.rented ? '#f3eeee' : '' }}
          >
            <List.Item.Meta
              title={<a onClick={() => showBikeDetailsModal(item)}>{item.name}</a>}
              description={item.rented ? 'Rented' : 'Available for rent'}
            />
          </List.Item>
        )}
      />
      {currentBike && (<BikeDetailsModal 
        bike={currentBike}
        isModalOpen={isBikeDetailsModalOpen}
        closeModal={closeBikeDetailsModal}
        updateBikesList={updateBikesList}
      />)}
    </>
  );
}

export default BikeList;