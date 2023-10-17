use bevy::prelude::{Component, Resource, Timer, TimerMode};

#[derive(Resource)]
pub struct PlayerTimer {
    pub red: Timer,
    pub black: Timer,
}

impl PlayerTimer {
    pub fn new(mut min: f32) -> Self {
        return Self {
            // 超时时间, 默认5分钟
            red: Timer::from_seconds(min * 60., TimerMode::Once),
            black: Timer::from_seconds(min * 60., TimerMode::Once),
        };
    }
}

impl Default for PlayerTimer {
    fn default() -> Self {
        return Self::new(5_f32);
    }
}

#[derive(Debug, Clone, Copy, Component, PartialEq, Eq)]
pub struct RedPlayerTimer;

#[derive(Debug, Clone, Copy, Component, PartialEq, Eq)]
pub struct BlackPlayerTimer;
