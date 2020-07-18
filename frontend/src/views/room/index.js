import React, { useContext } from 'react'
import { connect } from 'react-redux'
import PropTypes from 'prop-types'
import GroupGrid from '../../components/group_grid'
import { WebSocketContext } from '../../websocket'

const Room = ({ groups }) => {
  const ws = useContext(WebSocketContext)

  return <GroupGrid groups={groups} onGroupClick={(id) => ws.sendJoin(id)} />
}

Room.propTypes = {
  groups: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.string.isRequired,
      name: PropTypes.string.isRequire,
      // members: PropTypes.arrayOf(PropTypes.shape({}).isRequired),
      num_members: PropTypes.number.isRequired,
    }).isRequired
  ).isRequired,
}

const mapStateToProps = (state) => {
  const sortedGroups = state.room.groups
    .slice(0)
    .sort((left, right) => (left.name > right.name ? 1 : -1))

  return {
    groups: sortedGroups,
  }
}

export default connect(mapStateToProps)(Room)
