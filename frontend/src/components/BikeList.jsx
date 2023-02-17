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
  };

  const fetchBikes = async () => {
    const bikesResponse = await fetch('/api/bikes');
    const bikesData = await bikesResponse.json();
    setBikes(bikesData);
  }

  useEffect(() => {
    fetchBikes()
  }, [])

  return (
    <>
      <List
        itemLayout="horizontal"
        dataSource={bikes}
        renderItem={(item) => (
          <List.Item>
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
      />)}
    </>
  );
}

export default BikeList;