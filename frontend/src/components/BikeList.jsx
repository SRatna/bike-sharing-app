import { List } from 'antd';
import { useEffect, useState } from 'react';

const BikeList = () => {
  const [bikes, setBikes] = useState([]);

  const fetchBikes = async () => {
    const bikesResponse = await fetch('/api/bikes');
    const bikesData = await bikesResponse.json();
    setBikes(bikesData);
  }

  useEffect(() => {
    fetchBikes()
  }, [])

  return (
    <List
      itemLayout="horizontal"
      dataSource={bikes}
      renderItem={(item) => (
        <List.Item>
          <List.Item.Meta
            title={<a>{item.name}</a>}
            description={item.rented ? 'Rented' : 'Available for rent'}
          />
        </List.Item>
      )}
    />
  );
}

export default BikeList;