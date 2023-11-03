<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/NodeType.proto

namespace NodeType;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * ------------------------------------
 * -- Metrics
 * ------------------------------------
 * Represents The storage of all the metrics
 *
 * Generated from protobuf message <code>NodeType.Analyze</code>
 */
class Analyze extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.NodeType.Complexity complexity = 1;</code>
     */
    protected $complexity = null;
    /**
     * Generated from protobuf field <code>.NodeType.Volume volume = 2;</code>
     */
    protected $volume = null;
    /**
     * Generated from protobuf field <code>.NodeType.Maintainability maintainability = 3;</code>
     */
    protected $maintainability = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \NodeType\Complexity $complexity
     *     @type \NodeType\Volume $volume
     *     @type \NodeType\Maintainability $maintainability
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\NodeType::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.NodeType.Complexity complexity = 1;</code>
     * @return \NodeType\Complexity|null
     */
    public function getComplexity()
    {
        return $this->complexity;
    }

    public function hasComplexity()
    {
        return isset($this->complexity);
    }

    public function clearComplexity()
    {
        unset($this->complexity);
    }

    /**
     * Generated from protobuf field <code>.NodeType.Complexity complexity = 1;</code>
     * @param \NodeType\Complexity $var
     * @return $this
     */
    public function setComplexity($var)
    {
        GPBUtil::checkMessage($var, \NodeType\Complexity::class);
        $this->complexity = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.NodeType.Volume volume = 2;</code>
     * @return \NodeType\Volume|null
     */
    public function getVolume()
    {
        return $this->volume;
    }

    public function hasVolume()
    {
        return isset($this->volume);
    }

    public function clearVolume()
    {
        unset($this->volume);
    }

    /**
     * Generated from protobuf field <code>.NodeType.Volume volume = 2;</code>
     * @param \NodeType\Volume $var
     * @return $this
     */
    public function setVolume($var)
    {
        GPBUtil::checkMessage($var, \NodeType\Volume::class);
        $this->volume = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.NodeType.Maintainability maintainability = 3;</code>
     * @return \NodeType\Maintainability|null
     */
    public function getMaintainability()
    {
        return $this->maintainability;
    }

    public function hasMaintainability()
    {
        return isset($this->maintainability);
    }

    public function clearMaintainability()
    {
        unset($this->maintainability);
    }

    /**
     * Generated from protobuf field <code>.NodeType.Maintainability maintainability = 3;</code>
     * @param \NodeType\Maintainability $var
     * @return $this
     */
    public function setMaintainability($var)
    {
        GPBUtil::checkMessage($var, \NodeType\Maintainability::class);
        $this->maintainability = $var;

        return $this;
    }

}
