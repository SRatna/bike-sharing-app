import { Modal } from 'antd';
import { MapContainer, TileLayer, Marker } from 'react-leaflet'
import "leaflet/dist/leaflet.css"
import { useEffect, useRef } from 'react';

const BikeDetailsModal = ({ isModalOpen, closeModal, bike }) => {

  const handleOk = () => {
    closeModal();
  };
  
  const handleCancel = () => {
    closeModal();
  };

  const position = [bike.latitude, bike.longitude];
  const mapRef = useRef(null);

  useEffect(() => {
    setTimeout(() => {
      mapRef.current.invalidateSize()
    }, 150); 
  }, []);

  return (
    <Modal 
      title={bike.name} 
      open={isModalOpen} 
      onOk={handleOk} 
      onCancel={handleCancel}
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
  )
};

export default BikeDetailsModal;