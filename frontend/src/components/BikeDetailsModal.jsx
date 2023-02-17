import { Button, Modal } from 'antd';

const BikeDetailsModal = ({ isModalOpen, closeModal, bike }) => {

  const handleOk = () => {
    closeModal();
  };
  
  const handleCancel = () => {
    closeModal();
  };

  return (
    <Modal 
      title={bike.name} 
      open={isModalOpen} 
      onOk={handleOk} 
      onCancel={handleCancel}
      okText={bike.rented ? 'Return' : 'Rent'}
    >
      <p>Location: {bike.latitude}, {bike.longitude}</p>
    </Modal>
  )
};

export default BikeDetailsModal;