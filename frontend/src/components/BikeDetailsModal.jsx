import { Modal, message } from 'antd';
import { MapContainer, TileLayer, Marker } from 'react-leaflet'
import "leaflet/dist/leaflet.css"
import { useEffect, useRef } from 'react';

const BikeDetailsModal = ({ isModalOpen, closeModal, bike, updateBikesList }) => {
  const sessionId = sessionStorage.getItem('sessionId');
  const [messageApi, contextHolder] = message.useMessage();

  const updateBike = async () => {
    const { id, rented } = bike;
    const payload = {id, sessionId, rented: !rented };
    const rawResponse = await fetch('/api/bikes', {
      method: 'PATCH',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(payload)
    });
    if (rawResponse.status !== 200) {
      const msg = await rawResponse.text();
      messageApi.error({ content: `Sorry, you ${msg}.` });
      return;
    }
    closeModal();
    updateBikesList(payload);
    if (payload.rented) {
      messageApi.success({ content: 'Bike rented successfully.' });
    } else {
      messageApi.success({ content: 'Bike returned successfully.' });
    }
  }
  
  const position = [bike.latitude, bike.longitude];
  const mapRef = useRef(null);

  useEffect(() => {
    if (!isModalOpen) return;
    setTimeout(() => {
      mapRef.current.invalidateSize();
      mapRef.current.panTo(position);
    }, 150); 
  }, [isModalOpen]);

  return (
    <>
      {contextHolder}
      <Modal 
        title={bike.name} 
        open={isModalOpen} 
        onOk={updateBike} 
        onCancel={closeModal}
        okText={bike.rented ? 'Return' : 'Rent'}
      >
        <MapContainer ref={mapRef} style={{ height: 400 }} center={position} zoom={12}>
          <TileLayer
            attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
            url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          />
          <Marker position={position}>
          </Marker>
        </MapContainer>
      </Modal>
    </>
  )
};

export default BikeDetailsModal;